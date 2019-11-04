#!/bin/sh

CWD=`pwd`
cd $CWD

echo "building ucd ..."
sudo -s -u root
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
exit
cd $CWD
go install
cd cmd/ucctl 
go install
cd ../../../
sudo rm -rf ucwallet
echo "building ucwallet ..."
git clone https://github.com/UtopiaCoinOrg/ucwallet.git
cd ucwallet
sudo -s -u root
export GO111MODULE=on
rm -rf go.sum
rm -rf go.mod
go mod init
go mod vendor
rm -rf ./vendor/github.com/UtopiaCoinOrg/ucd
cp -rf ../ucd ./vendor/github.com/UtopiaCoinOrg/
exit
go install

