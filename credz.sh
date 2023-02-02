#! /bin/bash

export TSB_ADDRESS=$(cat mp.json | jq -r .tsb.address)
export TSB_ORGANIZATION=$(cat mp.json | jq -r .tsb.organization)
export TSB_USERNAME=admin
export TSB_PASSWORD=$(cat mp.json | jq -r .tsb.password)