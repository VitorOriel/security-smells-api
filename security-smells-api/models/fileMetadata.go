package models

type FileMetadata[T comparable] struct {
	Filename                      string `json:"filename"`
	WorkloadPositionMapByManifest map[T]int
}

func NewFileMetadata[T comparable]() *FileMetadata[T] {
	return &FileMetadata[T]{WorkloadPositionMapByManifest: make(map[T]int)}
}
