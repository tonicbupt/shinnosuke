package compress

import (
	"github.com/juju/errors"
)

type Compressor interface {
	Load(filename string) error
	Compress(size int64) error
	Dump(filename string) error
}

func GetCompressor(ext string) (Compressor, error) {
	switch ext {
	case ".jpeg", ".jpg", ".JPEG", ".JPG":
		return NewJPEGCompressor(), nil
	case ".png", ".PNG":
		return NewPNGCompressor(), nil
	default:
		return nil, errors.Errorf("Can't find compressor for %s", ext)
	}
}
