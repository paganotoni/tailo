# Tailo

Tailo is a Go wrapper for the common operations with the TailwindCSS binary. It is intended to automate the process of installing the TailwindCSS binary, running it, and cleaning up the generated files.

## Setup Command
The setup command is an easy way to invoke tailo from your application without creating a dependency on this package.

```sh
go run github.com/paganotoni/tailo/cmd/setup@latest // Using Latest Tailwind

// Or  Using a specific version
go run github.com/paganotoni/tailo/cmd/setup@latest -version=v3.4.6
```

## Build Command
Like the Setup command the Build Command allows to specify the version of Tailwind to be used.

```sh
go run ./cmd/build -version=v3.4.6
```
