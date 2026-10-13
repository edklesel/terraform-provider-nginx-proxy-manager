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
  - [Test Setup](#test-setup)
    - [Quirks](#quirks)

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

Tests can be run using the relevant command:
```bash
make test-sdk # For SDK testing
make test-acc # For TF provider testing
```

## Test Setup

Tests are set up using the following components, mirroring the setup in the main NGINX Proxy Manager repository:

- `NPM` / `MariaDB` - The main NGINX Proxy Manager service
- `PowerDNS` / `MariaDB` - DNS provider used for ACME DNS challenges
- `Step-CA` - ACME Certificate provider
- `DNSRouter` - Intercept DNS requests for the domain `example.com` and forward to PowerDNS.

### Quirks

- NPM has access to the step-ca volume as it needs to read the root CA for requests CA certificate verification.
- NPM locally mounts its data directory so the test suite has access to keys.json, which is used to decrypt the JWT tokens in tests.
- Step-CA, for some reason, creates its `root_ca.crt` file with `1000:1000` permissions, regardless of the user running the container. Therefore, the makefile sets the UID to this user.
- Similarly, NPM creates the `keys.json` file with only user rw permissions, so there's a command to allow group read permission during the `docker-up` make command so that it can be read by tests, in case the UID running the tests isn't 1000.
