# Terraform Provider for NGINX Proxy Manager

> [!IMPORTANT]
> This is a project which is mainly used for my own learning, and is **not** production ready.
> 
> I will try to maintain this project as much as possible but make no guarantees on fixes or
> updates to the provider. 
>
> Please feel free to make contributions and I will try to review them when I can.

## Contents

- [Terraform Provider for NGINX Proxy Manager](#terraform-provider-for-nginx-proxy-manager)
  - [Contents](#contents)
  - [Documentation](#documentation)
  - [Installation](#installation)
  - [Supported Versions](#supported-versions)
  - [Contributing](#contributing)
  - [Testing](#testing)

## Documentation

All documentation for this provider can be found on the Terraform Registry.

## Installation

This provider is installed via the following configuration block:

```terraform
terraform {
    required_providers {
        nginxproxymanager = {
            source = "edklesel/nginxproxymanager"
        }
    }
}
```

## Supported Versions

The provider is tested against the following versions in CI:
- 2.15.1
- 2.16.0

Due to API changes, the minimum required version for Nginx Proxy Manager is 2.15.0.

The provider is tested against Terraform versions 1.15.0 and above

## Contributing

For development, you should be able to run the below commands to get a test environment set up. **This required docker compose to be installed.**
```bash
make docker-up
```

To install the provider for local testing:
```bash
make install
```

Prior to creating a PR, please ensure that the following commands have been run to ensure formatting and linting is done.
```bash
make generate
make lint
```

## Testing