package clusters

import (
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	configPath = "clusters/configs"
)

type clusterConfig struct {
	ClusterName string   `yaml:"name"`
	Labels      []string `yaml:"labels"`
}
var Configs []clusterConfig

func ParseConfig() {
	files, err := os.ReadDir(configPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		path := filepath.Join(configPath, file.Name())
		configContent, err := os.ReadFile(path)
		if err != nil {
			log.Fatal("failed to read ", path)
		}

		var config clusterConfig
		err = yaml.Unmarshal(configContent, &config)
		if err != nil {
			log.Fatal("failed to unmarshal ", path)
		}

		Configs = append(Configs, config)
	}
}
