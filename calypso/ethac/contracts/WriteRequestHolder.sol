pragma solidity ^0.4.24;


import "./WriteRequest.sol";

contract WriteRequestHolder {

    WriteRequest[] private holder;
    mapping(bytes32 => WriteRequest) public wrmap;
    mapping(bytes32 => bool) public hasRequest;

    function addWriteRequest(WriteRequest wr) public {
        bytes32 h = wr.getID();
        wrmap[h] = wr;
        hasRequest[h] = true;
        holder.push(wr);
    }

    function getRequestByID(bytes32 b) public view returns (WriteRequest) {
        return wrmap[b];
    }
    function has() public view returns (bool) {
        return true;
    }
}