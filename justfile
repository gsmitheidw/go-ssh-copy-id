# justfile

set shell := ["powershell", "-NoProfile", "-Command"]

default:
    go build -o go-ssh-copy-id.exe -ldflags="-s -w" main.go

clean:
    if (Test-Path go-ssh-copy-id.exe) { Remove-Item go-ssh-copy-id.exe }
    if (Test-Path macos-amd64) { Remove-Item macos-amd64 -Recurse -Force }

release:
    go build -o go-ssh-copy-id.exe -ldflags="-s -w -X main.version=1.0.1"


build-macos:
    New-Item -ItemType Directory -Force -Path "macos-amd64" > $null;
    $env:GOOS = "darwin"; $env:GOARCH = "amd64"; go build -o "macos-amd64/go-ssh-copy-id" -ldflags="-s -w" main.go
