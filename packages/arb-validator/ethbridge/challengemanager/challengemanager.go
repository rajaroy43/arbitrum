// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package challengemanager

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

// ArbMachineABI is the input ABI used to generate the binding from.
const ArbMachineABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"instructionStackHash\",\"type\":\"bytes32\"},{\"name\":\"dataStackHash\",\"type\":\"bytes32\"},{\"name\":\"auxStackHash\",\"type\":\"bytes32\"},{\"name\":\"registerHash\",\"type\":\"bytes32\"},{\"name\":\"staticHash\",\"type\":\"bytes32\"},{\"name\":\"errHandlerHash\",\"type\":\"bytes32\"}],\"name\":\"machineHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbMachineFuncSigs maps the 4-byte function signature to its string representation.
var ArbMachineFuncSigs = map[string]string{
	"c1355b59": "machineHash(bytes32,bytes32,bytes32,bytes32,bytes32,bytes32)",
}

// ArbMachineBin is the compiled bytecode used for deploying new contracts.
var ArbMachineBin = "0x6101d6610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063c1355b591461003a575b600080fd5b610075600480360360c081101561005057600080fd5b5080359060208101359060408101359060608101359060808101359060a00135610087565b60408051918252519081900360200190f35b604080516101008101825260e081018881528152815160208181018452888252808301919091528251808201845287815282840152825180820184528681526060830152825180820184528581526080830152825190810190925282825260a0810191909152600060c08201819052906101009061010b565b979650505050505050565b600060028260c0015114156101225750600061019c565b60018260c0015114156101375750600161019c565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101205b91905056fea265627a7a723058209091de878570b6ba76aa703337c6d539580bc5cac1d728b699601cee642258e664736f6c634300050a0032"

// DeployArbMachine deploys a new Ethereum contract, binding an instance of ArbMachine to it.
func DeployArbMachine(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbMachine, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbMachineBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// ArbMachine is an auto generated Go binding around an Ethereum contract.
type ArbMachine struct {
	ArbMachineCaller     // Read-only binding to the contract
	ArbMachineTransactor // Write-only binding to the contract
	ArbMachineFilterer   // Log filterer for contract events
}

// ArbMachineCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbMachineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbMachineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbMachineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbMachineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbMachineSession struct {
	Contract     *ArbMachine       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbMachineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbMachineCallerSession struct {
	Contract *ArbMachineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ArbMachineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbMachineTransactorSession struct {
	Contract     *ArbMachineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ArbMachineRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbMachineRaw struct {
	Contract *ArbMachine // Generic contract binding to access the raw methods on
}

// ArbMachineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbMachineCallerRaw struct {
	Contract *ArbMachineCaller // Generic read-only contract binding to access the raw methods on
}

// ArbMachineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbMachineTransactorRaw struct {
	Contract *ArbMachineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbMachine creates a new instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachine(address common.Address, backend bind.ContractBackend) (*ArbMachine, error) {
	contract, err := bindArbMachine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbMachine{ArbMachineCaller: ArbMachineCaller{contract: contract}, ArbMachineTransactor: ArbMachineTransactor{contract: contract}, ArbMachineFilterer: ArbMachineFilterer{contract: contract}}, nil
}

// NewArbMachineCaller creates a new read-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineCaller(address common.Address, caller bind.ContractCaller) (*ArbMachineCaller, error) {
	contract, err := bindArbMachine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineCaller{contract: contract}, nil
}

// NewArbMachineTransactor creates a new write-only instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbMachineTransactor, error) {
	contract, err := bindArbMachine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbMachineTransactor{contract: contract}, nil
}

// NewArbMachineFilterer creates a new log filterer instance of ArbMachine, bound to a specific deployed contract.
func NewArbMachineFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbMachineFilterer, error) {
	contract, err := bindArbMachine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbMachineFilterer{contract: contract}, nil
}

// bindArbMachine binds a generic wrapper to an already deployed contract.
func bindArbMachine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbMachineABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.ArbMachineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.ArbMachineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbMachine *ArbMachineCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbMachine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbMachine *ArbMachineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbMachine *ArbMachineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbMachine.Contract.contract.Transact(opts, method, params...)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCaller) MachineHash(opts *bind.CallOpts, instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbMachine.contract.Call(opts, out, "machineHash", instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
	return *ret0, err
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// MachineHash is a free data retrieval call binding the contract method 0xc1355b59.
//
// Solidity: function machineHash(bytes32 instructionStackHash, bytes32 dataStackHash, bytes32 auxStackHash, bytes32 registerHash, bytes32 staticHash, bytes32 errHandlerHash) constant returns(bytes32)
func (_ArbMachine *ArbMachineCallerSession) MachineHash(instructionStackHash [32]byte, dataStackHash [32]byte, auxStackHash [32]byte, registerHash [32]byte, staticHash [32]byte, errHandlerHash [32]byte) ([32]byte, error) {
	return _ArbMachine.Contract.MachineHash(&_ArbMachine.CallOpts, instructionStackHash, dataStackHash, auxStackHash, registerHash, staticHash, errHandlerHash)
}

// ArbProtocolABI is the input ABI used to generate the binding from.
const ArbProtocolABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_data\",\"type\":\"bytes32\"},{\"name\":\"_tokenType\",\"type\":\"bytes21\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_destination\",\"type\":\"address\"}],\"name\":\"generateMessageStubHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_messageTokenNums\",\"type\":\"uint16[]\"},{\"name\":\"_messageAmounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateBeforeValues\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_afterHash\",\"type\":\"bytes32\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_firstMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_lastMessageHash\",\"type\":\"bytes32\"},{\"name\":\"_firstLogHash\",\"type\":\"bytes32\"},{\"name\":\"_lastLogHash\",\"type\":\"bytes32\"},{\"name\":\"_totalMessageValueAmounts\",\"type\":\"uint256[]\"}],\"name\":\"generateAssertionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_beforeHash\",\"type\":\"bytes32\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_beforeInbox\",\"type\":\"bytes32\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"generatePreconditionHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"}],\"name\":\"beforeBalancesValid\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbProtocolFuncSigs maps the 4-byte function signature to its string representation.
var ArbProtocolFuncSigs = map[string]string{
	"af17d922": "beforeBalancesValid(bytes21[],uint256[])",
	"0f89fbff": "calculateBeforeValues(bytes21[],uint16[],uint256[])",
	"20903721": "generateAssertionHash(bytes32,uint32,bytes32,bytes32,bytes32,bytes32,uint256[])",
	"004c28f6": "generateMessageStubHash(bytes32,bytes21,uint256,address)",
	"3e285598": "generatePreconditionHash(bytes32,uint64[2],bytes32,bytes21[],uint256[])",
}

// ArbProtocolBin is the compiled bytecode used for deploying new contracts.
var ArbProtocolBin = "0x611085610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100605760003560e01c80624c28f6146100655780630f89fbff146100bd57806320903721146102b25780633e2855981461037d578063af17d922146104db575b600080fd5b6100ab6004803603608081101561007b57600080fd5b5080359060208101356affffffffffffffffffffff191690604081013590606001356001600160a01b0316610612565b60408051918252519081900360200190f35b610262600480360360608110156100d357600080fd5b810190602081018135600160201b8111156100ed57600080fd5b8201836020820111156100ff57600080fd5b803590602001918460208302840111600160201b8311171561012057600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561016f57600080fd5b82018360208201111561018157600080fd5b803590602001918460208302840111600160201b831117156101a257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101f157600080fd5b82018360208201111561020357600080fd5b803590602001918460208302840111600160201b8311171561022457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610702945050505050565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561029e578181015183820152602001610286565b505050509050019250505060405180910390f35b6100ab600480360360e08110156102c857600080fd5b81359163ffffffff6020820135169160408201359160608101359160808201359160a08101359181019060e0810160c0820135600160201b81111561030c57600080fd5b82018360208201111561031e57600080fd5b803590602001918460208302840111600160201b8311171561033f57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506108e7945050505050565b6100ab600480360360c081101561039357600080fd5b6040805180820182528335939283019291606083019190602084019060029083908390808284376000920191909152509194833594909390925060408101915060200135600160201b8111156103e857600080fd5b8201836020820111156103fa57600080fd5b803590602001918460208302840111600160201b8311171561041b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561046a57600080fd5b82018360208201111561047c57600080fd5b803590602001918460208302840111600160201b8311171561049d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610986945050505050565b6105fe600480360360408110156104f157600080fd5b810190602081018135600160201b81111561050b57600080fd5b82018360208201111561051d57600080fd5b803590602001918460208302840111600160201b8311171561053e57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561058d57600080fd5b82018360208201111561059f57600080fd5b803590602001918460208302840111600160201b831117156105c057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610a71945050505050565b604080519115158252519081900360200190f35b60408051600480825260a0820190925260009160609190816020015b61063661101a565b81526020019060019003908161062e57905050905061065486610c76565b8160008151811061066157fe5b602002602001018190525061067e836001600160a01b0316610cd2565b8160018151811061068b57fe5b602002602001018190525061069f84610cd2565b816002815181106106ac57fe5b60209081029190910101526106ce6affffffffffffffffffffff198616610cd2565b816003815181106106db57fe5b60200260200101819052506106f76106f282610d2c565b610db4565b519695505050505050565b606060008351905060608551604051908082528060200260200182016040528015610737578160200160208202803883390190505b50905060005b828110156108dd57600086828151811061075357fe5b60200260200101519050878161ffff168151811061076d57fe5b602002602001015160146015811061078157fe5b1a60f81b6001600160f81b0319166107ce5785828151811061079f57fe5b6020026020010151838261ffff16815181106107b757fe5b6020026020010181815101915081815250506108d4565b828161ffff16815181106107de57fe5b602002602001015160001461083a576040805162461bcd60e51b815260206004820152601d60248201527f43616e277420696e636c756465204e465420746f6b656e207477696365000000604482015290519081900360640190fd5b85828151811061084657fe5b6020026020010151600014156108a3576040805162461bcd60e51b815260206004820152601f60248201527f4e465420746f6b656e206d7573742068617665206e6f6e2d7a65726f20696400604482015290519081900360640190fd5b8582815181106108af57fe5b6020026020010151838261ffff16815181106108c757fe5b6020026020010181815250505b5060010161073d565b5095945050505050565b600087878787878787604051602001808881526020018763ffffffff1663ffffffff1660e01b8152600401868152602001858152602001848152602001838152602001828051906020019060200280838360005b8381101561095357818101518382015260200161093b565b50505050905001975050505050505050604051602081830303815290604052805190602001209050979650505050505050565b600085858260200201518660016020020151868686604051602001808781526020018667ffffffffffffffff1667ffffffffffffffff1660c01b81526008018567ffffffffffffffff1667ffffffffffffffff1660c01b8152600801848152602001838051906020019060200280838360005b83811015610a115781810151838201526020016109f9565b50505050905001828051906020019060200280838360005b83811015610a41578181015183820152602001610a29565b50505050905001965050505050505060405160208183030381529060405280519060200120905095945050505050565b8151600090801580610a835750806001145b15610a92576001915050610c70565b60005b60018203811015610c23576000858281518110610aae57fe5b6020026020010151601460158110610ac257fe5b1a60f81b90506001600160f81b03198116610b2c57858281518110610ae357fe5b60200260200101516001600160581b031916868360010181518110610b0457fe5b60200260200101516001600160581b03191611610b275760009350505050610c70565b610c1a565b600160f81b6001600160f81b031982161415610c0e57858281518110610b4e57fe5b60200260200101516001600160581b031916868360010181518110610b6f57fe5b60200260200101516001600160581b0319161080610bfd5750858281518110610b9457fe5b60200260200101516001600160581b031916868360010181518110610bb557fe5b60200260200101516001600160581b031916148015610bfd5750848281518110610bdb57fe5b6020026020010151858360010181518110610bf257fe5b602002602001015111155b15610b275760009350505050610c70565b60009350505050610c70565b50600101610a95565b50600160f81b846001830381518110610c3857fe5b6020026020010151601460158110610c4c57fe5b1a60f81b6001600160f81b0319161115610c6a576000915050610c70565b60019150505b92915050565b610c7e61101a565b604080516060810182528381528151600080825260208281019094529192830191610cbf565b610cac61101a565b815260200190600190039081610ca45790505b508152600260209091015290505b919050565b610cda61101a565b604080516060810182528381528151600080825260208281019094529192830191610d1b565b610d0861101a565b815260200190600190039081610d005790505b508152600060209091015292915050565b610d3461101a565b610d3e8251610ea3565b610d8f576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b610dbc61103e565b6040820151600c60ff90911610610e0e576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff16610e3b576040518060200160405280610e328460000151610eaa565b90529050610ccd565b604082015160ff1660021415610e605750604080516020810190915281518152610ccd565b600360ff16826040015160ff1610158015610e8457506040820151600c60ff909116105b15610ea1576040518060200160405280610e328460200151610ece565bfe5b6008101590565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115610f1e576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610f4b578160200160208202803883390190505b50805190915060005b81811015610fa757610f6461103e565b610f80868381518110610f7357fe5b6020026020010151610db4565b90508060000151848381518110610f9357fe5b602090810291909101015250600101610f54565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610ff0578181015183820152602001610fd8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820b6e1c1b7fc73c26e24e077c00f32c02f63f47647b21d08dbf84d7799397dbe3b64736f6c634300050a0032"

// DeployArbProtocol deploys a new Ethereum contract, binding an instance of ArbProtocol to it.
func DeployArbProtocol(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbProtocol, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbProtocolBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// ArbProtocol is an auto generated Go binding around an Ethereum contract.
type ArbProtocol struct {
	ArbProtocolCaller     // Read-only binding to the contract
	ArbProtocolTransactor // Write-only binding to the contract
	ArbProtocolFilterer   // Log filterer for contract events
}

// ArbProtocolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbProtocolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbProtocolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbProtocolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbProtocolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbProtocolSession struct {
	Contract     *ArbProtocol      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbProtocolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbProtocolCallerSession struct {
	Contract *ArbProtocolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArbProtocolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbProtocolTransactorSession struct {
	Contract     *ArbProtocolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArbProtocolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbProtocolRaw struct {
	Contract *ArbProtocol // Generic contract binding to access the raw methods on
}

// ArbProtocolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbProtocolCallerRaw struct {
	Contract *ArbProtocolCaller // Generic read-only contract binding to access the raw methods on
}

// ArbProtocolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbProtocolTransactorRaw struct {
	Contract *ArbProtocolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbProtocol creates a new instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocol(address common.Address, backend bind.ContractBackend) (*ArbProtocol, error) {
	contract, err := bindArbProtocol(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbProtocol{ArbProtocolCaller: ArbProtocolCaller{contract: contract}, ArbProtocolTransactor: ArbProtocolTransactor{contract: contract}, ArbProtocolFilterer: ArbProtocolFilterer{contract: contract}}, nil
}

// NewArbProtocolCaller creates a new read-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolCaller(address common.Address, caller bind.ContractCaller) (*ArbProtocolCaller, error) {
	contract, err := bindArbProtocol(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolCaller{contract: contract}, nil
}

// NewArbProtocolTransactor creates a new write-only instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbProtocolTransactor, error) {
	contract, err := bindArbProtocol(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolTransactor{contract: contract}, nil
}

// NewArbProtocolFilterer creates a new log filterer instance of ArbProtocol, bound to a specific deployed contract.
func NewArbProtocolFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbProtocolFilterer, error) {
	contract, err := bindArbProtocol(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbProtocolFilterer{contract: contract}, nil
}

// bindArbProtocol binds a generic wrapper to an already deployed contract.
func bindArbProtocol(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbProtocolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.ArbProtocolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.ArbProtocolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbProtocol *ArbProtocolCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbProtocol.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbProtocol *ArbProtocolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbProtocol.Contract.contract.Transact(opts, method, params...)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCaller) BeforeBalancesValid(opts *bind.CallOpts, _tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "beforeBalancesValid", _tokenTypes, _beforeBalances)
	return *ret0, err
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// BeforeBalancesValid is a free data retrieval call binding the contract method 0xaf17d922.
//
// Solidity: function beforeBalancesValid(bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bool)
func (_ArbProtocol *ArbProtocolCallerSession) BeforeBalancesValid(_tokenTypes [][21]byte, _beforeBalances []*big.Int) (bool, error) {
	return _ArbProtocol.Contract.BeforeBalancesValid(&_ArbProtocol.CallOpts, _tokenTypes, _beforeBalances)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCaller) CalculateBeforeValues(opts *bind.CallOpts, _tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "calculateBeforeValues", _tokenTypes, _messageTokenNums, _messageAmounts)
	return *ret0, err
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// CalculateBeforeValues is a free data retrieval call binding the contract method 0x0f89fbff.
//
// Solidity: function calculateBeforeValues(bytes21[] _tokenTypes, uint16[] _messageTokenNums, uint256[] _messageAmounts) constant returns(uint256[])
func (_ArbProtocol *ArbProtocolCallerSession) CalculateBeforeValues(_tokenTypes [][21]byte, _messageTokenNums []uint16, _messageAmounts []*big.Int) ([]*big.Int, error) {
	return _ArbProtocol.Contract.CalculateBeforeValues(&_ArbProtocol.CallOpts, _tokenTypes, _messageTokenNums, _messageAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateAssertionHash(opts *bind.CallOpts, _afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateAssertionHash", _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
	return *ret0, err
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateAssertionHash is a free data retrieval call binding the contract method 0x20903721.
//
// Solidity: function generateAssertionHash(bytes32 _afterHash, uint32 _numSteps, bytes32 _firstMessageHash, bytes32 _lastMessageHash, bytes32 _firstLogHash, bytes32 _lastLogHash, uint256[] _totalMessageValueAmounts) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateAssertionHash(_afterHash [32]byte, _numSteps uint32, _firstMessageHash [32]byte, _lastMessageHash [32]byte, _firstLogHash [32]byte, _lastLogHash [32]byte, _totalMessageValueAmounts []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateAssertionHash(&_ArbProtocol.CallOpts, _afterHash, _numSteps, _firstMessageHash, _lastMessageHash, _firstLogHash, _lastLogHash, _totalMessageValueAmounts)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GenerateMessageStubHash(opts *bind.CallOpts, _data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generateMessageStubHash", _data, _tokenType, _value, _destination)
	return *ret0, err
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GenerateMessageStubHash is a free data retrieval call binding the contract method 0x004c28f6.
//
// Solidity: function generateMessageStubHash(bytes32 _data, bytes21 _tokenType, uint256 _value, address _destination) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GenerateMessageStubHash(_data [32]byte, _tokenType [21]byte, _value *big.Int, _destination common.Address) ([32]byte, error) {
	return _ArbProtocol.Contract.GenerateMessageStubHash(&_ArbProtocol.CallOpts, _data, _tokenType, _value, _destination)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCaller) GeneratePreconditionHash(opts *bind.CallOpts, _beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbProtocol.contract.Call(opts, out, "generatePreconditionHash", _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
	return *ret0, err
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// GeneratePreconditionHash is a free data retrieval call binding the contract method 0x3e285598.
//
// Solidity: function generatePreconditionHash(bytes32 _beforeHash, uint64[2] _timeBounds, bytes32 _beforeInbox, bytes21[] _tokenTypes, uint256[] _beforeBalances) constant returns(bytes32)
func (_ArbProtocol *ArbProtocolCallerSession) GeneratePreconditionHash(_beforeHash [32]byte, _timeBounds [2]uint64, _beforeInbox [32]byte, _tokenTypes [][21]byte, _beforeBalances []*big.Int) ([32]byte, error) {
	return _ArbProtocol.Contract.GeneratePreconditionHash(&_ArbProtocol.CallOpts, _beforeHash, _timeBounds, _beforeInbox, _tokenTypes, _beforeBalances)
}

// ArbValueABI is the input ABI used to generate the binding from.
const ArbValueABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"hashIntValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"getNextValidValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"immediateVal\",\"type\":\"bytes32\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointImmediateValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashEmptyTuple\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"opcode\",\"type\":\"uint8\"},{\"name\":\"nextCodePoint\",\"type\":\"bytes32\"}],\"name\":\"hashCodePointBasicValue\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"offset\",\"type\":\"uint256\"}],\"name\":\"deserializeValidValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deserializeValueHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"isValidTupleSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// ArbValueFuncSigs maps the 4-byte function signature to its string representation.
var ArbValueFuncSigs = map[string]string{
	"89df40da": "deserializeValidValueHash(bytes,uint256)",
	"8f346036": "deserializeValueHash(bytes)",
	"1f3d4d4e": "getNextValidValue(bytes,uint256)",
	"53409fab": "hashCodePointBasicValue(uint8,bytes32)",
	"264f384b": "hashCodePointImmediateValue(uint8,bytes32,bytes32)",
	"364df277": "hashEmptyTuple()",
	"1667b411": "hashIntValue(uint256)",
	"b2b9dc62": "isValidTupleSize(uint256)",
}

// ArbValueBin is the compiled bytecode used for deploying new contracts.
var ArbValueBin = "0x610d71610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100925760003560e01c806353409fab1161006557806353409fab1461022157806389df40da146102475780638f34603614610308578063b2b9dc62146103ae57610092565b80631667b411146100975780631f3d4d4e146100c6578063264f384b146101ed578063364df27714610219575b600080fd5b6100b4600480360360208110156100ad57600080fd5b50356103df565b60408051918252519081900360200190f35b61016e600480360360408110156100dc57600080fd5b8101906020810181356401000000008111156100f757600080fd5b82018360208201111561010957600080fd5b8035906020019184600183028401116401000000008311171561012b57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250610405915050565b6040518083815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101b1578181015183820152602001610199565b50505050905090810190601f1680156101de5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6100b46004803603606081101561020357600080fd5b5060ff813516906020810135906040013561049a565b6100b46104ec565b6100b46004803603604081101561023757600080fd5b5060ff813516906020013561055f565b6102ef6004803603604081101561025d57600080fd5b81019060208101813564010000000081111561027857600080fd5b82018360208201111561028a57600080fd5b803590602001918460018302840111640100000000831117156102ac57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050913592506105a6915050565b6040805192835260208301919091528051918290030190f35b6100b46004803603602081101561031e57600080fd5b81019060208101813564010000000081111561033957600080fd5b82018360208201111561034b57600080fd5b8035906020019184600183028401116401000000008311171561036d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610631945050505050565b6103cb600480360360208110156103c457600080fd5b50356106b5565b604080519115158252519081900360200190f35b60408051602080820184905282518083038201815291830190925280519101205b919050565b60006060600080610414610d06565b61041e87876106bc565b919450925090508215610478576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161048c888880840363ffffffff61081116565b945094505050509250929050565b60408051600160f81b60208083019190915260f89590951b6001600160f81b03191660218201526022810193909352604280840192909252805180840390920182526062909201909152805191012090565b6040805160008082526020808301808552600360f81b948401948552835192946003938593919260418501929091028083838a5b83811015610538578181015183820152602001610520565b50505050905001925050506040516020818303038152906040528051906020012091505090565b60408051600160f81b60208083019190915260f89490941b6001600160f81b0319166021820152602280820193909352815180820390930183526042019052805191012090565b6000806000806105b4610d06565b6105be87876106bc565b919450925090508215610618576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b8161062282610891565b51909890975095505050505050565b6000808061063d610d06565b6106488560006106bc565b9194509250905082156106a2576040805162461bcd60e51b815260206004820152601e60248201527f4d61727368616c6c65642076616c7565206d7573742062652076616c69640000604482015290519081900360640190fd5b6106ab81610891565b5195945050505050565b6008101590565b6000806106c7610d06565b8451841061071c576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b6000849050600086828151811061072f57fe5b016020015160019092019160f81c9050600081610771576107508884610980565b9093509050600083610761836109a7565b9197509550935061080a92505050565b60ff821660021415610798576107878884610980565b909350905060008361076183610a01565b600360ff8316108015906107af5750600c60ff8316105b156107eb576002198201606060006107c8838c88610a5b565b9097509250905080866107da84610b16565b98509850985050505050505061080a565b8160ff166127100160006107ff60006109a7565b919750955093505050505b9250925092565b60608183018451101561082357600080fd5b60608215801561083e57604051915060208201604052610888565b6040519150601f8416801560200281840101858101878315602002848b0101015b8183101561087757805183526020928301920161085f565b5050858452601f01601f1916604052505b50949350505050565b610899610d2a565b6040820151600c60ff909116106108eb576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff1661091857604051806020016040528061090f84600001516103df565b90529050610400565b604082015160ff166002141561093d5750604080516020810190915281518152610400565b600360ff16826040015160ff161015801561096157506040820151600c60ff909116105b1561097e57604051806020016040528061090f8460200151610b9e565bfe5b6000808281610995868363ffffffff610cea16565b60209290920196919550909350505050565b6109af610d06565b6040805160608101825283815281516000808252602082810190945291928301916109f0565b6109dd610d06565b8152602001906001900390816109d55790505b508152600060209091015292915050565b610a09610d06565b604080516060810182528381528151600080825260208281019094529192830191610a4a565b610a37610d06565b815260200190600190039081610a2f5790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015610aa657816020015b610a93610d06565b815260200190600190039081610a8b5790505b50905060005b8960ff168160ff161015610b0057610ac489856106bc565b8451859060ff8616908110610ad557fe5b6020908102919091010152945092508215610af857509094509092509050610b0d565b600101610aac565b5060009550919350909150505b93509350939050565b610b1e610d06565b610b2882516106b5565b610b79576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b6000600882511115610bee576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015610c1b578160200160208202803883390190505b50805190915060005b81811015610c7757610c34610d2a565b610c50868381518110610c4357fe5b6020026020010151610891565b90508060000151848381518110610c6357fe5b602090810291909101015250600101610c24565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015610cc0578181015183820152602001610ca8565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b60008160200183511015610cfd57600080fd5b50016020015190565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fea265627a7a72305820df9accba692330c50b5a0c6ce3f31382e1447f19a4cbff3ad62e18c40a9b21de64736f6c634300050a0032"

// DeployArbValue deploys a new Ethereum contract, binding an instance of ArbValue to it.
func DeployArbValue(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbValue, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbValueBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// ArbValue is an auto generated Go binding around an Ethereum contract.
type ArbValue struct {
	ArbValueCaller     // Read-only binding to the contract
	ArbValueTransactor // Write-only binding to the contract
	ArbValueFilterer   // Log filterer for contract events
}

// ArbValueCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbValueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbValueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbValueFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbValueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbValueSession struct {
	Contract     *ArbValue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbValueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbValueCallerSession struct {
	Contract *ArbValueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbValueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbValueTransactorSession struct {
	Contract     *ArbValueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbValueRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbValueRaw struct {
	Contract *ArbValue // Generic contract binding to access the raw methods on
}

// ArbValueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbValueCallerRaw struct {
	Contract *ArbValueCaller // Generic read-only contract binding to access the raw methods on
}

// ArbValueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbValueTransactorRaw struct {
	Contract *ArbValueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbValue creates a new instance of ArbValue, bound to a specific deployed contract.
func NewArbValue(address common.Address, backend bind.ContractBackend) (*ArbValue, error) {
	contract, err := bindArbValue(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbValue{ArbValueCaller: ArbValueCaller{contract: contract}, ArbValueTransactor: ArbValueTransactor{contract: contract}, ArbValueFilterer: ArbValueFilterer{contract: contract}}, nil
}

// NewArbValueCaller creates a new read-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueCaller(address common.Address, caller bind.ContractCaller) (*ArbValueCaller, error) {
	contract, err := bindArbValue(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueCaller{contract: contract}, nil
}

// NewArbValueTransactor creates a new write-only instance of ArbValue, bound to a specific deployed contract.
func NewArbValueTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbValueTransactor, error) {
	contract, err := bindArbValue(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbValueTransactor{contract: contract}, nil
}

// NewArbValueFilterer creates a new log filterer instance of ArbValue, bound to a specific deployed contract.
func NewArbValueFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbValueFilterer, error) {
	contract, err := bindArbValue(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbValueFilterer{contract: contract}, nil
}

// bindArbValue binds a generic wrapper to an already deployed contract.
func bindArbValue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbValueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.ArbValueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.ArbValueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbValue *ArbValueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ArbValue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbValue *ArbValueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbValue *ArbValueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbValue.Contract.contract.Transact(opts, method, params...)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValidValueHash(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "deserializeValidValueHash", data, offset)
	return *ret0, *ret1, err
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValidValueHash is a free data retrieval call binding the contract method 0x89df40da.
//
// Solidity: function deserializeValidValueHash(bytes data, uint256 offset) constant returns(uint256, bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValidValueHash(data []byte, offset *big.Int) (*big.Int, [32]byte, error) {
	return _ArbValue.Contract.DeserializeValidValueHash(&_ArbValue.CallOpts, data, offset)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) DeserializeValueHash(opts *bind.CallOpts, data []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "deserializeValueHash", data)
	return *ret0, err
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// DeserializeValueHash is a free data retrieval call binding the contract method 0x8f346036.
//
// Solidity: function deserializeValueHash(bytes data) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) DeserializeValueHash(data []byte) ([32]byte, error) {
	return _ArbValue.Contract.DeserializeValueHash(&_ArbValue.CallOpts, data)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCaller) GetNextValidValue(opts *bind.CallOpts, data []byte, offset *big.Int) (*big.Int, []byte, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _ArbValue.contract.Call(opts, out, "getNextValidValue", data, offset)
	return *ret0, *ret1, err
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// GetNextValidValue is a free data retrieval call binding the contract method 0x1f3d4d4e.
//
// Solidity: function getNextValidValue(bytes data, uint256 offset) constant returns(uint256, bytes)
func (_ArbValue *ArbValueCallerSession) GetNextValidValue(data []byte, offset *big.Int) (*big.Int, []byte, error) {
	return _ArbValue.Contract.GetNextValidValue(&_ArbValue.CallOpts, data, offset)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointBasicValue(opts *bind.CallOpts, opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointBasicValue", opcode, nextCodePoint)
	return *ret0, err
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointBasicValue is a free data retrieval call binding the contract method 0x53409fab.
//
// Solidity: function hashCodePointBasicValue(uint8 opcode, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointBasicValue(opcode uint8, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointBasicValue(&_ArbValue.CallOpts, opcode, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashCodePointImmediateValue(opts *bind.CallOpts, opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashCodePointImmediateValue", opcode, immediateVal, nextCodePoint)
	return *ret0, err
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashCodePointImmediateValue is a free data retrieval call binding the contract method 0x264f384b.
//
// Solidity: function hashCodePointImmediateValue(uint8 opcode, bytes32 immediateVal, bytes32 nextCodePoint) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashCodePointImmediateValue(opcode uint8, immediateVal [32]byte, nextCodePoint [32]byte) ([32]byte, error) {
	return _ArbValue.Contract.HashCodePointImmediateValue(&_ArbValue.CallOpts, opcode, immediateVal, nextCodePoint)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashEmptyTuple(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashEmptyTuple")
	return *ret0, err
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashEmptyTuple is a free data retrieval call binding the contract method 0x364df277.
//
// Solidity: function hashEmptyTuple() constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashEmptyTuple() ([32]byte, error) {
	return _ArbValue.Contract.HashEmptyTuple(&_ArbValue.CallOpts)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCaller) HashIntValue(opts *bind.CallOpts, val *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "hashIntValue", val)
	return *ret0, err
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// HashIntValue is a free data retrieval call binding the contract method 0x1667b411.
//
// Solidity: function hashIntValue(uint256 val) constant returns(bytes32)
func (_ArbValue *ArbValueCallerSession) HashIntValue(val *big.Int) ([32]byte, error) {
	return _ArbValue.Contract.HashIntValue(&_ArbValue.CallOpts, val)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCaller) IsValidTupleSize(opts *bind.CallOpts, size *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ArbValue.contract.Call(opts, out, "isValidTupleSize", size)
	return *ret0, err
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// IsValidTupleSize is a free data retrieval call binding the contract method 0xb2b9dc62.
//
// Solidity: function isValidTupleSize(uint256 size) constant returns(bool)
func (_ArbValue *ArbValueCallerSession) IsValidTupleSize(size *big.Int) (bool, error) {
	return _ArbValue.Contract.IsValidTupleSize(&_ArbValue.CallOpts, size)
}

// BisectionABI is the input ABI used to generate the binding from.
const BisectionABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assertionIndex\",\"type\":\"uint256\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bisectionHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"}],\"name\":\"BisectedAssertionFirst\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bisectionHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"prevOutputValues\",\"type\":\"uint256[]\"}],\"name\":\"BisectedAssertionOther\",\"type\":\"event\"}]"

// BisectionFuncSigs maps the 4-byte function signature to its string representation.
var BisectionFuncSigs = map[string]string{
	"353a5864": "bisectAssertionFirst(Challenge.Data storage,uint32,bytes32,bytes32,bytes32,bytes32[])",
	"8014c3e9": "bisectAssertionOther(Challenge.Data storage,bytes32[10],uint256[],uint32,uint256[],uint32,uint256[],bytes32[])",
	"110112ae": "continueChallenge(Challenge.Data storage,uint256,bytes,bytes32,bytes32)",
}

// BisectionBin is the compiled bytecode used for deploying new contracts.
var BisectionBin = "0x611640610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c8063110112ae14610050578063353a5864146101145780638014c3e9146101e7575b600080fd5b81801561005c57600080fd5b50610112600480360360a081101561007357600080fd5b813591602081013591810190606081016040820135600160201b81111561009957600080fd5b8201836020820111156100ab57600080fd5b803590602001918460018302840111600160201b831117156100cc57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135610471565b005b81801561012057600080fd5b50610112600480360360c081101561013757600080fd5b81359163ffffffff602082013516916040820135916060810135916080820135919081019060c0810160a0820135600160201b81111561017657600080fd5b82018360208201111561018857600080fd5b803590602001918460208302840111600160201b831117156101a957600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061074c945050505050565b8180156101f357600080fd5b50610112600480360361022081101561020b57600080fd5b604080516101408181019092528335939283019291610160830191906020840190600a90839083908082843760009201919091525091949392602081019250359050600160201b81111561025e57600080fd5b82018360208201111561027057600080fd5b803590602001918460208302840111600160201b8311171561029157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b8111156102ee57600080fd5b82018360208201111561030057600080fd5b803590602001918460208302840111600160201b8311171561032157600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b81111561037e57600080fd5b82018360208201111561039057600080fd5b803590602001918460208302840111600160201b831117156103b157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561040057600080fd5b82018360208201111561041257600080fd5b803590602001918460208302840111600160201b8311171561043357600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061089e945050505050565b846001015482146104b35760405162461bcd60e51b815260040180806020018281038252602b8152602001806115b2602b913960400191505060405180910390fd5b600585015467ffffffffffffffff16431115610516576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038501600101546001600160a01b031633146105645760405162461bcd60e51b815260040180806020018281038252602f8152602001806115dd602f913960400191505060405180910390fd5b73__$800fcb2f4a98daa165a5cdb21a355d7a15$__63b792d767848484886001016040518563ffffffff1660e01b81526004018080602001858152602001848152602001838152602001828103825286818151815260200191508051906020019080838360005b838110156105e35781810151838201526020016105cb565b50505050905090810190601f1680156106105780820380516001836020036101000a031916815260200191505b509550505050505060206040518083038186803b15801561063057600080fd5b505af4158015610644573d6000803e3d6000fd5b505050506040513d602081101561065a57600080fd5b50516106ad576040805162461bcd60e51b815260206004820152601a60248201527f496e76616c696420617373657274696f6e2073656c6563746564000000000000604482015290519081900360640190fd5b60058501805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160601b1793840416011667ffffffffffffffff19919091161790556001850181905584546004860154604080516001600160a01b03928316815260208101889052815192909316927f18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00929081900390910190a25050505050565b610755866109cc565b6001860154604080516020808201889052818301879052606082018690526001600160e01b031960e08a901b1660808301528251606481840301815260849092019092528051910120146107da5760405162461bcd60e51b815260040180806020018281038252603081526020018061154e6030913960400191505060405180910390fd5b6107e8868585858986610aee565b85546001600160a01b03167f06b770c15c61f5df9b9c6388352469b626ca74d653c33f394b93417c16bbc9e06003880160000154604080516001600160a01b0390921680835263ffffffff8a1691830191909152606060208084018281528751928501929092528651929387938c9391926080840191818701910280838360005b83811015610881578181015183820152602001610869565b5050505090500194505050505060405180910390a2505050505050565b6109c288604051806101e001604052808a6000600a81106108bb57fe5b602002015181526020018a6001600a81106108d257fe5b602002015181526020018981526020018a6002600a81106108ef57fe5b602002015181526020018a6003600a811061090657fe5b602002015181526020018a6004600a811061091d57fe5b602002015181526020018863ffffffff1681526020018a6005600a811061094057fe5b602002015181526020018a6006600a811061095757fe5b602002015181526020018781526020018a6007600a811061097457fe5b602002015181526020018663ffffffff1681526020018a6008600a811061099757fe5b602002015181526020018a6009600a81106109ae57fe5b602002015181526020018581525083610ec9565b5050505050505050565b600581015467ffffffffffffffff16431115610a2f576040805162461bcd60e51b815260206004820152601a60248201527f4368616c6c656e676520646561646c696e652065787069726564000000000000604482015290519081900360640190fd5b60038101600001546001600160a01b03163314610a93576040805162461bcd60e51b815260206004820181905260248201527f4f6e6c79206f7269676e616c2061737365727465722063616e20626973656374604482015290519081900360640190fd5b60016005820154600160601b900460ff166002811115610aaf57fe5b14610aeb5760405162461bcd60e51b815260040180806020018281038252603481526020018061157e6034913960400191505060405180910390fd5b50565b80516001016006811480610b18575060068363ffffffff16108015610b1857508263ffffffff1681145b610b69576040805162461bcd60e51b815260206004820152601960248201527f496e636f727265637420626973656374696f6e20636f756e7400000000000000604482015290519081900360640190fd5b60008163ffffffff168463ffffffff1681610b8057fe5b0490506000828563ffffffff1681610b9457fe5b061115610b9f576001015b606082600101604051908082528060200260200182016040528015610bce578160200160208202803883390190505b509050878785600081518110610be057fe5b602002602001015184604051602001808581526020018481526020018381526020018263ffffffff1663ffffffff1660e01b81526004019450505050506040516020818303038152906040528051906020012081600081518110610c4057fe5b602090810291909101015260015b60018403811015610d1657838663ffffffff1681610c6857fe5b06811415610c7857600019909201915b8888866001840381518110610c8957fe5b6020026020010151878481518110610c9d57fe5b602002602001015186604051602001808681526020018581526020018481526020018381526020018263ffffffff1663ffffffff1660e01b81526004019550505050505060405160208183030381529060405280519060200120828281518110610d0357fe5b6020908102919091010152600101610c4e565b50828563ffffffff1681610d2657fe5b06600184031415610d3957600019909101905b8787856001860381518110610d4a57fe5b60200260200101518885604051602001808681526020018581526020018481526020018381526020018263ffffffff1663ffffffff1660e01b81526004019550505050505060405160208183030381529060405280519060200120818481518110610db157fe5b60209081029190910181019190915260058a01805467ffffffffffffffff4363ffffffff600160401b60ff60601b19909416600160611b1793840416011667ffffffffffffffff19919091161790556040516309898dc160e41b81526004810182815283516024830152835173__$800fcb2f4a98daa165a5cdb21a355d7a15$__93639898dc10938693928392604490920191858101910280838360005b83811015610e67578181015183820152602001610e4f565b505050509050019250505060206040518083038186803b158015610e8a57600080fd5b505af4158015610e9e573d6000803e3d6000fd5b505050506040513d6020811015610eb457600080fd5b50516001909901989098555050505050505050565b610ed2836109cc565b610edc8383611187565b60008260c0015183610160015103905061106e8484602001518560a00151866040015160405160200180838152602001828051906020019060200280838360005b83811015610f35578181015183820152602001610f1d565b50505050905001925050506040516020818303038152906040528051906020012073__$9836fa7140e5a33041d4b827682e675a30$__6320903721886101400151878a606001518b6101a001518c608001518d6101a001518e6101c001516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015611012578181015183820152602001610ffa565b505050509050019850505050505050505060206040518083038186803b15801561103b57600080fd5b505af415801561104f573d6000803e3d6000fd5b505050506040513d602081101561106557600080fd5b50518587610aee565b83546001600160a01b03167f4f9f2ba0780165ec676b48b1d9a57ff77e0bd1e331e18b7888da37378200335d6003860160000160009054906101000a90046001600160a01b0316848487610120015160405180856001600160a01b03166001600160a01b03168152602001806020018463ffffffff1663ffffffff16815260200180602001838103835286818151815260200191508051906020019060200280838360005b8381101561112b578181015183820152602001611113565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561116a578181015183820152602001611152565b50505050905001965050505050505060405180910390a250505050565b816001015481602001518260000151836040015160405160200180838152602001828051906020019060200280838360005b838110156111d15781810151838201526020016111b9565b50505050905001925050506040516020818303038152906040528051906020012073__$9836fa7140e5a33041d4b827682e675a30$__63209037218560a001518660c0015187606001518860e0015189608001518a61010001518b61012001516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156112b0578181015183820152602001611298565b505050509050019850505050505050505060206040518083038186803b1580156112d957600080fd5b505af41580156112ed573d6000803e3d6000fd5b505050506040513d602081101561130357600080fd5b505161014085015161016086015160608701516101a088015160808901516101c08a0151604051632090372160e01b81526004810187815263ffffffff8716602483015260448201869052606482018590526084820184905260a4820185905260e060c48301908152835160e4840152835173__$9836fa7140e5a33041d4b827682e675a30$__9963209037219990989097909690959094869491939192610104909101906020858101910280838360005b838110156113cd5781810151838201526020016113b5565b505050509050019850505050505050505060206040518083038186803b1580156113f657600080fd5b505af415801561140a573d6000803e3d6000fd5b505050506040513d602081101561142057600080fd5b505160c08601516101608701516040805160208181019890985280820196909652606086019490945260808501929092526001600160e01b0319910360e01b1660a08301528051608481840301815260a490920190528051910120146114b75760405162461bcd60e51b815260040180806020018281038252603081526020018061154e6030913960400191505060405180910390fd5b60408101515160005b818110156115475782610120015181815181106114d957fe5b6020026020010151836040015182815181106114f157fe5b602002602001018181510391508181525050826101200151818151811061151457fe5b6020026020010151836101c00151828151811061152d57fe5b6020908102919091010180519190910390526001016114c0565b5050505056fe4269736563746f7220696e636f72726563746c792072657665616c656420626973656374696f6e207365676d656e747343616e206f6e6c792062697365637420617373657274696f6e20696e20726573706f6e736520746f2061206368616c6c656e6765636f6e74696e75654368616c6c656e67653a20496e636f72726563742070726576696f75732073746174654f6e6c79206f726967696e616c206368616c6c656e6765722063616e20636f6e74696e7565206368616c6c656e6765a265627a7a723058206a49747f24538a9e3d8829a801160295a199c62488232c0d1fb7ed57c5ac49f564736f6c634300050a0032"

// DeployBisection deploys a new Ethereum contract, binding an instance of Bisection to it.
func DeployBisection(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bisection, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	merkleLibAddr, _, _, _ := DeployMerkleLib(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$800fcb2f4a98daa165a5cdb21a355d7a15$__", merkleLibAddr.String()[2:], -1)

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	BisectionBin = strings.Replace(BisectionBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BisectionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// Bisection is an auto generated Go binding around an Ethereum contract.
type Bisection struct {
	BisectionCaller     // Read-only binding to the contract
	BisectionTransactor // Write-only binding to the contract
	BisectionFilterer   // Log filterer for contract events
}

// BisectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type BisectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BisectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BisectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BisectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BisectionSession struct {
	Contract     *Bisection        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BisectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BisectionCallerSession struct {
	Contract *BisectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// BisectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BisectionTransactorSession struct {
	Contract     *BisectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BisectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type BisectionRaw struct {
	Contract *Bisection // Generic contract binding to access the raw methods on
}

// BisectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BisectionCallerRaw struct {
	Contract *BisectionCaller // Generic read-only contract binding to access the raw methods on
}

// BisectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BisectionTransactorRaw struct {
	Contract *BisectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBisection creates a new instance of Bisection, bound to a specific deployed contract.
func NewBisection(address common.Address, backend bind.ContractBackend) (*Bisection, error) {
	contract, err := bindBisection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bisection{BisectionCaller: BisectionCaller{contract: contract}, BisectionTransactor: BisectionTransactor{contract: contract}, BisectionFilterer: BisectionFilterer{contract: contract}}, nil
}

// NewBisectionCaller creates a new read-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionCaller(address common.Address, caller bind.ContractCaller) (*BisectionCaller, error) {
	contract, err := bindBisection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionCaller{contract: contract}, nil
}

// NewBisectionTransactor creates a new write-only instance of Bisection, bound to a specific deployed contract.
func NewBisectionTransactor(address common.Address, transactor bind.ContractTransactor) (*BisectionTransactor, error) {
	contract, err := bindBisection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BisectionTransactor{contract: contract}, nil
}

// NewBisectionFilterer creates a new log filterer instance of Bisection, bound to a specific deployed contract.
func NewBisectionFilterer(address common.Address, filterer bind.ContractFilterer) (*BisectionFilterer, error) {
	contract, err := bindBisection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BisectionFilterer{contract: contract}, nil
}

// bindBisection binds a generic wrapper to an already deployed contract.
func bindBisection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BisectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.BisectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.BisectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bisection *BisectionCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bisection.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bisection *BisectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bisection *BisectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bisection.Contract.contract.Transact(opts, method, params...)
}

// BisectionBisectedAssertionFirstIterator is returned from FilterBisectedAssertionFirst and is used to iterate over the raw logs and unpacked data for BisectedAssertionFirst events raised by the Bisection contract.
type BisectionBisectedAssertionFirstIterator struct {
	Event *BisectionBisectedAssertionFirst // Event containing the contract specifics and raw log

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
func (it *BisectionBisectedAssertionFirstIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionBisectedAssertionFirst)
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
		it.Event = new(BisectionBisectedAssertionFirst)
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
func (it *BisectionBisectedAssertionFirstIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionBisectedAssertionFirstIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionBisectedAssertionFirst represents a BisectedAssertionFirst event raised by the Bisection contract.
type BisectionBisectedAssertionFirst struct {
	VmAddress       common.Address
	Bisecter        common.Address
	BisectionHashes [][32]byte
	NumSteps        uint32
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertionFirst is a free log retrieval operation binding the contract event 0x06b770c15c61f5df9b9c6388352469b626ca74d653c33f394b93417c16bbc9e0.
//
// Solidity: event BisectedAssertionFirst(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps)
func (_Bisection *BisectionFilterer) FilterBisectedAssertionFirst(opts *bind.FilterOpts, vmAddress []common.Address) (*BisectionBisectedAssertionFirstIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "BisectedAssertionFirst", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &BisectionBisectedAssertionFirstIterator{contract: _Bisection.contract, event: "BisectedAssertionFirst", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertionFirst is a free log subscription operation binding the contract event 0x06b770c15c61f5df9b9c6388352469b626ca74d653c33f394b93417c16bbc9e0.
//
// Solidity: event BisectedAssertionFirst(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps)
func (_Bisection *BisectionFilterer) WatchBisectedAssertionFirst(opts *bind.WatchOpts, sink chan<- *BisectionBisectedAssertionFirst, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "BisectedAssertionFirst", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionBisectedAssertionFirst)
				if err := _Bisection.contract.UnpackLog(event, "BisectedAssertionFirst", log); err != nil {
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

// ParseBisectedAssertionFirst is a log parse operation binding the contract event 0x06b770c15c61f5df9b9c6388352469b626ca74d653c33f394b93417c16bbc9e0.
//
// Solidity: event BisectedAssertionFirst(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps)
func (_Bisection *BisectionFilterer) ParseBisectedAssertionFirst(log types.Log) (*BisectionBisectedAssertionFirst, error) {
	event := new(BisectionBisectedAssertionFirst)
	if err := _Bisection.contract.UnpackLog(event, "BisectedAssertionFirst", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionBisectedAssertionOtherIterator is returned from FilterBisectedAssertionOther and is used to iterate over the raw logs and unpacked data for BisectedAssertionOther events raised by the Bisection contract.
type BisectionBisectedAssertionOtherIterator struct {
	Event *BisectionBisectedAssertionOther // Event containing the contract specifics and raw log

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
func (it *BisectionBisectedAssertionOtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionBisectedAssertionOther)
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
		it.Event = new(BisectionBisectedAssertionOther)
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
func (it *BisectionBisectedAssertionOtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionBisectedAssertionOtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionBisectedAssertionOther represents a BisectedAssertionOther event raised by the Bisection contract.
type BisectionBisectedAssertionOther struct {
	VmAddress        common.Address
	Bisecter         common.Address
	BisectionHashes  [][32]byte
	NumSteps         uint32
	PrevOutputValues []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertionOther is a free log retrieval operation binding the contract event 0x4f9f2ba0780165ec676b48b1d9a57ff77e0bd1e331e18b7888da37378200335d.
//
// Solidity: event BisectedAssertionOther(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps, uint256[] prevOutputValues)
func (_Bisection *BisectionFilterer) FilterBisectedAssertionOther(opts *bind.FilterOpts, vmAddress []common.Address) (*BisectionBisectedAssertionOtherIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "BisectedAssertionOther", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &BisectionBisectedAssertionOtherIterator{contract: _Bisection.contract, event: "BisectedAssertionOther", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertionOther is a free log subscription operation binding the contract event 0x4f9f2ba0780165ec676b48b1d9a57ff77e0bd1e331e18b7888da37378200335d.
//
// Solidity: event BisectedAssertionOther(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps, uint256[] prevOutputValues)
func (_Bisection *BisectionFilterer) WatchBisectedAssertionOther(opts *bind.WatchOpts, sink chan<- *BisectionBisectedAssertionOther, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "BisectedAssertionOther", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionBisectedAssertionOther)
				if err := _Bisection.contract.UnpackLog(event, "BisectedAssertionOther", log); err != nil {
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

// ParseBisectedAssertionOther is a log parse operation binding the contract event 0x4f9f2ba0780165ec676b48b1d9a57ff77e0bd1e331e18b7888da37378200335d.
//
// Solidity: event BisectedAssertionOther(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps, uint256[] prevOutputValues)
func (_Bisection *BisectionFilterer) ParseBisectedAssertionOther(log types.Log) (*BisectionBisectedAssertionOther, error) {
	event := new(BisectionBisectedAssertionOther)
	if err := _Bisection.contract.UnpackLog(event, "BisectedAssertionOther", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BisectionContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the Bisection contract.
type BisectionContinuedChallengeIterator struct {
	Event *BisectionContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *BisectionContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BisectionContinuedChallenge)
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
		it.Event = new(BisectionContinuedChallenge)
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
func (it *BisectionContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BisectionContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BisectionContinuedChallenge represents a ContinuedChallenge event raised by the Bisection contract.
type BisectionContinuedChallenge struct {
	VmAddress      common.Address
	Challenger     common.Address
	AssertionIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) FilterContinuedChallenge(opts *bind.FilterOpts, vmAddress []common.Address) (*BisectionContinuedChallengeIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.FilterLogs(opts, "ContinuedChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &BisectionContinuedChallengeIterator{contract: _Bisection.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *BisectionContinuedChallenge, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _Bisection.contract.WatchLogs(opts, "ContinuedChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BisectionContinuedChallenge)
				if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_Bisection *BisectionFilterer) ParseContinuedChallenge(log types.Log) (*BisectionContinuedChallenge, error) {
	event := new(BisectionContinuedChallenge)
	if err := _Bisection.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723058202a4b33a7382bd41fc01a7b408d0a3491e17cc0d680148cb3a692e8a4446bc59d64736f6c634300050a0032"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// ChallengeABI is the input ABI used to generate the binding from.
const ChallengeABI = "[]"

// ChallengeBin is the compiled bytecode used for deploying new contracts.
var ChallengeBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a72305820dc9fda101a1927d9f4ef42173156e74ac08667e572c4c6882ff858166c616da864736f6c634300050a0032"

// DeployChallenge deploys a new Ethereum contract, binding an instance of Challenge to it.
func DeployChallenge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Challenge, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// Challenge is an auto generated Go binding around an Ethereum contract.
type Challenge struct {
	ChallengeCaller     // Read-only binding to the contract
	ChallengeTransactor // Write-only binding to the contract
	ChallengeFilterer   // Log filterer for contract events
}

// ChallengeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeSession struct {
	Contract     *Challenge        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeCallerSession struct {
	Contract *ChallengeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ChallengeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeTransactorSession struct {
	Contract     *ChallengeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ChallengeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeRaw struct {
	Contract *Challenge // Generic contract binding to access the raw methods on
}

// ChallengeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeCallerRaw struct {
	Contract *ChallengeCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeTransactorRaw struct {
	Contract *ChallengeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallenge creates a new instance of Challenge, bound to a specific deployed contract.
func NewChallenge(address common.Address, backend bind.ContractBackend) (*Challenge, error) {
	contract, err := bindChallenge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Challenge{ChallengeCaller: ChallengeCaller{contract: contract}, ChallengeTransactor: ChallengeTransactor{contract: contract}, ChallengeFilterer: ChallengeFilterer{contract: contract}}, nil
}

// NewChallengeCaller creates a new read-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeCaller(address common.Address, caller bind.ContractCaller) (*ChallengeCaller, error) {
	contract, err := bindChallenge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeCaller{contract: contract}, nil
}

// NewChallengeTransactor creates a new write-only instance of Challenge, bound to a specific deployed contract.
func NewChallengeTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeTransactor, error) {
	contract, err := bindChallenge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeTransactor{contract: contract}, nil
}

// NewChallengeFilterer creates a new log filterer instance of Challenge, bound to a specific deployed contract.
func NewChallengeFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeFilterer, error) {
	contract, err := bindChallenge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeFilterer{contract: contract}, nil
}

// bindChallenge binds a generic wrapper to an already deployed contract.
func bindChallenge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.ChallengeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.ChallengeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Challenge *ChallengeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Challenge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Challenge *ChallengeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Challenge *ChallengeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Challenge.Contract.contract.Transact(opts, method, params...)
}

// ChallengeManagerABI is the input ABI used to generate the binding from.
const ChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_vmAddress\",\"type\":\"address\"},{\"name\":\"_beforeHashAndInbox\",\"type\":\"bytes32[2]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"},{\"name\":\"_afterHashAndMessages\",\"type\":\"bytes32[5]\"},{\"name\":\"_amounts\",\"type\":\"uint256[]\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProofFirst\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_escrows\",\"type\":\"uint128[2]\"},{\"name\":\"_challengePeriod\",\"type\":\"uint32\"},{\"name\":\"_challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmAddress\",\"type\":\"address\"}],\"name\":\"asserterTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmAddress\",\"type\":\"address\"},{\"name\":\"_fields\",\"type\":\"bytes32[10]\"},{\"name\":\"_timeBounds\",\"type\":\"uint64[2]\"},{\"name\":\"_tokenTypes\",\"type\":\"bytes21[]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"},{\"name\":\"_a1NumSteps\",\"type\":\"uint32\"},{\"name\":\"_a1OutputValues\",\"type\":\"uint256[]\"},{\"name\":\"_a2NumSteps\",\"type\":\"uint32\"},{\"name\":\"_a2OutputValues\",\"type\":\"uint256[]\"},{\"name\":\"_proof\",\"type\":\"bytes\"}],\"name\":\"oneStepProofOther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmAddress\",\"type\":\"address\"}],\"name\":\"challengerTimedOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_challengeId\",\"type\":\"address\"},{\"name\":\"_fields\",\"type\":\"bytes32[10]\"},{\"name\":\"_beforeBalances\",\"type\":\"uint256[]\"},{\"name\":\"_a1NumSteps\",\"type\":\"uint32\"},{\"name\":\"_a1OutputValues\",\"type\":\"uint256[]\"},{\"name\":\"_a2NumSteps\",\"type\":\"uint32\"},{\"name\":\"_a2OutputValues\",\"type\":\"uint256[]\"},{\"name\":\"_bisectionHashes\",\"type\":\"bytes32[]\"}],\"name\":\"bisectAssertionOther\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_challengeId\",\"type\":\"address\"},{\"name\":\"_numSteps\",\"type\":\"uint32\"},{\"name\":\"_preconditionPart1Hash\",\"type\":\"bytes32\"},{\"name\":\"_preconditionPart2Hash\",\"type\":\"bytes32\"},{\"name\":\"_assertionHash\",\"type\":\"bytes32\"},{\"name\":\"_bisectionHashes\",\"type\":\"bytes32[]\"}],\"name\":\"bisectAssertionFirst\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_vmAddress\",\"type\":\"address\"},{\"name\":\"_assertionToChallenge\",\"type\":\"uint256\"},{\"name\":\"_proof\",\"type\":\"bytes\"},{\"name\":\"_bisectionRoot\",\"type\":\"bytes32\"},{\"name\":\"_bisectionHash\",\"type\":\"bytes32\"}],\"name\":\"continueChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"assertionIndex\",\"type\":\"uint256\"}],\"name\":\"ContinuedChallenge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"preconditionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"bisectionHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"}],\"name\":\"BisectedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bisecter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"bisectionHashes\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"name\":\"numSteps\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"prevOutputValues\",\"type\":\"uint256[]\"}],\"name\":\"BisectedAssertionOther\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"proof\",\"type\":\"bytes\"}],\"name\":\"OneStepProofCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"vmAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"challengerWrong\",\"type\":\"bool\"}],\"name\":\"TimedOutChallenge\",\"type\":\"event\"}]"

// ChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var ChallengeManagerFuncSigs = map[string]string{
	"36ddd35d": "asserterTimedOut(address)",
	"e1e373a5": "bisectAssertionFirst(address,uint32,bytes32,bytes32,bytes32,bytes32[])",
	"d1665482": "bisectAssertionOther(address,bytes32[10],uint256[],uint32,uint256[],uint32,uint256[],bytes32[])",
	"bf06db66": "challengerTimedOut(address)",
	"eb94f27b": "continueChallenge(address,uint256,bytes,bytes32,bytes32)",
	"208e04d4": "initiateChallenge(address[2],uint128[2],uint32,bytes32)",
	"083b0c65": "oneStepProofFirst(address,bytes32[2],uint64[2],bytes21[],uint256[],bytes32[5],uint256[],bytes)",
	"643d43f9": "oneStepProofOther(address,bytes32[10],uint64[2],bytes21[],uint256[],uint32,uint256[],uint32,uint256[],bytes)",
}

// ChallengeManagerBin is the compiled bytecode used for deploying new contracts.
var ChallengeManagerBin = "0x608060405234801561001057600080fd5b50611d26806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063bf06db661161005b578063bf06db66146106d6578063d1665482146106fc578063e1e373a514610982578063eb94f27b14610a5157610088565b8063083b0c651461008d578063208e04d41461034c57806336ddd35d1461037c578063643d43f9146103a2575b600080fd5b61034a60048036036101c08110156100a457600080fd5b6040805180820182526001600160a01b0384351693928301929160608301919060208401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561012457600080fd5b82018360208201111561013657600080fd5b803590602001918460208302840111600160201b8311171561015757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156101a657600080fd5b8201836020820111156101b857600080fd5b803590602001918460208302840111600160201b831117156101d957600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160a0818101909252939695948181019493509150600590839083908082843760009201919091525091949392602081019250359050600160201b81111561025457600080fd5b82018360208201111561026657600080fd5b803590602001918460208302840111600160201b8311171561028757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102d657600080fd5b8201836020820111156102e857600080fd5b803590602001918460018302840111600160201b8311171561030957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610b0f945050505050565b005b61034a600480360360c081101561036257600080fd5b506040810163ffffffff60808301351660a0830135610e30565b61034a6004803603602081101561039257600080fd5b50356001600160a01b0316610fee565b61034a60048036036102808110156103b957600080fd5b604080516101408181019092526001600160a01b03843516939283019291610160830191906020840190600a908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561043e57600080fd5b82018360208201111561045057600080fd5b803590602001918460208302840111600160201b8311171561047157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156104c057600080fd5b8201836020820111156104d257600080fd5b803590602001918460208302840111600160201b831117156104f357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b81111561055057600080fd5b82018360208201111561056257600080fd5b803590602001918460208302840111600160201b8311171561058357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b8111156105e057600080fd5b8201836020820111156105f257600080fd5b803590602001918460208302840111600160201b8311171561061357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561066257600080fd5b82018360208201111561067457600080fd5b803590602001918460018302840111600160201b8311171561069557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611107945050505050565b61034a600480360360208110156106ec57600080fd5b50356001600160a01b031661146d565b61034a600480360361022081101561071357600080fd5b604080516101408181019092526001600160a01b03843516939283019291610160830191906020840190600a90839083908082843760009201919091525091949392602081019250359050600160201b81111561076f57600080fd5b82018360208201111561078157600080fd5b803590602001918460208302840111600160201b831117156107a257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b8111156107ff57600080fd5b82018360208201111561081157600080fd5b803590602001918460208302840111600160201b8311171561083257600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b81111561088f57600080fd5b8201836020820111156108a157600080fd5b803590602001918460208302840111600160201b831117156108c257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561091157600080fd5b82018360208201111561092357600080fd5b803590602001918460208302840111600160201b8311171561094457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611586945050505050565b61034a600480360360c081101561099857600080fd5b6001600160a01b038235169163ffffffff602082013516916040820135916060810135916080820135919081019060c0810160a0820135600160201b8111156109e057600080fd5b8201836020820111156109f257600080fd5b803590602001918460208302840111600160201b83111715610a1357600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061178d945050505050565b61034a600480360360a0811015610a6757600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b811115610a9657600080fd5b820183602082011115610aa857600080fd5b803590602001918460018302840111600160201b83111715610ac957600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505082359350505060200135611897565b60008060008a6001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f55f7f918072b72dcc999cdc8e581605f5$__63c4aa83c0828a8a8a8a8a8a8a6040518963ffffffff1660e01b81526004018089815260200188600260200280838360005b83811015610b96578181015183820152602001610b7e565b5050505090500187600260200280838360005b83811015610bc1578181015183820152602001610ba9565b505050920191505060208101604082018660a080838360005b83811015610bf2578181015183820152602001610bda565b50505050905001806020018060200185810385528a818151815260200191508051906020019060200280838360005b83811015610c39578181015183820152602001610c21565b50505050905001858103845289818151815260200191508051906020019060200280838360005b83811015610c78578181015183820152602001610c60565b50505050905001858103835287818151815260200191508051906020019060200280838360005b83811015610cb7578181015183820152602001610c9f565b50505050905001858103825286818151815260200191508051906020019080838360005b83811015610cf3578181015183820152602001610cdb565b50505050905090810190601f168015610d205780820380516001836020036101000a031916815260200191505b509c5050505050505050505050505060006040518083038186803b158015610d4757600080fd5b505af4158015610d5b573d6000803e3d6000fd5b50505050610d68816119b0565b886001600160a01b03167ffd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572338460405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610dea578181015183820152602001610dd2565b50505050905090810190601f168015610e175780820380516001836020036101000a031916815260200191505b50935050505060405180910390a2505050505050505050565b3360009081526020819052604090206001015415610e7f5760405162461bcd60e51b8152600401808060200182810382526023815260200180611ccf6023913960400191505060405180910390fd5b6040805160e081018252338152602081018390528151808301835290918281019190869060029083908390808284376000920191909152505050815260408051808201825260209092019190879060029083908390808284376000920191909152505050815263ffffffff841643810167ffffffffffffffff1660208301526040820152606001600190523360009081526020818152604091829020835181546001600160a01b0319166001600160a01b0390911617815590830151600182015590820151610f549060028084019190611b23565b506060820151610f6a9060038301906002611bc8565b50608082015160058201805460a085015163ffffffff1668010000000000000000026bffffffff00000000000000001967ffffffffffffffff90941667ffffffffffffffff1990921691909117929092169190911780825560c0840151919060ff60601b1916600160601b836002811115610fe157fe5b0217905550505050505050565b6001600160a01b038116600090815260208190526040902060016005820154600160601b900460ff16600281111561102257fe5b1461105e5760405162461bcd60e51b815260040180806020018281038252602e815260200180611c71602e913960400191505060405180910390fd5b600581015467ffffffffffffffff1643116110ba576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b6110c381611a94565b604080516001815290516001600160a01b038416917f2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35919081900360200190a25050565b60008060008c6001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f55f7f918072b72dcc999cdc8e581605f5$__63ab06a992828c8c8c8c8c8c8c8c8c6040518b63ffffffff1660e01b8152600401808b81526020018a600a60200280838360005b83811015611190578181015183820152602001611178565b5050505090500189600260200280838360005b838110156111bb5781810151838201526020016111a3565b5050505090500180602001806020018863ffffffff1663ffffffff168152602001806020018763ffffffff1663ffffffff168152602001806020018060200186810386528d818151815260200191508051906020019060200280838360005b8381101561123257818101518382015260200161121a565b5050505090500186810385528c818151815260200191508051906020019060200280838360005b83811015611271578181015183820152602001611259565b5050505090500186810384528a818151815260200191508051906020019060200280838360005b838110156112b0578181015183820152602001611298565b50505050905001868103835288818151815260200191508051906020019060200280838360005b838110156112ef5781810151838201526020016112d7565b50505050905001868103825287818151815260200191508051906020019080838360005b8381101561132b578181015183820152602001611313565b50505050905090810190601f1680156113585780820380516001836020036101000a031916815260200191505b509f5050505050505050505050505050505060006040518083038186803b15801561138257600080fd5b505af4158015611396573d6000803e3d6000fd5b505050506113a3816119b0565b8a6001600160a01b03167ffd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572338460405180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561142557818101518382015260200161140d565b50505050905090810190601f1680156114525780820380516001836020036101000a031916815260200191505b50935050505060405180910390a25050505050505050505050565b6001600160a01b038116600090815260208190526040902060026005820154600160601b900460ff1660028111156114a157fe5b146114dd5760405162461bcd60e51b8152600401808060200182810382526030815260200180611c9f6030913960400191505060405180910390fd5b600581015467ffffffffffffffff164311611539576040805162461bcd60e51b8152602060048201526017602482015276111958591b1a5b99481a185cdb89dd08195e1c1a5c9959604a1b604482015290519081900360640190fd5b611542816119b0565b604080516000815290516001600160a01b038416917f2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35919081900360200190a25050565b60008060008a6001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f5eea941ded5358daea4da7ea13a2128fd$__638014c3e9828a8a8a8a8a8a8a6040518963ffffffff1660e01b81526004018089815260200188600a60200280838360005b8381101561160d5781810151838201526020016115f5565b50505050905001806020018763ffffffff1663ffffffff168152602001806020018663ffffffff1663ffffffff168152602001806020018060200185810385528b818151815260200191508051906020019060200280838360005b83811015611680578181015183820152602001611668565b50505050905001858103845289818151815260200191508051906020019060200280838360005b838110156116bf5781810151838201526020016116a7565b50505050905001858103835287818151815260200191508051906020019060200280838360005b838110156116fe5781810151838201526020016116e6565b50505050905001858103825286818151815260200191508051906020019060200280838360005b8381101561173d578181015183820152602001611725565b505050509050019c5050505050505050505050505060006040518083038186803b15801561176a57600080fd5b505af415801561177e573d6000803e3d6000fd5b50505050505050505050505050565b6000806000886001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f5eea941ded5358daea4da7ea13a2128fd$__63353a58648288888888886040518763ffffffff1660e01b8152600401808781526020018663ffffffff1663ffffffff16815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b8381101561184e578181015183820152602001611836565b5050505090500197505050505050505060006040518083038186803b15801561187657600080fd5b505af415801561188a573d6000803e3d6000fd5b5050505050505050505050565b6000806000876001600160a01b03166001600160a01b03168152602001908152602001600020905073__$f5eea941ded5358daea4da7ea13a2128fd$__63110112ae82878787876040518663ffffffff1660e01b81526004018086815260200185815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561194257818101518382015260200161192a565b50505050905090810190601f16801561196f5780820380516001836020036101000a031916815260200191505b50965050505050505060006040518083038186803b15801561199057600080fd5b505af41580156119a4573d6000803e3d6000fd5b50505050505050505050565b80546040805180820182526002808501546001600160801b03808216600160801b909204811692909204011681526000602082015290516308b0246f60e21b81526001600160a01b03909216916322c091bc916003850191600481019060440183825b81546001600160a01b03168152600190910190602001808311611a135750839050604080838360005b83811015611a54578181015183820152602001611a3c565b5050505090500192505050600060405180830381600087803b158015611a7957600080fd5b505af1158015611a8d573d6000803e3d6000fd5b5050505050565b805460408051808201825260008152600280850154600160801b81046001600160801b03908116918116929092040116602082015290516308b0246f60e21b81526003840180546001600160a01b03908116600480850191825291909516946322c091bc9492939091604482019190880190602401808311611a13575050825181528260408083836020611a3c565b600183019183908215611bb85791602002820160005b83821115611b8357835183826101000a8154816001600160801b0302191690836001600160801b031602179055509260200192601001602081600f01049283019260010302611b39565b8015611bb65782816101000a8154906001600160801b030219169055601001602081600f01049283019260010302611b83565b505b50611bc4929150611c1c565b5090565b8260028101928215611c10579160200282015b82811115611c1057825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611bdb565b50611bc4929150611c4c565b611c4991905b80821115611bc45780546fffffffffffffffffffffffffffffffff19168155600101611c22565b90565b611c4991905b80821115611bc45780546001600160a01b0319168155600101611c5256fe43616e206f6e6c792074696d65206f7574206173736572746572206966206974206973207468656972207475726e43616e206f6e6c792074696d65206f7574206368616c6c656e676572206966206974206973207468656972207475726e5468657265206d757374206265206e6f206578697374696e67206368616c6c656e6765a265627a7a7230582052bd2640d87a5d0f2a26e880c19fd17a25c5a83665f57da3dafe4ded55bec83d64736f6c634300050a0032"

// DeployChallengeManager deploys a new Ethereum contract, binding an instance of ChallengeManager to it.
func DeployChallengeManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ChallengeManager, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	oneStepProofAddr, _, _, _ := DeployOneStepProof(auth, backend)
	ChallengeManagerBin = strings.Replace(ChallengeManagerBin, "__$f55f7f918072b72dcc999cdc8e581605f5$__", oneStepProofAddr.String()[2:], -1)

	bisectionAddr, _, _, _ := DeployBisection(auth, backend)
	ChallengeManagerBin = strings.Replace(ChallengeManagerBin, "__$f5eea941ded5358daea4da7ea13a2128fd$__", bisectionAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ChallengeManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// ChallengeManager is an auto generated Go binding around an Ethereum contract.
type ChallengeManager struct {
	ChallengeManagerCaller     // Read-only binding to the contract
	ChallengeManagerTransactor // Write-only binding to the contract
	ChallengeManagerFilterer   // Log filterer for contract events
}

// ChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeManagerSession struct {
	Contract     *ChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeManagerCallerSession struct {
	Contract *ChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeManagerTransactorSession struct {
	Contract     *ChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeManagerRaw struct {
	Contract *ChallengeManager // Generic contract binding to access the raw methods on
}

// ChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeManagerCallerRaw struct {
	Contract *ChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeManagerTransactorRaw struct {
	Contract *ChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeManager creates a new instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManager(address common.Address, backend bind.ContractBackend) (*ChallengeManager, error) {
	contract, err := bindChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeManager{ChallengeManagerCaller: ChallengeManagerCaller{contract: contract}, ChallengeManagerTransactor: ChallengeManagerTransactor{contract: contract}, ChallengeManagerFilterer: ChallengeManagerFilterer{contract: contract}}, nil
}

// NewChallengeManagerCaller creates a new read-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*ChallengeManagerCaller, error) {
	contract, err := bindChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerCaller{contract: contract}, nil
}

// NewChallengeManagerTransactor creates a new write-only instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeManagerTransactor, error) {
	contract, err := bindChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTransactor{contract: contract}, nil
}

// NewChallengeManagerFilterer creates a new log filterer instance of ChallengeManager, bound to a specific deployed contract.
func NewChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeManagerFilterer, error) {
	contract, err := bindChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerFilterer{contract: contract}, nil
}

// bindChallengeManager binds a generic wrapper to an already deployed contract.
func bindChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.ChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeManager *ChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeManager *ChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x36ddd35d.
//
// Solidity: function asserterTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactor) AsserterTimedOut(opts *bind.TransactOpts, _vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "asserterTimedOut", _vmAddress)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x36ddd35d.
//
// Solidity: function asserterTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerSession) AsserterTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// AsserterTimedOut is a paid mutator transaction binding the contract method 0x36ddd35d.
//
// Solidity: function asserterTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) AsserterTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.AsserterTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// BisectAssertionFirst is a paid mutator transaction binding the contract method 0xe1e373a5.
//
// Solidity: function bisectAssertionFirst(address _challengeId, uint32 _numSteps, bytes32 _preconditionPart1Hash, bytes32 _preconditionPart2Hash, bytes32 _assertionHash, bytes32[] _bisectionHashes) returns()
func (_ChallengeManager *ChallengeManagerTransactor) BisectAssertionFirst(opts *bind.TransactOpts, _challengeId common.Address, _numSteps uint32, _preconditionPart1Hash [32]byte, _preconditionPart2Hash [32]byte, _assertionHash [32]byte, _bisectionHashes [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "bisectAssertionFirst", _challengeId, _numSteps, _preconditionPart1Hash, _preconditionPart2Hash, _assertionHash, _bisectionHashes)
}

// BisectAssertionFirst is a paid mutator transaction binding the contract method 0xe1e373a5.
//
// Solidity: function bisectAssertionFirst(address _challengeId, uint32 _numSteps, bytes32 _preconditionPart1Hash, bytes32 _preconditionPart2Hash, bytes32 _assertionHash, bytes32[] _bisectionHashes) returns()
func (_ChallengeManager *ChallengeManagerSession) BisectAssertionFirst(_challengeId common.Address, _numSteps uint32, _preconditionPart1Hash [32]byte, _preconditionPart2Hash [32]byte, _assertionHash [32]byte, _bisectionHashes [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertionFirst(&_ChallengeManager.TransactOpts, _challengeId, _numSteps, _preconditionPart1Hash, _preconditionPart2Hash, _assertionHash, _bisectionHashes)
}

// BisectAssertionFirst is a paid mutator transaction binding the contract method 0xe1e373a5.
//
// Solidity: function bisectAssertionFirst(address _challengeId, uint32 _numSteps, bytes32 _preconditionPart1Hash, bytes32 _preconditionPart2Hash, bytes32 _assertionHash, bytes32[] _bisectionHashes) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) BisectAssertionFirst(_challengeId common.Address, _numSteps uint32, _preconditionPart1Hash [32]byte, _preconditionPart2Hash [32]byte, _assertionHash [32]byte, _bisectionHashes [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertionFirst(&_ChallengeManager.TransactOpts, _challengeId, _numSteps, _preconditionPart1Hash, _preconditionPart2Hash, _assertionHash, _bisectionHashes)
}

// BisectAssertionOther is a paid mutator transaction binding the contract method 0xd1665482.
//
// Solidity: function bisectAssertionOther(address _challengeId, bytes32[10] _fields, uint256[] _beforeBalances, uint32 _a1NumSteps, uint256[] _a1OutputValues, uint32 _a2NumSteps, uint256[] _a2OutputValues, bytes32[] _bisectionHashes) returns()
func (_ChallengeManager *ChallengeManagerTransactor) BisectAssertionOther(opts *bind.TransactOpts, _challengeId common.Address, _fields [10][32]byte, _beforeBalances []*big.Int, _a1NumSteps uint32, _a1OutputValues []*big.Int, _a2NumSteps uint32, _a2OutputValues []*big.Int, _bisectionHashes [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "bisectAssertionOther", _challengeId, _fields, _beforeBalances, _a1NumSteps, _a1OutputValues, _a2NumSteps, _a2OutputValues, _bisectionHashes)
}

// BisectAssertionOther is a paid mutator transaction binding the contract method 0xd1665482.
//
// Solidity: function bisectAssertionOther(address _challengeId, bytes32[10] _fields, uint256[] _beforeBalances, uint32 _a1NumSteps, uint256[] _a1OutputValues, uint32 _a2NumSteps, uint256[] _a2OutputValues, bytes32[] _bisectionHashes) returns()
func (_ChallengeManager *ChallengeManagerSession) BisectAssertionOther(_challengeId common.Address, _fields [10][32]byte, _beforeBalances []*big.Int, _a1NumSteps uint32, _a1OutputValues []*big.Int, _a2NumSteps uint32, _a2OutputValues []*big.Int, _bisectionHashes [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertionOther(&_ChallengeManager.TransactOpts, _challengeId, _fields, _beforeBalances, _a1NumSteps, _a1OutputValues, _a2NumSteps, _a2OutputValues, _bisectionHashes)
}

// BisectAssertionOther is a paid mutator transaction binding the contract method 0xd1665482.
//
// Solidity: function bisectAssertionOther(address _challengeId, bytes32[10] _fields, uint256[] _beforeBalances, uint32 _a1NumSteps, uint256[] _a1OutputValues, uint32 _a2NumSteps, uint256[] _a2OutputValues, bytes32[] _bisectionHashes) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) BisectAssertionOther(_challengeId common.Address, _fields [10][32]byte, _beforeBalances []*big.Int, _a1NumSteps uint32, _a1OutputValues []*big.Int, _a2NumSteps uint32, _a2OutputValues []*big.Int, _bisectionHashes [][32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.BisectAssertionOther(&_ChallengeManager.TransactOpts, _challengeId, _fields, _beforeBalances, _a1NumSteps, _a1OutputValues, _a2NumSteps, _a2OutputValues, _bisectionHashes)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xbf06db66.
//
// Solidity: function challengerTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ChallengerTimedOut(opts *bind.TransactOpts, _vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "challengerTimedOut", _vmAddress)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xbf06db66.
//
// Solidity: function challengerTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerSession) ChallengerTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// ChallengerTimedOut is a paid mutator transaction binding the contract method 0xbf06db66.
//
// Solidity: function challengerTimedOut(address _vmAddress) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ChallengerTimedOut(_vmAddress common.Address) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ChallengerTimedOut(&_ChallengeManager.TransactOpts, _vmAddress)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xeb94f27b.
//
// Solidity: function continueChallenge(address _vmAddress, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactor) ContinueChallenge(opts *bind.TransactOpts, _vmAddress common.Address, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "continueChallenge", _vmAddress, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xeb94f27b.
//
// Solidity: function continueChallenge(address _vmAddress, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerSession) ContinueChallenge(_vmAddress common.Address, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmAddress, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// ContinueChallenge is a paid mutator transaction binding the contract method 0xeb94f27b.
//
// Solidity: function continueChallenge(address _vmAddress, uint256 _assertionToChallenge, bytes _proof, bytes32 _bisectionRoot, bytes32 _bisectionHash) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) ContinueChallenge(_vmAddress common.Address, _assertionToChallenge *big.Int, _proof []byte, _bisectionRoot [32]byte, _bisectionHash [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.ContinueChallenge(&_ChallengeManager.TransactOpts, _vmAddress, _assertionToChallenge, _proof, _bisectionRoot, _bisectionHash)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, _players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "initiateChallenge", _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerSession) InitiateChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _players, _escrows, _challengePeriod, _challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] _players, uint128[2] _escrows, uint32 _challengePeriod, bytes32 _challengeRoot) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) InitiateChallenge(_players [2]common.Address, _escrows [2]*big.Int, _challengePeriod uint32, _challengeRoot [32]byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.InitiateChallenge(&_ChallengeManager.TransactOpts, _players, _escrows, _challengePeriod, _challengeRoot)
}

// OneStepProofFirst is a paid mutator transaction binding the contract method 0x083b0c65.
//
// Solidity: function oneStepProofFirst(address _vmAddress, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactor) OneStepProofFirst(opts *bind.TransactOpts, _vmAddress common.Address, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "oneStepProofFirst", _vmAddress, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof)
}

// OneStepProofFirst is a paid mutator transaction binding the contract method 0x083b0c65.
//
// Solidity: function oneStepProofFirst(address _vmAddress, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerSession) OneStepProofFirst(_vmAddress common.Address, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProofFirst(&_ChallengeManager.TransactOpts, _vmAddress, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof)
}

// OneStepProofFirst is a paid mutator transaction binding the contract method 0x083b0c65.
//
// Solidity: function oneStepProofFirst(address _vmAddress, bytes32[2] _beforeHashAndInbox, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, bytes32[5] _afterHashAndMessages, uint256[] _amounts, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) OneStepProofFirst(_vmAddress common.Address, _beforeHashAndInbox [2][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _afterHashAndMessages [5][32]byte, _amounts []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProofFirst(&_ChallengeManager.TransactOpts, _vmAddress, _beforeHashAndInbox, _timeBounds, _tokenTypes, _beforeBalances, _afterHashAndMessages, _amounts, _proof)
}

// OneStepProofOther is a paid mutator transaction binding the contract method 0x643d43f9.
//
// Solidity: function oneStepProofOther(address _vmAddress, bytes32[10] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, uint32 _a1NumSteps, uint256[] _a1OutputValues, uint32 _a2NumSteps, uint256[] _a2OutputValues, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactor) OneStepProofOther(opts *bind.TransactOpts, _vmAddress common.Address, _fields [10][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _a1NumSteps uint32, _a1OutputValues []*big.Int, _a2NumSteps uint32, _a2OutputValues []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.contract.Transact(opts, "oneStepProofOther", _vmAddress, _fields, _timeBounds, _tokenTypes, _beforeBalances, _a1NumSteps, _a1OutputValues, _a2NumSteps, _a2OutputValues, _proof)
}

// OneStepProofOther is a paid mutator transaction binding the contract method 0x643d43f9.
//
// Solidity: function oneStepProofOther(address _vmAddress, bytes32[10] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, uint32 _a1NumSteps, uint256[] _a1OutputValues, uint32 _a2NumSteps, uint256[] _a2OutputValues, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerSession) OneStepProofOther(_vmAddress common.Address, _fields [10][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _a1NumSteps uint32, _a1OutputValues []*big.Int, _a2NumSteps uint32, _a2OutputValues []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProofOther(&_ChallengeManager.TransactOpts, _vmAddress, _fields, _timeBounds, _tokenTypes, _beforeBalances, _a1NumSteps, _a1OutputValues, _a2NumSteps, _a2OutputValues, _proof)
}

// OneStepProofOther is a paid mutator transaction binding the contract method 0x643d43f9.
//
// Solidity: function oneStepProofOther(address _vmAddress, bytes32[10] _fields, uint64[2] _timeBounds, bytes21[] _tokenTypes, uint256[] _beforeBalances, uint32 _a1NumSteps, uint256[] _a1OutputValues, uint32 _a2NumSteps, uint256[] _a2OutputValues, bytes _proof) returns()
func (_ChallengeManager *ChallengeManagerTransactorSession) OneStepProofOther(_vmAddress common.Address, _fields [10][32]byte, _timeBounds [2]uint64, _tokenTypes [][21]byte, _beforeBalances []*big.Int, _a1NumSteps uint32, _a1OutputValues []*big.Int, _a2NumSteps uint32, _a2OutputValues []*big.Int, _proof []byte) (*types.Transaction, error) {
	return _ChallengeManager.Contract.OneStepProofOther(&_ChallengeManager.TransactOpts, _vmAddress, _fields, _timeBounds, _tokenTypes, _beforeBalances, _a1NumSteps, _a1OutputValues, _a2NumSteps, _a2OutputValues, _proof)
}

// ChallengeManagerBisectedAssertionIterator is returned from FilterBisectedAssertion and is used to iterate over the raw logs and unpacked data for BisectedAssertion events raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertionIterator struct {
	Event *ChallengeManagerBisectedAssertion // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerBisectedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerBisectedAssertion)
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
		it.Event = new(ChallengeManagerBisectedAssertion)
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
func (it *ChallengeManagerBisectedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerBisectedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerBisectedAssertion represents a BisectedAssertion event raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertion struct {
	VmAddress        common.Address
	Bisecter         common.Address
	PreconditionHash [32]byte
	BisectionHashes  [][32]byte
	NumSteps         uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertion is a free log retrieval operation binding the contract event 0xeb4d283214041cc36633c89fe29a2d72879a3e61097aafe62b80e25dbac82591.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32 preconditionHash, bytes32[] bisectionHashes, uint32 numSteps)
func (_ChallengeManager *ChallengeManagerFilterer) FilterBisectedAssertion(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerBisectedAssertionIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "BisectedAssertion", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerBisectedAssertionIterator{contract: _ChallengeManager.contract, event: "BisectedAssertion", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertion is a free log subscription operation binding the contract event 0xeb4d283214041cc36633c89fe29a2d72879a3e61097aafe62b80e25dbac82591.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32 preconditionHash, bytes32[] bisectionHashes, uint32 numSteps)
func (_ChallengeManager *ChallengeManagerFilterer) WatchBisectedAssertion(opts *bind.WatchOpts, sink chan<- *ChallengeManagerBisectedAssertion, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "BisectedAssertion", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerBisectedAssertion)
				if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
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

// ParseBisectedAssertion is a log parse operation binding the contract event 0xeb4d283214041cc36633c89fe29a2d72879a3e61097aafe62b80e25dbac82591.
//
// Solidity: event BisectedAssertion(address indexed vmAddress, address bisecter, bytes32 preconditionHash, bytes32[] bisectionHashes, uint32 numSteps)
func (_ChallengeManager *ChallengeManagerFilterer) ParseBisectedAssertion(log types.Log) (*ChallengeManagerBisectedAssertion, error) {
	event := new(ChallengeManagerBisectedAssertion)
	if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertion", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerBisectedAssertionOtherIterator is returned from FilterBisectedAssertionOther and is used to iterate over the raw logs and unpacked data for BisectedAssertionOther events raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertionOtherIterator struct {
	Event *ChallengeManagerBisectedAssertionOther // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerBisectedAssertionOtherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerBisectedAssertionOther)
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
		it.Event = new(ChallengeManagerBisectedAssertionOther)
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
func (it *ChallengeManagerBisectedAssertionOtherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerBisectedAssertionOtherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerBisectedAssertionOther represents a BisectedAssertionOther event raised by the ChallengeManager contract.
type ChallengeManagerBisectedAssertionOther struct {
	VmAddress        common.Address
	Bisecter         common.Address
	BisectionHashes  [][32]byte
	NumSteps         uint32
	PrevOutputValues []*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterBisectedAssertionOther is a free log retrieval operation binding the contract event 0x4f9f2ba0780165ec676b48b1d9a57ff77e0bd1e331e18b7888da37378200335d.
//
// Solidity: event BisectedAssertionOther(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps, uint256[] prevOutputValues)
func (_ChallengeManager *ChallengeManagerFilterer) FilterBisectedAssertionOther(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerBisectedAssertionOtherIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "BisectedAssertionOther", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerBisectedAssertionOtherIterator{contract: _ChallengeManager.contract, event: "BisectedAssertionOther", logs: logs, sub: sub}, nil
}

// WatchBisectedAssertionOther is a free log subscription operation binding the contract event 0x4f9f2ba0780165ec676b48b1d9a57ff77e0bd1e331e18b7888da37378200335d.
//
// Solidity: event BisectedAssertionOther(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps, uint256[] prevOutputValues)
func (_ChallengeManager *ChallengeManagerFilterer) WatchBisectedAssertionOther(opts *bind.WatchOpts, sink chan<- *ChallengeManagerBisectedAssertionOther, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "BisectedAssertionOther", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerBisectedAssertionOther)
				if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertionOther", log); err != nil {
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

// ParseBisectedAssertionOther is a log parse operation binding the contract event 0x4f9f2ba0780165ec676b48b1d9a57ff77e0bd1e331e18b7888da37378200335d.
//
// Solidity: event BisectedAssertionOther(address indexed vmAddress, address bisecter, bytes32[] bisectionHashes, uint32 numSteps, uint256[] prevOutputValues)
func (_ChallengeManager *ChallengeManagerFilterer) ParseBisectedAssertionOther(log types.Log) (*ChallengeManagerBisectedAssertionOther, error) {
	event := new(ChallengeManagerBisectedAssertionOther)
	if err := _ChallengeManager.contract.UnpackLog(event, "BisectedAssertionOther", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerContinuedChallengeIterator is returned from FilterContinuedChallenge and is used to iterate over the raw logs and unpacked data for ContinuedChallenge events raised by the ChallengeManager contract.
type ChallengeManagerContinuedChallengeIterator struct {
	Event *ChallengeManagerContinuedChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerContinuedChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerContinuedChallenge)
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
		it.Event = new(ChallengeManagerContinuedChallenge)
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
func (it *ChallengeManagerContinuedChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerContinuedChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerContinuedChallenge represents a ContinuedChallenge event raised by the ChallengeManager contract.
type ChallengeManagerContinuedChallenge struct {
	VmAddress      common.Address
	Challenger     common.Address
	AssertionIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterContinuedChallenge is a free log retrieval operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) FilterContinuedChallenge(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerContinuedChallengeIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "ContinuedChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerContinuedChallengeIterator{contract: _ChallengeManager.contract, event: "ContinuedChallenge", logs: logs, sub: sub}, nil
}

// WatchContinuedChallenge is a free log subscription operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) WatchContinuedChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerContinuedChallenge, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "ContinuedChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerContinuedChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
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

// ParseContinuedChallenge is a log parse operation binding the contract event 0x18bc06caad44fc005e4d92df184d8be472b3cba8a9df26c176a269393ed8fd00.
//
// Solidity: event ContinuedChallenge(address indexed vmAddress, address challenger, uint256 assertionIndex)
func (_ChallengeManager *ChallengeManagerFilterer) ParseContinuedChallenge(log types.Log) (*ChallengeManagerContinuedChallenge, error) {
	event := new(ChallengeManagerContinuedChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "ContinuedChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerOneStepProofCompletedIterator is returned from FilterOneStepProofCompleted and is used to iterate over the raw logs and unpacked data for OneStepProofCompleted events raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompletedIterator struct {
	Event *ChallengeManagerOneStepProofCompleted // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerOneStepProofCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerOneStepProofCompleted)
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
		it.Event = new(ChallengeManagerOneStepProofCompleted)
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
func (it *ChallengeManagerOneStepProofCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerOneStepProofCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerOneStepProofCompleted represents a OneStepProofCompleted event raised by the ChallengeManager contract.
type ChallengeManagerOneStepProofCompleted struct {
	VmAddress common.Address
	Asserter  common.Address
	Proof     []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOneStepProofCompleted is a free log retrieval operation binding the contract event 0xfd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572.
//
// Solidity: event OneStepProofCompleted(address indexed vmAddress, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) FilterOneStepProofCompleted(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerOneStepProofCompletedIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "OneStepProofCompleted", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerOneStepProofCompletedIterator{contract: _ChallengeManager.contract, event: "OneStepProofCompleted", logs: logs, sub: sub}, nil
}

// WatchOneStepProofCompleted is a free log subscription operation binding the contract event 0xfd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572.
//
// Solidity: event OneStepProofCompleted(address indexed vmAddress, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) WatchOneStepProofCompleted(opts *bind.WatchOpts, sink chan<- *ChallengeManagerOneStepProofCompleted, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "OneStepProofCompleted", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerOneStepProofCompleted)
				if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
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

// ParseOneStepProofCompleted is a log parse operation binding the contract event 0xfd6b3dfb79b0eff8bc6cb0b3b7b957066eeeb28dff458cc42ca1a9a1205e4572.
//
// Solidity: event OneStepProofCompleted(address indexed vmAddress, address asserter, bytes proof)
func (_ChallengeManager *ChallengeManagerFilterer) ParseOneStepProofCompleted(log types.Log) (*ChallengeManagerOneStepProofCompleted, error) {
	event := new(ChallengeManagerOneStepProofCompleted)
	if err := _ChallengeManager.contract.UnpackLog(event, "OneStepProofCompleted", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ChallengeManagerTimedOutChallengeIterator is returned from FilterTimedOutChallenge and is used to iterate over the raw logs and unpacked data for TimedOutChallenge events raised by the ChallengeManager contract.
type ChallengeManagerTimedOutChallengeIterator struct {
	Event *ChallengeManagerTimedOutChallenge // Event containing the contract specifics and raw log

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
func (it *ChallengeManagerTimedOutChallengeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeManagerTimedOutChallenge)
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
		it.Event = new(ChallengeManagerTimedOutChallenge)
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
func (it *ChallengeManagerTimedOutChallengeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeManagerTimedOutChallengeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeManagerTimedOutChallenge represents a TimedOutChallenge event raised by the ChallengeManager contract.
type ChallengeManagerTimedOutChallenge struct {
	VmAddress       common.Address
	ChallengerWrong bool
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterTimedOutChallenge is a free log retrieval operation binding the contract event 0x2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35.
//
// Solidity: event TimedOutChallenge(address indexed vmAddress, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) FilterTimedOutChallenge(opts *bind.FilterOpts, vmAddress []common.Address) (*ChallengeManagerTimedOutChallengeIterator, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.FilterLogs(opts, "TimedOutChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeManagerTimedOutChallengeIterator{contract: _ChallengeManager.contract, event: "TimedOutChallenge", logs: logs, sub: sub}, nil
}

// WatchTimedOutChallenge is a free log subscription operation binding the contract event 0x2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35.
//
// Solidity: event TimedOutChallenge(address indexed vmAddress, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) WatchTimedOutChallenge(opts *bind.WatchOpts, sink chan<- *ChallengeManagerTimedOutChallenge, vmAddress []common.Address) (event.Subscription, error) {

	var vmAddressRule []interface{}
	for _, vmAddressItem := range vmAddress {
		vmAddressRule = append(vmAddressRule, vmAddressItem)
	}

	logs, sub, err := _ChallengeManager.contract.WatchLogs(opts, "TimedOutChallenge", vmAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeManagerTimedOutChallenge)
				if err := _ChallengeManager.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
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

// ParseTimedOutChallenge is a log parse operation binding the contract event 0x2b79ef590eb1a8f7c1a7551b57e7c229503020a6ebd3a18ad3bf8563a0d5cb35.
//
// Solidity: event TimedOutChallenge(address indexed vmAddress, bool challengerWrong)
func (_ChallengeManager *ChallengeManagerFilterer) ParseTimedOutChallenge(log types.Log) (*ChallengeManagerTimedOutChallenge, error) {
	event := new(ChallengeManagerTimedOutChallenge)
	if err := _ChallengeManager.contract.UnpackLog(event, "TimedOutChallenge", log); err != nil {
		return nil, err
	}
	return event, nil
}

// DebugPrintABI is the input ABI used to generate the binding from.
const DebugPrintABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"b32\",\"type\":\"bytes32\"}],\"name\":\"bytes32string\",\"outputs\":[{\"name\":\"out\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// DebugPrintFuncSigs maps the 4-byte function signature to its string representation.
var DebugPrintFuncSigs = map[string]string{
	"252fb38d": "bytes32string(bytes32)",
}

// DebugPrintBin is the compiled bytecode used for deploying new contracts.
var DebugPrintBin = "0x610202610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100355760003560e01c8063252fb38d1461003a575b600080fd5b6100576004803603602081101561005057600080fd5b50356100cc565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610091578181015183820152602001610079565b50505050905090810190601f1680156100be5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60408051818152606081810183529182919060208201818038833901905050905060005b602081101561019457600084826020811061010757fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b61012d8261019b565b85856002028151811061013c57fe5b60200101906001600160f81b031916908160001a90535061015c8161019b565b85856002026001018151811061016e57fe5b60200101906001600160f81b031916908160001a90535050600190920191506100f09050565b5092915050565b6000600a60f883901c10156101bb578160f81c60300160f81b90506101c8565b8160f81c60570160f81b90505b91905056fea265627a7a7230582042bb954870bc145926365d8ed4ccaa3d4de4ff516a5ff93e2d9a62f98eef530664736f6c634300050a0032"

// DeployDebugPrint deploys a new Ethereum contract, binding an instance of DebugPrint to it.
func DeployDebugPrint(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DebugPrint, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DebugPrintBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// DebugPrint is an auto generated Go binding around an Ethereum contract.
type DebugPrint struct {
	DebugPrintCaller     // Read-only binding to the contract
	DebugPrintTransactor // Write-only binding to the contract
	DebugPrintFilterer   // Log filterer for contract events
}

// DebugPrintCaller is an auto generated read-only Go binding around an Ethereum contract.
type DebugPrintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DebugPrintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DebugPrintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DebugPrintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DebugPrintSession struct {
	Contract     *DebugPrint       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DebugPrintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DebugPrintCallerSession struct {
	Contract *DebugPrintCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DebugPrintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DebugPrintTransactorSession struct {
	Contract     *DebugPrintTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DebugPrintRaw is an auto generated low-level Go binding around an Ethereum contract.
type DebugPrintRaw struct {
	Contract *DebugPrint // Generic contract binding to access the raw methods on
}

// DebugPrintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DebugPrintCallerRaw struct {
	Contract *DebugPrintCaller // Generic read-only contract binding to access the raw methods on
}

// DebugPrintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DebugPrintTransactorRaw struct {
	Contract *DebugPrintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDebugPrint creates a new instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrint(address common.Address, backend bind.ContractBackend) (*DebugPrint, error) {
	contract, err := bindDebugPrint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DebugPrint{DebugPrintCaller: DebugPrintCaller{contract: contract}, DebugPrintTransactor: DebugPrintTransactor{contract: contract}, DebugPrintFilterer: DebugPrintFilterer{contract: contract}}, nil
}

// NewDebugPrintCaller creates a new read-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintCaller(address common.Address, caller bind.ContractCaller) (*DebugPrintCaller, error) {
	contract, err := bindDebugPrint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintCaller{contract: contract}, nil
}

// NewDebugPrintTransactor creates a new write-only instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintTransactor(address common.Address, transactor bind.ContractTransactor) (*DebugPrintTransactor, error) {
	contract, err := bindDebugPrint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DebugPrintTransactor{contract: contract}, nil
}

// NewDebugPrintFilterer creates a new log filterer instance of DebugPrint, bound to a specific deployed contract.
func NewDebugPrintFilterer(address common.Address, filterer bind.ContractFilterer) (*DebugPrintFilterer, error) {
	contract, err := bindDebugPrint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DebugPrintFilterer{contract: contract}, nil
}

// bindDebugPrint binds a generic wrapper to an already deployed contract.
func bindDebugPrint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DebugPrintABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.DebugPrintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.DebugPrintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DebugPrint *DebugPrintCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DebugPrint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DebugPrint *DebugPrintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DebugPrint *DebugPrintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DebugPrint.Contract.contract.Transact(opts, method, params...)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCaller) Bytes32string(opts *bind.CallOpts, b32 [32]byte) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _DebugPrint.contract.Call(opts, out, "bytes32string", b32)
	return *ret0, err
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// Bytes32string is a free data retrieval call binding the contract method 0x252fb38d.
//
// Solidity: function bytes32string(bytes32 b32) constant returns(string out)
func (_DebugPrint *DebugPrintCallerSession) Bytes32string(b32 [32]byte) (string, error) {
	return _DebugPrint.Contract.Bytes32string(&_DebugPrint.CallOpts, b32)
}

// IChallengeManagerABI is the input ABI used to generate the binding from.
const IChallengeManagerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"players\",\"type\":\"address[2]\"},{\"name\":\"escrows\",\"type\":\"uint128[2]\"},{\"name\":\"challengePeriod\",\"type\":\"uint32\"},{\"name\":\"challengeRoot\",\"type\":\"bytes32\"}],\"name\":\"initiateChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IChallengeManagerFuncSigs maps the 4-byte function signature to its string representation.
var IChallengeManagerFuncSigs = map[string]string{
	"208e04d4": "initiateChallenge(address[2],uint128[2],uint32,bytes32)",
}

// IChallengeManager is an auto generated Go binding around an Ethereum contract.
type IChallengeManager struct {
	IChallengeManagerCaller     // Read-only binding to the contract
	IChallengeManagerTransactor // Write-only binding to the contract
	IChallengeManagerFilterer   // Log filterer for contract events
}

// IChallengeManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IChallengeManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IChallengeManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IChallengeManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IChallengeManagerSession struct {
	Contract     *IChallengeManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IChallengeManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IChallengeManagerCallerSession struct {
	Contract *IChallengeManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IChallengeManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IChallengeManagerTransactorSession struct {
	Contract     *IChallengeManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IChallengeManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IChallengeManagerRaw struct {
	Contract *IChallengeManager // Generic contract binding to access the raw methods on
}

// IChallengeManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IChallengeManagerCallerRaw struct {
	Contract *IChallengeManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IChallengeManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IChallengeManagerTransactorRaw struct {
	Contract *IChallengeManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIChallengeManager creates a new instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManager(address common.Address, backend bind.ContractBackend) (*IChallengeManager, error) {
	contract, err := bindIChallengeManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IChallengeManager{IChallengeManagerCaller: IChallengeManagerCaller{contract: contract}, IChallengeManagerTransactor: IChallengeManagerTransactor{contract: contract}, IChallengeManagerFilterer: IChallengeManagerFilterer{contract: contract}}, nil
}

// NewIChallengeManagerCaller creates a new read-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerCaller(address common.Address, caller bind.ContractCaller) (*IChallengeManagerCaller, error) {
	contract, err := bindIChallengeManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerCaller{contract: contract}, nil
}

// NewIChallengeManagerTransactor creates a new write-only instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IChallengeManagerTransactor, error) {
	contract, err := bindIChallengeManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerTransactor{contract: contract}, nil
}

// NewIChallengeManagerFilterer creates a new log filterer instance of IChallengeManager, bound to a specific deployed contract.
func NewIChallengeManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IChallengeManagerFilterer, error) {
	contract, err := bindIChallengeManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IChallengeManagerFilterer{contract: contract}, nil
}

// bindIChallengeManager binds a generic wrapper to an already deployed contract.
func bindIChallengeManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IChallengeManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.IChallengeManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.IChallengeManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IChallengeManager *IChallengeManagerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IChallengeManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IChallengeManager *IChallengeManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IChallengeManager.Contract.contract.Transact(opts, method, params...)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactor) InitiateChallenge(opts *bind.TransactOpts, players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.contract.Transact(opts, "initiateChallenge", players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// InitiateChallenge is a paid mutator transaction binding the contract method 0x208e04d4.
//
// Solidity: function initiateChallenge(address[2] players, uint128[2] escrows, uint32 challengePeriod, bytes32 challengeRoot) returns()
func (_IChallengeManager *IChallengeManagerTransactorSession) InitiateChallenge(players [2]common.Address, escrows [2]*big.Int, challengePeriod uint32, challengeRoot [32]byte) (*types.Transaction, error) {
	return _IChallengeManager.Contract.InitiateChallenge(&_IChallengeManager.TransactOpts, players, escrows, challengePeriod, challengeRoot)
}

// IVMTrackerABI is the input ABI used to generate the binding from.
const IVMTrackerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_players\",\"type\":\"address[2]\"},{\"name\":\"_rewards\",\"type\":\"uint128[2]\"}],\"name\":\"completeChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IVMTrackerFuncSigs maps the 4-byte function signature to its string representation.
var IVMTrackerFuncSigs = map[string]string{
	"22c091bc": "completeChallenge(address[2],uint128[2])",
}

// IVMTracker is an auto generated Go binding around an Ethereum contract.
type IVMTracker struct {
	IVMTrackerCaller     // Read-only binding to the contract
	IVMTrackerTransactor // Write-only binding to the contract
	IVMTrackerFilterer   // Log filterer for contract events
}

// IVMTrackerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVMTrackerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVMTrackerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVMTrackerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVMTrackerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVMTrackerSession struct {
	Contract     *IVMTracker       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVMTrackerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVMTrackerCallerSession struct {
	Contract *IVMTrackerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IVMTrackerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVMTrackerTransactorSession struct {
	Contract     *IVMTrackerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IVMTrackerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVMTrackerRaw struct {
	Contract *IVMTracker // Generic contract binding to access the raw methods on
}

// IVMTrackerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVMTrackerCallerRaw struct {
	Contract *IVMTrackerCaller // Generic read-only contract binding to access the raw methods on
}

// IVMTrackerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVMTrackerTransactorRaw struct {
	Contract *IVMTrackerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVMTracker creates a new instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTracker(address common.Address, backend bind.ContractBackend) (*IVMTracker, error) {
	contract, err := bindIVMTracker(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVMTracker{IVMTrackerCaller: IVMTrackerCaller{contract: contract}, IVMTrackerTransactor: IVMTrackerTransactor{contract: contract}, IVMTrackerFilterer: IVMTrackerFilterer{contract: contract}}, nil
}

// NewIVMTrackerCaller creates a new read-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerCaller(address common.Address, caller bind.ContractCaller) (*IVMTrackerCaller, error) {
	contract, err := bindIVMTracker(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerCaller{contract: contract}, nil
}

// NewIVMTrackerTransactor creates a new write-only instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerTransactor(address common.Address, transactor bind.ContractTransactor) (*IVMTrackerTransactor, error) {
	contract, err := bindIVMTracker(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerTransactor{contract: contract}, nil
}

// NewIVMTrackerFilterer creates a new log filterer instance of IVMTracker, bound to a specific deployed contract.
func NewIVMTrackerFilterer(address common.Address, filterer bind.ContractFilterer) (*IVMTrackerFilterer, error) {
	contract, err := bindIVMTracker(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVMTrackerFilterer{contract: contract}, nil
}

// bindIVMTracker binds a generic wrapper to an already deployed contract.
func bindIVMTracker(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVMTrackerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.IVMTrackerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.IVMTrackerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVMTracker *IVMTrackerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IVMTracker.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVMTracker *IVMTrackerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVMTracker *IVMTrackerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVMTracker.Contract.contract.Transact(opts, method, params...)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactor) CompleteChallenge(opts *bind.TransactOpts, _players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.contract.Transact(opts, "completeChallenge", _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _players, _rewards)
}

// CompleteChallenge is a paid mutator transaction binding the contract method 0x22c091bc.
//
// Solidity: function completeChallenge(address[2] _players, uint128[2] _rewards) returns()
func (_IVMTracker *IVMTrackerTransactorSession) CompleteChallenge(_players [2]common.Address, _rewards [2]*big.Int) (*types.Transaction, error) {
	return _IVMTracker.Contract.CompleteChallenge(&_IVMTracker.TransactOpts, _players, _rewards)
}

// MerkleLibABI is the input ABI used to generate the binding from.
const MerkleLibABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_addresses\",\"type\":\"address[]\"}],\"name\":\"generateAddressRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_hashes\",\"type\":\"bytes32[]\"}],\"name\":\"generateRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"proof\",\"type\":\"bytes\"},{\"name\":\"root\",\"type\":\"bytes32\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"verifyProof\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// MerkleLibFuncSigs maps the 4-byte function signature to its string representation.
var MerkleLibFuncSigs = map[string]string{
	"6a2dda67": "generateAddressRoot(address[])",
	"9898dc10": "generateRoot(bytes32[])",
	"b792d767": "verifyProof(bytes,bytes32,bytes32,uint256)",
}

// MerkleLibBin is the compiled bytecode used for deploying new contracts.
var MerkleLibBin = "0x610575610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004b5760003560e01c80636a2dda67146100505780639898dc1014610105578063b792d767146101a8575b600080fd5b6100f36004803603602081101561006657600080fd5b81019060208101813564010000000081111561008157600080fd5b82018360208201111561009357600080fd5b803590602001918460208302840111640100000000831117156100b557600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061026d945050505050565b60408051918252519081900360200190f35b6100f36004803603602081101561011b57600080fd5b81019060208101813564010000000081111561013657600080fd5b82018360208201111561014857600080fd5b8035906020019184602083028401116401000000008311171561016a57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610301945050505050565b610259600480360360808110156101be57600080fd5b8101906020810181356401000000008111156101d957600080fd5b8201836020820111156101eb57600080fd5b8035906020019184600183028401116401000000008311171561020d57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550508235935050506020810135906040013561043f565b604080519115158252519081900360200190f35b60006060825160405190808252806020026020018201604052801561029c578160200160208202803883390190505b50905060005b83518110156102f0578381815181106102b757fe5b602002602001015160601b6bffffffffffffffffffffffff19168282815181106102dd57fe5b60209081029190910101526001016102a2565b506102fa81610301565b9392505050565b6000815b600181511115610422576060600282516001018161031f57fe5b04604051908082528060200260200182016040528015610349578160200160208202803883390190505b50905060005b815181101561041a5782518160020260010110156103e25782816002028151811061037657fe5b602002602001015183826002026001018151811061039057fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001208282815181106103d157fe5b602002602001018181525050610412565b8281600202815181106103f157fe5b602002602001015182828151811061040557fe5b6020026020010181815250505b60010161034f565b509050610305565b8060008151811061042f57fe5b6020026020010151915050919050565b600080838160205b88518111610532578089015193506020818a51036020018161046557fe5b0491505b60008211801561047c5750600286066001145b801561048a57508160020a86115b1561049d57600286046001019550610469565b600286066104e85783836040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209250600286816104e057fe5b04955061052a565b828460405160200180838152602001828152602001925050506040516020818303038152906040528051906020012092506002868161052357fe5b0460010195505b602001610447565b50509094149594505050505056fea265627a7a72305820a30d7a9fc682b6c10a7f71432e6ecd99e509b21738d16aed5eb432041d0451b764736f6c634300050a0032"

// DeployMerkleLib deploys a new Ethereum contract, binding an instance of MerkleLib to it.
func DeployMerkleLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleLib, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// MerkleLib is an auto generated Go binding around an Ethereum contract.
type MerkleLib struct {
	MerkleLibCaller     // Read-only binding to the contract
	MerkleLibTransactor // Write-only binding to the contract
	MerkleLibFilterer   // Log filterer for contract events
}

// MerkleLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleLibSession struct {
	Contract     *MerkleLib        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleLibCallerSession struct {
	Contract *MerkleLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MerkleLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleLibTransactorSession struct {
	Contract     *MerkleLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MerkleLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleLibRaw struct {
	Contract *MerkleLib // Generic contract binding to access the raw methods on
}

// MerkleLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleLibCallerRaw struct {
	Contract *MerkleLibCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleLibTransactorRaw struct {
	Contract *MerkleLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleLib creates a new instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLib(address common.Address, backend bind.ContractBackend) (*MerkleLib, error) {
	contract, err := bindMerkleLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleLib{MerkleLibCaller: MerkleLibCaller{contract: contract}, MerkleLibTransactor: MerkleLibTransactor{contract: contract}, MerkleLibFilterer: MerkleLibFilterer{contract: contract}}, nil
}

// NewMerkleLibCaller creates a new read-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibCaller(address common.Address, caller bind.ContractCaller) (*MerkleLibCaller, error) {
	contract, err := bindMerkleLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibCaller{contract: contract}, nil
}

// NewMerkleLibTransactor creates a new write-only instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleLibTransactor, error) {
	contract, err := bindMerkleLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleLibTransactor{contract: contract}, nil
}

// NewMerkleLibFilterer creates a new log filterer instance of MerkleLib, bound to a specific deployed contract.
func NewMerkleLibFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleLibFilterer, error) {
	contract, err := bindMerkleLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleLibFilterer{contract: contract}, nil
}

// bindMerkleLib binds a generic wrapper to an already deployed contract.
func bindMerkleLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.MerkleLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.MerkleLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleLib *MerkleLibCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleLib *MerkleLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleLib *MerkleLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleLib.Contract.contract.Transact(opts, method, params...)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateAddressRoot(opts *bind.CallOpts, _addresses []common.Address) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateAddressRoot", _addresses)
	return *ret0, err
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateAddressRoot is a free data retrieval call binding the contract method 0x6a2dda67.
//
// Solidity: function generateAddressRoot(address[] _addresses) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateAddressRoot(_addresses []common.Address) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateAddressRoot(&_MerkleLib.CallOpts, _addresses)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCaller) GenerateRoot(opts *bind.CallOpts, _hashes [][32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "generateRoot", _hashes)
	return *ret0, err
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// GenerateRoot is a free data retrieval call binding the contract method 0x9898dc10.
//
// Solidity: function generateRoot(bytes32[] _hashes) constant returns(bytes32)
func (_MerkleLib *MerkleLibCallerSession) GenerateRoot(_hashes [][32]byte) ([32]byte, error) {
	return _MerkleLib.Contract.GenerateRoot(&_MerkleLib.CallOpts, _hashes)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCaller) VerifyProof(opts *bind.CallOpts, proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MerkleLib.contract.Call(opts, out, "verifyProof", proof, root, hash, index)
	return *ret0, err
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// VerifyProof is a free data retrieval call binding the contract method 0xb792d767.
//
// Solidity: function verifyProof(bytes proof, bytes32 root, bytes32 hash, uint256 index) constant returns(bool)
func (_MerkleLib *MerkleLibCallerSession) VerifyProof(proof []byte, root [32]byte, hash [32]byte, index *big.Int) (bool, error) {
	return _MerkleLib.Contract.VerifyProof(&_MerkleLib.CallOpts, proof, root, hash, index)
}

// OneStepProofABI is the input ABI used to generate the binding from.
const OneStepProofABI = "[]"

// OneStepProofFuncSigs maps the 4-byte function signature to its string representation.
var OneStepProofFuncSigs = map[string]string{
	"c4aa83c0": "oneStepProofFirst(Challenge.Data storage,bytes32[2],uint64[2],bytes21[],uint256[],bytes32[5],uint256[],bytes)",
	"ab06a992": "oneStepProofOther(Challenge.Data storage,bytes32[10],uint64[2],bytes21[],uint256[],uint32,uint256[],uint32,uint256[],bytes)",
}

// OneStepProofBin is the compiled bytecode used for deploying new contracts.
var OneStepProofBin = "0x614596610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c8063ab06a99214610045578063c4aa83c014610372575b600080fd5b610370600480360361028081101561005c57600080fd5b604080516101408181019092528335939283019291610160830191906020840190600a908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b8111156100d857600080fd5b8201836020820111156100ea57600080fd5b803590602001918460208302840111600160201b8311171561010b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561015a57600080fd5b82018360208201111561016c57600080fd5b803590602001918460208302840111600160201b8311171561018d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b8111156101ea57600080fd5b8201836020820111156101fc57600080fd5b803590602001918460208302840111600160201b8311171561021d57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929563ffffffff853516959094909350604081019250602001359050600160201b81111561027a57600080fd5b82018360208201111561028c57600080fd5b803590602001918460208302840111600160201b831117156102ad57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156102fc57600080fd5b82018360208201111561030e57600080fd5b803590602001918460018302840111600160201b8311171561032f57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610626945050505050565b005b61037060048036036101c081101561038957600080fd5b604080518082018252833593928301929160608301919060208401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561040057600080fd5b82018360208201111561041257600080fd5b803590602001918460208302840111600160201b8311171561043357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561048257600080fd5b82018360208201111561049457600080fd5b803590602001918460208302840111600160201b831117156104b557600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250506040805160a0818101909252939695948181019493509150600590839083908082843760009201919091525091949392602081019250359050600160201b81111561053057600080fd5b82018360208201111561054257600080fd5b803590602001918460208302840111600160201b8311171561056357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156105b257600080fd5b8201836020820111156105c457600080fd5b803590602001918460018302840111600160201b831117156105e557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506107ef945050505050565b6107e38a604051806101e001604052808c6000600a811061064357fe5b602002015181526020018c6001600a811061065a57fe5b60200201518c8c6040516020018084815260200183600260200280838360005b8381101561069257818101518382015260200161067a565b50505050905001828051906020019060200280838360005b838110156106c25781810151838201526020016106aa565b5050505090500193505050506040516020818303038152906040528051906020012081526020018981526020018c6002600a81106106fc57fe5b602002015181526020018c6003600a811061071357fe5b602002015181526020018c6004600a811061072a57fe5b602002015181526020018863ffffffff1681526020018c6005600a811061074d57fe5b602002015181526020018c6006600a811061076457fe5b602002015181526020018781526020018c6007600a811061078157fe5b602002015181526020018663ffffffff1681526020018c6008600a81106107a457fe5b602002015181526020018c6009600a81106107bb57fe5b60200201518152602001858152508b6001600a81106107d657fe5b60200201518b8b86610c68565b50505050505050505050565b60016005890154600160601b900460ff16600281111561080b57fe5b146108475760405162461bcd60e51b81526004018080602001828103825260398152602001806145296039913960400191505060405180910390fd5b600588015467ffffffffffffffff164311156108aa576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6520737465702070726f6f66206d697373656420646561646c696e650000604482015290519081900360640190fd5b60018089015490889060200201518760006020020151886001602002015188604051602001808581526020018467ffffffffffffffff1667ffffffffffffffff1660c01b81526008018367ffffffffffffffff1667ffffffffffffffff1660c01b8152600801828051906020019060200280838360005b83811015610939578181015183820152602001610921565b50505050905001945050505050604051602081830303815290604052805190602001208860006002811061096957fe5b60200201518660405160200180838152602001828051906020019060200280838360005b838110156109a557818101518382015260200161098d565b50505050905001925050506040516020818303038152906040528051906020012073__$9836fa7140e5a33041d4b827682e675a30$__6320903721876000600581106109ed57fe5b60200201516001898160200201518a600260200201518b600360200201518c600460200201518c6040518863ffffffff1660e01b8152600401808881526020018763ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b83811015610a8d578181015183820152602001610a75565b505050509050019850505050505050505060206040518083038186803b158015610ab657600080fd5b505af4158015610aca573d6000803e3d6000fd5b505050506040513d6020811015610ae057600080fd5b505160408051602081810195909552808201939093526060830191909152600160e01b60808301528051606481840301815260849092019052805191012014610b5a5760405162461bcd60e51b81526004018080602001828103825260268152602001806143d56026913960400191505060405180910390fd5b6000610c126040518060e001604052808a600060028110610b7757fe5b602002015181526020018a600160028110610b8e57fe5b6020020151815260200186600060058110610ba557fe5b6020020151815260200186600160058110610bbc57fe5b6020020151815260200186600260058110610bd357fe5b6020020151815260200186600360058110610bea57fe5b6020020151815260200186600460058110610c0157fe5b602002015190528888888787610de8565b90508015610c5d576040805162461bcd60e51b8152602060048201526013602482015272141c9bdbd9881dd85cc81a5b98dbdc9c9958dd606a1b604482015290519081900360640190fd5b505050505050505050565b60016005870154600160601b900460ff166002811115610c8457fe5b14610cc05760405162461bcd60e51b81526004018080602001828103825260398152602001806145296039913960400191505060405180910390fd5b600586015467ffffffffffffffff16431115610d23576040805162461bcd60e51b815260206004820152601e60248201527f4f6e6520737465702070726f6f66206d697373656420646561646c696e650000604482015290519081900360640190fd5b610d2d8686611180565b6000610d946040518060e001604052808860a00151815260200187815260200188610140015181526020018860e00151815260200188610180015181526020018861010001518152602001886101a00151815250858589604001518a6101c0015187610de8565b90508015610ddf576040805162461bcd60e51b8152602060048201526013602482015272141c9bdbd9881dd85cc81a5b98dbdc9c9958dd606a1b604482015290519081900360640190fd5b50505050505050565b608086015160608701516000918291829182911480159060001990610ffd5760005b88518167ffffffffffffffff161015610e9f57888167ffffffffffffffff1681518110610e3357fe5b6020026020010151600014610e97578160070b60001914610e93576040805162461bcd60e51b81526020600482015260156024820152746d756c7469706c65206f7574206d6573736167657360581b604482015290519081900360640190fd5b8091505b600101610e0a565b508060070b60001914610ff857878160070b81518110610ebb57fe5b60200260200101519350898160070b81518110610ed457fe5b6020026020010151945060019250898160070b81518110610ef157fe5b6020026020010151601460158110610f0557fe5b1a60f81b6001600160f81b031916600160f81b1415610f8d5783898260070b81518110610f2e57fe5b602002602001015114610f88576040805162461bcd60e51b815260206004820152601a60248201527f707265636f6e646974696f6e206d7573742068617665206e6674000000000000604482015290519081900360640190fd5b610ff8565b888160070b81518110610f9c57fe5b6020026020010151841115610ff8576040805162461bcd60e51b815260206004820152601c60248201527f707265636f6e646974696f6e206d75737420686176652076616c756500000000604482015290519081900360640190fd5b61108f565b60005b88518167ffffffffffffffff16101561108d57888167ffffffffffffffff168151811061102957fe5b6020026020010151600014611085576040805162461bcd60e51b815260206004820152601b60248201527f4d7573742068617665206e6f206d6573736167652076616c7565730000000000604482015290519081900360640190fd5b600101611000565b505b6111706040518061018001604052808e6000600781106110ab57fe5b602002015181526020018d81526020018e6001600781106110c857fe5b602002015181526020018e6002600781106110df57fe5b602002015181526020018e6003600781106110f657fe5b602002015181526020018e60046007811061110d57fe5b602002015181526020018e60056007811061112457fe5b602002015181526020018e60066007811061113b57fe5b60200201518152602001876affffffffffffffffffffff19168152602001868152602001851515815260200189815250611546565b9c9b505050505050505050505050565b816001015481602001518260000151836040015160405160200180838152602001828051906020019060200280838360005b838110156111ca5781810151838201526020016111b2565b50505050905001925050506040516020818303038152906040528051906020012073__$9836fa7140e5a33041d4b827682e675a30$__63209037218560a001518660c0015187606001518860e0015189608001518a61010001518b61012001516040518863ffffffff1660e01b8152600401808881526020018763ffffffff1663ffffffff16815260200186815260200185815260200184815260200183815260200180602001828103825283818151815260200191508051906020019060200280838360005b838110156112a9578181015183820152602001611291565b505050509050019850505050505050505060206040518083038186803b1580156112d257600080fd5b505af41580156112e6573d6000803e3d6000fd5b505050506040513d60208110156112fc57600080fd5b505161014085015161016086015160608701516101a088015160808901516101c08a0151604051632090372160e01b81526004810187815263ffffffff8716602483015260448201869052606482018590526084820184905260a4820185905260e060c48301908152835160e4840152835173__$9836fa7140e5a33041d4b827682e675a30$__9963209037219990989097909690959094869491939192610104909101906020858101910280838360005b838110156113c65781810151838201526020016113ae565b505050509050019850505050505050505060206040518083038186803b1580156113ef57600080fd5b505af4158015611403573d6000803e3d6000fd5b505050506040513d602081101561141957600080fd5b505160c08601516101608701516040805160208181019890985280820196909652606086019490945260808501929092526001600160e01b0319910360e01b1660a08301528051608481840301815260a490920190528051910120146114b05760405162461bcd60e51b815260040180806020018281038252603081526020018061441d6030913960400191505060405180910390fd5b60408101515160005b818110156115405782610120015181815181106114d257fe5b6020026020010151836040015182815181106114ea57fe5b602002602001018181510391508181525050826101200151818151811061150d57fe5b6020026020010151836101c00151828151811061152657fe5b6020908102919091010180519190910390526001016114b9565b50505050565b6000808080606061155561433d565b61155d61433d565b611566886123f0565b939950929650909450925090506001600060ff88168214156115bc576115b5838660008151811061159357fe5b6020026020010151876001815181106115a857fe5b6020026020010151612830565b9150612244565b60ff8816600214156115fb576115b583866000815181106115d957fe5b6020026020010151876001815181106115ee57fe5b6020026020010151612880565b60ff88166003141561163a576115b5838660008151811061161857fe5b60200260200101518760018151811061162d57fe5b60200260200101516128c1565b60ff881660041415611679576115b5838660008151811061165757fe5b60200260200101518760018151811061166c57fe5b6020026020010151612902565b60ff8816600514156116b8576115b5838660008151811061169657fe5b6020026020010151876001815181106116ab57fe5b6020026020010151612943565b60ff8816600614156116f7576115b583866000815181106116d557fe5b6020026020010151876001815181106116ea57fe5b6020026020010151612984565b60ff881660071415611736576115b5838660008151811061171457fe5b60200260200101518760018151811061172957fe5b60200260200101516129c5565b60ff88166008141561178a576115b5838660008151811061175357fe5b60200260200101518760018151811061176857fe5b60200260200101518860028151811061177d57fe5b6020026020010151612a06565b60ff8816600914156117de576115b583866000815181106117a757fe5b6020026020010151876001815181106117bc57fe5b6020026020010151886002815181106117d157fe5b6020026020010151612a5f565b60ff8816600a141561181d576115b583866000815181106117fb57fe5b60200260200101518760018151811061181057fe5b6020026020010151612aa7565b60ff88166010141561185c576115b5838660008151811061183a57fe5b60200260200101518760018151811061184f57fe5b6020026020010151612ae8565b60ff88166011141561189b576115b5838660008151811061187957fe5b60200260200101518760018151811061188e57fe5b6020026020010151612b29565b60ff8816601214156118da576115b583866000815181106118b857fe5b6020026020010151876001815181106118cd57fe5b6020026020010151612b6a565b60ff8816601314156118f7576115b583866000815181106118b857fe5b60ff881660141415611936576115b5838660008151811061191457fe5b60200260200101518760018151811061192957fe5b6020026020010151612bab565b60ff881660151415611960576115b5838660008151811061195357fe5b6020026020010151612bd7565b60ff88166016141561199f576115b5838660008151811061197d57fe5b60200260200101518760018151811061199257fe5b6020026020010151612c1d565b60ff8816601714156119de576115b583866000815181106119bc57fe5b6020026020010151876001815181106119d157fe5b6020026020010151612c5e565b60ff881660181415611a1d576115b583866000815181106119fb57fe5b602002602001015187600181518110611a1057fe5b6020026020010151612c9f565b60ff881660191415611a47576115b58386600081518110611a3a57fe5b6020026020010151612ce0565b60ff8816601a1415611a86576115b58386600081518110611a6457fe5b602002602001015187600181518110611a7957fe5b6020026020010151612d16565b60ff8816601b1415611ac5576115b58386600081518110611aa357fe5b602002602001015187600181518110611ab857fe5b6020026020010151612d57565b60ff881660201415611aef576115b58386600081518110611ae257fe5b6020026020010151612d98565b60ff881660211415611b19576115b58386600081518110611b0c57fe5b6020026020010151612db4565b60ff881660301415611b43576115b58386600081518110611b3657fe5b6020026020010151612dcf565b60ff881660311415611b58576115b583612dd7565b60ff881660321415611b6d576115b583612df8565b60ff881660331415611b97576115b58386600081518110611b8a57fe5b6020026020010151612e11565b60ff881660341415611bc1576115b58386600081518110611bb457fe5b6020026020010151612e2a565b60ff881660351415611c00576115b58386600081518110611bde57fe5b602002602001015187600181518110611bf357fe5b6020026020010151612e40565b60ff881660361415611c15576115b583612e73565b60ff881660371415611c2f576115b5838560000151612ea5565b60ff881660381415611c59576115b58386600081518110611c4c57fe5b6020026020010151612eb7565b60ff881660391415611ce657611c6d61439e565b611c7c8b610160015188612ec9565b919950975090508715611cc05760405162461bcd60e51b81526004018080602001828103825260218152602001806145086021913960400191505060405180910390fd5b611cd0858263ffffffff61301e16565b611ce0848263ffffffff61304016565b50612244565b60ff8816603b1415611cf757612244565b60ff881660401415611d21576115b58386600081518110611d1457fe5b602002602001015161305d565b60ff881660411415611d60576115b58386600081518110611d3e57fe5b602002602001015187600181518110611d5357fe5b602002602001015161307f565b60ff881660421415611db4576115b58386600081518110611d7d57fe5b602002602001015187600181518110611d9257fe5b602002602001015188600281518110611da757fe5b60200260200101516130b1565b60ff881660431415611df3576115b58386600081518110611dd157fe5b602002602001015187600181518110611de657fe5b60200260200101516130f3565b60ff881660441415611e47576115b58386600081518110611e1057fe5b602002602001015187600181518110611e2557fe5b602002602001015188600281518110611e3a57fe5b6020026020010151613105565b60ff881660501415611e86576115b58386600081518110611e6457fe5b602002602001015187600181518110611e7957fe5b6020026020010151613127565b60ff881660511415611eda576115b58386600081518110611ea357fe5b602002602001015187600181518110611eb857fe5b602002602001015188600281518110611ecd57fe5b602002602001015161319e565b60ff881660521415611f04576115b58386600081518110611ef757fe5b6020026020010151613217565b60ff881660601415611f19576115b58361324a565b60ff88166061141561200357611f438386600081518110611f3657fe5b6020026020010151613250565b60e08c015160c08d0151604080516020808201939093528082018590528151808203830181526060909101909152805191012092945090925014611fb85760405162461bcd60e51b81526004018080602001828103825260258152602001806144946025913960400191505060405180910390fd5b8960a001518a6080015114611ffe5760405162461bcd60e51b81526004018080602001828103825260278152602001806144b96027913960400191505060405180910390fd5b612244565b60ff88166070141561210157600080612030858860008151811061202357fe5b6020026020010151613272565b809450819550829650839750505050508b60a001518c6080015184604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120146120b45760405162461bcd60e51b81526004018080602001828103825260288152602001806144e06028913960400191505060405180910390fd5b8b60e001518c60c00151146120fa5760405162461bcd60e51b815260040180806020018281038252602681526020018061444d6026913960400191505060405180910390fd5b5050612244565b60ff88166071141561212157600080612030858860008151811061202357fe5b60ff8816607214156121dd576040805160028082526060828101909352816020015b61214b61439e565b81526020019060019003908161214357505060208c01519091506121809060005b602002015167ffffffffffffffff16613449565b8160008151811061218d57fe5b60200260200101819052506121ac8b6020015160016002811061216c57fe5b816001815181106121b957fe5b6020026020010181905250611ce06121d0826134a3565b859063ffffffff61304016565b60ff88166073141561221a576115b583866000815181106121fa57fe5b602002602001015160405180602001604052808e6040015181525061352b565b60ff88166074141561222f5760009150612244565b60ff881660751415612244576122448361359d565b806122d5578960a001518a608001511461228f5760405162461bcd60e51b81526004018080602001828103825260278152602001806144b96027913960400191505060405180910390fd5b8960e001518a60c00151146122d55760405162461bcd60e51b815260040180806020018281038252602681526020018061444d6026913960400191505060405180910390fd5b816123375760408051600160f81b602080830191909152600060218301819052602280840191909152835180840390910181526042909201909252805191012060a084015151141561232f5761232a836135a7565b612337565b60a083015183525b612340846135b1565b8a511461237e5760405162461bcd60e51b81526004018080602001828103825260228152602001806143fb6022913960400191505060405180910390fd5b612387836135b1565b8a60600151146123de576040805162461bcd60e51b815260206004820181905260248201527f50726f6f6620686164206e6f6e206d61746368696e6720656e64207374617465604482015290519081900360640190fd5b6000985050505050505050505b919050565b600060606123fc61433d565b61240461433d565b6000808061241061433d565b61241981613646565b61242889610160015184613650565b909450909250905061243861433d565b61244182613755565b905060008a6101600151858151811061245657fe5b602001015160f81c60f81b60f81c905060008b6101600151866001018151811061247c57fe5b016020015160f81c90506000612491826137b3565b90506060816040519080825280602002602001820160405280156124cf57816020015b6124bc61439e565b8152602001906001900390816124b45790505b5090506002880197508360ff16600014806124ed57508360ff166001145b61253e576040805162461bcd60e51b815260206004820152601c60248201527f50726f6f662068616420626164206f7065726174696f6e207479706500000000604482015290519081900360640190fd5b60ff84166125e1576040805160208082018084528951516353409fab60e01b90915260ff87166024840152604483015291519091829173__$d969135829891f807aa9c34494da4ecd99$__916353409fab916064808601929190818703018186803b1580156125ac57600080fd5b505af41580156125c0573d6000803e3d6000fd5b505050506040513d60208110156125d657600080fd5b505190528652612738565b6125e961439e565b6125f88f61016001518a612ec9565b909a5090985090508715612653576040805162461bcd60e51b815260206004820152601d60248201527f50726f6f66206861642062616420696d6d6564696174652076616c7565000000604482015290519081900360640190fd5b821561267757808260008151811061266757fe5b6020026020010181905250612687565b612687868263ffffffff61304016565b604051806020016040528073__$d969135829891f807aa9c34494da4ecd99$__63264f384b876126b6866137cd565b518c5151604080516001600160e01b031960e087901b16815260ff909416600485015260248401929092526044830152516064808301926020929190829003018186803b15801561270657600080fd5b505af415801561271a573d6000803e3d6000fd5b505050506040513d602081101561273057600080fd5b505190528752505b60ff84165b828110156127cc576127548f61016001518a612ec9565b845185908590811061276257fe5b60209081029190910101529950975087156127c4576040805162461bcd60e51b815260206004820152601960248201527f50726f6f66206861642062616420737461636b2076616c756500000000000000604482015290519081900360640190fd5b60010161273d565b815115612819575060005b8460ff16825103811015612819576128118282600185510303815181106127fa57fe5b60200260200101518861304090919063ffffffff16565b6001016127d7565b50919d919c50939a50919850939650945050505050565b600061283b836138bc565b158061284d575061284b826138bc565b155b1561285a57506000612879565b82518251808201612871878263ffffffff6138c716565b600193505050505b9392505050565b600061288b836138bc565b158061289d575061289b826138bc565b155b156128aa57506000612879565b82518251808202612871878263ffffffff6138c716565b60006128cc836138bc565b15806128de57506128dc826138bc565b155b156128eb57506000612879565b82518251808203612871878263ffffffff6138c716565b600061290d836138bc565b158061291f575061291d826138bc565b155b1561292c57506000612879565b82518251808204612871878263ffffffff6138c716565b600061294e836138bc565b1580612960575061295e826138bc565b155b1561296d57506000612879565b82518251808205612871878263ffffffff6138c716565b600061298f836138bc565b15806129a1575061299f826138bc565b155b156129ae57506000612879565b82518251808206612871878263ffffffff6138c716565b60006129d0836138bc565b15806129e257506129e0826138bc565b155b156129ef57506000612879565b82518251808207612871878263ffffffff6138c716565b6000612a11846138bc565b1580612a235750612a21836138bc565b155b15612a3057506000612a57565b8351835183516000818385089050612a4e898263ffffffff6138c716565b60019450505050505b949350505050565b6000612a6a846138bc565b1580612a7c5750612a7a836138bc565b155b15612a8957506000612a57565b8351835183516000818385099050612a4e898263ffffffff6138c716565b6000612ab2836138bc565b1580612ac45750612ac2826138bc565b155b15612ad157506000612879565b8251825180820a612871878263ffffffff6138c716565b6000612af3836138bc565b1580612b055750612b03826138bc565b155b15612b1257506000612879565b82518251808210612871878263ffffffff6138c716565b6000612b34836138bc565b1580612b465750612b44826138bc565b155b15612b5357506000612879565b82518251808211612871878263ffffffff6138c716565b6000612b75836138bc565b1580612b875750612b85826138bc565b155b15612b9457506000612879565b82518251808212612871878263ffffffff6138c716565b6000612bcd6121d0612bbc846137cd565b51612bc6866137cd565b51146138db565b5060019392505050565b6000612be2826138bc565b612bfc57612bf783600063ffffffff6138c716565b612c13565b81518015612c10858263ffffffff6138c716565b50505b5060015b92915050565b6000612c28836138bc565b1580612c3a5750612c38826138bc565b155b15612c4757506000612879565b82518251808216612871878263ffffffff6138c716565b6000612c69836138bc565b1580612c7b5750612c79826138bc565b155b15612c8857506000612879565b82518251808217612871878263ffffffff6138c716565b6000612caa836138bc565b1580612cbc5750612cba826138bc565b155b15612cc957506000612879565b82518251808218612871878263ffffffff6138c716565b6000612ceb826138bc565b612cf757506000612c17565b81518019612d0b858263ffffffff6138c716565b506001949350505050565b6000612d21836138bc565b1580612d335750612d31826138bc565b155b15612d4057506000612879565b8251825180821a612871878263ffffffff6138c716565b6000612d62836138bc565b1580612d745750612d72826138bc565b155b15612d8157506000612879565b8251825180820b612871878263ffffffff6138c716565b6000612c13612da6836137cd565b51849063ffffffff6138c716565b6000612c13612dc283613904565b849063ffffffff61304016565b600192915050565b6000612df082608001518361398d90919063ffffffff16565b506001919050565b6000612df082606001518361398d90919063ffffffff16565b6000612e1c826137cd565b606084015250600192915050565b6000612e35826137cd565b835250600192915050565b6000612e4b826138bc565b612e5757506000612879565b825115612bcd57612e67836137cd565b84525060019392505050565b6000612df0612e98612e8b612e8661399b565b6137cd565b51602085015151146138db565b839063ffffffff61304016565b6000612c13838363ffffffff61398d16565b6000612c13838363ffffffff61301e16565b600080612ed461439e565b84518410612f29576040805162461bcd60e51b815260206004820152601960248201527f44617461206f6666736574206f7574206f6620626f756e647300000000000000604482015290519081900360640190fd5b60008490506000868281518110612f3c57fe5b016020015160019092019160f81c9050600081612f7e57612f5d88846139f5565b9093509050600083612f6e83613449565b9197509550935061301792505050565b60ff821660021415612fa557612f9488846139f5565b9093509050600083612f6e83613a1c565b600360ff831610801590612fbc5750600c60ff8316105b15612ff857600219820160606000612fd5838c88613a76565b909750925090508086612fe7846134a3565b985098509850505050505050613017565b8160ff1661271001600061300c6000613449565b919750955093505050505b9250925092565b613034826040015161302f836137cd565b613b31565b82604001819052505050565b613051826020015161302f836137cd565b82602001819052505050565b600061306f838363ffffffff61304016565b612c13838363ffffffff61304016565b6000613091848363ffffffff61304016565b6130a1848463ffffffff61304016565b612bcd848363ffffffff61304016565b60006130c3858363ffffffff61304016565b6130d3858463ffffffff61304016565b6130e3858563ffffffff61304016565b612d0b858363ffffffff61304016565b60006130a1848463ffffffff61304016565b6000613117858563ffffffff61304016565b6130e3858463ffffffff61304016565b6000613132836138bc565b1580613144575061314282613be7565b155b1561315157506000612879565b61315a82613bf6565b60ff168360000151111561317057506000612879565b612bcd826020015184600001518151811061318757fe5b60200260200101518561304090919063ffffffff16565b60006131a983613be7565b15806131bb57506131b9846138bc565b155b156131c857506000612a57565b6131d183613bf6565b60ff16846000015111156131e757506000612a57565b8183602001518560000151815181106131fc57fe5b6020908102919091010152612d0b858463ffffffff61304016565b600061322282613be7565b61322e57506000612c17565b612c1361323a83613bf6565b849060ff1663ffffffff6138c716565b50600190565b60008061325b6143c2565b613264846137cd565b516001969095509350505050565b600080600080600080600061328688613be7565b61329a576000965094509092509050613440565b6132bb88602001516001815181106132ae57fe5b60200260200101516138bc565b6132cf576000965094509092509050613440565b6132e388602001516002815181106132ae57fe5b6132f7576000965094509092509050613440565b61330b88602001516003815181106132ae57fe5b61331f576000965094509092509050613440565b876020015160018151811061333057fe5b60200260200101516000015160001b9250876020015160028151811061335257fe5b602002602001015160000151915073__$9836fa7140e5a33041d4b827682e675a30$__624c28f66133828a6137cd565b6000015185858c6020015160038151811061339957fe5b60209081029190910181015151604080516001600160e01b031960e089901b16815260048101969096526affffffffffffffffffffff199094166024860152604485019290925260609190911c60648401529051608480840193829003018186803b15801561340757600080fd5b505af415801561341b573d6000803e3d6000fd5b505050506040513d602081101561343157600080fd5b50516001975095509193509150505b92959194509250565b61345161439e565b604080516060810182528381528151600080825260208281019094529192830191613492565b61347f61439e565b8152602001906001900390816134775790505b508152600060209091015292915050565b6134ab61439e565b6134b58251613c05565b613506576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b50604080516060810182526000815260208101839052915160030160ff169082015290565b8051600090613539846137cd565b51141561358d576040805162461bcd60e51b815260206004820152601d60248201527f496e626f7820696e737472756374696f6e2077617320626c6f636b6564000000604482015290519081900360640190fd5b612bcd848363ffffffff61398d16565b600260c090910152565b600160c090910152565b600060028260c0015114156135c8575060006123eb565b60018260c0015114156135dd575060016123eb565b508051516020808301515160408085015151606080870151516080808901515160a0808b0151518751808b019b909b528a8801989098529389019490945287015285015260c0808501929092528051808503909201825260e090930190925281519101206123eb565b600060c090910152565b60008061365b61433d565b61366361433d565b600060c082018190526136768787613c0c565b845296509050801561368e5793508492509050613017565b6136988787613c0c565b60208501529650905080156136b35793508492509050613017565b6136bd8787613c0c565b60408501529650905080156136d85793508492509050613017565b6136e28787613c0c565b60608501529650905080156136fd5793508492509050613017565b6137078787613c0c565b60808501529650905080156137225793508492509050613017565b61372c8787613c0c565b60a08501529650905080156137475793508492509050613017565b506000969495509392505050565b61375d61433d565b6040518060e0016040528083600001518152602001836020015181526020018360400151815260200183606001518152602001836080015181526020018360a0015181526020018360c001518152509050919050565b60008060006137c48460ff16613c4a565b50949350505050565b6137d56143c2565b6040820151600c60ff90911610613827576040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b604082015160ff1661385457604051806020016040528061384b84600001516140f1565b905290506123eb565b604082015160ff166002141561387957506040805160208101909152815181526123eb565b600360ff16826040015160ff161015801561389d57506040820151600c60ff909116105b156138ba57604051806020016040528061384b8460200151614115565bfe5b6040015160ff161590565b613051826020015161302f612e8684613449565b6138e361439e565b81156138fa576138f36001613449565b90506123eb565b6138f36000613449565b61390c61439e565b816040015160ff16600214156139535760405162461bcd60e51b81526004018080602001828103825260218152602001806144736021913960400191505060405180910390fd5b604082015160ff16613969576138f36000613449565b816040015160ff1660011415613983576138f36001613449565b6138f36003613449565b613051826020015182613b31565b6139a361439e565b604080516060810182526000808252825181815260208181019094529192830191906139e5565b6139d261439e565b8152602001906001900390816139ca5790505b5081526003602090910152905090565b6000808281613a0a868363ffffffff61426116565b60209290920196919550909350505050565b613a2461439e565b604080516060810182528381528151600080825260208281019094529192830191613a65565b613a5261439e565b815260200190600190039081613a4a5790505b508152600260209091015292915050565b60008060606000849050600060608860ff16604051908082528060200260200182016040528015613ac157816020015b613aae61439e565b815260200190600190039081613aa65790505b50905060005b8960ff168160ff161015613b1b57613adf8985612ec9565b8451859060ff8616908110613af057fe5b6020908102919091010152945092508215613b1357509094509092509050613b28565b600101613ac7565b5060009550919350909150505b93509350939050565b613b396143c2565b6040805160028082526060828101909352816020015b613b576143c2565b815260200190600190039081613b4f5790505090508281600081518110613b7a57fe5b60200260200101819052508381600181518110613b9357fe5b60200260200101819052506040518060200160405280613bdd6040518060400160405280613bc48860000151613a1c565b8152602001613bd68960000151613a1c565b905261427d565b9052949350505050565b6000612c1782604001516142fc565b6000612c17826040015161431a565b6008101590565b600080613c176143c2565b836000613c2a878363ffffffff61426116565b604080516020808201909252918252600099930197509550909350505050565b6000806001831415613c6257506002905060016140ec565b6002831415613c7757506002905060016140ec565b6003831415613c8c57506002905060016140ec565b6004831415613ca157506002905060016140ec565b6005831415613cb657506002905060016140ec565b6006831415613ccb57506002905060016140ec565b6007831415613ce057506002905060016140ec565b6008831415613cf557506003905060016140ec565b6009831415613d0a57506003905060016140ec565b600a831415613d1f57506002905060016140ec565b6010831415613d3457506002905060016140ec565b6011831415613d4957506002905060016140ec565b6012831415613d5e57506002905060016140ec565b6013831415613d7357506002905060016140ec565b6014831415613d8857506002905060016140ec565b6015831415613d9c575060019050806140ec565b6016831415613db157506002905060016140ec565b6017831415613dc657506002905060016140ec565b6018831415613ddb57506002905060016140ec565b6019831415613def575060019050806140ec565b601a831415613e0457506002905060016140ec565b601b831415613e1957506002905060016140ec565b6020831415613e2d575060019050806140ec565b6021831415613e41575060019050806140ec565b6030831415613e5657506001905060006140ec565b6031831415613e6b57506000905060016140ec565b6032831415613e8057506000905060016140ec565b6033831415613e9557506001905060006140ec565b6034831415613eaa57506001905060006140ec565b6035831415613ebf57506002905060006140ec565b6036831415613ed457506000905060016140ec565b6037831415613ee957506000905060016140ec565b6038831415613efe57506001905060006140ec565b6039831415613f1357506000905060016140ec565b603a831415613f2857506000905060016140ec565b603b831415613f3c575060009050806140ec565b603c831415613f5157506000905060016140ec565b603d831415613f6657506001905060006140ec565b6040831415613f7b57506001905060026140ec565b6041831415613f9057506002905060036140ec565b6042831415613fa557506003905060046140ec565b6043831415613fb9575060029050806140ec565b6044831415613fcd575060039050806140ec565b6050831415613fe257506002905060016140ec565b6051831415613ff757506003905060016140ec565b605283141561400b575060019050806140ec565b606083141561401f575060009050806140ec565b606183141561403457506001905060006140ec565b607083141561404957506001905060006140ec565b607183141561405d575060019050806140ec565b607283141561407257506000905060016140ec565b6073831415614086575060019050806140ec565b607483141561409a575060009050806140ec565b60758314156140ae575060009050806140ec565b6040805162461bcd60e51b815260206004820152600e60248201526d496e76616c6964206f70636f646560901b604482015290519081900360640190fd5b915091565b60408051602080820193909352815180820384018152908201909152805191012090565b6000600882511115614165576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015614192578160200160208202803883390190505b50805190915060005b818110156141ee576141ab6143c2565b6141c78683815181106141ba57fe5b60200260200101516137cd565b905080600001518483815181106141da57fe5b60209081029190910101525060010161419b565b508351600360ff160182604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b8381101561423757818101518382015260200161421f565b50505050905001925050506040516020818303038152906040528051906020012092505050919050565b6000816020018351101561427457600080fd5b50016020015190565b60408051600280825260608281019093526000929190816020015b6142a061439e565b815260200190600190039081614298575050805190915060005b818110156142f2578481600281106142ce57fe5b60200201518382815181106142df57fe5b60209081029190910101526001016142ba565b50612a5782614115565b6000600c60ff8316108015612c17575050600360ff91909116101590565b6000614325826142fc565b15614335575060021981016123eb565b5060016123eb565b6040518060e001604052806143506143c2565b815260200161435d6143c2565b815260200161436a6143c2565b81526020016143776143c2565b81526020016143846143c2565b81526020016143916143c2565b8152602001600081525090565b60405180606001604052806000815260200160608152602001600060ff1681525090565b6040805160208101909152600081529056fe4f6e6520737465702070726f6f66207769746820696e76616c6964207072657620737461746550726f6f6620686164206e6f6e206d61746368696e672073746172742073746174654269736563746f7220696e636f72726563746c792072657665616c656420626973656374696f6e207365676d656e74734c6f67206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f56616c7565206d757374206861766520612076616c6964207479706520636f64654c6f676765642076616c756520646f65736e2774206d61746368206f7574707574206c6f6753656e64206e6f742063616c6c65642c20627574206d657373616765206973206e6f6e7a65726f73656e74206d65737361676520646f65736e2774206d61746368206f7574707574206d657361676550726f6f66206f6620617578706f702068616420626164206175782076616c756543616e206f6e6c79206f6e6520737465702070726f6f6620666f6c6c6f77696e6720612073696e676c652073746570206368616c6c656e6765a265627a7a7230582056729687dfa9e013654ec8d4dae80334230fe27d3733dfbd0e8bb9f2ea6b577264736f6c634300050a0032"

// DeployOneStepProof deploys a new Ethereum contract, binding an instance of OneStepProof to it.
func DeployOneStepProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OneStepProof, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	arbProtocolAddr, _, _, _ := DeployArbProtocol(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$9836fa7140e5a33041d4b827682e675a30$__", arbProtocolAddr.String()[2:], -1)

	arbValueAddr, _, _, _ := DeployArbValue(auth, backend)
	OneStepProofBin = strings.Replace(OneStepProofBin, "__$d969135829891f807aa9c34494da4ecd99$__", arbValueAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OneStepProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// OneStepProof is an auto generated Go binding around an Ethereum contract.
type OneStepProof struct {
	OneStepProofCaller     // Read-only binding to the contract
	OneStepProofTransactor // Write-only binding to the contract
	OneStepProofFilterer   // Log filterer for contract events
}

// OneStepProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type OneStepProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OneStepProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OneStepProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OneStepProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OneStepProofSession struct {
	Contract     *OneStepProof     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OneStepProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OneStepProofCallerSession struct {
	Contract *OneStepProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OneStepProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OneStepProofTransactorSession struct {
	Contract     *OneStepProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OneStepProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type OneStepProofRaw struct {
	Contract *OneStepProof // Generic contract binding to access the raw methods on
}

// OneStepProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OneStepProofCallerRaw struct {
	Contract *OneStepProofCaller // Generic read-only contract binding to access the raw methods on
}

// OneStepProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OneStepProofTransactorRaw struct {
	Contract *OneStepProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOneStepProof creates a new instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProof(address common.Address, backend bind.ContractBackend) (*OneStepProof, error) {
	contract, err := bindOneStepProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OneStepProof{OneStepProofCaller: OneStepProofCaller{contract: contract}, OneStepProofTransactor: OneStepProofTransactor{contract: contract}, OneStepProofFilterer: OneStepProofFilterer{contract: contract}}, nil
}

// NewOneStepProofCaller creates a new read-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofCaller(address common.Address, caller bind.ContractCaller) (*OneStepProofCaller, error) {
	contract, err := bindOneStepProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofCaller{contract: contract}, nil
}

// NewOneStepProofTransactor creates a new write-only instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofTransactor(address common.Address, transactor bind.ContractTransactor) (*OneStepProofTransactor, error) {
	contract, err := bindOneStepProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OneStepProofTransactor{contract: contract}, nil
}

// NewOneStepProofFilterer creates a new log filterer instance of OneStepProof, bound to a specific deployed contract.
func NewOneStepProofFilterer(address common.Address, filterer bind.ContractFilterer) (*OneStepProofFilterer, error) {
	contract, err := bindOneStepProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OneStepProofFilterer{contract: contract}, nil
}

// bindOneStepProof binds a generic wrapper to an already deployed contract.
func bindOneStepProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OneStepProofABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.OneStepProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.OneStepProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OneStepProof *OneStepProofCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OneStepProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OneStepProof *OneStepProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OneStepProof *OneStepProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OneStepProof.Contract.contract.Transact(opts, method, params...)
}