pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

contract WriteRequest {
    bytes public data;

    bytes public extraData;

    bytes public LTSID;
    
    bytes public U;

    bytes public Cs;

    address[] private policy;

    //This is to know how to split the array
    //Because constructing multi-dimensional 
    //arrays is kind of clunky
    int64 public split;

    constructor(bytes d, bytes ed, bytes l, address[] p, bytes u, bytes cs, int64 s) public {
        data = d;
        extraData = ed;
        LTSID = l;
        policy = p;
        U = u;
        Cs = cs;
        split = s;
    }

    function getPolicySize() public returns (uint256) {
        return policy.length;
    }
}