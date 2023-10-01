package awx

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zcubbs/go-k8s/kubernetes"
	"github.com/zcubbs/x/must"
	"kwx/pkg/helm"
	"os"
	"time"
)

var (
	kubeconfig string
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "install awx operator & instance",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		verbose := cmd.Flag("verbose").Value.String() == "true"
		must.Succeed(install(verbose))
	},
}

func addInstance(values instanceTmplValues, _ string, debug bool) error {
	err := kubernetes.ApplyManifest(adminPasswordSecretTmpl, values, debug)
	if err != nil {
		return fmt.Errorf("failed to apply awx admin password secret \n %w", err)
	}

	err = kubernetes.ApplyManifest(instanceTmpl, values, debug)
	if err != nil {
		return fmt.Errorf("failed to apply awx instance \n %w", err)
	}
	return nil
}

type instanceTmplValues struct {
	Name          string
	Namespace     string
	IsIpv6        bool
	IsNodePort    bool
	NodePort      int
	AdminUser     string
	AdminPassword string
	NoLog         bool
}

// #nosec G101
var instanceTmpl = `
apiVersion: awx.ansible.com/v1beta1
kind: AWX
metadata:
  name: {{ .Name }}
  namespace: {{ .Namespace }}
spec:
  {{- if .IsIpv6 }}
  ipv6_enabled: true
  {{- end }}
  {{- if .IsNodePort }}
  service_type: NodePort
  nodeport_port: {{ .NodePort }}
  {{- else }}
  service_type: ClusterIP
  {{- end }}
  ingress_type: none
  no_log: {{ .NoLog }}
  admin_user: {{ .AdminUser }}

`

// #nosec G101
var adminPasswordSecretTmpl = `
apiVersion: v1
kind: Secret
metadata:
  name: {{ .Name }}-admin-password
  namespace: {{ .Namespace }}
stringData:
  password: {{ .AdminPassword }}

`

func install(verbose bool) error {
	cfg, err := Load(awxConfigPath, verbose)
	if err != nil {
		return err
	}

	kc, err := getKubeConfig(kubeconfig, verbose)
	if err != nil {
		return err
	}

	if verbose {
		printConfig(*cfg)
	}

	client := helm.NewClient()
	client.Settings.SetNamespace("awx")

	args := map[string]string{}

	_ = os.Setenv("HELM_NAMESPACE", cfg.Namespace)

	client.RepoAdd(cfg.Operator.HelmRepoName, cfg.Operator.HelmRepoURL)
	client.RepoUpdate()
	// create namespace
	if err := createNamespace(kc, cfg.Namespace); err != nil {
		return fmt.Errorf("failed to create namespace: %w", err)
	}
	client.InstallChart(cfg.Operator.HelmRelease, cfg.Operator.HelmRepoName, cfg.Operator.HelmChartName, args)

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 3*time.Minute)
	defer cancel()
	err = kubernetes.IsDeploymentReady(
		ctxWithTimeout,
		kc,
		"awx",
		[]string{
			"awx-operator-controller-manager",
		},
		verbose,
	)
	if err != nil {
		return fmt.Errorf("failed to wait for awx-operator to be ready \n %w", err)
	}

	//apply awx instance
	err = addInstance(instanceTmplValues{
		Name:          cfg.Instance.Name,
		Namespace:     cfg.Namespace,
		IsIpv6:        false,
		IsNodePort:    cfg.Instance.IsNodePort,
		NodePort:      cfg.Instance.NodePort,
		AdminUser:     cfg.Instance.AdminUser,
		AdminPassword: cfg.Instance.AdminPass,
		NoLog:         cfg.Instance.NoLog,
	}, kubeconfig, verbose)
	if err != nil {
		return fmt.Errorf("failed to apply awx instance \n %w", err)
	}

	return nil
}

func createNamespace(kc string, namespace string) error {
	return kubernetes.CreateNamespace(kc, []string{namespace})
}

func init() {
	installCmd.Flags().StringVarP(&kubeconfig, "kubeconfig", "k", "", "kubeconfig file path (default is $HOME/.kube/config)")

	Cmd.AddCommand(installCmd)
}
