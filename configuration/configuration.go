package configuration

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/goccy/go-yaml"
)

type Device struct {
	Name      string   `yaml:"name" validate:"required"`
	Path      string   `yaml:"path" validate:"required"`
	Allowlist []string `yaml:"allowlist"`
}

type Config struct {
	ConfigPath   string
	PlaylistsDir string   `yaml:"playlists_dir" validate:"required"`
	Devices      []Device `yaml:"devices"`
}

func ReadFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	check(err)
	return string(data), nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadYaml(path string) (Config, error) {
	var v Config
	validate := validator.New()
	file, _ := ReadFile(path)
	dec := yaml.NewDecoder(
		strings.NewReader(file),
		yaml.Validator(validate),
		yaml.Strict(),
	)
	absPath, _ := filepath.Abs(path)
	v.ConfigPath = absPath
	if err := dec.Decode(&v); err != nil {
		return v, err
	} else {
		fmt.Println(v)
		return v, nil
	}
}
