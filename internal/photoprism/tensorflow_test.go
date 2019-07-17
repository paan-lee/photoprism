package photoprism

import (
	"io/ioutil"
	"testing"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestTensorFlow_LoadLabelRules(t *testing.T) {
	t.Run("labels.yml exists", func(t *testing.T) {
		conf := config.NewTestConfig()

		tensorFlow := NewTensorFlow(conf)

		result := tensorFlow.loadLabelRules()
		assert.Nil(t, result)
	})
	t.Run("labels.yml not existing in config path", func(t *testing.T) {
		conf := config.NewTestErrorConfig()

		tensorFlow := NewTensorFlow(conf)

		result := tensorFlow.loadLabelRules()
		assert.Contains(t, result.Error(), "label rules file not found")
	})
}

func TestTensorFlow_LabelsFromFile(t *testing.T) {
	t.Run("/chameleon_lime.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := NewTensorFlow(conf)

		result, err := tensorFlow.LabelsFromFile(conf.ExamplesPath() + "/chameleon_lime.jpg")

		assert.Nil(t, err)

		if err != nil {
			t.Log(err.Error())
			t.Fail()
		}

		assert.NotNil(t, result)
		assert.IsType(t, Labels{}, result)
		assert.Equal(t, 1, len(result))

		t.Log(result)

		assert.Equal(t, "chameleon", result[0].Name)

		assert.Equal(t, 7, result[0].Uncertainty)
	})
	t.Run("not existing file", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := NewTensorFlow(conf)

		result, err := tensorFlow.LabelsFromFile(conf.ExamplesPath() + "/notexisting.jpg")
		assert.Contains(t, err.Error(), "no such file or directory")
		assert.Empty(t, result)
	})
}

func TestTensorFlow_Labels(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	t.Run("/chameleon_lime.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := NewTensorFlow(conf)

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/chameleon_lime.jpg"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)

			t.Log(result)

			assert.NotNil(t, result)

			assert.Nil(t, err)
			assert.IsType(t, Labels{}, result)
			assert.Equal(t, 1, len(result))

			assert.Equal(t, "chameleon", result[0].Name)

			assert.Equal(t, 100-93, result[0].Uncertainty)
		}
	})
	t.Run("/dog_orange.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := NewTensorFlow(conf)

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/dog_orange.jpg"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)

			t.Log(result)

			assert.NotNil(t, result)

			assert.Nil(t, err)
			assert.IsType(t, Labels{}, result)
			assert.Equal(t, 2, len(result))

			assert.Equal(t, "chihuahua dog", result[0].Name)
			assert.Equal(t, "pembroke dog", result[1].Name)

			assert.Equal(t, 34, result[0].Uncertainty)
			assert.Equal(t, 91, result[1].Uncertainty)
		}
	})
	t.Run("/Random.docx", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := NewTensorFlow(conf)

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/Random.docx"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)
			assert.Empty(t, result)
			assert.Contains(t, err.Error(), "invalid image")
		}
	})
	t.Run("/6720px_white.jpg", func(t *testing.T) {
		conf := config.TestConfig()

		tensorFlow := NewTensorFlow(conf)

		if imageBuffer, err := ioutil.ReadFile(conf.ExamplesPath() + "/6720px_white.jpg"); err != nil {
			t.Error(err)
		} else {
			result, err := tensorFlow.Labels(imageBuffer)
			assert.Empty(t, result)
			assert.Nil(t, err)
		}
	})
}

func TestTensorFlow_LoadModel(t *testing.T) {
	t.Run("model path exists", func(t *testing.T) {
		conf := config.NewTestConfig()

		tensorFlow := NewTensorFlow(conf)

		result := tensorFlow.loadModel()
		assert.Nil(t, result)
	})
	t.Run("model path does not exist", func(t *testing.T) {
		conf := config.NewTestErrorConfig()

		tensorFlow := NewTensorFlow(conf)

		result := tensorFlow.loadModel()
		assert.Contains(t, result.Error(), "Could not find SavedModel")
	})
}
