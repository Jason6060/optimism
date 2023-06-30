// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// LegacyOVMETHMetaData contains all meta data concerning the LegacyOVMETH contract.
var LegacyOVMETHMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BRIDGE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REMOTE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remoteToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101206040523480156200001257600080fd5b5073420000000000000000000000000000000000001060006040518060400160405280600581526020016422ba3432b960d91b815250604051806040016040528060048152602001630ae8aa8960e31b815250600160006001848481600390816200007e91906200015c565b5060046200008d82826200015c565b50505060809290925260a05260c05250506001600160a01b0390811660e052166101005262000228565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620000e257607f821691505b6020821081036200010357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200015757600081815260208120601f850160051c81016020861015620001325750805b601f850160051c820191505b8181101562000153578281556001016200013e565b5050505b505050565b81516001600160401b03811115620001785762000178620000b7565b6200019081620001898454620000cd565b8462000109565b602080601f831160018114620001c85760008415620001af5750858301515b600019600386901b1c1916600185901b17855562000153565b600085815260208120601f198616915b82811015620001f957888601518255948401946001909101908401620001d8565b5085821015620002185787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b60805160a05160c05160e051610100516116cb62000289600039600081816102f50152818161038a015281816105cf01526107a90152600081816101a9015261031b015260006107380152600061070f015260006106e601526116cb6000f3fe608060405234801561001057600080fd5b50600436106101775760003560e01c806370a08231116100d8578063ae1f6aaf1161008c578063dd62ed3e11610066578063dd62ed3e1461033f578063e78cea92146102f3578063ee9a31a21461038557600080fd5b8063ae1f6aaf146102f3578063c01e1bd614610319578063d6c0b2c41461031957600080fd5b80639dc29fac116100bd5780639dc29fac146102ba578063a457c2d7146102cd578063a9059cbb146102e057600080fd5b806370a082311461027c57806395d89b41146102b257600080fd5b806323b872dd1161012f5780633950935111610114578063395093511461024c57806340c10f191461025f57806354fd4d501461027457600080fd5b806323b872dd1461022a578063313ce5671461023d57600080fd5b806306fdde031161016057806306fdde03146101f0578063095ea7b31461020557806318160ddd1461021857600080fd5b806301ffc9a71461017c578063033964be146101a4575b600080fd5b61018f61018a366004611307565b6103ac565b60405190151581526020015b60405180910390f35b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200161019b565b6101f861049d565b60405161019b919061137c565b61018f6102133660046113f6565b61052f565b6002545b60405190815260200161019b565b61018f610238366004611420565b610547565b6040516012815260200161019b565b61018f61025a3660046113f6565b61056b565b61027261026d3660046113f6565b6105b7565b005b6101f86106df565b61021c61028a36600461145c565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b6101f8610782565b6102726102c83660046113f6565b610791565b61018f6102db3660046113f6565b6108a8565b61018f6102ee3660046113f6565b610979565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b7f00000000000000000000000000000000000000000000000000000000000000006101cb565b61021c61034d366004611477565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6101cb7f000000000000000000000000000000000000000000000000000000000000000081565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007f1d1d8b63000000000000000000000000000000000000000000000000000000007fec4fc8e3000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000851683148061046557507fffffffff00000000000000000000000000000000000000000000000000000000858116908316145b8061049457507fffffffff00000000000000000000000000000000000000000000000000000000858116908216145b95945050505050565b6060600380546104ac906114aa565b80601f01602080910402602001604051908101604052809291908181526020018280546104d8906114aa565b80156105255780601f106104fa57610100808354040283529160200191610525565b820191906000526020600020905b81548152906001019060200180831161050857829003601f168201915b5050505050905090565b60003361053d818585610987565b5060019392505050565b600033610555858285610b3b565b610560858585610c12565b506001949350505050565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff8716845290915281205490919061053d90829086906105b290879061152c565b610987565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610681576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e00000000000000000000000060648201526084015b60405180910390fd5b61068b8282610ec5565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885826040516106d391815260200190565b60405180910390a25050565b606061070a7f0000000000000000000000000000000000000000000000000000000000000000610fe5565b6107337f0000000000000000000000000000000000000000000000000000000000000000610fe5565b61075c7f0000000000000000000000000000000000000000000000000000000000000000610fe5565b60405160200161076e93929190611544565b604051602081830303815290604052905090565b6060600480546104ac906114aa565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610856576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f4f7074696d69736d4d696e7461626c6545524332303a206f6e6c79206272696460448201527f67652063616e206d696e7420616e64206275726e0000000000000000000000006064820152608401610678565b6108608282611122565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5826040516106d391815260200190565b33600081815260016020908152604080832073ffffffffffffffffffffffffffffffffffffffff871684529091528120549091908381101561096c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760448201527f207a65726f0000000000000000000000000000000000000000000000000000006064820152608401610678565b6105608286868403610987565b60003361053d818585610c12565b73ffffffffffffffffffffffffffffffffffffffff8316610a29576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff8216610acc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f20616464726560448201527f73730000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591015b60405180910390a3505050565b73ffffffffffffffffffffffffffffffffffffffff8381166000908152600160209081526040808320938616835292905220547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114610c0c5781811015610bff576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610678565b610c0c8484848403610987565b50505050565b73ffffffffffffffffffffffffffffffffffffffff8316610cb5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f20616460448201527f64726573730000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff8216610d58576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201527f65737300000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604090205481811015610e0e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e742065786365656473206260448201527f616c616e636500000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff808516600090815260208190526040808220858503905591851681529081208054849290610e5290849061152c565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610eb891815260200190565b60405180910390a3610c0c565b73ffffffffffffffffffffffffffffffffffffffff8216610f42576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610678565b8060026000828254610f54919061152c565b909155505073ffffffffffffffffffffffffffffffffffffffff821660009081526020819052604081208054839290610f8e90849061152c565b909155505060405181815273ffffffffffffffffffffffffffffffffffffffff8316906000907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9060200160405180910390a35050565b60608160000361102857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115611052578061103c816115ba565b915061104b9050600a83611621565b915061102c565b60008167ffffffffffffffff81111561106d5761106d611635565b6040519080825280601f01601f191660200182016040528015611097576020820181803683370190505b5090505b841561111a576110ac600183611664565b91506110b9600a8661167b565b6110c490603061152c565b60f81b8183815181106110d9576110d961168f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611113600a86611621565b945061109b565b949350505050565b73ffffffffffffffffffffffffffffffffffffffff82166111c5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360448201527f73000000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff82166000908152602081905260409020548181101561127b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60448201527f63650000000000000000000000000000000000000000000000000000000000006064820152608401610678565b73ffffffffffffffffffffffffffffffffffffffff831660009081526020819052604081208383039055600280548492906112b7908490611664565b909155505060405182815260009073ffffffffffffffffffffffffffffffffffffffff8516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90602001610b2e565b60006020828403121561131957600080fd5b81357fffffffff000000000000000000000000000000000000000000000000000000008116811461134957600080fd5b9392505050565b60005b8381101561136b578181015183820152602001611353565b83811115610c0c5750506000910152565b602081526000825180602084015261139b816040850160208701611350565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b803573ffffffffffffffffffffffffffffffffffffffff811681146113f157600080fd5b919050565b6000806040838503121561140957600080fd5b611412836113cd565b946020939093013593505050565b60008060006060848603121561143557600080fd5b61143e846113cd565b925061144c602085016113cd565b9150604084013590509250925092565b60006020828403121561146e57600080fd5b611349826113cd565b6000806040838503121561148a57600080fd5b611493836113cd565b91506114a1602084016113cd565b90509250929050565b600181811c908216806114be57607f821691505b6020821081036114f7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561153f5761153f6114fd565b500190565b60008451611556818460208901611350565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551611592816001850160208a01611350565b600192019182015283516115ad816002840160208801611350565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115eb576115eb6114fd565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082611630576116306115f2565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015611676576116766114fd565b500390565b60008261168a5761168a6115f2565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a",
}

// LegacyOVMETHABI is the input ABI used to generate the binding from.
// Deprecated: Use LegacyOVMETHMetaData.ABI instead.
var LegacyOVMETHABI = LegacyOVMETHMetaData.ABI

// LegacyOVMETHBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LegacyOVMETHMetaData.Bin instead.
var LegacyOVMETHBin = LegacyOVMETHMetaData.Bin

// DeployLegacyOVMETH deploys a new Ethereum contract, binding an instance of LegacyOVMETH to it.
func DeployLegacyOVMETH(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LegacyOVMETH, error) {
	parsed, err := LegacyOVMETHMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LegacyOVMETHBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LegacyOVMETH{LegacyOVMETHCaller: LegacyOVMETHCaller{contract: contract}, LegacyOVMETHTransactor: LegacyOVMETHTransactor{contract: contract}, LegacyOVMETHFilterer: LegacyOVMETHFilterer{contract: contract}}, nil
}

// LegacyOVMETH is an auto generated Go binding around an Ethereum contract.
type LegacyOVMETH struct {
	LegacyOVMETHCaller     // Read-only binding to the contract
	LegacyOVMETHTransactor // Write-only binding to the contract
	LegacyOVMETHFilterer   // Log filterer for contract events
}

// LegacyOVMETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type LegacyOVMETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyOVMETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LegacyOVMETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyOVMETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LegacyOVMETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LegacyOVMETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LegacyOVMETHSession struct {
	Contract     *LegacyOVMETH     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LegacyOVMETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LegacyOVMETHCallerSession struct {
	Contract *LegacyOVMETHCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// LegacyOVMETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LegacyOVMETHTransactorSession struct {
	Contract     *LegacyOVMETHTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// LegacyOVMETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type LegacyOVMETHRaw struct {
	Contract *LegacyOVMETH // Generic contract binding to access the raw methods on
}

// LegacyOVMETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LegacyOVMETHCallerRaw struct {
	Contract *LegacyOVMETHCaller // Generic read-only contract binding to access the raw methods on
}

// LegacyOVMETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LegacyOVMETHTransactorRaw struct {
	Contract *LegacyOVMETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLegacyOVMETH creates a new instance of LegacyOVMETH, bound to a specific deployed contract.
func NewLegacyOVMETH(address common.Address, backend bind.ContractBackend) (*LegacyOVMETH, error) {
	contract, err := bindLegacyOVMETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETH{LegacyOVMETHCaller: LegacyOVMETHCaller{contract: contract}, LegacyOVMETHTransactor: LegacyOVMETHTransactor{contract: contract}, LegacyOVMETHFilterer: LegacyOVMETHFilterer{contract: contract}}, nil
}

// NewLegacyOVMETHCaller creates a new read-only instance of LegacyOVMETH, bound to a specific deployed contract.
func NewLegacyOVMETHCaller(address common.Address, caller bind.ContractCaller) (*LegacyOVMETHCaller, error) {
	contract, err := bindLegacyOVMETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETHCaller{contract: contract}, nil
}

// NewLegacyOVMETHTransactor creates a new write-only instance of LegacyOVMETH, bound to a specific deployed contract.
func NewLegacyOVMETHTransactor(address common.Address, transactor bind.ContractTransactor) (*LegacyOVMETHTransactor, error) {
	contract, err := bindLegacyOVMETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETHTransactor{contract: contract}, nil
}

// NewLegacyOVMETHFilterer creates a new log filterer instance of LegacyOVMETH, bound to a specific deployed contract.
func NewLegacyOVMETHFilterer(address common.Address, filterer bind.ContractFilterer) (*LegacyOVMETHFilterer, error) {
	contract, err := bindLegacyOVMETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETHFilterer{contract: contract}, nil
}

// bindLegacyOVMETH binds a generic wrapper to an already deployed contract.
func bindLegacyOVMETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LegacyOVMETHMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyOVMETH *LegacyOVMETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyOVMETH.Contract.LegacyOVMETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyOVMETH *LegacyOVMETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.LegacyOVMETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyOVMETH *LegacyOVMETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.LegacyOVMETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LegacyOVMETH *LegacyOVMETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _LegacyOVMETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LegacyOVMETH *LegacyOVMETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LegacyOVMETH *LegacyOVMETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.contract.Transact(opts, method, params...)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCaller) BRIDGE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "BRIDGE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHSession) BRIDGE() (common.Address, error) {
	return _LegacyOVMETH.Contract.BRIDGE(&_LegacyOVMETH.CallOpts)
}

// BRIDGE is a free data retrieval call binding the contract method 0xee9a31a2.
//
// Solidity: function BRIDGE() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) BRIDGE() (common.Address, error) {
	return _LegacyOVMETH.Contract.BRIDGE(&_LegacyOVMETH.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCaller) REMOTETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "REMOTE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHSession) REMOTETOKEN() (common.Address, error) {
	return _LegacyOVMETH.Contract.REMOTETOKEN(&_LegacyOVMETH.CallOpts)
}

// REMOTETOKEN is a free data retrieval call binding the contract method 0x033964be.
//
// Solidity: function REMOTE_TOKEN() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) REMOTETOKEN() (common.Address, error) {
	return _LegacyOVMETH.Contract.REMOTETOKEN(&_LegacyOVMETH.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LegacyOVMETH.Contract.Allowance(&_LegacyOVMETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _LegacyOVMETH.Contract.Allowance(&_LegacyOVMETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LegacyOVMETH.Contract.BalanceOf(&_LegacyOVMETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _LegacyOVMETH.Contract.BalanceOf(&_LegacyOVMETH.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHSession) Bridge() (common.Address, error) {
	return _LegacyOVMETH.Contract.Bridge(&_LegacyOVMETH.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) Bridge() (common.Address, error) {
	return _LegacyOVMETH.Contract.Bridge(&_LegacyOVMETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyOVMETH *LegacyOVMETHCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyOVMETH *LegacyOVMETHSession) Decimals() (uint8, error) {
	return _LegacyOVMETH.Contract.Decimals(&_LegacyOVMETH.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) Decimals() (uint8, error) {
	return _LegacyOVMETH.Contract.Decimals(&_LegacyOVMETH.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCaller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHSession) L1Token() (common.Address, error) {
	return _LegacyOVMETH.Contract.L1Token(&_LegacyOVMETH.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) L1Token() (common.Address, error) {
	return _LegacyOVMETH.Contract.L1Token(&_LegacyOVMETH.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCaller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHSession) L2Bridge() (common.Address, error) {
	return _LegacyOVMETH.Contract.L2Bridge(&_LegacyOVMETH.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) L2Bridge() (common.Address, error) {
	return _LegacyOVMETH.Contract.L2Bridge(&_LegacyOVMETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHSession) Name() (string, error) {
	return _LegacyOVMETH.Contract.Name(&_LegacyOVMETH.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) Name() (string, error) {
	return _LegacyOVMETH.Contract.Name(&_LegacyOVMETH.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCaller) RemoteToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "remoteToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHSession) RemoteToken() (common.Address, error) {
	return _LegacyOVMETH.Contract.RemoteToken(&_LegacyOVMETH.CallOpts)
}

// RemoteToken is a free data retrieval call binding the contract method 0xd6c0b2c4.
//
// Solidity: function remoteToken() view returns(address)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) RemoteToken() (common.Address, error) {
	return _LegacyOVMETH.Contract.RemoteToken(&_LegacyOVMETH.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_LegacyOVMETH *LegacyOVMETHCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_LegacyOVMETH *LegacyOVMETHSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _LegacyOVMETH.Contract.SupportsInterface(&_LegacyOVMETH.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _LegacyOVMETH.Contract.SupportsInterface(&_LegacyOVMETH.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHSession) Symbol() (string, error) {
	return _LegacyOVMETH.Contract.Symbol(&_LegacyOVMETH.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) Symbol() (string, error) {
	return _LegacyOVMETH.Contract.Symbol(&_LegacyOVMETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHSession) TotalSupply() (*big.Int, error) {
	return _LegacyOVMETH.Contract.TotalSupply(&_LegacyOVMETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) TotalSupply() (*big.Int, error) {
	return _LegacyOVMETH.Contract.TotalSupply(&_LegacyOVMETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _LegacyOVMETH.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHSession) Version() (string, error) {
	return _LegacyOVMETH.Contract.Version(&_LegacyOVMETH.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_LegacyOVMETH *LegacyOVMETHCallerSession) Version() (string, error) {
	return _LegacyOVMETH.Contract.Version(&_LegacyOVMETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Approve(&_LegacyOVMETH.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Approve(&_LegacyOVMETH.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_LegacyOVMETH *LegacyOVMETHTransactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_LegacyOVMETH *LegacyOVMETHSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Burn(&_LegacyOVMETH.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_LegacyOVMETH *LegacyOVMETHTransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Burn(&_LegacyOVMETH.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.DecreaseAllowance(&_LegacyOVMETH.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.DecreaseAllowance(&_LegacyOVMETH.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.IncreaseAllowance(&_LegacyOVMETH.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.IncreaseAllowance(&_LegacyOVMETH.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_LegacyOVMETH *LegacyOVMETHTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_LegacyOVMETH *LegacyOVMETHSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Mint(&_LegacyOVMETH.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_LegacyOVMETH *LegacyOVMETHTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Mint(&_LegacyOVMETH.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Transfer(&_LegacyOVMETH.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.Transfer(&_LegacyOVMETH.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.TransferFrom(&_LegacyOVMETH.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_LegacyOVMETH *LegacyOVMETHTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _LegacyOVMETH.Contract.TransferFrom(&_LegacyOVMETH.TransactOpts, from, to, amount)
}

// LegacyOVMETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LegacyOVMETH contract.
type LegacyOVMETHApprovalIterator struct {
	Event *LegacyOVMETHApproval // Event containing the contract specifics and raw log

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
func (it *LegacyOVMETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyOVMETHApproval)
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
		it.Event = new(LegacyOVMETHApproval)
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
func (it *LegacyOVMETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyOVMETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyOVMETHApproval represents a Approval event raised by the LegacyOVMETH contract.
type LegacyOVMETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyOVMETH *LegacyOVMETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LegacyOVMETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETHApprovalIterator{contract: _LegacyOVMETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyOVMETH *LegacyOVMETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LegacyOVMETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyOVMETHApproval)
				if err := _LegacyOVMETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LegacyOVMETH *LegacyOVMETHFilterer) ParseApproval(log types.Log) (*LegacyOVMETHApproval, error) {
	event := new(LegacyOVMETHApproval)
	if err := _LegacyOVMETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyOVMETHBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the LegacyOVMETH contract.
type LegacyOVMETHBurnIterator struct {
	Event *LegacyOVMETHBurn // Event containing the contract specifics and raw log

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
func (it *LegacyOVMETHBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyOVMETHBurn)
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
		it.Event = new(LegacyOVMETHBurn)
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
func (it *LegacyOVMETHBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyOVMETHBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyOVMETHBurn represents a Burn event raised by the LegacyOVMETH contract.
type LegacyOVMETHBurn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_LegacyOVMETH *LegacyOVMETHFilterer) FilterBurn(opts *bind.FilterOpts, account []common.Address) (*LegacyOVMETHBurnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.FilterLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETHBurnIterator{contract: _LegacyOVMETH.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_LegacyOVMETH *LegacyOVMETHFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *LegacyOVMETHBurn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.WatchLogs(opts, "Burn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyOVMETHBurn)
				if err := _LegacyOVMETH.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed account, uint256 amount)
func (_LegacyOVMETH *LegacyOVMETHFilterer) ParseBurn(log types.Log) (*LegacyOVMETHBurn, error) {
	event := new(LegacyOVMETHBurn)
	if err := _LegacyOVMETH.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyOVMETHMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the LegacyOVMETH contract.
type LegacyOVMETHMintIterator struct {
	Event *LegacyOVMETHMint // Event containing the contract specifics and raw log

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
func (it *LegacyOVMETHMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyOVMETHMint)
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
		it.Event = new(LegacyOVMETHMint)
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
func (it *LegacyOVMETHMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyOVMETHMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyOVMETHMint represents a Mint event raised by the LegacyOVMETH contract.
type LegacyOVMETHMint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_LegacyOVMETH *LegacyOVMETHFilterer) FilterMint(opts *bind.FilterOpts, account []common.Address) (*LegacyOVMETHMintIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.FilterLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETHMintIterator{contract: _LegacyOVMETH.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_LegacyOVMETH *LegacyOVMETHFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *LegacyOVMETHMint, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.WatchLogs(opts, "Mint", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyOVMETHMint)
				if err := _LegacyOVMETH.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed account, uint256 amount)
func (_LegacyOVMETH *LegacyOVMETHFilterer) ParseMint(log types.Log) (*LegacyOVMETHMint, error) {
	event := new(LegacyOVMETHMint)
	if err := _LegacyOVMETH.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LegacyOVMETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LegacyOVMETH contract.
type LegacyOVMETHTransferIterator struct {
	Event *LegacyOVMETHTransfer // Event containing the contract specifics and raw log

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
func (it *LegacyOVMETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LegacyOVMETHTransfer)
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
		it.Event = new(LegacyOVMETHTransfer)
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
func (it *LegacyOVMETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LegacyOVMETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LegacyOVMETHTransfer represents a Transfer event raised by the LegacyOVMETH contract.
type LegacyOVMETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyOVMETH *LegacyOVMETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LegacyOVMETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LegacyOVMETHTransferIterator{contract: _LegacyOVMETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyOVMETH *LegacyOVMETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LegacyOVMETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LegacyOVMETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LegacyOVMETHTransfer)
				if err := _LegacyOVMETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LegacyOVMETH *LegacyOVMETHFilterer) ParseTransfer(log types.Log) (*LegacyOVMETHTransfer, error) {
	event := new(LegacyOVMETHTransfer)
	if err := _LegacyOVMETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
