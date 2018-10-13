pragma solidity ^0.4.24;


contract Owners {
    address[] private owners;
    mapping(address => bool) private ownersMap;

    constructor(address[] a) public {
        owners = a;
        for(uint32 i = 0; i < a.length;i++) {
            ownersMap[a[i]] = true;
        }
    }

    function getNrOwners() public view returns (uint256) {
        return owners.length;
    }

    function canWrite(address a) public view returns (bool) {
        return ownersMap[a];
    }
}