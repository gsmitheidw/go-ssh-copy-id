# go-ssh-copy-id

go-ssh-copy-id is an implementation of ssh-copy-id for Windows. 
It is a static binary with no dependencies other than native ssh client
that already is in modern Windows.

## Usage

```
Usage: go-ssh-copy-id user@host[:port] [-i=/path/to/key.pub] [-p=22]

Options:
-i      Path to your public SSH key (optional)
-p      SSH port (optional, default is 22)
-h      Display this help message
-v      Display version information

```


```mermaid
 graph TD
    A[fa:fa-terminal go-ssh-copy-id] --> B{fa:fa-key Check for ed25519 key}
    B -->|Yes| C[fa:fa-file Read ed25519 key]
    B -->|No| D{fa:fa-key Check for rsa key}
    D -->|Yes| E[fa:fa-file Read rsa key]
    D -->|No| F[fa:fa-exclamation-triangle Show error and exit]
    C --> G[fa:fa-broom Clean key data]
    E --> G
    G --> H[fa:fa-cogs Prepare SSH command]
    H --> I[fa:fa-network-wired Execute SSH command]
    I --> J{fa:fa-check SSH command successful?}
    J -->|Yes| K[fa:fa-thumbs-up Show success message]
    J -->|No| L[fa:fa-thumbs-down Show error message]
    K --> M[fa:fa-flag-checkered End]
    L --> M
    I --> N[fa:fa-server Append key to authorized_keys]
    N --> M 
```

