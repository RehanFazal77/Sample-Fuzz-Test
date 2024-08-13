# Sample-Fuzz-Test
# Go Fuzzing Example

This project demonstrates how to implement and run fuzz testing for a Go function using OSS-Fuzz.

## Project Overview

This project contains a simple Go function `ParseName` that splits a full name into first and last names. We've implemented a fuzz test for this function to identify potential edge cases and unexpected behaviors.

## Contents

- `main.go`: Contains the `ParseName` function
- `fuzz_test.go`: Contains the fuzz test for `ParseName`
- `Dockerfile`: Used to set up the OSS-Fuzz environment for running the fuzz tests

## Prerequisites

- Go 1.18 or later
- Docker

## Running the Fuzz Test

1. Run this before making the docker image
   ```
   go mod init fuzz-example
   ```
2. Build the Docker image:
   ```
   docker build -t fuzz-test .
   ```

4. Run the fuzz test:
   ```
   docker run --rm -it fuzz-test go test -fuzz=FuzzParseName -fuzztime=30s
   ```

## Understanding the Results

The fuzz test will run for 30 seconds (as specified by `-fuzztime=30s`). It will output information about the fuzzing process, including any failing cases it finds.

A successful fuzz test often results in a failure, as it identifies inputs that cause unexpected behavior. In this project, the fuzzer identifies that a single space input (" ") causes the `ParseName` function to return empty strings without an error, which is caught by our test assertions.
