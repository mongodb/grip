package message

import (
	"testing"

	"github.com/mongodb/grip/level"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFieldsLevelMutability(t *testing.T) {
	assert := assert.New(t) // nolint

	m := Fields{"message": "hello world"}
	c := ConvertToComposer(level.Error, m)

	r := c.Raw().(Fields)
	assert.Equal(level.Error, c.Priority())
	assert.Equal(level.Error, r["metadata"].(*Base).Level)

	c = ConvertToComposer(level.Info, m)
	r = c.Raw().(Fields)
	assert.Equal(level.Info, c.Priority())
	assert.Equal(level.Info, r["metadata"].(*Base).Level)
}

func TestFields(t *testing.T) {
	t.Run("NewFieldsMessageCollectsBasicMetadata", func(t *testing.T) {
		m := Fields{"message": "hello world"}
		c := NewFields(level.Error, m)
		r, ok := c.Raw().(Fields)
		require.True(t, ok)
		assert.Equal(t, level.Error, c.Priority())
		base, ok := r["metadata"].(*Base)
		require.True(t, ok)
		assert.Equal(t, level.Error, base.Level)
		assert.Zero(t, base.Context)
		assert.Zero(t, base.Hostname)
		assert.Zero(t, base.Pid)
		assert.Zero(t, base.Process)
		assert.Zero(t, base.Time)
	})
	t.Run("NewExtendedFieldsCollectsExtendedMetadata", func(t *testing.T) {
		m := Fields{"message": "hello world"}
		c := NewExtendedFields(level.Error, m)
		r, ok := c.Raw().(Fields)
		require.True(t, ok)
		assert.Equal(t, level.Error, c.Priority())
		base, ok := r["metadata"].(*Base)
		require.True(t, ok)
		assert.Equal(t, level.Error, base.Level)
		assert.Zero(t, base.Context)
		assert.NotZero(t, base.Hostname)
		assert.NotZero(t, base.Pid)
		assert.NotZero(t, base.Process)
		assert.NotZero(t, base.Time)
	})
	t.Run("NewSimpleFieldsHasNoMetadata", func(t *testing.T) {
		m := Fields{"message": "hello world"}
		c := NewSimpleFields(level.Error, m)
		r, ok := c.Raw().(Fields)
		require.True(t, ok)
		assert.Equal(t, level.Error, c.Priority())
		assert.Zero(t, r["metadata"])
	})
}
