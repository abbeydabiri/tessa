pragma solidity >=0.4.0 <0.7.0;

contract SafeMath {
    
    function safeAdd(uint a, uint b) internal pure returns (uint c) {
        c = a + b;
        require(c >= a);
    }
    function safeSub(uint a, uint b) internal pure returns (uint c) {
        require(b <= a);
        c = a - b;
    }
    function safeMul(uint a, uint b) internal pure returns (uint c) {
        if (a == 0) {
            return 0;
        }
        
        c = a * b;
        require(c / a == b);
    }
    function safeDiv(uint a, uint b) internal pure returns (uint c) {
        require(b > 0);
        c = a / b;
    }
}

contract ERC20 {
    function totalSupply() public view returns (uint);
    function balanceOf(address tokenOwner) public view returns (uint balance);
    function allowance(address tokenOwner, address spender) public view returns (uint remaining);
    function transfer(address to, uint tokens) public returns (bool success);
    function approve(address spender, uint tokens) public returns (bool success);
    function transferFrom(address from, address to, uint tokens) public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);
    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}


contract Owned {
    address payable owner;
    address payable newOwner;

    event OwnershipTransferred(address indexed _from, address indexed _to);

    constructor() public {
        owner = msg.sender;
    }

    modifier onlyOwner {
        require(msg.sender == owner, "Only contract owner can perform this transaction");
        _;
    }
    
    function transferOwnership(address payable _newOwner) public onlyOwner {
        newOwner = _newOwner;
    }
    function acceptOwnership() public {
        require(msg.sender == newOwner, "Only new owner can perform this transaction");
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
        newOwner = address(0);
    }
}


contract SmartToken is ERC20, Owned, SafeMath {
    string public name;
    string public symbol;
    uint8 public decimals;
    uint public curTotalSupply;
    uint public maxTotalSupply;

    mapping(address => uint) balances;
    mapping(address => bool) public blacklist;
    mapping(address => mapping(address => uint)) allowed;


    constructor( string memory _symbol, string memory _name, uint _maxTotalSupply, uint _seed, address _recipient  ) public {
        decimals = 18;
    
        require( bytes(_name).length > 3, "_name must be 4 characters minimum");
        require( bytes(_symbol).length > 2, "_symbol must be 3 characters minimum");
        require( _maxTotalSupply > 1, "_maxTotalSupply must be 1 units minimum");
        require( _recipient == address(0), "_recipient must be set");

        name = _name;
        symbol = _symbol;
        curTotalSupply = _seed*(10**uint(decimals));
        maxTotalSupply = _maxTotalSupply*(10**uint(decimals));
        balances[_recipient] = curTotalSupply;
        emit Transfer(address(0), _recipient, curTotalSupply);
    }


    //minting functionality
    function mintTokens(address to, uint tokens) public onlyOwner returns (bool success) {
        require(!blacklist[to], "This account has been blacklisted");

        tokens = tokens*(10**uint(decimals));
        curTotalSupply = safeAdd(curTotalSupply, tokens);
        require(curTotalSupply <= maxTotalSupply);

        balances[to] = safeAdd(balances[to], tokens);
        emit Transfer(owner, to, tokens);
        return true;
    }
    //minting functionality


    function () external payable {
        // owner.transfer(msg.value);
        msg.sender.transfer(msg.value); //return the value back to sender
    }

    function totalSupply() public view returns (uint) {
        return curTotalSupply - balances[address(0)];
    }

    function balanceOf(address tokenOwner) public view returns (uint balance) {
        return balances[tokenOwner];
    }

    function transfer(address to, uint tokens) public returns (bool success) {
        require(!blacklist[msg.sender], "This account has been blacklisted");

        balances[msg.sender] = safeSub(balances[msg.sender], tokens);
        balances[to] = safeAdd(balances[to], tokens);
        emit Transfer(msg.sender, to, tokens);
        return true;
    }

    function approve(address spender, uint tokens) public returns (bool success) {
        require(spender == address(0), "Spender address cannot be a zero");
        require(!blacklist[spender], "This account has been blacklisted");

        allowed[msg.sender][spender] = tokens;
        emit Approval(msg.sender, spender, tokens);
        return true;
    }


    function transferFrom(address from, address to, uint tokens) public returns (bool success) {
        require(!blacklist[msg.sender], "Transaction sender has been blacklisted");
        require(!blacklist[from], "From address been blacklisted");
        require(!blacklist[to], "To address has been blacklisted");

        balances[from] = safeSub(balances[from], tokens);
        if (msg.sender != owner) {
            allowed[from][msg.sender] = safeSub(allowed[from][msg.sender], tokens);
        }

        balances[to] = safeAdd(balances[to], tokens);
        emit Transfer(from, to, tokens);
        return true;
    }

    function allowance(address tokenOwner, address spender) public view returns (uint remaining) {
        return allowed[tokenOwner][spender];
    }


    function approveAndCall(address spender, uint tokens, bytes memory data) public returns (bool success) {
        require(!blacklist[msg.sender], "This account has been blacklisted");
        allowed[msg.sender][spender] = tokens;
        emit Approval(msg.sender, spender, tokens);
        approveAndCallFallBack(spender).receiveApproval(msg.sender, tokens, address(this), data);
        return true;
    }

    function addToBlacklist(address user) public onlyOwner returns (bool success) {
        blacklist[user] = true;
        return true;
    }
    
    function removeFromBlacklist(address user) public onlyOwner returns (bool success) {
        blacklist[user] = false;
        return true;
    }

    
}   
    
contract approveAndCallFallBack {
    function receiveApproval(address from, uint256 tokens, address token, bytes memory data) public;
}