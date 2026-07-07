---
name: updating-tetrate-dependency
description: Use when bumping the github.com/tetrateio/tetrate module version in this repo, after the upstream TSB monorepo's API changes, or when tsb_* resources/fields are missing, stale, or fail to generate.
---

# Updating the tetrateio/tetrate dependency

## Overview

`github.com/tetrateio/tetrate` is the private monorepo vendoring the TSB v2 API
proto descriptors. `cmd/gen-provider` reads them via reflection to generate
`internal/provider/*_gen.go`. Bumping the dependency and regenerating is the
only supported way to pick up new/changed TSB resources — never hand-edit
`*_gen.go`.

## Step 0 — check for a local replace

```sh
grep replace go.mod
```

- **Found** `replace github.com/tetrateio/tetrate => /some/local/path` — this is
  a dev-only local checkout; go.mod's `require` version is ignored while it's
  active. "Updating" means updating that checkout instead (`git -C <path>
  fetch && git -C <path> checkout <ref>`), then skip straight to Step 2. Do
  **not** bump the `require` line while the replace is present — it silently
  does nothing.
- **Not found** — normal versioned dependency, do Step 1.

## Step 1 — bump the module

This is a private repo; use the same auth CI uses
(`.github/workflows/go.yml`):

```sh
git config --global url."git@github.com:".insteadOf "https://github.com/"
export GOPRIVATE=github.com/tetrateio/*
go get github.com/tetrateio/tetrate@<tag-or-commit-sha>
go mod tidy
```

Prefer a real tag or a full 40-char SHA. Watch what else `go mod tidy` shifts
(`google.golang.org/genproto`, `grpc`, `protobuf`, `istio.io/api`) — a version
skew between this repo's pins and what the new commit itself depends on causes
duplicate-proto-registration panics at generate/runtime, not a clean compile
error.

## Step 2 — check for new API groups

`cmd/gen-provider/imports.go` blank-imports every TSB API *group* package.
Discovery of *resources* within an already-imported group is fully automatic,
but a **brand-new group package needs a manual import line here** — nothing
else will notice it exists. Diff the vendored module's `api/tsb/*/v2` package
list against `imports.go`'s import block.

## Step 3 — dry-run before writing anything

```sh
go run ./cmd/gen-provider -dry-run
```

Compare the printed resource list — and any `gen-provider: skipping candidate
...` warnings on stderr — against the current `ls internal/provider/*_resource_gen.go`.
Discovery silently *skips* non-conforming services/messages; a resource that
disappears here is a breaking change for existing Terraform users, not a build
error, so it needs investigating before you regenerate for real.

## Step 4 — regenerate and verify

```sh
make check    # generate + docs + format, fails if that produced any diff
make test
```

- `make docs` needs a real `terraform` binary on `PATH` — a failure here is
  unrelated to the dependency bump.
- `make check`'s dirty-tree check fails on *any* uncommitted file, not just
  generated ones — start from a clean `git status` before running it.
- Review `internal/provider/*_gen.go` diffs for attribute renames/removals on
  existing resources (breaking for users), not just "did it compile".

## Step 5 — commit together

Commit `go.mod`, `go.sum`, `internal/provider/**`, and `docs/**` in one
changeset — splitting them makes `make check` fail on the intermediate commit.

## Step 6 — ask whether to cut a release

Once the bump commit is merged to `main` and CI (`.github/workflows/go.yml`'s
`check`/`build` jobs) is green there, **ask the user whether they want a new
release cut now** — do not tag/push on your own judgment, this step publishes
public artifacts and cannot be undone by force-pushing over a tag.

If they say yes:

1. Find the latest tag and propose the next version:
   ```sh
   git fetch --tags
   git tag -l | sort -V | tail -1
   ```
   Default to a patch bump (`vX.Y.Z+1`). If Step 4 turned up attribute
   renames/removals on existing resources, propose a minor/major bump instead
   per whatever versioning convention the user confirms — state this
   explicitly rather than silently defaulting to patch.
2. Confirm the exact version string with the user before tagging.
3. From a clean checkout of `main` at the merged commit:
   ```sh
   git tag -a vX.Y.Z -m "vX.Y.Z"
   git push origin vX.Y.Z
   ```
   Pushing the tag triggers `.github/workflows/release.yml`: GoReleaser
   cross-compiles for freebsd/windows/linux/darwin, GPG-signs the checksums
   file, and publishes a GitHub Release (config in `.goreleaser.yml`) — there
   is no separate "publish" step to run locally.
4. Watch it land: `gh run watch` (or check the Actions tab), then
   `gh release view vX.Y.Z` once the workflow finishes.

## Common mistakes

| Mistake | Result |
|---|---|
| Bump `require` while a local `replace` is still active | New version is never actually fetched or used; no error |
| Skip `GOPRIVATE`/SSH `insteadOf` setup | `go get`/`go mod tidy` fail with an opaque 404/410 on the private repo |
| Assume `go build` succeeding means nothing broke | Misses silently-skipped resources (Step 3) and renamed attributes |
| New API group added upstream, `imports.go` not touched | Its resources are never generated, with no error at all |
| Hand-edit a `*_gen.go` file instead of regenerating | Overwritten next `make generate`; diverges from what CI's `make check` expects |
| Tag/push a release without asking, or before the bump commit is merged and green on `main` | Publishes an unreviewed or broken build as a public release; a pushed tag isn't safely undoable |
