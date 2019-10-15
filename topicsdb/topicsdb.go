package topicsdb

import (
	"github.com/Fantom-foundation/go-lachesis/common/bigendian"
	"github.com/Fantom-foundation/go-lachesis/kvdb"
	"github.com/Fantom-foundation/go-lachesis/kvdb/table"
)

type TopicsDb struct {
	db    kvdb.KeyValueStore
	table struct {
		// topic+topicN+blockN+recordID -> pair_count
		Topic kvdb.KeyValueStore `table:"topic"`
		// recordID+N -> topic, data
		Record kvdb.KeyValueStore `table:"record"`
	}

	fetchMethod func(cc ...Condition) (res []*Logrec, err error)
}

func New(db kvdb.KeyValueStore) *TopicsDb {
	tt := &TopicsDb{
		db: db,
	}

	tt.fetchMethod = tt.fetchAsync

	table.MigrateTables(&tt.table, tt.db)

	return tt
}

func (tt *TopicsDb) Find(cc ...Condition) (res []*Logrec, err error) {
	return tt.fetchMethod(cc...)
}

func (tt *TopicsDb) Push(rec *Logrec) error {
	count := bigendian.Int32ToBytes(uint32(len(
		rec.Topics)))

	for n, topic := range rec.Topics {
		key := topicKey(topic, n, rec.BlockN, rec.Id)
		err := tt.table.Topic.Put(key, count)
		if err != nil {
			return err
		}
	}

	for n, topic := range rec.Topics {
		key := logrecKey(rec, n)
		err := tt.table.Record.Put(key, topic.Bytes())
		if err != nil {
			return err
		}
	}

	return nil
}
