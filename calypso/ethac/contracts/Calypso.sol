pragma solidity ^0.4.24;

import "./Owners.sol";
import "./WriteRequestHolder.sol";
import "./WriteRequest.sol";
import "./ReadRequest.sol";
import "./ReadRequestHolder.sol";


contract Calypso {

    Owners private owners;
    WriteRequestHolder private wrHolder;
    mapping(bytes32 => bool) hasReq;
    ReadRequestHolder private rrHolder;


    //The addresses of the owners of the contract
    //This is a mapping that maps the hash of a request to the
    //associated WriteRequest of that hash
    
    constructor(address a, address holder, address rrholder) public {
        owners = Owners(a);
        wrHolder = WriteRequestHolder(holder);
        rrHolder = ReadRequestHolder(rrholder);
    }

    function canWrite(address a) public view returns (bool) {
        return owners.canWrite(a);
    }

    function addWriteRequest(address a) public {
        //require(canWrite(msg.sender), "You are not a registered owner");
        wrHolder.addWriteRequest(a);
    }

    function AddReadRequest(address a) public {
        ReadRequest rr = ReadRequest(a);
        address write = rr.writeRequest();
        require(wrHolder.canRead(write), "There is no corresponding write request");
        rrHolder.addReadRequest(a);
    }

    function checkIfRequestIsValid(address a) public returns (bool) {
        return wrHolder.canRead(a);
    }

    //This is also purely a testing method. Way too hard to
    //pass a contract as a parameter outside of a solidity contract
    function getIDOfWriteRequest(address a) public returns (bytes32) {
        WriteRequest wr = WriteRequest(a);
        return wr.getID();
    }

    function isValidReadRequest(address a)  public view returns (bool) {
        return rrHolder.hasReadRequest(a);
    }

}



