package concurrent

import (
	"sync"

	"github.com/Tixcash-foundation/lachesis-base/hash"
	"github.com/Tixcash-foundation/lachesis-base/inter/idx"
)

type ValidatorEventsSet struct {
	sync.RWMutex
	Val map[idx.ValidatorID]hash.Event
}

func WrapValidatorEventsSet(v map[idx.ValidatorID]hash.Event) *ValidatorEventsSet {
	return &ValidatorEventsSet{
		RWMutex: sync.RWMutex{},
		Val:     v,
	}
}
