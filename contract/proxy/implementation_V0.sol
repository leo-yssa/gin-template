// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

contract implementation_V0 {
    uint public result = 10;
    function multi(uint a) public returns (bool, uint) {
        result = a * 1;
        return (true, a * 1);
    }
}
