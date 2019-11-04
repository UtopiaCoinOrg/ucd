#!/bin/sh

CWD=`pwd`
cd $CWD

rm -rf $GOBIN/ucd
rm -rf $GOBIN/ucwallet
rm -rf $GOBIN/ucctl

echo "go mod ucd ..."

export GO111MODULE=on
rm -rf go.sum
rm -rf go.mod
go mod init
go mod vendor
cd ./chaincfg/chainhash/x19rhash/x19r_c_lib/
make clean
echo "ucd make ..."
make
echo "ucd make over..."

echo "go mod ucwallet ..."
cd ../
sudo rm -rf ucwallet
git clone https://github.com/UtopiaCoinOrg/ucwallet.git
cd ucwallet
rm -rf go.sum
rm -rf go.mod
go mod init
go mod vendor
rm -rf ./vendor/github.com/UtopiaCoinOrg/ucd
cp -rf ../ucd ./vendor/github.com/UtopiaCoinOrg/
export GO111MODULE=off

echo "go mod ucwallet over ..."
go install
cd ../ucd
go install
cd cmd/ucctl
go install

