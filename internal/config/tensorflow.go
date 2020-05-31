package config

import (
	"path/filepath"

	tf "github.com/tensorflow/tensorflow/tensorflow/go"
)

// TensorFlowVersion returns the TenorFlow framework version.
func (c *Config) TensorFlowVersion() string {
	return tf.Version()
}

// TensorFlowOff returns true if TensorFlow should NOT be used for image classification (or anything else).
func (c *Config) TensorFlowOff() bool {
	return c.params.TensorFlowOff
}

// TensorFlowModelPath returns the TensorFlow model path.
func (c *Config) TensorFlowModelPath() string {
	return filepath.Join(c.AssetsPath(), "nasnet")
}

// NSFWModelPath returns the "not safe for work" TensorFlow model path.
func (c *Config) NSFWModelPath() string {
	return filepath.Join(c.AssetsPath(), "nsfw")
}
