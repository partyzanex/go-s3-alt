package storage

import (
	"context"
	"github.com/partyzanex/go-s3-alt/internal/domains/file"
	"io"
)

type FileRepository interface {
	GetStatus(ctx context.Context, fileID file.ID) (file.Status, error)
	GetFileInfo(ctx context.Context, fileID file.ID) (*file.Info, error)
}

type Downloader interface {
	DownloadParts(ctx context.Context, fileID file.ID, partsInfo []file.PartInfo) (io.ReadCloser, error)
}

type Service struct {
	repository FileRepository
	downloader Downloader
}

func (s *Service) GetFileStatus(ctx context.Context, fileID file.ID) (file.Status, error) {
	return s.repository.GetStatus(ctx, fileID)
}

func (s *Service) DownloadFile(ctx context.Context, fileID file.ID) (*file.Info, io.ReadCloser, error) {
	fileInfo, err := s.repository.GetFileInfo(ctx, fileID)
	if err != nil {
		return nil, nil, err
	}

	body, err := s.downloader.DownloadParts(ctx, fileID, fileInfo.Parts)
	if err != nil {
		return nil, nil, err
	}

	return fileInfo, body, nil
}
