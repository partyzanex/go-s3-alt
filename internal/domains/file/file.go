package file

import (
	"github.com/google/uuid"
	"github.com/partyzanex/go-s3-alt/internal/domains/node"
	"io"
	"time"
)

type (
	ID     = uuid.UUID
	PartID = uuid.UUID
)

type Info struct {
	ID         ID
	Name       string
	Size       int64
	Parts      []PartInfo
	UploadedAt time.Time
}

type PartInfo struct {
	ID     PartID
	Size   int64
	Order  uint
	NodeID node.ID
}

type Part struct {
	Info PartInfo
	Body io.ReadCloser
}

type Status int8
