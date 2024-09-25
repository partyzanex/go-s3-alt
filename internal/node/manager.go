package node

import (
	"context"
	"github.com/partyzanex/go-s3-alt/internal/domains/file"
	"github.com/partyzanex/go-s3-alt/internal/domains/node"
	"io"
)

type Repository interface {
	GetNodes(ctx context.Context) ([]*node.Node, error)
	SetNode(ctx context.Context, node *node.Node) error
}

type Manager struct {
	repository Repository
}

func (m *Manager) GetAvailableNodes(ctx context.Context) ([]*node.Node, error) {
	nodes, err := m.repository.GetNodes(ctx)
	if err != nil {
		return nil, err
	}

	available := make([]*node.Node, 0, len(nodes))

	for _, nodeInfo := range nodes {
		if nodeInfo.Online {
			available = append(available, nodeInfo)
		}
	}

	return available, nil
}

func (m *Manager) RegisterNode(ctx context.Context, node *node.Node) error {
	return nil
}

func (m *Manager) Upload(ctx context.Context, nodeID node.ID, r io.ReadCloser) error {
	return nil
}

func (m *Manager) Download(ctx context.Context, nodeID node.ID, partID file.PartID) (io.ReadCloser, error) {
	return nil, nil
}
