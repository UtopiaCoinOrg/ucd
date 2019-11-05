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
rm -rf vendor
go mod init
go mod vendor
cd ./chaincfg/chainhash/x19rhash/x19r_c_lib/
make clean
echo "ucd make ..."
make
echo "ucd make over..."

echo "go mod ucwallet ..."
cd $CWD/../
sudo rm -rf ucwallet
git clone https://github.com/UtopiaCoinOrg/ucwallet.git
cd ucwallet
rm -rf go.sum
rm -rf go.mod
rm -rf vendor
go mod init
go mod vendor
rm -rf ./vendor/github.com/UtopiaCoinOrg/ucd
cp -rf ../ucd ./vendor/github.com/UtopiaCoinOrg/
rm -rf ./vendor/github.com/UtopiaCoinOrg/ucd/vendor
export GO111MODULE=off
echo $pwd
echo "go mod ucwallet over ..."
echo "install ucwallet ..."
go install
cd ../ucd
echo "install ucd ..."
go install
cd cmd/ucctl
go install
echo "install over ..."

screen_ucd=$"mainnet_ucd"
screen -dmS screen_ucd
echo "create screen ucd  ..."
ucd_cmd=$"ucd -u admin -P 123 --addrindex --txindex";
screen -x -S screen_ucd -p 0 -X stuff "ucd_cmd"
screen -x -S screen_ucd -p 0 -X stuff $'\n'

screen_ucwallet=$"mainnet_ucwallet"
screen -dmS screen_ucwallet
echo "create screen ucwallet  ..."
#ucd_cmd=$"ucwallet -u admin -P 123 --addrindex --txindex";
#screen -x -S screen_ucd -p 0 -X stuff "ucd_cmd"
#screen -x -S screen_ucd -p 0 -X stuff $'\n'

echo "All over..."
