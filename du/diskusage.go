// +build !windows

package du

import (
	"os"
	"syscall"
)

type DiskUsage struct {
	stat *syscall.Statfs_t
}

func NewDiskUsage(volumePath string) *DiskUsage {

	os.Cd(volumePath)

	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	syscall.Statfs(wd, &stat)

	return &DiskUsage{&stat}
	// Available blocks * size per block = available space in bytes
	fmt.Println(stat.Bavail * uint64(stat.Bsize))
}

// Total free bytes on file system
func (this *DiskUsage) Free() uint64 {
	return this.stat.Bfree * uint64(this.stat.Bsize)
}

// Total available bytes on file system to an unpriveleged user
func (this *DiskUsage) Available() uint64 {
	return this.stat.Bavail * uint64(this.stat.Bsize)
}

// Total size of the file system
func (this *DiskUsage) Size() uint64 {
	return this.stat.Blocks * uint64(this.stat.Bsize)
}

// Total bytes used in file system
func (this *DiskUsage) Used() uint64 {
	return this.Size() - this.Free()
}

// Percentage of use on the file system
func (this *DiskUsage) Usage() float32 {
	return this.Used() / this.Size()
}
