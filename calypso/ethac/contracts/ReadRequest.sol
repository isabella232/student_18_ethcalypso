pragma solidity ^0.4.24;

contract ReadRequest {
    bytes32 public writeRequest;
    address public xc;

    constructor(bytes32 b32, address x) public {
        writeRequest = b32;
        xc = x;
    }

    function CompareReadWrite(bytes32 wr) public view returns (bool) {
        return wr == writeRequest;
    }


    //Don't really need this ID per say but it can help me
    //find this read request in the blockchain since every 
    //address writeRequest ID should be unique and therefore lead to a unique hash.
    function getID() public view returns (bytes32) {
        bytes memory packet = encode();
        return sha256(packet);
    }

    //Makes changing the ReadRequest and getting the
    //ID of the read request more seperate since it only
    //calls this function the get the packet to hash.
    function encode() private view returns (bytes) {
        return abi.encodePacked(writeRequest, xc);
    }
}