#!/bin/bash

revisions=(
  "32b2d1f137ef5742e4e6320b3e1e38fa318a56f6" # 1.1.3
)

for v in ${revisions[@]}; do
  rm -rf Cargo.lock
  sed -i "s!^rev =.*!rev = \"$v\"!g" Cargo.toml
  cargo vendor --versioned-dirs --no-delete .
done
