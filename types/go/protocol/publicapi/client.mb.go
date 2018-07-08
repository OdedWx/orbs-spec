// AUTO GENERATED FILE (by membufc proto compiler v0.0.11)
package publicapi

import (
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
)

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionInput

// reader

type SendTransactionInput struct {
	message membuffers.Message
}

var _SendTransactionInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var _SendTransactionInput_Unions = [][]membuffers.FieldType{}

func SendTransactionInputReader(buf []byte) *SendTransactionInput {
	x := &SendTransactionInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _SendTransactionInput_Scheme, _SendTransactionInput_Unions)
	return x
}

func (x *SendTransactionInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendTransactionInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendTransactionInput) SignedTransaction() *protocol.SignedTransaction {
	b, s := x.message.GetMessage(0)
	return protocol.SignedTransactionReader(b[:s])
}

func (x *SendTransactionInput) RawSignedTransaction() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type SendTransactionInputBuilder struct {
	builder membuffers.Builder
	SignedTransaction *protocol.SignedTransactionBuilder
}

func (w *SendTransactionInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.SignedTransaction)
	if err != nil {
		return
	}
	return nil
}

func (w *SendTransactionInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendTransactionInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendTransactionInputBuilder) Build() *SendTransactionInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendTransactionInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message SendTransactionOutput

// reader

type SendTransactionOutput struct {
	message membuffers.Message
}

var _SendTransactionOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,}
var _SendTransactionOutput_Unions = [][]membuffers.FieldType{}

func SendTransactionOutputReader(buf []byte) *SendTransactionOutput {
	x := &SendTransactionOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _SendTransactionOutput_Scheme, _SendTransactionOutput_Unions)
	return x
}

func (x *SendTransactionOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *SendTransactionOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *SendTransactionOutput) Receipt() *protocol.TransactionReceipt {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *SendTransactionOutput) RawReceipt() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *SendTransactionOutput) TransactionStatus() protocol.TransactionStatus {
	return protocol.TransactionStatus(x.message.GetUint16(1))
}

func (x *SendTransactionOutput) RawTransactionStatus() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *SendTransactionOutput) MutateTransactionStatus(v protocol.TransactionStatus) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *SendTransactionOutput) BlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *SendTransactionOutput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *SendTransactionOutput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *SendTransactionOutput) BlockTimestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *SendTransactionOutput) RawBlockTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *SendTransactionOutput) MutateBlockTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type SendTransactionOutputBuilder struct {
	builder membuffers.Builder
	Receipt *protocol.TransactionReceiptBuilder
	TransactionStatus protocol.TransactionStatus
	BlockHeight uint64
	BlockTimestamp uint64
}

func (w *SendTransactionOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Receipt)
	if err != nil {
		return
	}
	w.builder.WriteUint16(buf, uint16(w.TransactionStatus))
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *SendTransactionOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *SendTransactionOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *SendTransactionOutputBuilder) Build() *SendTransactionOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return SendTransactionOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodInput

// reader

type CallMethodInput struct {
	message membuffers.Message
}

var _CallMethodInput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,}
var _CallMethodInput_Unions = [][]membuffers.FieldType{}

func CallMethodInputReader(buf []byte) *CallMethodInput {
	x := &CallMethodInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _CallMethodInput_Scheme, _CallMethodInput_Unions)
	return x
}

func (x *CallMethodInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CallMethodInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CallMethodInput) Transaction() *protocol.Transaction {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionReader(b[:s])
}

func (x *CallMethodInput) RawTransaction() []byte {
	return x.message.RawBufferForField(0, 0)
}

// builder

type CallMethodInputBuilder struct {
	builder membuffers.Builder
	Transaction *protocol.TransactionBuilder
}

func (w *CallMethodInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Transaction)
	if err != nil {
		return
	}
	return nil
}

func (w *CallMethodInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CallMethodInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CallMethodInputBuilder) Build() *CallMethodInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CallMethodInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message CallMethodOutput

// reader

type CallMethodOutput struct {
	message membuffers.Message
}

var _CallMethodOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray,membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,}
var _CallMethodOutput_Unions = [][]membuffers.FieldType{}

func CallMethodOutputReader(buf []byte) *CallMethodOutput {
	x := &CallMethodOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _CallMethodOutput_Scheme, _CallMethodOutput_Unions)
	return x
}

func (x *CallMethodOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *CallMethodOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *CallMethodOutput) OutputArgumentIterator() *CallMethodOutputOutputArgumentIterator {
	return &CallMethodOutputOutputArgumentIterator{iterator: x.message.GetMessageArrayIterator(0)}
}

type CallMethodOutputOutputArgumentIterator struct {
	iterator *membuffers.Iterator
}

func (i *CallMethodOutputOutputArgumentIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *CallMethodOutputOutputArgumentIterator) NextOutputArgument() *protocol.MethodArgument {
	b, s := i.iterator.NextMessage()
	return protocol.MethodArgumentReader(b[:s])
}

func (x *CallMethodOutput) RawOutputArgumentArray() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *CallMethodOutput) CallStatus() protocol.CallMethodStatus {
	return protocol.CallMethodStatus(x.message.GetUint16(1))
}

func (x *CallMethodOutput) RawCallStatus() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *CallMethodOutput) MutateCallStatus(v protocol.CallMethodStatus) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *CallMethodOutput) BlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *CallMethodOutput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *CallMethodOutput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *CallMethodOutput) BlockTimestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *CallMethodOutput) RawBlockTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *CallMethodOutput) MutateBlockTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type CallMethodOutputBuilder struct {
	builder membuffers.Builder
	OutputArgument []*protocol.MethodArgumentBuilder
	CallStatus protocol.CallMethodStatus
	BlockHeight uint64
	BlockTimestamp uint64
}

func (w *CallMethodOutputBuilder) arrayOfOutputArgument() []membuffers.MessageBuilder {
	res := make([]membuffers.MessageBuilder, len(w.OutputArgument))
	for i, v := range w.OutputArgument {
		res[i] = v
	}
	return res
}

func (w *CallMethodOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessageArray(buf, w.arrayOfOutputArgument())
	if err != nil {
		return
	}
	w.builder.WriteUint16(buf, uint16(w.CallStatus))
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *CallMethodOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *CallMethodOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *CallMethodOutputBuilder) Build() *CallMethodOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return CallMethodOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusInput

// reader

type GetTransactionStatusInput struct {
	message membuffers.Message
}

var _GetTransactionStatusInput_Scheme = []membuffers.FieldType{membuffers.TypeBytes,membuffers.TypeUint64,}
var _GetTransactionStatusInput_Unions = [][]membuffers.FieldType{}

func GetTransactionStatusInputReader(buf []byte) *GetTransactionStatusInput {
	x := &GetTransactionStatusInput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionStatusInput_Scheme, _GetTransactionStatusInput_Unions)
	return x
}

func (x *GetTransactionStatusInput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionStatusInput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionStatusInput) Txid() []byte {
	return x.message.GetBytes(0)
}

func (x *GetTransactionStatusInput) RawTxid() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusInput) MutateTxid(v []byte) error {
	return x.message.SetBytes(0, v)
}

func (x *GetTransactionStatusInput) Timestamp() uint64 {
	return x.message.GetUint64(1)
}

func (x *GetTransactionStatusInput) RawTimestamp() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetTransactionStatusInput) MutateTimestamp(v uint64) error {
	return x.message.SetUint64(1, v)
}

// builder

type GetTransactionStatusInputBuilder struct {
	builder membuffers.Builder
	Txid []byte
	Timestamp uint64
}

func (w *GetTransactionStatusInputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	w.builder.WriteBytes(buf, w.Txid)
	w.builder.WriteUint64(buf, w.Timestamp)
	return nil
}

func (w *GetTransactionStatusInputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionStatusInputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionStatusInputBuilder) Build() *GetTransactionStatusInput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionStatusInputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// message GetTransactionStatusOutput

// reader

type GetTransactionStatusOutput struct {
	message membuffers.Message
}

var _GetTransactionStatusOutput_Scheme = []membuffers.FieldType{membuffers.TypeMessage,membuffers.TypeUint16,membuffers.TypeUint64,membuffers.TypeUint64,}
var _GetTransactionStatusOutput_Unions = [][]membuffers.FieldType{}

func GetTransactionStatusOutputReader(buf []byte) *GetTransactionStatusOutput {
	x := &GetTransactionStatusOutput{}
	x.message.Init(buf, membuffers.Offset(len(buf)), _GetTransactionStatusOutput_Scheme, _GetTransactionStatusOutput_Unions)
	return x
}

func (x *GetTransactionStatusOutput) IsValid() bool {
	return x.message.IsValid()
}

func (x *GetTransactionStatusOutput) Raw() []byte {
	return x.message.RawBuffer()
}

func (x *GetTransactionStatusOutput) Receipt() *protocol.TransactionReceipt {
	b, s := x.message.GetMessage(0)
	return protocol.TransactionReceiptReader(b[:s])
}

func (x *GetTransactionStatusOutput) RawReceipt() []byte {
	return x.message.RawBufferForField(0, 0)
}

func (x *GetTransactionStatusOutput) Status() protocol.TransactionStatus {
	return protocol.TransactionStatus(x.message.GetUint16(1))
}

func (x *GetTransactionStatusOutput) RawStatus() []byte {
	return x.message.RawBufferForField(1, 0)
}

func (x *GetTransactionStatusOutput) MutateStatus(v protocol.TransactionStatus) error {
	return x.message.SetUint16(1, uint16(v))
}

func (x *GetTransactionStatusOutput) BlockHeight() uint64 {
	return x.message.GetUint64(2)
}

func (x *GetTransactionStatusOutput) RawBlockHeight() []byte {
	return x.message.RawBufferForField(2, 0)
}

func (x *GetTransactionStatusOutput) MutateBlockHeight(v uint64) error {
	return x.message.SetUint64(2, v)
}

func (x *GetTransactionStatusOutput) BlockTimestamp() uint64 {
	return x.message.GetUint64(3)
}

func (x *GetTransactionStatusOutput) RawBlockTimestamp() []byte {
	return x.message.RawBufferForField(3, 0)
}

func (x *GetTransactionStatusOutput) MutateBlockTimestamp(v uint64) error {
	return x.message.SetUint64(3, v)
}

// builder

type GetTransactionStatusOutputBuilder struct {
	builder membuffers.Builder
	Receipt *protocol.TransactionReceiptBuilder
	Status protocol.TransactionStatus
	BlockHeight uint64
	BlockTimestamp uint64
}

func (w *GetTransactionStatusOutputBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w.builder.Reset()
	err = w.builder.WriteMessage(buf, w.Receipt)
	if err != nil {
		return
	}
	w.builder.WriteUint16(buf, uint16(w.Status))
	w.builder.WriteUint64(buf, w.BlockHeight)
	w.builder.WriteUint64(buf, w.BlockTimestamp)
	return nil
}

func (w *GetTransactionStatusOutputBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w.builder.GetSize()
}

func (w *GetTransactionStatusOutputBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w.builder.GetSize()
}

func (w *GetTransactionStatusOutputBuilder) Build() *GetTransactionStatusOutput {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return GetTransactionStatusOutputReader(buf)
}

/////////////////////////////////////////////////////////////////////////////
// enums

