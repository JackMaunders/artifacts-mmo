# Artifacts MMO Controller - Server

A Golang API that makes calls to the [Artifacts MMO](https://artifactsmmo.com/) API to allow in game movements & action to be performed

## Getting Started

These instructions will get you up and running on your local machine for development and testing purposes

## Installation (Optional)

To install dependencies, run:

```sh
$ go mod download
```
**Note**: This step is optional, the dependencies will download automatically before a run or build

## Usage

### Running the server locally

```sh
$ go run src/main.go
```

### Building

```sh
$ go build src/main.go
```

This will compile the code into an executable

### Running the build locally

```sh
$ ./main
```

*Note* this requires [Building](#building) first.

### Linting

To lint the project locally you will need to install the linter, this is easily done using Brew (more ways to install [here](https://golangci-lint.run/welcome/install/)):

```sh
brew install golangci-lint
```

Once installed, you can run:

```sh
golangci-lint run
```
