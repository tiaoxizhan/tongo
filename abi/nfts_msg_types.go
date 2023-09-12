package abi

// Code autogenerated. DO NOT EDIT.

import (
	"errors"
	"github.com/tonkeeper/tongo/boc"
	"github.com/tonkeeper/tongo/tlb"
)

func (j *NFTPayload) UnmarshalTLB(cell *boc.Cell, decoder *tlb.Decoder) error {
	if cell.BitsAvailableForRead() == 0 && cell.RefsAvailableForRead() == 0 {
		return nil
	}
	tempCell := cell.CopyRemaining()
	op64, err := tempCell.ReadUint(32)
	if errors.Is(err, boc.ErrNotEnoughBits) {
		j.SumType = UnknownNFTOp
		j.Value = cell.CopyRemaining()
		return nil
	}
	op := uint32(op64)
	j.OpCode = &op
	switch op {

	}
	j.SumType = UnknownNFTOp
	j.Value = cell.CopyRemaining()

	return nil
}

const ()

var KnownNFTTypes = map[string]any{}
var nftOpCodes = map[NFTOpName]NFTOpCode{}
