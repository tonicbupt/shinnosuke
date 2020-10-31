package job

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/tonicbupt/shinnosuke/pkg/compress"
)

type worker struct {
	ch   chan string
	size int64
	wg   *sync.WaitGroup
}

func (w *worker) work() {
	defer w.wg.Done()
	for input := range w.ch {
		logrus.Infof("get input %s", input)
		ext := filepath.Ext(input)
		compressor, err := compress.GetCompressor(ext)
		if err != nil {
			continue
		}

		if err := compressor.Load(input); err != nil {
			logrus.Errorf("Error loading file %s, %s", input, err)
			return
		}

		if err := compressor.Compress(w.size); err != nil {
			logrus.Errorf("Error compressing file %s, %s", input, err)
			return
		}

		out := fmt.Sprintf("%s.compressed%s", input, ext)
		if err := compressor.Dump(out); err != nil {
			logrus.Errorf("Error dumping file %s, %s", input, err)
			return
		}
	}
}

type CompressJob struct {
	dir  string
	max  int
	size int64
}

func NewCompressJob(dir string, max int, size int64) *CompressJob {
	return &CompressJob{
		dir:  dir,
		max:  max,
		size: size,
	}
}

func (c *CompressJob) walkfiles() chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		filepath.Walk(c.dir, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				logrus.Infof("filename %s", info.Name())
				out <- info.Name()
			}
			return nil
		})
	}()
	return out
}

func (c *CompressJob) Do() {
	filenames := c.walkfiles()
	wg := &sync.WaitGroup{}

	for i := 0; i < c.max; i++ {
		wg.Add(1)
		w := &worker{
			ch:   filenames,
			size: c.size,
			wg:   wg,
		}
		go w.work()
	}

	wg.Wait()
}
