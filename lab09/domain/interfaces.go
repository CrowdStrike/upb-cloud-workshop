package domain

import "context"

type FilesMetadataStorage interface {
	SaveFile(ctx context.Context, metadata *FileMetadata) error
	RetrieveFile(ctx context.Context, sha256 string) (*FileMetadata, error)
}
