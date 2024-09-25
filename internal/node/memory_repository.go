package node

import (
	"context"
	ttlcache "github.com/jellydator/ttlcache/v3"
	"github.com/partyzanex/go-s3-alt/internal/domains/node"
	"time"
)

type MemoryRepository struct {
	cache *ttlcache.Cache[node.ID, *node.Node]
}

func NewMemoryRepository(ttl time.Duration) *MemoryRepository {
	return &MemoryRepository{
		cache: ttlcache.New(
			ttlcache.WithTTL[node.ID, *node.Node](ttl),
		),
	}
}

func (r *MemoryRepository) GetNodes(_ context.Context) ([]*node.Node, error) {
	nodes := make([]*node.Node, r.cache.Len())
	i := 0

	r.cache.Range(func(item *ttlcache.Item[node.ID, *node.Node]) bool {
		nodes[i] = item.Value()
		i++

		return true
	})

	return nodes, nil
}

func (r *MemoryRepository) SetNode(_ context.Context, nodeInfo *node.Node) error {
	r.cache.Set(nodeInfo.ID, nodeInfo, 0)

	return nil
}
