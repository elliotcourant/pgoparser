package keywords

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewKeyword(t *testing.T) {
	t.Run("panics on invalid", func(t *testing.T) {
		assert.Panics(t, func() {
			NewKeyword("ajhgklass")
		})
	})

	t.Run("succeeds on valid", func(t *testing.T) {
		keyword := NewKeyword("select")
		assert.IsType(t, SELECT{}, keyword)
	})
}
