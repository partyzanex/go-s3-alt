package uploader

import (
	"context"
	"github.com/partyzanex/go-s3-alt/internal/domains/file"
)

type Temporary struct {
}

func (u *Temporary) UploadFile(ctx context.Context, params UploadParams) (*file.Info, error) {
	return nil, nil
}

func (u *Temporary) Clean(ctx context.Context, fileIDs ...file.ID) error {
	return nil
}
