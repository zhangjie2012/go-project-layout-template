package cache

import "github.com/zhangjie2012/go-project-layout-template/cmd/options"

type Cache struct{}

func NewCache(opt *options.AppOption) (*Cache, error) {
	return &Cache{}, nil
}

func (c *Cache) Close() error {
	return nil
}
