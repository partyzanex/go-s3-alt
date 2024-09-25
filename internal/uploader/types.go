package uploader

import (
	"github.com/partyzanex/go-s3-alt/internal/domains/file"
	"io"
)

type UploadParams struct {
	Body     io.Reader
	FileID   file.ID
	FileName string
}
