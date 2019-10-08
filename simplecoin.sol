pragma solidity ^0.4.0;
contract SimpleCoin{

    string name;
    uint total;
    mapping (address => uint) balances;

    constructor(string _name, uint _total) {
        total = _total;
        name = _name;
        balances[msg.sender] = total;
    }


    function totalSupply() constant returns (uint256) {
        return total;
    }

    function balanceOf(address _owner) constant returns (uint256) {
        return balances[_owner];
    }

    function transfer(address _to, uint256 _value) returns (bool) {
        if (balances[msg.sender] >= _value) {
            balances[msg.sender] -= _value;
            balances[_to] += _value;
            emit Transfer(msg.sender, _to, _value);
            return true;
        } else {
            return false;
        }

    }
    event Transfer(address indexed _from, address indexed _to, uint256 _value);
}
