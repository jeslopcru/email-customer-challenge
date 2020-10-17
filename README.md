# customerimporter

## Instructions
package customerimporter reads from the given customers.csv file and returns a
sorted (data structure of your choice) of email domains along with the number
of customers with e-mail addresses for each domain.  Any errors should be
logged (or handled). Performance matters (this is only ~3k lines, but *could*
be 1m lines or run on a small machine).

## How to run

```bash
go run cmd/importer-cli/main.go customers
```
## How to build

```bash
go build go run cmd/importer-cli/main.go
```

## How to run tests

```bash
go test ./... -v
```

## Files

- cmd/importer-cli/main.go - **Entry point of the program**
I added Cobra to create a command line app
- internal/cli/customer.go 
Cobra function is the entrypoint of business logic 
- internal/interview.go - **Business logic of the program**
- internal/interview_test.go - **Unit Tests of business logic**

## Improvements
- Make a docker image
- Improve the way the tests are written to mock some func when necessary.
- Increase the test coverage
