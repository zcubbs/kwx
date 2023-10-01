package awx

import (
	"github.com/spf13/cobra"
	"github.com/zcubbs/x/must"
)

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "uninstall awx operator & instance",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		verbose := cmd.Flag("verbose").Value.String() == "true"
		must.Succeed(uninstall(verbose))
	},
}

func uninstall(verbose bool) error {

	return nil
}

func init() {
	uninstallCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "kubeconfig file path (default is $HOME/.kube/config)")

	Cmd.AddCommand(uninstallCmd)
}
