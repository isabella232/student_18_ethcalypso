pragma solidity ^0.4.24;

contract WriteRequest {
    bytes public data;

    bytes public extraData;

    bytes public LTSID;
    
    bytes public U;

    bytes public Ubar;

    bytes public E;

    bytes public F;

    bytes public X;

    bytes public Cs;

    address[] private policy;

    mapping(address => bool) public policyMap;

    //This is to know how to split the array
    //Because constructing multi-dimensional 
    //arrays is kind of clunky
    int64 public split;
    int64[] public slice;

    constructor(bytes d, bytes ed, bytes l, address[] p, bytes u, bytes cs, int64[] sl, bytes f, bytes e, bytes ub) public {
        data = d;
        extraData = ed;
        LTSID = l;
        policy = p;
        U = u;
        Cs = cs;
        split = 1;
        for(uint32 i = 0;i < p.length;i++) {
            policyMap[p[i]] = true;
        }
        slice = sl;
        F = f;
        E = e;
        Ubar = ub;
    }

    function getPolicySize() public returns (uint256) {
        return policy.length;
    }

    function CanRead(address a) public view returns(bool) {
        return policyMap[a];
    }

    function GetSlice() public view returns(int64[]) {
        return slice;
    }
}