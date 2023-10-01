package awx

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/zcubbs/x/pretty"
	"path/filepath"
)

func Load(path string, debug bool) (*Config, error) {
	var cfg Config

	initViperPresets(path)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("unable to load config file path=%s err=%s", path, err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("could not decode config into struct err=%s", err)
	}

	err = validate(&cfg)
	if err != nil {
		return nil, fmt.Errorf("config validation failed err=%s", err)
	}

	if debug {
		printConfig(cfg)
	}

	return &cfg, nil
}

func initViperPresets(path string) {
	dir := filepath.Dir(path)
	file := filepath.Base(path)
	viper.AddConfigPath(dir)
	viper.SetConfigName(file)
	viper.SetConfigType("yaml")

	for k, v := range defaults {
		viper.SetDefault(k, v)
	}
}

func validate(_ *Config) error {
	return nil
}

func printConfig(cfg Config) {
	fmt.Println("======================config======================")
	if cfg.Instance.AdminPass != "" {
		cfg.Instance.AdminPass = "********"
	}
	pretty.PrintJson(cfg)
	fmt.Println("==================================================")
}

const defaultOperatorName = "awx-operator"
const defaultNamespace = "awx"

var defaults = map[string]interface{}{
	"namespace":                  defaultNamespace,
	"operator.helm_repo_url":     "https://ansible.github.io/awx-operator/",
	"operator.helm_repo_name":    defaultOperatorName,
	"operator.helm_chart_name":   defaultOperatorName,
	"operator.helm_release_name": defaultOperatorName,
	"instance.name":              defaultNamespace,
	"instance.admin_user":        "admin",
	"instance.admin_pass":        "admin",
	"instance.is_node_port":      false,
	"instance.node_port":         30080,
	"instance.no_log":            true,
}
