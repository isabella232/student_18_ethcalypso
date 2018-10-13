pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;
import ".././contracts/Calypso.sol";
import "truffle/DeployedAddresses.sol";
import "truffle/Assert.sol";

/*contract TestCalypso {
    
    
    Calypso c;
    function testBeforeAll() public {
        c = new Calypso();
    }
    
    function testAddHash() public returns (bool){
        bytes memory d = "\x72\x74\x69";
        bytes memory ed = "\x69\x79";
        bytes memory l ="\x69\x79\x89";
        address[] memory accs;
        bytes32 h = c.newWriteRequest(d, ed, l, accs);
        return c.hasWriteRequest(h);
    }

    function testNotInMap() public returns (bool){
        bytes32 ba = "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA";
        return !c.hasWriteRequest(ba);
    }

    function testDoesEverythingPass() returns (bool){
        return false;
    }
}*/