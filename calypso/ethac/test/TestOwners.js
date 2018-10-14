const  assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require("web3");
const web3 = new Web3(ganache.provider())

//Instance of a Smart contract
let Calypso;

//The accounts on the deployed blockchain.
let accounts;

const json = require("./../build/contracts/Owners.json");

const interface = json['abi'];

const bytecode = json['bytecode'];

beforeEach(async () => {
    accounts = await web3.eth.getAccounts();
    const acc1 = accounts[0];
    const acc2 = accounts[1];
    const acc3 = accounts[2];
    const acc4 = accounts[3];

    const depAccounts = [acc1, acc2, acc3, acc4];
    Calypso = await new web3.eth.Contract(interface).deploy({
        data : bytecode,
        arguments : [depAccounts],
    }).send({from : acc1, gas : 1000000})
});


it("Can register an owner", async () => {
    accounts = await web3.eth.getAccounts();
    const nrOwners = await Calypso.methods.getNrOwners().call();
    console.log(nrOwners);
    assert.equal(nrOwners, 4);
});

it("Can check that an owner can write to the contract", async () => {
    const acc1 = accounts[1];
    const canWrite = await Calypso.methods.canWrite(acc1).call();
    assert.equal(canWrite, true);
});

it("Can check that an owner that does not have permission can't write to the contract", async () => {
    const acc1 = accounts[5];
    const canWrite = await Calypso.methods.canWrite(acc1).call();
    assert.equal(!canWrite, true);
});