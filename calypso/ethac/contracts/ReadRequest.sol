pragma solidity ^0.4.24;

contract ReadRequest {
    address public writeRequest;
    bytes public xc;
    address public owner;

    constructor(address a, bytes x) public {
        owner = msg.sender;
        writeRequest = a;
        xc = x;
    }

    function CompareReadWrite(address wr) public view returns (bool) {
        return wr == writeRequest;
    }
}