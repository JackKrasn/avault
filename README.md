# avault

Decrypt a yaml file  that has been decrypted with Ansible Vault.

## Getting started

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/avault
```

The secret phrase for decryption can be specified by define environment AVAULT_PASSWORD=xxxx

```bash
export AVAULT_PASSWORD=<password_phrase>
```

### Testing

``make test``


### Build

``make build``

or build for different platforms

``make snaphost``


### Release

Build and publish the avault with [goreleaser](https://goreleaser.com).

Goreleaser and autotag must be installed.

```bash
make release
```

This command run `autotag`  and then `goreleaser`.

At first `autotag`  determine what the next tag should be and then creates the tag by executing git tag.

[autog](https://github.com/pantheon-systems/autotag)

Then `goreleaser` build project and publish artifacts.

In order to release to GitHub, you'll need to export a GITHUB_TOKEN environment variable, 
which should contain a valid GitHub token with the repo scope. 

It will be used to deploy releases to your GitHub repository.

[Read more](https://goreleaser.com/quick-start/)

## Decrypt file

```bash
avault decrypt file.yaml -p <password phrase>
```

Or specify environment variable at first `AVAULT_PASSWORD=<password phrase>`

And then run `avault decrypt file.yaml`

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
