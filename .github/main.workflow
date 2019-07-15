workflow "build and test" {
  on = "push"
  resolves = ["test", "lint"]
}

action "lint" {
  uses = "docker://golangci/golangci-lint"
  args = "golangci-lint run ./..."
}

action "test" {
  uses = "docker://golang:1.12.7"
  args = "go test -v ./..."
}
