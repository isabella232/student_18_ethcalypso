const  assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require("web3");
const web3 = new Web3(ganache.provider());
const cal = artifacts.require("./Calypso.sol");


let accounts;

let Calypso;

const json = require("./../build/contracts/Calypso.json");
const wrJSON = require("./../build/contracts/WriteRequest.json");
const holderJSON = require("./../build/contracts/WriteRequestHolder.json");
const readJSON = require("./../build/contracts/ReadRequest.json");
const rrHolderJSON = require("./../build/contracts/ReadRequestHolder.json");

const interface = json['abi'];
const wrInterface = wrJSON['abi'];
const holderInterface = holderJSON['abi'];
const readInterface = readJSON['abi'];
const rrHolderInterface = rrHolderJSON['abi'];

const bytecode = json['bytecode'];
const wrBytecode = wrJSON['bytecode'];
const holderBytecode = holderJSON['bytecode'];
const readBytecode = readJSON['bytecode'];
const rrHolderBytecode = rrHolderJSON['bytecode'];

beforeEach(async () => {
    accounts = await web3.eth.getAccounts();
    const acc1 = accounts[1];
    const acc2 = accounts[2];
    const acc3 = accounts[3];
    const acc4 = accounts[4];

    const owners = [acc1, acc2, acc3, acc4];
    const holder = await createHolder();
    console.log("Holder address is: ", holder.options.address);
    const rrHolder = await createReadHolder();
    console.log("rrholder address is: ", rrHolder.options.address)

    Calypso = await new web3.eth.Contract(interface).deploy(
        {
            data : bytecode,
            arguments : [owners, holder.options.address, rrHolder.options.address],
        }
    ).send({ from : accounts[0], gas : 4712388});

});

it("Making sure that the owners contract and the Calypso contract integrate properly" , async () => {
    const canWrite = await Calypso.methods.canWrite(accounts[1]).call();
    assert.equal(canWrite, true);
});

it("Adding a writeRequest and then check that the transaction went through", async () => {
    const acc1 = accounts[1];
    const acc2 = accounts[2];
    const acc3 = accounts[3];
    const acc4 = accounts[4];
    const d = "0x726984";
    const ed = "0x548792";
    const l = "0x578952";
    const owners = [acc1, acc2, acc3, acc4];
    const wr1 = await CreateWR(d, ed, l, owners);
    const address = wr1.options.address;
    
    Calypso.methods.addWriteRequest(address).send({from : accounts[1], gas: 4712388})
    const h = await Calypso.methods.getIDOfWriteRequest(address).call();
    const isValid = await Calypso.methods.checkIfRequestIsValid(address).call();
    assert.equal(true, isValid);
});

it("Adding a ReadRequest with a corresponding WriteRequest", async () => {
    const acc1 = accounts[1];
    const acc2 = accounts[2];
    const acc3 = accounts[3];
    const acc4 = accounts[4];
    const d = "0x726984";
    const ed = "0x548792";
    const l = "0x578952";
    const owners = [acc1, acc2, acc3, acc4];
    const wr = await CreateWR(d, ed, l, owners);
    const wrAddress = wr.options.address;
    const rr = await createReadRequest(wrAddress, accounts[0]);
    const rrAddress = rr.options.address;
    await Calypso.methods.addWriteRequest(wrAddress).send({from : accounts[1], gas: 4712388});
    await Calypso.methods.AddReadRequest(rrAddress).send({from : accounts[1], gas: 4712388});
    const canRead = await Calypso.methods.isValidReadRequest(rrAddress).call();
    assert.equal(canRead, true);
})


const CreateWR = async (d, ed, l, a) => {
    const data = [d, ed, l, a];
    const wr1 = await new web3.eth.Contract(wrInterface).deploy(
        {
            data : wrBytecode,
            arguments : data,
        }
    ).send({from : accounts[0], gas : 4712388})
    return wr1;
}

const createHolder = async () => {
    const h1 = await new web3.eth.Contract(holderInterface).deploy(
        {
            data : holderBytecode,
        }
    ).send({from : accounts[0], gas : 4712388});
    return h1;
}

const createReadRequest = async (a1, a2) => {
    const rr = await new web3.eth.Contract(readInterface).deploy(
        {
            data : readBytecode,
            arguments : [a1, a2],
        }
    ).send({from : accounts[0], gas : 4712388});
    return rr;
}

const createReadHolder = async () => {
    const rrHolder = await new web3.eth.Contract(rrHolderInterface).deploy(
        {
            data : rrHolderBytecode,
        }
    ).send({from : accounts[0], gas : 4712388});
    return rrHolder;
}