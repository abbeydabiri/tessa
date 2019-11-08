pragma solidity >=0.4.0 <0.7.0;

contract approveAndCallFallBack {
    function receiveApproval(address from, uint256 tokens, address token, bytes memory data) public;
}