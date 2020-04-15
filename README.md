## Client

[![Lint Status](https://github.com/dansku/microservice_example_client/workflows/golangci-lint/badge.svg)](https://github.com/dansku/microservice_example_client/actions) [![Test Status](https://github.com/dansku/microservice_example_client/workflows/code-test/badge.svg)](https://github.com/dansku/microservice_example_client/actions)

CLI that sends a list of `float32` comma separated numbers to a server that will return with some simple calculations.

Options:
* `-host` - Selects the host server IP address.
* `-port` - The port that the server is listening to for the requests.
* `-operations` - Math operation to run on the list of numbers. Currently possible `sum`, `subtract`, `multiply`, and `divide`.
* `-numbers` - List of numbers comma separated.

Example:
`go run main.go -host=127.0.0.1 -port=5300 -operation subtract -numbers 10.2,19,13.3`
