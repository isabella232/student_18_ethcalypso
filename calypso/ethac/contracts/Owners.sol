pragma solidity ^0.4.24;


contract Owners {
    address[] private owners;
    mapping(address => bool) private ownersMap;

    constructor() public {
        
    }

    function getNrOwners() public view returns (uint256) {
        return owners.length;
    }

    function canWrite(address a) public view returns (bool) {
        return ownersMap[a];
    }

    function AddOwner(address a) public {
        ownersMap[a] = true;
    }
}