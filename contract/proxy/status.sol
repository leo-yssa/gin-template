// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "./beacon.sol";
import "../@openzeppelin/contracts/token/ERC20/ERC20.sol";

contract Status {
    uint public result = 1;
    address public beaconAddress;

    constructor(address _beacon) {
        beaconAddress = _beacon;
    }

    function getLogicAddress() public view returns (address) {
        return Beacon(beaconAddress).implementaionAddress();
    }

    function multi(uint a) public returns (bool, uint) {
        address logic = Beacon(beaconAddress).implementaionAddress();
        (bool success, ) = logic.delegatecall(
            abi.encodeWithSignature("multi(uint256)", a)
        );
        return (success, result);
    }
}
