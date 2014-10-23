# Monolith1

This README outlines the details of collaborating on this Ember & Go application.

A short introduction of this app could easily go here.

## Prerequisites

You will need the following things properly installed on your computer.

* [Git](http://git-scm.com/)
* [Node.js](http://nodejs.org/) (with NPM) and [Bower](http://bower.io/)
* [Go](http://golang.org)

## Installation

* `git clone <repository-url>` this repository
* change into the new directory
* `npm install`
* `bower install`
* `go get github.com/GeertJohan/go.rice`
* `go get github.com/GeertJohan/go.rice/rice`
* `go get`


## Running / Development

From the `<project root folder>/monolith` run

    go run main.go

and point your browser at [http://localhost:8080](http://localhost:8080)

### Running Tests

* `ember test`
* `ember test --server`

### Deploying

From the project root run the following command

    bin/release

This will create a single binary in `<project root folder>/monolith1` called `monolith1`. Run the binary and point your browser at [http://localhost:8080](http://localhost:8080)

    monolith1/monolith1
