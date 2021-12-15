package du

import "syscall"

// DiskUsage contains usage data and provides user-friendly access methods
type DiskUsage struct {
	stat *syscall.Statfs_t
}

// NewDiskUsages returns an object holding the disk usage of volumePath
// or nil in case of error (invalid path, etc)
func NewDiskUsage(volumePath string) *DiskUsage {

	var stat syscall.Statfs_t
	syscall.Statfs(volumePath, &stat)
	return &DiskUsage{&stat}
}

// Free returns total free bytes on file system
func (du *DiskUsage) Free() uint64 {
	return du.stat.F_bfree * uint64(du.stat.F_bsize)
}

// Available return total available bytes on file system to an unprivileged user
func (du *DiskUsage) Available() uint64 {
	return uint64(du.stat.F_bavail) * uint64(du.stat.F_bsize)
}

// Size returns total size of the file system
func (du *DiskUsage) Size() uint64 {
	return du.stat.F_blocks * uint64(du.stat.F_bsize)
}

// Used returns total bytes used in file system
func (du *DiskUsage) Used() uint64 {
	return du.Size() - du.Free()
}

// Usage returns percentage of use on the file system
func (du *DiskUsage) Usage() float32 {
	return float32(du.Used()) / float32(du.Size())
}
