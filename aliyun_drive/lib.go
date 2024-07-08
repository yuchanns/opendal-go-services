package aliyun_drive

import (
	"fmt"

	services "github.com/yuchanns/opendal-go-services"
)

var Scheme = scheme{}

type scheme struct{}

func (scheme) Scheme() string {
	return "aliyun_drive"
}

func (scheme) Path() string {
	return path
}

var path string

func init() {
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
		panic(fmt.Errorf("opendal/services/aliyun_drive: %s", err))
	}
}
