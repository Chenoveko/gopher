#!/bin/bash

# Display the installed Go version
go version

# Display Go's environment variables and configuration
go env

# Format Go source code
go fmt main.go

# Compiles and execute Go source code
go run main.go

# Compiles Go source code
go build main.go

# Initialize a new Go module and create a go.mod file
go mod init mod_manager

# Compile and install the current package as an executable
# This requires the package to be a command, usually package main
go install 

# Add an external dependency to the current module
# This may update go.mod and go.sum
go get rsc.io/quote

# Runs any tests associated with the current project
go test

# Standard sequence for a new Go project
go mod init mod_manager # Initialize the Go module
go get rsc.io/quote # Add an external dependency to the module
go run . # Compile and run the current package
go build .  # Compile the current package and create an executable
go install . # Compile and install the current package
go test ./... # Run tests for all packages in the module and its subdirectories