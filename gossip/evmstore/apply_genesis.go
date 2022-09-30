package evmstore

import (
	"github.com/Tixcash-foundation/lachesis-base/kvdb"

	"github.com/Tixcash-foundation/go-opera/opera/genesis"
)

// ApplyGenesis writes initial state.
func (s *Store) ApplyGenesis(g genesis.Genesis) (err error) {
	batch := s.EvmDb.NewBatch()
	defer batch.Reset()
	g.RawEvmItems.ForEach(func(key, value []byte) bool {
		if err != nil {
			return false
		}
		err = batch.Put(key, value)
		if err != nil {
			return false
		}
		if batch.ValueSize() > kvdb.IdealBatchSize {
			err = batch.Write()
			if err != nil {
				return false
			}
			batch.Reset()
		}
		return true
	})
	if err != nil {
		return err
	}
	return batch.Write()
}
