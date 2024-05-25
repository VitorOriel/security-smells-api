package models

type FileMetadata[T comparable] struct {
	Filename                      string `json:"filename"`
	WorkloadPositionMapByManifest map[T][]int
}

func NewFileMetadata[T comparable]() *FileMetadata[T] {
	return &FileMetadata[T]{WorkloadPositionMapByManifest: make(map[T][]int)}
}

func (f *FileMetadata[T]) AppendWorkloadPosition(workload T, position int) {
	f.WorkloadPositionMapByManifest[workload] = append(f.WorkloadPositionMapByManifest[workload], position)
}
