const WR = artifacts.require("./WriteRequest.sol");
const Web3 = require("web3");
const ganache = require('ganache-cli');
const web3 = new Web3(ganache.provider());

module.exports = (deployer) => {
    deployWR(deployer);
};

const deployWR =  async (deployer, [data]) => {
    deployer.deploy(WR, data);
}