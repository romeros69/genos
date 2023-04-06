package folders

import (
	"fmt"
	"genos/internal/service"
	"io/fs"
	"os"
)

type FolderSource struct {
}

func NewFolderSource() *FolderSource {
	return &FolderSource{}
}

var _ service.FolderSourceWorker = (*FolderSource)(nil)

func (f *FolderSource) CreateDir(path string, mode fs.FileMode) error {
	err := os.Mkdir(path, mode)
	if err != nil {
		return fmt.Errorf("error in creating directory %s: %w", path, err)
	}
	return nil
}
