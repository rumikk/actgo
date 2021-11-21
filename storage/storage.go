package storage

import (
	"encoding/binary"
	bolt "go.etcd.io/bbolt"
	"gopkg.in/yaml.v3"
	"time"
)

type Storage struct {
	Db *bolt.DB
}

type Entry struct {
	CreatedAt int64  `yaml:"created_at"`
	Name      string `yaml:"name"`
	Selected  string `yaml:"selected"`
	Extracted string `yaml:"extracted"`
}

func (s *Storage) Init() error {
	return s.Db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("entries"))
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *Storage) AddEntry(entry *Entry) error {
	return s.Db.Update(func(tx *bolt.Tx) error {
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

func (s *Storage) FindEntry(extracted string) *Entry {
	var result Entry
	s.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		b.ForEach(func(k, v []byte) error {
			if result.CreatedAt != 0 {
				return nil
			}
			var entry Entry
			err := yaml.Unmarshal(v, &entry)
			if err != nil {
				return err
			}
			if entry.Extracted == extracted {
				result = entry
				return nil
			}
			entry = Entry{}
			return nil
		})
		return nil
	})
	return &result
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
