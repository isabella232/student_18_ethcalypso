
apt-get install curl python-software-properties

curl -sL https://deb.nodesource.com/setup_10.x | sudo -E bash -

apt-get install nodejs

npm install -g truffle

npm install -g ganache-cli

npm install -g mocha

go get -u ./...

cd $HOME/go/src/github.com/ethereum/go-ethereum ; make

cd $HOME/go/src/github.com/ethereum/go-ethereum ; make devtools

abigen --sol ethac/contracts/Calypso.sol --pkg gocontracts > ethac/gocontracts/Calypso.go