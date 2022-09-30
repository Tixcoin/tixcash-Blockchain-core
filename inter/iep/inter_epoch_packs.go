package iep

import (
	"github.com/Tixcash-foundation/go-opera/inter"
	"github.com/Tixcash-foundation/go-opera/inter/ier"
)

type LlrEpochPack struct {
	Votes  []inter.LlrSignedEpochVote
	Record ier.LlrIdxFullEpochRecord
}
