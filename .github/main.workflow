workflow "build and test" {
  on = "push"
  resolves = ["build"]
}

action "lint" {
  uses = "docker://golangci/golangci-lint"
  args = "golangci-lint run ./..."
}

action "test" {
  uses = "docker://golang:1.12.7"
  args = "go test -v ./..."
}

action "build" {
  needs = ["lint", "test"]
  uses = "docker://golang:1.12.7"
  args = "go build ./..."
}
