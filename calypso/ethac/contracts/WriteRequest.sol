pragma solidity ^0.4.24;

contract WriteRequest {
    bytes public data;

    bytes public extraData;

    bytes public LTSID;
    
    bytes public U;

    address[] private policy;

    constructor(bytes d, bytes ed, bytes l, address[] p, bytes u) public {
        data = d;
        extraData = ed;
        LTSID = l;
        policy = p;
        U = u;
    }

    function getPolicySize() public returns (uint256) {
        return policy.length;
    }
}