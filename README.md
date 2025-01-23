# Encoding Utility Project

This project provides utilities for encoding and decoding data using different encoding schemes. It supports Base16, Base32, and Base64 encoding, and includes a command-line tool to run encoding operations.

## Project Structure

```
├── Makefile
├── base16util
│ ├── base16util.go
│ └── base16util_test.go
├── base32util
│ ├── base32util.go
│ └── base32util_test.go
├── base64util
│ ├── base64util.go
│ └── base64util_test.go
├── encoder
│ ├── encoder.go
│ └── encoder_test.go
├── go.mod
├── internal
│ ├── delpadd
│ │ └── delpadd.go
│ └── errchecker
│ └── errchecker.go
└── main.go
```
 - `base16util`, `base32util`, and `base64util`: Modules containing the implementation for each encoding format, along with their corresponding test files.
- `encoder`: Contains the core encoder logic that works with different encoding schemes.
- `internal`: Includes utility functions like `delpadd` and `errchecker`.
- `main.go`: Entry point for the application.
- `Makefile`: Includes build and run instructions for the project.

## Features

- Supports **Base16**, **Base32**, and **Base64** encoding.
- Includes a command-line interface (CLI) to encode data with the specified encoding scheme.
- Allows users to specify the encoding format via the `-encoding` parameter.

## Usage

You can specify the encoding format when running the program by using the `-encoding` parameter. Available encoding options are:
- `base16`
- `base32`
- `base64`

#### Example:

`make run ENCODING=base32`
 This command will use the Base64 encoding format.

## Commands

To build the application, run:

`make build`

 To run the application, run:

`make run`

To run tests:

`make test`

## Dependencies

The project uses the Go programming language. Ensure you have Go installed and your GOPATH set up properly.

## License

This project is licensed under the MIT License - see the LICENSE file for details.