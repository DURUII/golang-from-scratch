package ch00

import (
	"advanced/ch00/set"
	"advanced/ch00/undo"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetUndo(t *testing.T) {
	s := set.NewSet[string]()
	s.Add("武汉科技大学")
	assert.Equal(t, 1, s.Size())
	s.Undo()
	assert.Equal(t, 0, s.Size())
	s.Undo()
	s.Undo()
	assert.Equal(t, undo.ErrNoUndo, s.Err())
	s.Add("武汉大学")
	assert.Equal(t, 1, s.Size())
	s.Add("华中科技大学")
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains("武汉大学"))
	assert.False(t, s.Contains("武汉科技大学"))
	s.Undo()
	s.Undo()
	assert.False(t, s.Contains("华中科技大学"))
	s.Add("武汉科技大学")
	assert.True(t, s.Contains("武汉科技大学"))
	s.Delete("武汉科技大学")
	assert.False(t, s.Contains("武汉科技大学"))
	s.Undo()
	assert.True(t, s.Contains("武汉科技大学"))
}
