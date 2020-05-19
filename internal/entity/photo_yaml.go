package entity

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/photoprism/photoprism/pkg/fs"
	"gopkg.in/yaml.v2"
)

// Yaml returns photo data as YAML string.
func (m *Photo) Yaml() ([]byte, error) {
	out, err := yaml.Marshal(m)

	if err != nil {
		return []byte{}, err
	}

	return out, err
}

// SaveAsYaml saves photo data as YAML file.
func (m *Photo) SaveAsYaml(fileName string) error {
	data, err := m.Yaml()

	if err != nil {
		return err
	}

	// Make sure directory exists.
	if err := os.MkdirAll(filepath.Dir(fileName), os.ModePerm); err != nil {
		return err
	}

	// Write YAML data to file.
	if err := ioutil.WriteFile(fileName, data, os.ModePerm); err != nil {
		return err
	}

	return nil
}

// LoadFromYaml photo data from a YAML file.
func (m *Photo) LoadFromYaml(fileName string) error {
	data, err := ioutil.ReadFile(fileName)

	if err != nil {
		return err
	}

	hadDesc := m.HasDescription()

	if err := yaml.Unmarshal(data, m); err != nil {
		return err
	}

	if m.HasDescription() && !hadDesc && m.DescriptionSrc == "" {
		m.DescriptionSrc = SrcYaml
	}

	return nil
}

// YamlFileName returns the YAML backup file name.
func (m *Photo) YamlFileName(originalsPath string, hidden bool) string {
	if hidden {
		return filepath.Join(originalsPath, m.PhotoPath, fs.HiddenPath, m.PhotoName) + fs.YamlExt
	}

	return filepath.Join(originalsPath, m.PhotoPath, m.PhotoName) + fs.YamlExt
}
