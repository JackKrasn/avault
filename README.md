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

Build and publish the avault with [goreleaser](https://goreleaser.com)

```bash
make release
```

## Decrypt file

```bash
avault decrypt file.yaml -p <password phrase>
```

Or specify environment variable at first `AVAULT_PASSWORD=<password phrase>`

And then run `avault decrypt file.yaml`