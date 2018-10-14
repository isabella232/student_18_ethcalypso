pragma solidity ^0.4.24;
import "./ReadRequest.sol";

contract ReadRequestHolder {
    ReadRequest[] private requests;
    mapping(bytes32 => bool) private hasReq;

    constructor(ReadRequest[] rra) public {
        requests = rra;
    }

    function getReadRequestsByKey(address a) public view returns (ReadRequest[]) {
        ReadRequest[] memory temp = new ReadRequest[](requests.length);
        uint256 counter = 0;
        for(uint256 i = 0; i < requests.length;i++) {
            ReadRequest rq = requests[i];
            if(rq.xc() == a) {
                temp[counter] = rq;
                counter++;
            }
        }
        ReadRequest[] memory val = new ReadRequest[](counter);
        for(uint256 j = 0; j < counter;j++) {
            val[j] = temp[j];
        }
        return val;
    }

    function addReadRequest(ReadRequest rr) public {
        requests.push(rr);
        bytes32 h = rr.getID();
        hasReq[h] = true;
    }

    function hasReadRequest(bytes32 id) public view returns (bool) {
        return hasReq[id];
    }
}