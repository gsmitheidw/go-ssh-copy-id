# justfile

set shell := ["powershell", "-NoProfile", "-Command"]

default:
    go build -o go-ssh-copy-id.exe -ldflags="-s -w" main.go

clean:
    remove-item go-ssh-copy-id.exe

release:
    go build -o go-ssh-copy-id.exe -ldflags="-s -w -X main.version=1.0.0"

