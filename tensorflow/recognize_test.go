package tensorflow

import (
	"testing"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func TestRecognizeImage(t *testing.T) {
	if imageBuffer, err := ioutil.ReadFile("cat.jpg"); err != nil {
		t.Error(err)
	} else {
		result, err := RecognizeImage(string(imageBuffer))

		assert.NotNil(t, result)
		assert.Nil(t, err)
		assert.IsType(t, []LabelResult{}, result)
		assert.Equal(t, 5, len(result))

		assert.Equal(t, "tabby", result[0].Label)
		assert.Equal(t, "tiger cat", result[1].Label)

		assert.Equal(t, float32(0.23251747), result[1].Probability)
	}
}