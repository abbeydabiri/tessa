// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package smartcontracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SmarttokenABI is the input ABI used to generate the binding from.
const SmarttokenABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_maxTotalSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seed\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"_totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"addToBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"blacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"mintTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"removeFromBlacklist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SmarttokenBin is the compiled bytecode used for deploying new contracts.
var SmarttokenBin = "0x60806040523480156200001157600080fd5b50604051620022b2380380620022b2833981810160405260808110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b506040526020018051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506012600460006101000a81548160ff021916908360ff160217905550600483511162000284576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180620022676022913960400191505060405180910390fd5b6002845111620002e0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526024815260200180620022436024913960400191505060405180910390fd5b600a82116200033b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526029815260200180620022896029913960400191505060405180910390fd5b82600290805190602001906200035392919062000461565b5083600390805190602001906200036c92919062000461565b50600460009054906101000a900460ff1660ff16600a0a8102600581905550600460009054906101000a900460ff1660ff16600a0a8202600681905550600554600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef6005546040518082815260200191505060405180910390a35050505062000510565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620004a457805160ff1916838001178555620004d5565b82800160010185558215620004d5579182015b82811115620004d4578251825591602001919060010190620004b7565b5b509050620004e49190620004e8565b5090565b6200050d91905b8082111562000509576000816000905550600101620004ef565b5090565b90565b611d2380620005206000396000f3fe6080604052600436106101095760003560e01c806370a0823111610095578063cae9ca5111610064578063cae9ca51146105eb578063dd62ed3e146106f5578063f0dda65c1461077a578063f2fde38b146107ed578063f9f92be41461083e57610109565b806370a082311461046c57806379ba5097146104d157806395d89b41146104e8578063a9059cbb1461057857610109565b80632ab4d052116100dc5780632ab4d05214610313578063313ce5671461033e5780633eaaf86b1461036f57806344337ea11461039a578063537df3b61461040357610109565b806306fdde0314610152578063095ea7b3146101e257806318160ddd1461025557806323b872dd14610280575b3373ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f1935050505015801561014f573d6000803e3d6000fd5b50005b34801561015e57600080fd5b506101676108a7565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101a757808201518184015260208101905061018c565b50505050905090810190601f1680156101d45780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ee57600080fd5b5061023b6004803603604081101561020557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610945565b604051808215151515815260200191505060405180910390f35b34801561026157600080fd5b5061026a610ada565b6040518082815260200191505060405180910390f35b34801561028c57600080fd5b506102f9600480360360608110156102a357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610b25565b604051808215151515815260200191505060405180910390f35b34801561031f57600080fd5b50610328610e0a565b6040518082815260200191505060405180910390f35b34801561034a57600080fd5b50610353610e10565b604051808260ff1660ff16815260200191505060405180910390f35b34801561037b57600080fd5b50610384610e23565b6040518082815260200191505060405180910390f35b3480156103a657600080fd5b506103e9600480360360208110156103bd57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e29565b604051808215151515815260200191505060405180910390f35b34801561040f57600080fd5b506104526004803603602081101561042657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f31565b604051808215151515815260200191505060405180910390f35b34801561047857600080fd5b506104bb6004803603602081101561048f57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611039565b6040518082815260200191505060405180910390f35b3480156104dd57600080fd5b506104e6611082565b005b3480156104f457600080fd5b506104fd61126b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561053d578082015181840152602081019050610522565b50505050905090810190601f16801561056a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561058457600080fd5b506105d16004803603604081101561059b57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611309565b604051808215151515815260200191505060405180910390f35b3480156105f757600080fd5b506106db6004803603606081101561060e57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001909291908035906020019064010000000081111561065557600080fd5b82018360208201111561066757600080fd5b8035906020019184600183028401116401000000008311171561068957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050611535565b604051808215151515815260200191505060405180910390f35b34801561070157600080fd5b506107646004803603604081101561071857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061180b565b6040518082815260200191505060405180910390f35b34801561078657600080fd5b506107d36004803603604081101561079d57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611892565b604051808215151515815260200191505060405180910390f35b3480156107f957600080fd5b5061083c6004803603602081101561081057600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611b35565b005b34801561084a57600080fd5b5061088d6004803603602081101561086157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611c1e565b604051808215151515815260200191505060405180910390f35b60028054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561093d5780601f106109125761010080835404028352916020019161093d565b820191906000526020600020905b81548152906001019060200180831161092057829003601f168201915b505050505081565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156109ea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611ca36021913960400191505060405180910390fd5b81600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925846040518082815260200191505060405180910390a36001905092915050565b6000600760008073ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205460055403905090565b6000610b70600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611c3e565b600760008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610d0e57610c8d600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611c3e565b600960008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505b610d57600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611c58565b600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a3600190509392505050565b60065481565b600460009054906101000a900460ff1681565b60055481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ed0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180611c736030913960400191505060405180910390fd5b6001600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610fd8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180611c736030913960400191505060405180910390fd5b6000600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b6000600760008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611128576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b815260200180611cc4602b913960400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113015780601f106112d657610100808354040283529160200191611301565b820191906000526020600020905b8154815290600101906020018083116112e457829003601f168201915b505050505081565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156113ae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611ca36021913960400191505060405180910390fd5b6113f7600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611c3e565b600760003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550611483600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611c58565b600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000600860003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156115da576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611ca36021913960400191505060405180910390fd5b82600960003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040518082815260200191505060405180910390a38373ffffffffffffffffffffffffffffffffffffffff16638f4ffcb1338530866040518563ffffffff1660e01b8152600401808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018481526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561179957808201518184015260208101905061177e565b50505050905090810190601f1680156117c65780820380516001836020036101000a031916815260200191505b5095505050505050600060405180830381600087803b1580156117e857600080fd5b505af11580156117fc573d6000803e3d6000fd5b50505050600190509392505050565b6000600960008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611939576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180611c736030913960400191505060405180910390fd5b600860008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16156119dc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526021815260200180611ca36021913960400191505060405180910390fd5b600460009054906101000a900460ff1660ff16600a0a82029150611a0260055483611c58565b6005819055506006546005541115611a1957600080fd5b611a62600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205483611c58565b600760008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040518082815260200191505060405180910390a36001905092915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611bda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526030815260200180611c736030913960400191505060405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60086020528060005260406000206000915054906101000a900460ff1681565b600082821115611c4d57600080fd5b818303905092915050565b6000818301905082811015611c6c57600080fd5b9291505056fe4f6e6c7920636f6e7472616374206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6e54686973206163636f756e7420686173206265656e20626c61636b6c69737465644f6e6c79206e6577206f776e65722063616e20706572666f726d2074686973207472616e73616374696f6ea265627a7a72315820030e4cf986bfd013619456385cb1c228bf3ced5fb9f3f22afad06a33b356d08264736f6c634300050c00325f73796d626f6c206d75737420626520332063686172616374657273206d696e696d756d5f6e616d65206d75737420626520342063686172616374657273206d696e696d756d5f6d6178546f74616c537570706c79206d7573742062652031303020756e697473206d696e696d756d"

// DeploySmarttoken deploys a new Ethereum contract, binding an instance of Smarttoken to it.
func DeploySmarttoken(auth *bind.TransactOpts, backend bind.ContractBackend, _symbol string, _name string, _maxTotalSupply *big.Int, _seed *big.Int) (common.Address, *types.Transaction, *Smarttoken, error) {
	parsed, err := abi.JSON(strings.NewReader(SmarttokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SmarttokenBin), backend, _symbol, _name, _maxTotalSupply, _seed)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Smarttoken{SmarttokenCaller: SmarttokenCaller{contract: contract}, SmarttokenTransactor: SmarttokenTransactor{contract: contract}, SmarttokenFilterer: SmarttokenFilterer{contract: contract}}, nil
}

// Smarttoken is an auto generated Go binding around an Ethereum contract.
type Smarttoken struct {
	SmarttokenCaller     // Read-only binding to the contract
	SmarttokenTransactor // Write-only binding to the contract
	SmarttokenFilterer   // Log filterer for contract events
}

// SmarttokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmarttokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmarttokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmarttokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmarttokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmarttokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmarttokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmarttokenSession struct {
	Contract     *Smarttoken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SmarttokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmarttokenCallerSession struct {
	Contract *SmarttokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SmarttokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmarttokenTransactorSession struct {
	Contract     *SmarttokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SmarttokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmarttokenRaw struct {
	Contract *Smarttoken // Generic contract binding to access the raw methods on
}

// SmarttokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmarttokenCallerRaw struct {
	Contract *SmarttokenCaller // Generic read-only contract binding to access the raw methods on
}

// SmarttokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmarttokenTransactorRaw struct {
	Contract *SmarttokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmarttoken creates a new instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttoken(address common.Address, backend bind.ContractBackend) (*Smarttoken, error) {
	contract, err := bindSmarttoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Smarttoken{SmarttokenCaller: SmarttokenCaller{contract: contract}, SmarttokenTransactor: SmarttokenTransactor{contract: contract}, SmarttokenFilterer: SmarttokenFilterer{contract: contract}}, nil
}

// NewSmarttokenCaller creates a new read-only instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttokenCaller(address common.Address, caller bind.ContractCaller) (*SmarttokenCaller, error) {
	contract, err := bindSmarttoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmarttokenCaller{contract: contract}, nil
}

// NewSmarttokenTransactor creates a new write-only instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SmarttokenTransactor, error) {
	contract, err := bindSmarttoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmarttokenTransactor{contract: contract}, nil
}

// NewSmarttokenFilterer creates a new log filterer instance of Smarttoken, bound to a specific deployed contract.
func NewSmarttokenFilterer(address common.Address, filterer bind.ContractFilterer) (*SmarttokenFilterer, error) {
	contract, err := bindSmarttoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmarttokenFilterer{contract: contract}, nil
}

// bindSmarttoken binds a generic wrapper to an already deployed contract.
func bindSmarttoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmarttokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smarttoken *SmarttokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smarttoken.Contract.SmarttokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smarttoken *SmarttokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smarttoken.Contract.SmarttokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smarttoken *SmarttokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smarttoken.Contract.SmarttokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Smarttoken *SmarttokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Smarttoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Smarttoken *SmarttokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smarttoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Smarttoken *SmarttokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Smarttoken.Contract.contract.Transact(opts, method, params...)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "_totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenSession) TotalSupply() (*big.Int, error) {
	return _Smarttoken.Contract.TotalSupply(&_Smarttoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x3eaaf86b.
//
// Solidity: function _totalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Smarttoken.Contract.TotalSupply(&_Smarttoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) constant returns(uint256 remaining)
func (_Smarttoken *SmarttokenCaller) Allowance(opts *bind.CallOpts, tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "allowance", tokenOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) constant returns(uint256 remaining)
func (_Smarttoken *SmarttokenSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Smarttoken.Contract.Allowance(&_Smarttoken.CallOpts, tokenOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address tokenOwner, address spender) constant returns(uint256 remaining)
func (_Smarttoken *SmarttokenCallerSession) Allowance(tokenOwner common.Address, spender common.Address) (*big.Int, error) {
	return _Smarttoken.Contract.Allowance(&_Smarttoken.CallOpts, tokenOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256 balance)
func (_Smarttoken *SmarttokenCaller) BalanceOf(opts *bind.CallOpts, tokenOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "balanceOf", tokenOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256 balance)
func (_Smarttoken *SmarttokenSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Smarttoken.Contract.BalanceOf(&_Smarttoken.CallOpts, tokenOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address tokenOwner) constant returns(uint256 balance)
func (_Smarttoken *SmarttokenCallerSession) BalanceOf(tokenOwner common.Address) (*big.Int, error) {
	return _Smarttoken.Contract.BalanceOf(&_Smarttoken.CallOpts, tokenOwner)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) constant returns(bool)
func (_Smarttoken *SmarttokenCaller) Blacklist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "blacklist", arg0)
	return *ret0, err
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) constant returns(bool)
func (_Smarttoken *SmarttokenSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Smarttoken.Contract.Blacklist(&_Smarttoken.CallOpts, arg0)
}

// Blacklist is a free data retrieval call binding the contract method 0xf9f92be4.
//
// Solidity: function blacklist(address ) constant returns(bool)
func (_Smarttoken *SmarttokenCallerSession) Blacklist(arg0 common.Address) (bool, error) {
	return _Smarttoken.Contract.Blacklist(&_Smarttoken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Smarttoken *SmarttokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Smarttoken *SmarttokenSession) Decimals() (uint8, error) {
	return _Smarttoken.Contract.Decimals(&_Smarttoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_Smarttoken *SmarttokenCallerSession) Decimals() (uint8, error) {
	return _Smarttoken.Contract.Decimals(&_Smarttoken.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenCaller) MaxTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "maxTotalSupply")
	return *ret0, err
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenSession) MaxTotalSupply() (*big.Int, error) {
	return _Smarttoken.Contract.MaxTotalSupply(&_Smarttoken.CallOpts)
}

// MaxTotalSupply is a free data retrieval call binding the contract method 0x2ab4d052.
//
// Solidity: function maxTotalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenCallerSession) MaxTotalSupply() (*big.Int, error) {
	return _Smarttoken.Contract.MaxTotalSupply(&_Smarttoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Smarttoken *SmarttokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Smarttoken *SmarttokenSession) Name() (string, error) {
	return _Smarttoken.Contract.Name(&_Smarttoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Smarttoken *SmarttokenCallerSession) Name() (string, error) {
	return _Smarttoken.Contract.Name(&_Smarttoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Smarttoken *SmarttokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Smarttoken *SmarttokenSession) Symbol() (string, error) {
	return _Smarttoken.Contract.Symbol(&_Smarttoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Smarttoken *SmarttokenCallerSession) Symbol() (string, error) {
	return _Smarttoken.Contract.Symbol(&_Smarttoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Smarttoken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenSession) TotalSupply() (*big.Int, error) {
	return _Smarttoken.Contract.TotalSupply(&_Smarttoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Smarttoken *SmarttokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Smarttoken.Contract.TotalSupply(&_Smarttoken.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Smarttoken *SmarttokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Smarttoken *SmarttokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _Smarttoken.Contract.AcceptOwnership(&_Smarttoken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_Smarttoken *SmarttokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _Smarttoken.Contract.AcceptOwnership(&_Smarttoken.TransactOpts)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address user) returns(bool success)
func (_Smarttoken *SmarttokenTransactor) AddToBlacklist(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "addToBlacklist", user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address user) returns(bool success)
func (_Smarttoken *SmarttokenSession) AddToBlacklist(user common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.AddToBlacklist(&_Smarttoken.TransactOpts, user)
}

// AddToBlacklist is a paid mutator transaction binding the contract method 0x44337ea1.
//
// Solidity: function addToBlacklist(address user) returns(bool success)
func (_Smarttoken *SmarttokenTransactorSession) AddToBlacklist(user common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.AddToBlacklist(&_Smarttoken.TransactOpts, user)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "approve", spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.Approve(&_Smarttoken.TransactOpts, spender, tokens)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactorSession) Approve(spender common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.Approve(&_Smarttoken.TransactOpts, spender, tokens)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Smarttoken *SmarttokenTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "approveAndCall", spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Smarttoken *SmarttokenSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Smarttoken.Contract.ApproveAndCall(&_Smarttoken.TransactOpts, spender, tokens, data)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 tokens, bytes data) returns(bool success)
func (_Smarttoken *SmarttokenTransactorSession) ApproveAndCall(spender common.Address, tokens *big.Int, data []byte) (*types.Transaction, error) {
	return _Smarttoken.Contract.ApproveAndCall(&_Smarttoken.TransactOpts, spender, tokens, data)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactor) MintTokens(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "mintTokens", to, tokens)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenSession) MintTokens(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.MintTokens(&_Smarttoken.TransactOpts, to, tokens)
}

// MintTokens is a paid mutator transaction binding the contract method 0xf0dda65c.
//
// Solidity: function mintTokens(address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactorSession) MintTokens(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.MintTokens(&_Smarttoken.TransactOpts, to, tokens)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(address user) returns(bool success)
func (_Smarttoken *SmarttokenTransactor) RemoveFromBlacklist(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "removeFromBlacklist", user)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(address user) returns(bool success)
func (_Smarttoken *SmarttokenSession) RemoveFromBlacklist(user common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.RemoveFromBlacklist(&_Smarttoken.TransactOpts, user)
}

// RemoveFromBlacklist is a paid mutator transaction binding the contract method 0x537df3b6.
//
// Solidity: function removeFromBlacklist(address user) returns(bool success)
func (_Smarttoken *SmarttokenTransactorSession) RemoveFromBlacklist(user common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.RemoveFromBlacklist(&_Smarttoken.TransactOpts, user)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "transfer", to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.Transfer(&_Smarttoken.TransactOpts, to, tokens)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactorSession) Transfer(to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.Transfer(&_Smarttoken.TransactOpts, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "transferFrom", from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.TransferFrom(&_Smarttoken.TransactOpts, from, to, tokens)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokens) returns(bool success)
func (_Smarttoken *SmarttokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Smarttoken.Contract.TransferFrom(&_Smarttoken.TransactOpts, from, to, tokens)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Smarttoken *SmarttokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Smarttoken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Smarttoken *SmarttokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.TransferOwnership(&_Smarttoken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Smarttoken *SmarttokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Smarttoken.Contract.TransferOwnership(&_Smarttoken.TransactOpts, _newOwner)
}

// SmarttokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Smarttoken contract.
type SmarttokenApprovalIterator struct {
	Event *SmarttokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmarttokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmarttokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmarttokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmarttokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmarttokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmarttokenApproval represents a Approval event raised by the Smarttoken contract.
type SmarttokenApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Smarttoken *SmarttokenFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*SmarttokenApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Smarttoken.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SmarttokenApprovalIterator{contract: _Smarttoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Smarttoken *SmarttokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SmarttokenApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Smarttoken.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmarttokenApproval)
				if err := _Smarttoken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Smarttoken *SmarttokenFilterer) ParseApproval(log types.Log) (*SmarttokenApproval, error) {
	event := new(SmarttokenApproval)
	if err := _Smarttoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmarttokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Smarttoken contract.
type SmarttokenOwnershipTransferredIterator struct {
	Event *SmarttokenOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmarttokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmarttokenOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmarttokenOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmarttokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmarttokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmarttokenOwnershipTransferred represents a OwnershipTransferred event raised by the Smarttoken contract.
type SmarttokenOwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Smarttoken *SmarttokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*SmarttokenOwnershipTransferredIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Smarttoken.contract.FilterLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &SmarttokenOwnershipTransferredIterator{contract: _Smarttoken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Smarttoken *SmarttokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmarttokenOwnershipTransferred, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Smarttoken.contract.WatchLogs(opts, "OwnershipTransferred", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmarttokenOwnershipTransferred)
				if err := _Smarttoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed _from, address indexed _to)
func (_Smarttoken *SmarttokenFilterer) ParseOwnershipTransferred(log types.Log) (*SmarttokenOwnershipTransferred, error) {
	event := new(SmarttokenOwnershipTransferred)
	if err := _Smarttoken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SmarttokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Smarttoken contract.
type SmarttokenTransferIterator struct {
	Event *SmarttokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SmarttokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmarttokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SmarttokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SmarttokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmarttokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmarttokenTransfer represents a Transfer event raised by the Smarttoken contract.
type SmarttokenTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Smarttoken *SmarttokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SmarttokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Smarttoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SmarttokenTransferIterator{contract: _Smarttoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Smarttoken *SmarttokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SmarttokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Smarttoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmarttokenTransfer)
				if err := _Smarttoken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Smarttoken *SmarttokenFilterer) ParseTransfer(log types.Log) (*SmarttokenTransfer, error) {
	event := new(SmarttokenTransfer)
	if err := _Smarttoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
