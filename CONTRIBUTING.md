# Contributing to goldmark

Thank you for considering contributing to goldmark! We welcome contributions from the community and are grateful for your support. This document will guide you through the process of setting up your development environment, running tests, and submitting pull requests.

## Table of Contents

- [Getting Started](#getting-started)
- [Development Environment](#development-environment)
- [Running Tests](#running-tests)
- [Submitting Pull Requests](#submitting-pull-requests)
- [Code Style](#code-style)
- [Code of Conduct](#code-of-conduct)

## Getting Started

To get started with contributing to goldmark, follow these steps:

1. Fork the repository on GitHub.
2. Clone your forked repository to your local machine.
3. Create a new branch for your changes.

```bash
git clone https://github.com/your-username/goldmark.git
cd goldmark
git checkout -b your-branch-name
```

## Development Environment

To set up your development environment, you will need to have Go installed on your machine. You can download and install Go from the official website: [https://golang.org/dl/](https://golang.org/dl/).

Once you have Go installed, you can install the project dependencies by running the following command:

```bash
go mod tidy
```

## Running Tests

Before submitting your changes, make sure to run the tests to ensure that everything is working correctly. You can run the tests using the following command:

```bash
go test -v ./...
```

Additionally, you can run the benchmark tests to measure the performance of the library:

```bash
go test -bench ./...
```

## Submitting Pull Requests

When you are ready to submit your changes, follow these steps:

1. Commit your changes with a descriptive commit message.
2. Push your changes to your forked repository.
3. Create a pull request on GitHub.

```bash
git add .
git commit -m "Description of your changes"
git push origin your-branch-name
```

In your pull request, provide a clear description of the changes you have made and any relevant information that will help the maintainers review your contribution.

## Code Style

Please follow the Go coding style and conventions when contributing to goldmark. You can use the `golangci-lint` tool to check your code for any style issues. To run the linter, use the following command:

```bash
golangci-lint run
```

## Code of Conduct

We expect all contributors to adhere to the [Code of Conduct](CODE_OF_CONDUCT.md). Please read it to understand the expectations for behavior when contributing to this project.

Thank you for your contributions and support!
