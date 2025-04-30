package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const version = "1.0.0"

func main() {
	// Flags
	keyPath := flag.String("i", "", "Path to your public SSH key (optional)")
	portFlag := flag.Int("p", 0, "SSH port (optional)")
	helpFlag := flag.Bool("h", false, "Show help")
	versionFlag := flag.Bool("v", false, "Show version information")

	// Parse flags
	flag.Parse()

	// Show version if -v is provided
	if *versionFlag {
		showVersion()
		return
	}

	// Show help message if -h is provided
	if *helpFlag {
		showHelp()
		return
	}

	// no input
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go-ssh-copy-id user@host[:port] [-i=\\path\\to\\key.pub] [-p=22]")
		os.Exit(1)
	}

	rawHost := args[0]
	userHost, parsedPort := splitHostPort(rawHost)

	// If -p flag is given, use instead of parsedPort
	port := 22
	if parsedPort != 0 {
		port = parsedPort
	}
	if *portFlag != 0 {
		port = *portFlag
	}

	// get key from ~ 
	finalKeyPath := *keyPath
	if finalKeyPath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error finding home directory:", err)
			os.Exit(1)
		}
		ed25519 := filepath.Join(home, ".ssh", "id_ed25519.pub")
		rsa := filepath.Join(home, ".ssh", "id_rsa.pub")
		if fileExists(ed25519) {
			finalKeyPath = ed25519
		} else if fileExists(rsa) {
			finalKeyPath = rsa
		} else {
			fmt.Println("No default public key found. Use -i to specify one.")
			os.Exit(1)
		}
	}

	// Read the public key file
	keyData, err := ioutil.ReadFile(finalKeyPath)
	if err != nil {
		fmt.Println("Error reading key file:", err)
		os.Exit(1)
	}

	// clean carriage returns
	keyData = []byte(strings.Replace(string(keyData), "\r", "", -1))

	// ssh command
	sshArgs := []string{userHost}
	if port != 22 {
		sshArgs = append([]string{"-p", strconv.Itoa(port)}, sshArgs...)
	}
	sshArgs = append(sshArgs, "mkdir -p ~/.ssh && cat >> ~/.ssh/authorized_keys && chmod 600 ~/.ssh/authorized_keys && chmod 700 ~/.ssh")

	cmd := exec.Command("ssh", sshArgs...)
	cmd.Stdin = strings.NewReader(string(keyData))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("Sending public key via single ssh session...")
	if err := cmd.Run(); err != nil {
		fmt.Println("SSH command failed:", err)
		os.Exit(1)
	}

	fmt.Println("✅ Public key installed successfully.")
}

 
func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// splits user@host[:port] into user@host and port
func splitHostPort(input string) (string, int) {
	parts := strings.Split(input, ":")
	if len(parts) == 2 {
		port, err := strconv.Atoi(parts[1])
		if err == nil {
			return parts[0], port
		}
	}
	return input, 0
}


func showVersion() {
	fmt.Printf("go-ssh-copy-id version %s\n", version)
}

 
func showHelp() {
	fmt.Println("Usage: go-ssh-copy-id user@host[:port] [-i=\\path\\to\\key.pub] [-p=22]")
	fmt.Println()
	fmt.Println("Options:")
	fmt.Println("  -i      Path to your public SSH key (optional)")
	fmt.Println("  -p      SSH port (optional, default is 22)")
	fmt.Println("  -h      Display this help message")
	fmt.Println("  -v      Display version information")
}

