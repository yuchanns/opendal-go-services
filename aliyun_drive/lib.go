package aliyun_drive

import (
	"fmt"
	"sync"

	services "github.com/yuchanns/opendal-go-services"
)

var Scheme = scheme{once: &sync.Once{}}

type scheme struct {
	once *sync.Once
}

func (scheme) Scheme() string {
	return "aliyun_drive"
}

func (s scheme) Path() (string, error) {
	var err error
	s.once.Do(func() {
		err = load()
	})
	return path, err
}

var path string

func load() error {
	var (
		err  error
		data []byte
	)
	var loop = true
	for loop {
		loop = false

		data, err = services.DecompressLib(libopendalZst)
		if err != nil {
			break
		}
		path, err = services.WriteTempExec("libopendal_c*.so", data)
		if err != nil {
			break
		}
	}
	if err != nil {
		return fmt.Errorf("opendal/services/aliyun_drive: %s", err)
	}
	return nil
}
