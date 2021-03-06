// AUTO GENERATED FILE (by membufc proto compiler v0.0.19)
package services

import (
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// service CrosschainConnector

type CrosschainConnector interface {
	EthereumCallContract(input *EthereumCallContractInput) (*EthereumCallContractOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput (non serializable)

type EthereumCallContractInput struct {
	BlockHeight primitives.BlockHeight
	EthereumContractAddress string
	EthereumFunctionCanonicalForm string
	InputArguments []*protocol.MethodArgument
	EthereumBlockHeight uint64
}

func (x *EthereumCallContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockHeight:%s,EthereumContractAddress:%s,EthereumFunctionCanonicalForm:%s,InputArguments:%s,EthereumBlockHeight:%s,}", x.StringBlockHeight(), x.StringEthereumContractAddress(), x.StringEthereumFunctionCanonicalForm(), x.StringInputArguments(), x.StringEthereumBlockHeight())
}

func (x *EthereumCallContractInput) StringBlockHeight() (res string) {
	res = fmt.Sprintf("%s", x.BlockHeight)
	return
}

func (x *EthereumCallContractInput) StringEthereumContractAddress() (res string) {
	res = fmt.Sprintf(x.EthereumContractAddress)
	return
}

func (x *EthereumCallContractInput) StringEthereumFunctionCanonicalForm() (res string) {
	res = fmt.Sprintf(x.EthereumFunctionCanonicalForm)
	return
}

func (x *EthereumCallContractInput) StringInputArguments() (res string) {
	res = "["
		for _, v := range x.InputArguments {
		res += v.String() + ","
  }
	res += "]"
	return
}

func (x *EthereumCallContractInput) StringEthereumBlockHeight() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockHeight)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput (non serializable)

type EthereumCallContractOutput struct {
	OutputArguments []*protocol.MethodArgument
	CallResult protocol.ExecutionResult
	EthereumBlockHeight uint64
	EthereumBlockTimestamp uint64
}

func (x *EthereumCallContractOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{OutputArguments:%s,CallResult:%s,EthereumBlockHeight:%s,EthereumBlockTimestamp:%s,}", x.StringOutputArguments(), x.StringCallResult(), x.StringEthereumBlockHeight(), x.StringEthereumBlockTimestamp())
}

func (x *EthereumCallContractOutput) StringOutputArguments() (res string) {
	res = "["
		for _, v := range x.OutputArguments {
		res += v.String() + ","
  }
	res += "]"
	return
}

func (x *EthereumCallContractOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

func (x *EthereumCallContractOutput) StringEthereumBlockHeight() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockHeight)
	return
}

func (x *EthereumCallContractOutput) StringEthereumBlockTimestamp() (res string) {
	res = fmt.Sprintf("%x", x.EthereumBlockTimestamp)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums

