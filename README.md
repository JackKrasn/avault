# Avault

Avault is a tool for decrypting yaml files. Yaml files are encrypted by Ansible Vault. 

Use Avault to:

- Decrypt `values.yaml` for helm. Use as a secret driver in [helm-secret](https://github.com/JackKrasn/helm-secrets) plugin. 
- Decrypt any yaml file which are encrypted by Ansible Vault. 

## Install


Binary downloads of the Avault utility can be found on [the Releases page](https://github.com/JackKrasn/helm-secrets/releases/latest).

Unpack the `avault` binary and add it to your PATH and you are good to go!


## Usage

For decrypting  `file.yaml` use:

```bash
avault decrypt file.yaml -p <password phrase>
```

`file.yaml.dec` it's result of the work avault utility. The encrypted values in file will be decrypted. 

The decrypted file saves in the same path where the original file is located.

The secret phrase for decryption can be specified by define environment variable `AVAULT_PASSWORD`

```bash
export AVAULT_PASSWORD='<password_phrase>'
```

## CLI Reference

```bash
The utility decrypts yaml files.
Yaml files encrypted by Ansible Vault.

Usage:
  avault [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  decrypt     Decrypt yaml file encrypted by Ansible Vault
  help        Help about any command
  version     Print the version number of generated code example

Flags:
      --debug             enable verbose output
  -h, --help              help for avault
  -p, --password string   password phrase for decryption

Use "avault [command] --help" for more information about a command.
```

## Testing

``make test``

## Build

``make build``

or build for different platforms

``make snaphost``

## Release

Build and publish the avault with [goreleaser](https://goreleaser.com).

Goreleaser and autotag must be installed.

```bash
make release
```

This command run `autotag`  and then `goreleaser`.

At first `autotag`  determine what the next tag should be and then creates the tag by executing git tag.

[autotag](https://github.com/pantheon-systems/autotag)

Then `goreleaser` build project and publish artifacts.

In order to release to GitHub, you'll need to export a GITHUB_TOKEN environment variable, 
which should contain a valid GitHub token with the repo scope. 

It will be used to deploy releases to your GitHub repository.

[Read more](https://goreleaser.com/quick-start/)
