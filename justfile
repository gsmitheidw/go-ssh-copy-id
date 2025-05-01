# justfile

set shell := ["powershell", "-NoProfile", "-Command"]

default:
    go build -o go-ssh-copy-id.exe -ldflags="-s -w" main.go

clean:
    remove-item go-ssh-copy-id.exe

release:
    go build -o go-ssh-copy-id.exe -ldflags="-s -w -X main.version=1.0.0"

build-macos:
    New-Item -ItemType Directory -Force -Path "dist/macos-amd64" > $null;
    $env:GOOS = "darwin"; $env:GOARCH = "amd64";
    go build -o "dist/macos-amd64/go-ssh-copy-id" main.go

