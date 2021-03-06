package chezmoi

import (
	"os"
	"os/exec"

	vfs "github.com/twpayne/go-vfs"
)

// A ReadOnlySystem is a system that may only be read from.
type ReadOnlySystem struct {
	noUpdateSystemMixin
	system System
}

// NewReadOnlySystem returns a new ReadOnlySystem that wraps system.
func NewReadOnlySystem(system System) *ReadOnlySystem {
	return &ReadOnlySystem{
		system: system,
	}
}

// Glob implements System.Glob.
func (s *ReadOnlySystem) Glob(pattern string) ([]string, error) {
	return s.system.Glob(pattern)
}

// IdempotentCmdOutput implements System.IdempotentCmdOutput.
func (s *ReadOnlySystem) IdempotentCmdOutput(cmd *exec.Cmd) ([]byte, error) {
	return s.system.IdempotentCmdOutput(cmd)
}

// Lstat implements System.Lstat.
func (s *ReadOnlySystem) Lstat(filename AbsPath) (os.FileInfo, error) {
	return s.system.Lstat(filename)
}

// RawPath implements System.RawPath.
func (s *ReadOnlySystem) RawPath(path AbsPath) (AbsPath, error) {
	return s.system.RawPath(path)
}

// ReadDir implements System.ReadDir.
func (s *ReadOnlySystem) ReadDir(dirname AbsPath) ([]os.FileInfo, error) {
	return s.system.ReadDir(dirname)
}

// ReadFile implements System.ReadFile.
func (s *ReadOnlySystem) ReadFile(filename AbsPath) ([]byte, error) {
	return s.system.ReadFile(filename)
}

// Readlink implements System.Readlink.
func (s *ReadOnlySystem) Readlink(name AbsPath) (string, error) {
	return s.system.Readlink(name)
}

// Stat implements System.Stat.
func (s *ReadOnlySystem) Stat(name AbsPath) (os.FileInfo, error) {
	return s.system.Stat(name)
}

// UnderlyingFS implements System.UnderlyingFS.
func (s *ReadOnlySystem) UnderlyingFS() vfs.FS {
	return s.system.UnderlyingFS()
}
