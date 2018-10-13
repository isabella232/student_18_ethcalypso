const  assert = require('assert');
const ganache = require('ganache-cli');
const Web3 = require("web3");
const web3 = new Web3(ganache.provider());

let accounts;

let writeRequest;

const json = require("./../build/contracts/WriteRequest.json");

const interface = json['abi'];

const bytecode = json['bytecode'];

beforeEach(async () => {
    accounts = await web3.eth.getAccounts();
    const acc1 = accounts[1];
    const acc2 = accounts[2];
    const acc3 = accounts[3];
    const acc4 = accounts[4];


    //These are basically byte arrays
    const data = "0x726984";
    const extraData = "0x548792";
    const LTSID = "0x578952";

    const policy = [acc1, acc2, acc3, acc4];

    writeRequest = await new web3.eth.Contract(interface).deploy(
        {
            data: bytecode,
            arguments : [data, extraData, LTSID, policy],
        }
    ).send({from : accounts[0], gas : 1000000});
});

it("Make sure that the same writeRequest gives the same hash", async () => {
    const h1 = writeRequest.methods.getID().call();
    const h2 = writeRequest.methods.getID().call();
    assert.deepEqual(h1, h2, "These should both return the same bytes32 array");
});

it("Make sure that different WriteRequests give a differentHash", async () => {
    //These are basically byte arrays
    const data = "0x554433";
    const extraData = "0x998877";
    const LTSID = "0x112233";

    const policy = [accounts[7], accounts[5]];
    const neq = await writeRequest.methods.compare(data, extraData, LTSID, policy).call();
    assert.equal(true, !neq);
});


