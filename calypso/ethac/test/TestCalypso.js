const  assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require("web3");
const web3 = new Web3(ganache.provider());


let accounts;

let Calypso;

const json = require("./../build/contracts/Calypso.json");

const interface = json['abi'];

const bytecode = json['bytecode'];

beforeEach(async () => {
    accounts = await web3.eth.getAccounts();
    const acc1 = accounts[1];
    const acc2 = accounts[2];
    const acc3 = accounts[3];
    const acc4 = accounts[4];

    const owners = [acc1, acc2, acc3, acc4];

    Calypso = await new web3.eth.Contract(interface).deploy(
        {
            data : bytecode,
            arguments : [owners],
        }
    ).send({ from : accounts[0], gas : 3000000});

});

it("Making sure that the owners contract and the Calypso contract integrate properly" , async () => {
    const canWrite = await Calypso.methods.canWrite(accounts[1]).call();
    assert.equal(canWrite, true);
});

it("Adding a writeRequest and then check that the transaction went through", async () => {
    const d = "0x726984";
    const ed = "0x548792";
    const l = "0x578952";
    const acc1 = accounts[1];
    const acc2 = accounts[2];
    const acc3 = accounts[3];
    const acc4 = accounts[4];

    const p = [acc1, acc2, acc3, acc4];
    await Calypso.methods.addWriteRequest(d, ed, l, p);
    console.log('Writerequest added');
    const h = await Calypso.methods.getIDOfWriteRequest(d, ed, l, p).call(
        {
            gas : 4000000,
        }
    );
    console.log("ID of writeRequest", h);
    const isValid = await Calypso.methods.checkIfRequestIsValid(d, ed, l, p).call(
        {
            gas : 6000000,
        }
    );
    console.log("is valid");
    assert.equal(true, isValid);
});

