# Simple Secrets Encryptor

[![Maintainability](https://api.codeclimate.com/v1/badges/c5bf10a5bdc27cf25567/maintainability)](https://codeclimate.com/github/dimw/simple-secrets-encryptor/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/c5bf10a5bdc27cf25567/test_coverage)](https://codeclimate.com/github/dimw/simple-secrets-encryptor/test_coverage)

This project can be used for encoding and decoding of a bunch of structured files containing key-value maps (e.g. YAML) 
using asymmetric encryption ([RSA]). 
This encryption allows the holders of the public key to perform the one way encryption without the possibility 
to retrieve the original value. The holders of the private key, however, can perform the decryption (and also 
[derive the public key](https://stackoverflow.com/questions/696472/given-a-private-key-is-it-possible-to-derive-its-public-key) 
for encryption).

## Use-cases 

### CI/CD

Often, a lot of secrets are required during the phases of CI/CD. 
Holding the secrets in the unencrypted format as part of the code can be considered as bad practice. 
However, by encrypting the secrets before pushing them to the repository makes them useless for eavesdropper 
which do not have the private key required for encryption.

A step of the CI/CD pipeline (e.g. in [GitHub Action]) can inject the public key (e.g. from the [GitHub Secrets]) and 
decrypt the required secrets using the _Simple Secret Encoder_.

### GitHub Secrets Management

[GitHub Secrets] can be managed via the GitHub UI, via the [REST API](https://docs.github.com/en/rest/reference/actions#secrets), 
 by using a 3rd party tool like [Terraform](https://registry.terraform.io/providers/integrations/github/latest/docs/resources/actions_secret), etc.
For easier secret management it can be considered storing the encrypted secrets in a file (which can be pushed to a 
repository) and decrypt it before executing Terraform.

## How to Build

[Go] 1.15+ is required to build the project. 
Checkout the repository and trigger the build on Unix-like systems by executing the following command:

```
$ go build -o sse
```

On Windows:

```
> go build -o sse.exe
```

Both commands will generate the executable binary named `sse`. 
This executable will be used in the following examples. 
However, `go run main.go` can be used if the binary is missing.

## Commands

The following command will retrieve the supported commands.

```bash
$ sse help
```

### Generation of Keys

The well-known [ssh-keygen] tool can be used for generation of the RSA key pair. 
However, the _Simple Secret Encryptor_ provides a shortcut for generating the private and public keys. 
Per default, the output will be stored in the files `private.key` and `public.pem` accordingly. 

```bash
$ sse generate-keys
```

### Encryption

The `encrypt` can be used to encrypt secrets. The process is structured as following:
1. It traverses through the working directory (default: `./`), 
2. Visits each dictionary-like file (default: `*.{yml,yaml}`), 
3. Detects keys referring to secrets (those containing `secret`, `password`, or `token` keywords)
4. Performs encryption of those secrets using the provided public key (default: `public.pem`).
5. The original files are replaces.

Example:
```bash
$ mkdir ./secrets
$ echo "foo-token: bar" > ./secrets/secret.yml
$ sse encrypt --workdir=./secrets
$ cat ./secrets/secret.yml
foo-token: ENC[rsa,data:OG4ZLQYbAthyx8...Es4kcZSDXy50iQ==]
```

The values which are already encrypted (similar to `ENC[...]`) are omitted. 
The command can be considered being idempotent.

### Decryption

The decryption can be performed when the private key (default file: `private.key`) is present. 
The process reads all files of the working directory, searches the encrypted keys, and replaces them in-place.

Example:
```bash
$ mkdir ./secrets
$ echo "foo-token: ENC[rsa,data:OG4ZLQYbAthyx8...Es4kcZSDXy50iQ==]" > ./secrets/secret.yml
$ sse encrypt --workdir=./secrets
$ cat ./secrets/secret.yml
foo-token: bar
```

## Supported Secret File Formats

Single-level key-value pairs can be encrypted and decrypted only. The secret files can be in the following formats: 

- YAML

## Known Limitations
- Only key-value maps without nesting are currently supported for encoding.
- Use private key protected by a passphrase.
- Add option for selecting a dedicated folder for output.
- Add option for outputting to a different structured format (e.g. JSON). 

[RSA]: https://en.wikipedia.org/wiki/RSA_(cryptosystem)
[ssh-keygen]: https://www.ssh.com/ssh/keygen/
[Go]: https://golang.org/
[GitHub Actions]: https://docs.github.com/en/actions
[GitHub Secrets]: https://docs.github.com/en/actions/reference/encrypted-secrets
