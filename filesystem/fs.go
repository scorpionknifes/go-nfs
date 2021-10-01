package filesystem

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

var ErrUnsupported = fmt.Errorf("unsupported operation")

func Join(fs fs.FS, elem ...string) string {
	if FS, ok := fs.(JoinFS); ok {
		return FS.Join(elem...)
	}

	return filepath.Join(elem...)
}

func ReadDir(fs fs.FS, name string) ([]fs.DirEntry, error) {
	if FS, ok := fs.(ReadDirFS); ok {
		return FS.ReadDir(name)
	}

	return nil, fmt.Errorf("readdir %s: operation not supported", name)
}

func MkdirAll(fs fs.FS, path string, perm fs.FileMode) error {
	if FS, ok := fs.(MkdirAllFS); ok {
		return FS.MkdirAll(path, perm)
	}

	return fmt.Errorf("mkdirall %s: operation not supported", path)
}

func Stat(fs fs.FS, filename string) (fs.FileInfo, error) {
	f, err := fs.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Open: %w", err)
	}

	return f.Stat()
}

func Create(fs fs.FS, name string) (fs.File, error) {
	if FS, ok := fs.(CreateFS); ok {
		return FS.Create(name)
	}

	return nil, fmt.Errorf("create %s: operation not supported", name)
}

func Lstat(fs fs.FS, name string) (fs.File, error) {
	if FS, ok := fs.(CreateFS); ok {
		return FS.Create(name)
	}

	return nil, fmt.Errorf("create %s: operation not supported", name)
}

func Rename(fs fs.FS, oldpath, newpath string) error {
	if FS, ok := fs.(RenameFS); ok {
		return FS.Rename(oldpath, newpath)
	}

	return fmt.Errorf("rename %s to %s: operation not supported", oldpath, newpath)
}

func OpenFile(fs fs.FS, name string, flag int, perm fs.FileMode) (fs.File, error) {
	if FS, ok := fs.(OpenFileFS); ok {
		return FS.OpenFile(name, flag, perm)
	}

	return nil, fmt.Errorf("openfile %s: operation not supported", name)
}

func Symlink(fs fs.FS, oldname string, newname string) error {
	if FS, ok := fs.(SymlinkFS); ok {
		return FS.Symlink(oldname, newname)
	}

	return fmt.Errorf("symlink %s to %s: operation not supported", oldname, newname)
}

func Readlink(fs fs.FS, name string) (string, error) {
	if FS, ok := fs.(ReadlinkFS); ok {
		return FS.Readlink(name)
	}

	return "", fmt.Errorf("readlink %s: operation not supported", name)
}

func Remove(fs fs.FS, name string) error {
	if FS, ok := fs.(RemoveFS); ok {
		return FS.Remove(name)
	}

	return fmt.Errorf("remove %s: operation not supported", name)
}

func WriteCapabilityCheck(fs fs.FS) bool {
	_, ok := fs.(WriteFileFS)
	return ok
}
