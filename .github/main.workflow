workflow "build and test" {
  on = "push"
  resolves = ["test"]
}

action "test" {
  uses = "docker://golang:1.12.7"
  args = "go test -v ./..."
}
