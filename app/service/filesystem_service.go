package service

import (
	"errors"
	"strings"

	"github.com/bitwormhole/gie/app/web/dto"
	"github.com/bitwormhole/starter/io/fs"
	"github.com/bitwormhole/starter/lang"
	"github.com/bitwormhole/starter/markup"
	"github.com/bitwormhole/starter/util"
)

// FileSystemService ...
type FileSystemService interface {
	QueryRoots(param *dto.Dir, result *dto.Dir) error
	QueryDir(param *dto.Dir, result *dto.Dir) error
	QueryFile(param *dto.File, result *dto.File) error
}

////////////////////////////////////////////////////////////////////////////////

// FileSystemServiceImpl ...
type FileSystemServiceImpl struct {
	markup.Component `id:"filesystem-service"`

	VFS   VFSService         `inject:"#vfs-service"`
	Types ContentTypeService `inject:"#content-type-service"`
}

func (inst *FileSystemServiceImpl) init() FileSystemService {
	return inst
}

func (inst *FileSystemServiceImpl) isVirtual(uri lang.URI) bool {
	return uri.Scheme() == "vfs"
}

func (inst *FileSystemServiceImpl) resolveQueryURI(param *dto.Dir) (lang.URI, error) {

	path := param.Path
	url := param.URL
	lenURL := len(url)
	lenPath := len(path)
	if lenURL < 1 && lenPath < 1 {
		return lang.ParseURI("vfs:/")
	}

	str := ""
	if lenURL < lenPath {
		str = path
	} else {
		str = url
	}

	pb := util.PathBuilder{}
	pb.AppendPath(str)
	str = pb.String()

	if strings.HasPrefix(str, "vfs:") {
		return lang.ParseURI(str)
	} else if strings.HasPrefix(str, "file:") {
		return lang.ParseURI(str)
	}
	return lang.ParseURI("file:///" + str)
}

// QueryRoots ...
func (inst *FileSystemServiceImpl) QueryRoots(param *dto.Dir, result *dto.Dir) error {

	roots := fs.Default().Roots()

	result.Path = "roots:/"
	result.Name = "roots"
	result.URL = "roots:/"
	result.IsDir = false
	result.IsFile = false
	result.IsVirtual = true
	result.IsSymlink = false

	// make items

	src := roots
	dst := make([]*dto.DirItem, 0)
	for _, item1 := range src {
		item2 := &dto.DirItem{}
		item2.ContentType = "directory"
		item2.Exists = item1.Exists()
		item2.IsDir = item1.IsDir()
		item2.IsFile = item1.IsFile()
		item2.IsSymlink = item1.IsSymlink()
		item2.Name = item1.Path()
		item2.URL = item1.URI().String()
		dst = append(dst, item2)
	}

	result.Items = dst
	return nil
}

// QueryDir ...
func (inst *FileSystemServiceImpl) QueryDir(param *dto.Dir, result *dto.Dir) error {
	uri, err := inst.resolveQueryURI(param)
	if err != nil {
		return err
	}
	if inst.isVirtual(uri) {
		return inst.queryDirVFS(uri, result)
	}
	return inst.queryDirFS(uri, result)
}

func (inst *FileSystemServiceImpl) queryDirVFS(uri lang.URI, result *dto.Dir) error {
	return inst.VFS.LoadURI(uri, result)
}

func (inst *FileSystemServiceImpl) queryDirFS(uri lang.URI, result *dto.Dir) error {

	dir, err := fs.Default().GetPathByURI(uri)
	if err != nil {
		return err
	}
	uri = dir.URI()

	result.Path = dir.Path()
	result.Name = dir.Name()
	result.URL = uri.String()
	result.IsDir = dir.IsDir()
	result.IsFile = dir.IsFile()
	result.IsVirtual = false
	result.IsSymlink = dir.IsSymlink()

	if !result.IsDir {
		return errors.New("the path is not a dir, path=" + dir.Path())
	}

	// make items
	ctFinder := inst.Types.NewFinder()
	src := dir.ListItems()
	dst := make([]*dto.DirItem, 0)
	for _, item1 := range src {
		item2, err := inst.makeItemForDir(item1, ctFinder)
		if err == nil {
			dst = append(dst, item2)
		}
	}

	result.Items = dst
	return nil
}

func (inst *FileSystemServiceImpl) makeItemForDir(item1 fs.Path, ctf ContentTypeFinder) (*dto.DirItem, error) {

	item2 := &dto.DirItem{}
	item2.Name = item1.Name()
	item2.IsDir = item1.IsDir()
	item2.IsFile = item1.IsFile()
	item2.Exists = item1.Exists()
	item2.UpdatedAt = item1.LastModTime()
	item2.Size = item1.Size()

	if item2.IsDir {
		item2.ContentType = "directory"
	} else {
		suffix := ctf.GetFileNameSuffix(item2.Name)
		mime, err := ctf.FindContentType(suffix)
		if err == nil {
			item2.ContentType = mime
		}
	}

	return item2, nil
}

// QueryFile ...
func (inst *FileSystemServiceImpl) QueryFile(param *dto.File, result *dto.File) error {
	return nil
}
