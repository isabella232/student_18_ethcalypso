pragma solidity ^0.4.24;

import "./Owners.sol";
import "./WriteRequestHolder.sol";
import "./WriteRequest.sol";


contract Calypso {

    Owners private owners;
    WriteRequestHolder private wrHolder;
    mapping(bytes32 => bool) hasReq;


    //The addresses of the owners of the contract
    //This is a mapping that maps the hash of a request to the
    //associated WriteRequest of that hash
    
    constructor(address[] a) public {
        owners = new Owners(a);
    }

    function canWrite(address a) public view returns (bool) {
        return owners.canWrite(a);
    }

    function addWriteRequest(bytes d, bytes ed, bytes l, address[] p) public {
        require(canWrite(msg.sender), "You are not a registered owner");
        WriteRequest wr = newWriteRequest(d, ed, l, p);
        wrHolder.addWriteRequest(wr);
        hasReq[wr.getID()] = true;
    }

    function getRequestByID(bytes32 b) public view returns (WriteRequest) {
        WriteRequest wr = wrHolder.getRequestByID(b);
        return wr;
    }

    //This  functions is basically only there for testing purposes.
    //Passing contracts to truffle or js frameworks is a
    //pain but this allows me to be able to check if for a
    //given ID I have stored a contract. This is only used
    //for testing purposes and will probably not go to
    //production.
    function checkIfRequestIsValid(bytes d, bytes ed, bytes l, address[] p) public returns (bool) {
        WriteRequest wr = newWriteRequest(d, ed, l, p);
        return hasReq[wr.getID()];
    }

    function newWriteRequest(bytes d, bytes ed, bytes l, address[] p) public returns (WriteRequest) {
        WriteRequest wr = new WriteRequest(d, ed, l, p);
        return wr;
    }

    //This is also purely a testing method. Way too hard to
    //pass a contract as a parameter outside of a solidity contract
    function getIDOfWriteRequest(bytes d, bytes ed, bytes l, address[] p) public returns (bytes32) {
        WriteRequest wr = new WriteRequest(d, ed, l, p);
        return wr.getID();
    }

}



