package filesystem

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

func NewDirFSWrapper(root string) fs.FS {
	return &dirFS{os.DirFS(root), root}
}

type dirFS struct {
	fs.FS
	root string
}

func (dfs *dirFS) Open(name string) (fs.File, error) {
	file, err := dfs.FS.Open(name)
	if err != nil {
		return nil, err
	}

	return NewDirFileWrapper(dfs, file, dfs.root)
}

func (dfs *dirFS) Join(elem ...string) string {
	return filepath.Join(elem...)
}

func (dfs *dirFS) Stat(name string) (os.FileInfo, error) {
	return os.Stat(dfs.Join(dfs.root, name))
}

func (dfs *dirFS) ReadDir(name string) ([]fs.DirEntry, error) {
	return os.ReadDir(dfs.Join(dfs.root, name))
}

func (dfs *dirFS) WriteFile(name string, data []byte, mode fs.FileMode) error {
	return os.WriteFile(dfs.Join(dfs.root, name), data, mode)
}

func (dfs *dirFS) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(dfs.Join(dfs.root, path), perm)
}

func (dfs *dirFS) Create(name string) (fs.File, error) {
	file, err := os.Create(dfs.Join(dfs.root, name))
	if err != nil {
		return nil, err
	}

	return NewDirFileWrapper(dfs, file, dfs.root)
}

func (dfs *dirFS) OpenFile(name string, flag int, perm fs.FileMode) (fs.File, error) {
	file, err := os.OpenFile(dfs.Join(dfs.root, name), flag, perm)
	if err != nil {
		return nil, err
	}

	return NewDirFileWrapper(dfs, file, dfs.root)
}

func (dfs *dirFS) Chmod(name string, mode fs.FileMode) error {
	return os.Chmod(dfs.Join(dfs.root, name), mode)
}

func (dfs *dirFS) Lchown(name string, uid int, gid int) error {
	return os.Lchown(dfs.Join(dfs.root, name), uid, gid)
}

func (dfs *dirFS) Chown(name string, uid int, gid int) error {
	return os.Chown(dfs.Join(dfs.root, name), uid, gid)
}

func (dfs *dirFS) Chtimes(name string, atime, mtime time.Time) error {
	return os.Chtimes(dfs.Join(dfs.root, name), atime, mtime)
}

func (dfs *dirFS) Readlink(name string) (string, error) {
	return os.Readlink(dfs.Join(dfs.root, name))
}

func (dfs *dirFS) Remove(name string) error {
	return os.Remove(dfs.Join(dfs.root, name))
}

func (dfs *dirFS) Rename(oldname, newname string) error {
	return os.Rename(dfs.Join(dfs.root, oldname), dfs.Join(dfs.root, newname))
}

func (dfs *dirFS) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(dfs.Join(dfs.root, name))
}

func (dfs *dirFS) Symlink(oldname, newname string) error {
	return os.Symlink(oldname, dfs.Join(dfs.root, newname))
}

func (dfs *dirFS) RemoveAll(path string) error {
	return os.RemoveAll(dfs.Join(dfs.root, path))
}

// Create new file wrapper with write functionality
func NewDirFileWrapper(fs *dirFS, file fs.File, root string) (fs.File, error) {
	if osFile, ok := file.(*os.File); ok {
		return &dirFile{osFile, fs.Join(root, osFile.Name())}, nil
	}

	return nil, errors.New("Not an os.File")
}

type dirFile struct {
	fs.File
	path string
}

func (df *dirFile) ReadAt(b []byte, off int64) (int, error) {
	if osFile, ok := df.File.(*os.File); ok {
		return osFile.ReadAt(b, off)
	}
	return 0, errors.New("Not an os.File")
}

func (df *dirFile) Seek(offset int64, whence int) (int64, error) {
	if osFile, ok := df.File.(*os.File); ok {
		return osFile.Seek(offset, whence)
	}
	return 0, errors.New("Not an os.File")
}

func (df *dirFile) WriteAt(p []byte, off int64) (int, error) {
	if osFile, ok := df.File.(*os.File); ok {
		return osFile.WriteAt(p, off)
	}
	return 0, errors.New("Not an os.File")
}

func (df *dirFile) Truncate(size int64) error {
	if osFile, ok := df.File.(*os.File); ok {
		return osFile.Truncate(size)
	}
	return errors.New("Not an os.File")
}

func (df *dirFile) Write(b []byte) (int, error) {
	if osFile, ok := df.File.(*os.File); ok {
		return osFile.Write(b)
	}
	return 0, errors.New("Not an os.File")
}
