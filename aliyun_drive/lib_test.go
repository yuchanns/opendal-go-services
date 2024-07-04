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

	scheme := aliyun_drive.Scheme{}
	assert.Equal(scheme.Scheme(), "aliyun_drive")
	assert.NotEmpty(scheme.Path())
	t.Logf("path: %s", scheme.Path())
	os.Remove(scheme.Path())
}
