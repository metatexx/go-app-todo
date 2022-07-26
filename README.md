# Go-App sample ToDo application

We are using a 'Magefile' which is a way to use Go for creating complex makefiles.
Checkout [Mage](https://magefile.org/) for more information. 

| Action   | Command with Zero install Mage | Mage         |
|----------|--------------------------------|--------------|
| Build    | `go run mage.go build`         | `mage build` |
| Run      | `go run mage.go run`           | `mage run`   |
| Clean | `go run mage.go clean`         | `mage clean` |

We also added minimal `makefile` if you like that more. It supports `make run` and `make build`.

This work is published under the [MIT-License](LICENSE) and copyright (c) 2022 by METATEXX GmbH.