package ier

import (
	"github.com/Tixcash-foundation/lachesis-base/hash"
	"github.com/Tixcash-foundation/lachesis-base/inter/idx"

	"github.com/Tixcash-foundation/go-opera/inter/iblockproc"
)

type LlrFullEpochRecord struct {
	BlockState iblockproc.BlockState
	EpochState iblockproc.EpochState
}

type LlrIdxFullEpochRecord struct {
	LlrFullEpochRecord
	Idx idx.Epoch
}

func (er LlrFullEpochRecord) Hash() hash.Hash {
	return hash.Of(er.BlockState.Hash().Bytes(), er.EpochState.Hash().Bytes())
}
