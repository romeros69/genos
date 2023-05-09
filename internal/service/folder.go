package service

import (
	"fmt"
	"io/fs"
	"os"
)

const perm fs.FileMode = 0777

type FolderUC struct {
	fWork FolderSourceWorker
}

func NewFolderUC(fWork FolderSourceWorker) *FolderUC {
	return &FolderUC{fWork: fWork}
}

var _ FolderContract = (*FolderUC)(nil)

func (f *FolderUC) CreateFolder(nameProject, path string) error {
	err := f.fWork.CreateDir(nameProject, perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = os.Chdir(path)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("cmd/main", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("configs", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("docs", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("internal", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("internal/app", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("internal/controller", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("internal/controller/http", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("internal/entity", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("internal/usecase", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("internal/usecase/repo", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("pkg", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("pkg/httpserver", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	err = f.fWork.CreateDir("pkg/postgres", perm)
	if err != nil {
		return fmt.Errorf("error in CreateFolder: %w", err)
	}
	return nil
}
