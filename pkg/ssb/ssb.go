package ssb

import (
	"fmt"
	"time"

	"golang.org/x/crypto/ssh"
)

// New defines SSH client configurations
func New(username string, password string, timeout time.Duration) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         timeout,
		User:            username,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
	}
}

// Connect for dialing to SSH server
func Connect(host string, port int, config *ssh.ClientConfig) (bool, error) {
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", host, port), config)
	if err != nil {
		return false, err
	}
	_ = client.Close()
	return true, nil
}
