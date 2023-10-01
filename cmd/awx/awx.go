package awx

import "github.com/spf13/cobra"

var (
	awxConfigPath string
)

var Cmd = &cobra.Command{
	Use:   "awx",
	Short: "awx commands",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		if err != nil {
			panic(err)
		}
	},
}

type Config struct {
	Namespace string         `mapstructure:"namespace" json:"namespace"`
	Operator  OperatorConfig `mapstructure:"operator" json:"operator"`
	Instance  InstanceConfig `mapstructure:"instance" json:"instance"`
}

type OperatorConfig struct {
	HelmRepoURL   string `mapstructure:"helm_repo_url" json:"helm_repo_url"`
	HelmRepoName  string `mapstructure:"helm_repo_name" json:"helm_repo_name"`
	HelmChartName string `mapstructure:"helm_chart_name" json:"helm_chart_name"`
	HelmRelease   string `mapstructure:"helm_release_name" json:"helm_release_name"`
}

type InstanceConfig struct {
	Name       string `mapstructure:"name" json:"name"`
	AdminUser  string `mapstructure:"admin_user" json:"admin_user"`
	AdminPass  string `mapstructure:"admin_pass" json:"admin_pass"`
	IsNodePort bool   `mapstructure:"is_node_port" json:"is_node_port"`
	NodePort   int    `mapstructure:"node_port" json:"node_port"`
	NoLog      bool   `mapstructure:"no_log" json:"no_log"`
}

func init() {
	Cmd.PersistentFlags().StringVarP(&awxConfigPath, "config", "c", "", "yaml config file path")
}
