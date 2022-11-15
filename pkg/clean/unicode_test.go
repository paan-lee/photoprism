package clean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnicode(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		assert.Equal(t, "Naïve bonds and futures surge as inflation eases 🚀🚀🚀", Unicode("Naïve bonds and futures surge as inflation eases 🚀🚀🚀"))
	})
	t.Run("Empty", func(t *testing.T) {
		assert.Equal(t, "", Unicode(""))
	})
}
