package awx

import (
	"fmt"
	"os"
)

func getKubeConfig(path string, debug bool) (string, error) {
	if path != "" {
		return path, nil
	}

	hd, err := getUserHomeDir()
	if err != nil {
		return "", err
	}

	kc := fmt.Sprintf("%s/.kube/config", hd)
	fi, err := os.Stat(kc)
	os.IsNotExist(err)
	if err != nil && debug {
		fmt.Printf("kubeconfig not found in default location %s\n", kc)
	}

	if fi != nil {
		fmt.Printf("kubeconfig found in default location %s\n", kc)
		return kc, nil
	}

	kc = "/etc/rancher/k3s/k3s.yaml"
	fi, err = os.Stat(kc)
	os.IsNotExist(err)
	if err != nil && debug {
		fmt.Printf("kubeconfig not found in default location %s\n", kc)
	}

	if fi != nil {
		fmt.Printf("kubeconfig found in default location %s\n", kc)
		return kc, nil
	}

	kc = os.Getenv("KUBECONFIG")
	if kc == "" {
		return "", fmt.Errorf("kubeconfig not found")
	}

	fmt.Printf("kubeconfig found in default location %s\n", kc)
	return kc, nil
}

func getUserHomeDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home dir")
	}
	return home, nil
}
