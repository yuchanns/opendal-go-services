package aliyun_drive_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuchanns/opendal-go-services/aliyun_drive"
)

func TestSchemeAliyunDrive(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	assert.Equal(aliyun_drive.Scheme.Name(), "aliyun_drive")
	assert.Nil(aliyun_drive.Scheme.LoadOnce())
	path := aliyun_drive.Scheme.Path()
	assert.NotEmpty(path)
	assert.Nil(os.Remove(path))
}
