package storage

import (
	"actgo/parser"
	"encoding/binary"
	bolt "go.etcd.io/bbolt"
	"gopkg.in/yaml.v3"
	"time"
)

type Storage struct{}

type Entry struct {
	CreatedAt int64          `yaml:"created_at"`
	Process   parser.Process `yaml:"process"`
}

func Init(db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("entries"))
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *Storage) AddEntry(entry *Entry, db *bolt.DB) error {
	return db.Update(func(tx *bolt.Tx) error {
		entry.CreatedAt = time.Now().Unix()
		b := tx.Bucket([]byte("entries"))
		id, err := b.NextSequence()
		if err != nil {
			return err
		}
		entryYaml, err := yaml.Marshal(entry)
		if err != nil {
			return err
		}
		return b.Put(itob(id), entryYaml)
	})
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
