package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kwx/cmd/awx"
	"os"
)

var (
	Version string
	Commit  string
	Date    string
)

var (
	rootCmd = &cobra.Command{
		Use:   "",
		Short: "",
		Long:  "",
	}

	viperCommand = &cobra.Command{
		Run: func(c *cobra.Command, args []string) {
			fmt.Println(viper.GetString("Flag"))
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(getVersion())
		},
	}

	aboutCmd = &cobra.Command{
		Use:   "about",
		Short: "Print the info about the cli",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			About()
		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	rootCmd.AddCommand(aboutCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(viperCommand)

	addSubCommandPalettes()
}

func addSubCommandPalettes() {
	rootCmd.AddCommand(awx.Cmd)
}

func About() {
	fmt.Print(getFullVersion())
	fmt.Print(getDescription())
	fmt.Printf("Author: zcubbs \n")
	fmt.Println("Repository: https://github.com/zcubbs/kwx")
}

func getVersion() string {
	return fmt.Sprintf("v%s", Version)
}

func getFullVersion() string {
	return fmt.Sprintf(`
Version: v%s
Commit: %s
Date: %s`, Version, Commit, Date)
}

func getDescription() string {
	return `
/kwx/ is a cli tool that drives the bootstrapping
of an AWX Instance on K8s.
`
}
