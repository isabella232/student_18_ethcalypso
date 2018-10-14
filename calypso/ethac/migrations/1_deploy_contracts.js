const Calypso = artifacts.require("./Calypso.sol");
const Web3 = require("web3");
const ganache = require('ganache-cli');
const web3 = new Web3(ganache.provider());

module.exports = (deployer) => {
    deployCalypso(deployer);
};

const deployCalypso = async (deployer) => {
    const accounts = web3.eth.getAccounts();
    const account0 = accounts[0];
    const account1 = accounts[1];
    const account2 = accounts[2];
    const accs = [account0, account1, account2];
    deployer.deploy(Calypso, accs);
}