# kWx CLI

`kWx` is a Command-Line Interface (CLI) tool designed to simplify the deployment of AWX on Kubernetes. With just a few commands, you can install the AWX operator, set up an AWX instance, or uninstall them both.

---
<p align="center">
</p>
<p align="center">
  <img width="350" src="docs/assets/logo.png">
</p>

---

## Features

- **Install AWX-Operator:** Quickly deploy the AWX-Operator on your Kubernetes cluster.
- **Deploy AWX Instance:** Set up an AWX instance using default or customized configurations.
- **Clean Uninstallation:** Easily uninstall both the operator and instance when needed.

## Installation
```bash
curl -sfL https://raw.githubusercontent.com/zcubbs/kwx/main/scripts/install.sh | bash
```

## Configuration

KWx optionally accepts a configuration file to customize your AWX deployment. The configuration is in YAML format. By default, the CLI uses the following settings:

```yaml
namespace: awx
operator:
    helm_repo_url: https://ansible.github.io/awx-operator/
    helm_repo_name: awx-operator
    helm_chart_name: awx-operator
    helm_release_name: awx-operator
instance:
    name: awx
    admin_user: admin
    admin_pass: admin
    is_node_port: false
    node_port: 30080
    no_log: true
```

### Config Struct

Here's a brief overview of the configuration structure:

- **Namespace:** Kubernetes namespace where the deployment will occur.
- **Operator:** Settings for the AWX-Operator.
    - **HelmRepoURL:** URL of the Helm repository.
    - **HelmRepoName:** Name of the Helm repository.
    - **HelmChartName:** Name of the Helm chart.
    - **HelmRelease:** Helm release name.
- **Instance:** Configuration for the AWX instance.
    - **Name:** Instance name.
    - **AdminUser & AdminPass:** AWX admin credentials.
    - **IsNodePort:** Whether to use a NodePort service. If `true`, makes AWX accessible on a static port of the node.
    - **NodePort:** The static port to expose AWX if `IsNodePort` is `true`.
    - **NoLog:** If `true`, logging will be disabled.

## Usage

1. **Install AWX**:
    ```bash
    kwx awx install -c /path/to/config.yaml
    ```

2. **Uninstall AWX**:

    ```bash
    kwx awx uninstall -c /path/to/config.yaml
    ```
For more detailed options and commands, use the `--help` flag.

## Contributing

Contributions are welcome! If you find any issues, have suggestions, or would like to contribute code, please open an issue or a pull request on our GitHub page.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
