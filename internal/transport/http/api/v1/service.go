package v1

import (
	"context"
	"github.com/partyzanex/go-s3-alt/internal/domains/file"
	v1 "github.com/partyzanex/go-s3-alt/internal/transport/http/api/v1/gen"
	"github.com/partyzanex/go-s3-alt/internal/uploader"
	"io"
)

type Uploader interface {
	UploadFile(ctx context.Context, params uploader.UploadParams) (*file.Info, error)
	Clean(ctx context.Context, fileIDs ...file.ID) error
}

type Storage interface {
	GetFileStatus(ctx context.Context, fileID file.ID) (file.Status, error)
	DownloadFile(ctx context.Context, fileID file.ID) (*file.Info, io.ReadCloser, error)
}

type Service struct {
	v1.Handler

	uploader Uploader
	storage  Storage
}
