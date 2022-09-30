package iblockproc

import (
	"github.com/Tixcash-foundation/lachesis-base/hash"
	"github.com/Tixcash-foundation/lachesis-base/inter/idx"
	"github.com/Tixcash-foundation/lachesis-base/inter/pos"

	"github.com/Tixcash-foundation/go-opera/inter"
	"github.com/Tixcash-foundation/go-opera/opera"
)

type ValidatorEpochStateV0 struct {
	GasRefund      uint64
	PrevEpochEvent hash.Event
}

type EpochStateV0 struct {
	Epoch          idx.Epoch
	EpochStart     inter.Timestamp
	PrevEpochStart inter.Timestamp

	EpochStateRoot hash.Hash

	Validators        *pos.Validators
	ValidatorStates   []ValidatorEpochStateV0
	ValidatorProfiles ValidatorProfiles

	Rules opera.Rules
}
