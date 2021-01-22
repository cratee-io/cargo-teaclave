#!/bin/bash

workingDir=$PWD
outDir=$workingDir/testing
rev=99f12d997a985e57a279728b6134269888944e86
remote="https://github.com/cratee-io/testing"
redundants=(
  $workingDir/third_party/rsgx-assets/vendor/sgx_edl
  $workingDir/Cargo.toml
)

# remove old files
rm -rf $(ls -A | grep -v "README.md" | grep -v "renew.sh")
#rm -rf $(ls -A | grep -v "README.md" | grep -v "renew.sh" | grep -v "testing")

git clone $remote $outDir

cd $outDir
git checkout $rev
cp -r tests/sgx/* $workingDir
cd -

rm -rf $outDir ${redundants[@]}

echo "$rev" > rev.txt
