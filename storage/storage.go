package storage

import "actgo/parser"

type Storage struct{}

type Entry struct {
	CreatedAt int64
	Process   parser.Process
}
