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
	if s.FindEntry(entry.Name, entry.Extracted).Extracted != "" {
		return nil
	}
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

func (s *Storage) FindEntry(name string, extracted string) *Entry {
	var entry Entry
	s.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("entries"))
		b.ForEach(func(k, v []byte) error {
			err := yaml.Unmarshal(v, &entry)
			if err != nil {
				return err
			}
			if entry.Name == name && entry.Extracted == extracted {
				return nil
			}
			entry = Entry{}
			return nil
		})
		return nil
	})
	return &entry
}

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, v)
	return b
}
