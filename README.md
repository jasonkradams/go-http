# go-gttp

## Introduction
`go-gttp` is a lightweight HTTP server written in Go that serves a specified file as `/index.html` with customizable HTTP headers and status codes. This server allows you to quickly set up a simple HTTP server with precise control over the response headers and status codes.

## Table of Contents
- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Configuration](#configuration)
- [Examples](#examples)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

## Getting Started

### Installation
1. Clone the repository:
    ```sh
    git clone https://github.com/jasonkradams/go-gttp.git
    ```
2. Navigate to the project directory:
    ```sh
    cd go-gttp
    ```
3. Build the project:
    ```sh
    go build
    ```

## Usage

### Running the Server

To run the server, use the following command:

```sh
./go-gttp -file <file-to-serve> -header "<Header-Name>: <Header-Value>" -status <status-code>
```

### Command Line Arguments

```sh
-file: Specify the file to serve as /index.html. This argument is required.
-header: Add custom headers to the HTTP response. This argument can be repeated for multiple headers.
-status: Specify the status response code (supports 200, 404, 500).
```

## Configuration

### File Configuration

Specify the file to serve with the -file argument:

```sh
./go-gttp -file assets/hello-world.html
```

### Custom Headers

Add custom headers using the -header argument:

```sh
./go-gttp -header "Cache-Control: no-store" -header "Pragma: no-cache"
```

### Status Codes

Specify the status code with the -status argument:

```sh
./go-gttp -status 200
```

## Examples

### Basic Example

Serve a file as /index.html without any custom headers:

```sh
./go-gttp -file assets/hello-world.html
```

### Advanced Example

Serve a file with custom headers and a status code of 200:

```
./go-gttp -file assets/hello-world.html -header "Cache-Control: no-store" -header "Pragma: no-cache" -status 200
```

## Troubleshooting

If you encounter any issues, ensure that:

* The specified file exists.
* The file path is correct.
* The status code is one of the supported values (200, 404, 500).

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.
