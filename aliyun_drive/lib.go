package aliyun_drive

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/klauspost/compress/zstd"
)

type Scheme struct{}

func (Scheme) Scheme() string {
	return "aliyun_drive"
}

func (Scheme) Path() string {
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

		data, err = decompressLib(libopendalZst)
		if err != nil {
			break
		}
		path, err = writeTempExec("libopendal_c*.so", data)
		if err != nil {
			break
		}
	}
	if err != nil {
		panic(fmt.Errorf("opendal/services/aliyun_drive: %s", err))
	}
}

func writeTempExec(pattern string, binary []byte) (path string, err error) {
	f, err := os.CreateTemp("", pattern)
	if err != nil {
		err = fmt.Errorf("cannot create temp file: %s", err)
		return
	}
	defer f.Close()
	_, err = f.Write(binary)
	if err != nil {
		err = fmt.Errorf("cannot write binary: %s", err)
		return
	}
	err = f.Chmod(os.ModePerm)
	if err != nil {
		err = fmt.Errorf("cannot chmod: %s", err)
		return
	}
	path = f.Name()
	return
}

func decompressLib(raw []byte) (data []byte, err error) {
	decoder, err := zstd.NewReader(bytes.NewReader(raw))
	if err != nil {
		err = fmt.Errorf("cannot create zstd reader: %s", err)
		return
	}
	defer decoder.Close()
	data, err = io.ReadAll(decoder)
	if err != nil {
		err = fmt.Errorf("cannot reading decompressed data: %s", err)
	}
	return
}
