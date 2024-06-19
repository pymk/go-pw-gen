# go-pw-gen

This is a CLI tool for generating random passwords.

> [!WARNING]
> It does not use the ["crypto/rand"](https://pkg.go.dev/crypto/rand) package.

## Usage

```sh
go run .

# To generate more than 1 password, pass an integer:
go run . 10
```
