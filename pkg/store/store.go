package store

import "github.com/zhangjie2012/go-project-layout-template/cmd/options"

type Store struct{}

func NewStore(config *options.AppOption) (*Store, error) {
	return &Store{}, nil
}

func (s *Store) Close() error {
	return nil
}
