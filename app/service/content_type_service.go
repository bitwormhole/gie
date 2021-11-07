package service

import (
	"errors"
	"strings"

	"github.com/bitwormhole/starter/application"
	"github.com/bitwormhole/starter/collection"
	"github.com/bitwormhole/starter/markup"
)

type ContentTypeFinder interface {
	// return '.xxx'
	GetFileNameSuffix(filename string) string

	FindContentType(suffix string) (string, error)
}

type ContentTypeService interface {
	NewFinder() ContentTypeFinder
}

////////////////////////////////////////////////////////////////////////////////

type ContentTypeServiceImpl struct {
	markup.Component `id:"content-type-service" initMethod:"Init"`

	Context application.Context `inject:"context"`

	innerFinder *InnerMimeTypesFinder
}

func (inst *ContentTypeServiceImpl) _Impl() ContentTypeService {
	return inst
}

func (inst *ContentTypeServiceImpl) Init() error {
	finder := &InnerMimeTypesFinder{}
	err := finder.init(inst)
	if err != nil {
		return err
	}
	inst.innerFinder = finder
	return nil
}

func (inst *ContentTypeServiceImpl) NewFinder() ContentTypeFinder {
	finder := &contentTypeFinderImpl{}
	return finder.init(inst)
}

////////////////////////////////////////////////////////////////////////////////

type contentTypeFinderImpl struct {
	service       *ContentTypeServiceImpl
	mimeTypeCache map[string]string // map[suffix] mineType
}

func (inst *contentTypeFinderImpl) init(service *ContentTypeServiceImpl) ContentTypeFinder {
	inst.service = service
	return inst
}

func (inst *contentTypeFinderImpl) GetFileNameSuffix(filename string) string {
	suffix := ""
	index := strings.LastIndexByte(filename, '.')
	if index < 0 {
		suffix = filename
	} else {
		suffix = filename[index:]
	}
	if strings.ContainsAny(suffix, "/\\") {
		suffix = strings.ReplaceAll(suffix, "\\", "/")
		i2 := strings.LastIndex(suffix, "/")
		suffix = suffix[i2+1:]
	}
	suffix = strings.ToLower(suffix)
	return suffix
}

func (inst *contentTypeFinderImpl) FindContentType(suffix string) (string, error) {

	cache := inst.getCache()
	mime := cache[suffix]
	if mime != "" {
		return mime, nil
	}

	mime, err := inst.find(suffix)
	if err != nil {
		return "", err
	}

	cache[suffix] = mime
	return mime, nil
}

func (inst *contentTypeFinderImpl) getCache() map[string]string {
	cache := inst.mimeTypeCache
	if cache == nil {
		cache = make(map[string]string)
		inst.mimeTypeCache = cache
	}
	return cache
}

func (inst *contentTypeFinderImpl) find(suffix string) (string, error) {
	t, err := inst.findInner(suffix)
	if err == nil {
		return t, nil
	}
	return inst.findExtends(suffix)
}

func (inst *contentTypeFinderImpl) findInner(suffix string) (string, error) {
	finder := inst.service.innerFinder
	return finder.findTypeNameBySuffix(suffix)
}

func (inst *contentTypeFinderImpl) findExtends(suffix string) (string, error) {
	return "", errors.New("todo: no impl")
}

////////////////////////////////////////////////////////////////////////////////

type InnerMimeTypesFinder struct {
	props collection.Properties
}

func (inst *InnerMimeTypesFinder) init(service *ContentTypeServiceImpl) error {

	const resfile = "content-types.properties"

	text, err := service.Context.GetResources().GetText(resfile)
	if err != nil {
		return err
	}

	props, err := collection.ParseProperties(text, nil)
	if err != nil {
		return err
	}

	inst.props = props
	return nil
}

func (inst *InnerMimeTypesFinder) findTypeNameBySuffix(suffix string) (string, error) {
	key := "type" + suffix + ".name"
	value, err := inst.props.GetPropertyRequired(key)
	if err != nil {
		return "", errors.New("no type info for file-name-suffix: " + suffix)
	}
	return value, nil
}
