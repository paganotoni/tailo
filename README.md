# Tailo
Tailo is a Go wrapper for the Tailwind CSS standalone binary. It is intended to help the downloading and usage of the TailwindCSS binary.

## Download
To download the Tailwind CSS binary, you can use the `download` command.

```sh
// Downloading latest Tailwind CSS version
go tool tailo download

// Or  Using a specific version
go tool tailo download -v v3.4.6 --musl

// OR invoke via go run
go run github.com/paganotoni/tailo/@latest download -v v3.4.6 --musl
```

This will download the Tailwind CSS binary into `bin/tailwindcss`.

## Calling the Tailwind CSS binary

If the tool is invoked with other thing than 'download' it passes the arguments to the Tailwind CSS binary. For example:

```sh
go tool tailo --input <input_file> --output <output_file>
go tool tailo --watch -i tailwind.css --o application.css

// or via go run
go run github.com/paganotoni/tailo/@latest --watch -i tailwind.css --o application.css
```

Invoking the Tailwind CSS binary assumes the binary has already been downloaded and is available in the `bin` directory.
