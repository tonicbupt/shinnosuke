package compress

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"io/ioutil"
	"os"

	"github.com/juju/errors"
)

type PNGCompressor struct {
	i image.Image
	b *bytes.Buffer
}

func NewPNGCompressor() *PNGCompressor {
	return &PNGCompressor{}
}

func (p *PNGCompressor) Load(filename string) error {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Annotatef(err, "failed to open file %s", filename)
	}

	image, err := png.Decode(bytes.NewBuffer(bs))
	if err != nil {
		return errors.Annotatef(err, "failed to load image %s", filename)
	}

	p.i = image
	p.b = bytes.NewBuffer(bs)
	return nil
}

func (p *PNGCompressor) Compress(size int64) error {
	q := 10
	for int64(p.b.Len()) > size && q > 0 {
		if err := p.compress(); err != nil {
			return errors.Annotate(err, "failed to compress")
		}
		q--
	}
	return nil
}

func (p *PNGCompressor) compress() error {
	p.b.Reset()
	encoder := &png.Encoder{CompressionLevel: png.BestCompression}
	return encoder.Encode(p.b, p.i)
}

func (p *PNGCompressor) Dump(filename string) error {
	out, err := os.Create(filename)
	if err != nil {
		return errors.Annotatef(err, "failed to create file %s", filename)
	}
	defer out.Close()

	_, err = io.Copy(out, p.b)
	return errors.Annotatef(err, "failed to dump file %s", filename)
}
