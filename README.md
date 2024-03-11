# poc-go-ginkgo
<!-- TOC -->

- [poc-go-ginkgo](#poc-go-ginkgo)
    - [Introduction](#introduction)
    - [Getting Started](#getting-started)
    - [Build and Test](#build-and-test)
    - [Contribute](#contribute)

<!-- /TOC -->
## Introduction
After some tests using *GO* to develop some tests against an API, doing end-2-end against the quality assurance environment, 
After some work done using *GO* to do end-2-end tests against an API, and using the base *GO* test package [testing](https://pkg.go.dev/testing) and two other packages: [httpexpect](https://pkg.go.dev/github.com/gavv/httpexpect/v2) and [testify](https://pkg.go.dev/github.com/stretchr/testify/assert).

Using *testing* as the base framework, *httpexpect* to build, execute and the base assertion. The *testify* package has the *assert* that helps to do some extra assertions. All the packages are useful and easy to use, but there are some missing functionalities, some hard work to implement functionalities for example the tags/labels. The output of the tests execution is cloud be better done, exposing every test and the result and one of the most relevant for me easy to read and understand tests.

After more investigation of other packages available to build tests, the package [Ginkgo](https://onsi.github.io/ginkgo/) has some nice features in regarding the readability of the tests and the tests contexts.

This package can be seen as a Behavior Driven framework since it has a node tree with the ability to build the node tree giving the ability to create a readable context to the tests. In my case that used [Cypress](https://www.cypress.io/) to create tests to webpages *Ginkgo* gives a almost matching way to organize the tests.

The objective of this project is to experiment *Ginkgo* as main testing framework, to know the compatibility of this framework with other packages, what are the advantages and the disadvantages. The experiments done here will ease the future work when selecting a tool to use.

## Getting Started
To use this project the [GO](https://go.dev/) development environment and have some knowledge about the language, the way to run tests, and to manage the dependencies. Beside the base knowledge it is good to have some knowledge regarding building tests and run tests using *GO*.

The example tests done in the project are end-2-end API tests against some two public APIs, [meowfacts](https://github.com/wh-iterabb-it/meowfacts) and [catfacts](https://alexwohlbruck.github.io/cat-facts/docs/endpoints/facts.html). In case of questions regarding the end-2-end (E2E) API tests, instead of looking to a project with a GUI is easy to think that an E2E test needs to have the GUI but, thinking that the interface of the product is the API we could take the same approach used to develop E3E tests to the GUI.

## Build and Test

## Contribute

[Unit testing in Go with Ginkgo: Part 1](https://medium.com/boldly-going/unit-testing-in-go-with-ginkgo-part-1-ce6ff06eb17f)
[meowfacts](https://github.com/wh-iterabb-it/meowfacts)
[How to do unit testing of HTTP requests using Ginkgo?](https://stackoverflow.com/questions/45434849/how-to-do-unit-testing-of-http-requests-using-ginkgo)
[Gomega](https://onsi.github.io/gomega/#ghttp-testing-http-clients)
[Testing With Ginkgo and Gomega](https://medium.com/@dees3g/testing-with-ginkgo-and-gomega-1f1ecc8407a8)

`go get -u github.com/onsi/ginkgo/ginkgo`

`go get -u github.com/onsi/gomega/`

`go get github.com/gavv/httpexpect/v2`

`go install github.com/onsi/ginkgo/ginkgo`

`ginkgo bootstrap`

`ginkgo generate`

Recursive
`ginkgo -r`

Recursive with verbosity
`ginkgo -v -r`

Recursive with max verbosity
`ginkgo run -r -vv`

Recursive and only tests with **params** label
`ginkgo run -r --label-filter="params"`

Recursive and only tests with **noparams** or **id** labels
`ginkgo run -r --label-filter="noparams||id"`

Recursive and only tests with **params** and **1** labels
`ginkgo run -r --label-filter="params&&1"`


APIs
[meowfacts](https://github.com/wh-iterabb-it/meowfacts)
[catfacts](https://alexwohlbruck.github.io/cat-facts/docs/endpoints/facts.html)

We can have only one test file in the package, where the test suite definition and the tests leave. This can be a good thing for organization and reducing the number of files to maintain but, using the VS Codium to work and not adding any extra extension (besides the go support ones) as we scroll throw the tests VS Codium shows at the to the hierarchy of the test that we are located (describe, context, when, and it). This give a very important information to understand the context of what we are working on. But if we add to the test file the test suite definition Codium won't show that information at the top.