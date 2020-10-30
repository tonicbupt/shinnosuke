package compress

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"io/ioutil"
	"os"

	"github.com/juju/errors"
)

type JPEGCompressor struct {
	i image.Image
	b *bytes.Buffer
}

func NewJPEGCompressor() *JPEGCompressor {
	return &JPEGCompressor{}
}

func (j *JPEGCompressor) Load(filename string) error {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Annotatef(err, "failed to open file %s", filename)
	}

	image, err := jpeg.Decode(bytes.NewBuffer(bs))
	if err != nil {
		return errors.Annotatef(err, "failed to load image %s", filename)
	}

	j.i = image
	j.b = bytes.NewBuffer(bs)
	return nil
}

func (j *JPEGCompressor) Compress(size int64) error {
	q := 95
	for int64(j.b.Len()) > size && q > 0 {
		if err := j.compress(q); err != nil {
			return errors.Annotate(err, "failed to compress")
		}
		q -= 5
	}
	return nil
}

func (j *JPEGCompressor) compress(quality int) error {
	j.b.Reset()
	return jpeg.Encode(j.b, j.i, &jpeg.Options{Quality: quality})
}

func (j *JPEGCompressor) Dump(filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return errors.Annotatef(err, "failed to create file %s", filename)
	}
	defer out.Close()

	_, err = io.Copy(out, j.b)
	return errors.Annotatef(err, "failed to dump file %s", filename)
}
