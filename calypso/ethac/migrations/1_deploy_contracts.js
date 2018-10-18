const Calypso = artifacts.require("./Calypso.sol");
const Web3 = require("web3");
const ganache = require('ganache-cli');
const web3 = new Web3(ganache.provider());

module.exports = (deployer) => {
    deployCalypso(deployer);
};

const deployCalypso = async (deployer, [owners]) => {
    deployer.deploy(Calypso, owners);
}