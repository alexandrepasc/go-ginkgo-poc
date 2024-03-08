# poc-go-ginkgo

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