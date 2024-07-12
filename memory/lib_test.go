package memory_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yuchanns/opendal-go-services/memory"
)

func TestSchemeMemory(t *testing.T) {
	t.Parallel()
	assert := require.New(t)

	assert.Equal(memory.Scheme.Name(), "memory")
	assert.Nil(memory.Scheme.LoadOnce())
	path := memory.Scheme.Path()
	assert.NotEmpty(path)
	assert.Nil(os.Remove(path))
}
