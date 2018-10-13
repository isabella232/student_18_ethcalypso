const Calypso = artifacts.require("./Calypso.sol");

module.exports = (deployer) => {
    deployer.deploy(Calypso, []);
};