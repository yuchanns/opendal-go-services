package aliyun_drive_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yuchanns/opendal-go-services/aliyun_drive"
)

func TestSchemeAliyunDrive(t *testing.T) {
	t.Parallel()
	assert := assert.New(t)

	assert.Equal(aliyun_drive.Scheme.Scheme(), "aliyun_drive")
	assert.NotEmpty(aliyun_drive.Scheme.Path())
	assert.Nil(os.Remove(aliyun_drive.Scheme.Path()))
}
