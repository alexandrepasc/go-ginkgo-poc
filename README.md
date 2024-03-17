# poc-go-ginkgo
<!-- TOC -->

- [poc-go-ginkgo](#poc-go-ginkgo)
    - [Introduction](#introduction)
    - [Getting Started](#getting-started)
    - [Build and Test](#build-and-test)
    - [Contribute](#contribute)
    - [Appendix](#appendix)
        - [Links](#links)
        - [Commands](#commands)

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
Install the project dependencies that are defined in the file `go.mod` localized in the root folder of the project.
- `go get ./...`

The tests are located in the `tests` folder, sub-folders one for each endpoint, and in the sub-folders the test file(s) with the tests. This organization approach might not be the best to create unit tests, but since we are approaching this as E2E black box tests there are no code to test so tried to have a similar organization as other E2E tools (ex. *Cypress*).

To create a package test suite and test file we can use 2 commands from *Ginkgo*, we can open the package respective folder and run the following commands:
- `ginkgo bootstrap` Generate the test suite file
- `ginkgo generate` Generate the test file

To run the tests we can use directly the `test` command from *GO*, or using the `ginkgo` by installing it with `go install github.com/onsi/ginkgo/ginkgo`.

- `go test ./...`
- `ginkgo -r`
- `ginkgo -v -r`
- `ginkgo run -r -vv`

One functionality that the package *testing* has but it's very hard to maintain in a large test project, is the tags used to filter tests to execute. Using *Ginkgo* is very easy to add the tags, that in this framework is called `Label`s and can be added to tests (`It`), context, and describe. To execute the tests using the `Label`s to filter is:
- `ginkgo run -r --label-filter="params"`
- `ginkgo run -r --label-filter="noparams||id"`
- `ginkgo run -r --label-filter="params&&1"`

Executing tests with *testing* or *Ginkgo* there are no guarantees of the order that the tests are executed. *Ginkgo* adds the `Ordered` decoration that sets the block to execute the specs sequentially by the order that they are written, the `Ordered` specs may be run in parallel with respect respect to other specs but they will always run sequentially in the same parallel process.

When a spec fails the normal behaviour is to run the next spec, but using an `Ordered`container all the specs are considered not independent so *Ginkgo* will skip all the next specs. This can be override using the decoration `ContinueOnFailure`.

In case we need to a container by itself, with no other parallel specs running, the decoration `Serial` can be used to identify the container.

## Contribute
If you want to change, correct, improve the project create an `issue` in the project `Issues` screen with the proposal and the necessary documentation. If the proposal or correction has already the implementation developed link the branch with the change in the `issue`.

## Appendix
### Links
- [Unit testing in Go with Ginkgo: Part 1](https://medium.com/boldly-going/unit-testing-in-go-with-ginkgo-part-1-ce6ff06eb17f)
- [How to do unit testing of HTTP requests using Ginkgo?](https://stackoverflow.com/questions/45434849/how-to-do-unit-testing-of-http-requests-using-ginkgo)
- [Gomega](https://onsi.github.io/gomega/#ghttp-testing-http-clients)
- [Testing With Ginkgo and Gomega](https://medium.com/@dees3g/testing-with-ginkgo-and-gomega-1f1ecc8407a8)
- [meowfacts](https://github.com/wh-iterabb-it/meowfacts)
- [catfacts](https://alexwohlbruck.github.io/cat-facts/docs/endpoints/facts.html)

### Commands
- Get *Ginkgo* in the project
```
go get -u github.com/onsi/ginkgo/ginkgo
```

- Get *Gomega* in the project
```
go get -u github.com/onsi/gomega/
```

- Get *Httpexpect* in the project
```
go get github.com/gavv/httpexpect/v2
```

- Install *Ginkgo*
```
go install github.com/onsi/ginkgo/ginkgo
```

- Create a package test suite
```
ginkgo bootstrap
```

- Create the testing file for the package
```
ginkgo generate
```

- Run *Ginkgo* recursively
```
ginkgo -r
```

- Recursive with verbosity
```
ginkgo -v -r
```

- Recursive with max verbosity
```
ginkgo run -r -vv
```

- Recursive and only tests with **params** label
```
ginkgo run -r --label-filter="params"
```

- Recursive and only tests with **noparams** or **id** labels
```
ginkgo run -r --label-filter="noparams||id"
```

Recursive and only tests with **params** and **1** labels
```
ginkgo run -r --label-filter="params&&1"
```