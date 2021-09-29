# Phaidra Task

## Usage
- Golang is required to run this. You can install Go by following the instructions (here)[https://golang.org/doc/install]
- Run the cli using `go run ./cmd/service`
- It expects inputs in the form of `EquipmentName StartTime EndTime`, for example `Equipment2 10:30 19:15`.
- Default data is given. If you want to use your own json data, you can provide the path to the file using `-filepath` flag. For example: `go run ./cmd/service -filepath ./custom-data.json`

## Running Tests
- Tests can be run using `go test ./...`