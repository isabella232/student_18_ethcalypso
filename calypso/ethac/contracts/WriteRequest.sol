pragma solidity ^0.4.24;

contract WriteRequest {
    bytes private data;

    bytes private extraData;

    bytes private LTSID;

    address[] private policy;

    constructor(bytes d, bytes ed, bytes l, address[] p) public {
        data = d;
        extraData = ed;
        LTSID = l;
        policy = p;
    }

    function getID() public returns (bytes32) {
        bytes memory ba = encPack(data, extraData, LTSID, policy);
        return sha256(ba);
    }

    function encPack(bytes d, bytes ed, bytes l, address[] p) private returns (bytes) {
        return abi.encodePacked(d, ed, l, p);
    }

    //Basically passing a contract was too hard so I just all the data of a contract.
    function compare(bytes d, bytes ed, bytes l, address[] p) public returns (bool) {
        bytes memory ba = encPack(d, ed, l, p);
        return this.getID() == sha256(ba);
    }

    function getPolicySize() public returns (uint256) {
        return policy.length;
    }
}