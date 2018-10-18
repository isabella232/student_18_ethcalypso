pragma solidity ^0.4.24;

contract ReadRequest {
    address public writeRequest;
    address public xc;

    constructor(address a, address x) public {
        writeRequest = a;
        xc = x;
    }

    function CompareReadWrite(address wr) public view returns (bool) {
        return wr == writeRequest;
    }
}