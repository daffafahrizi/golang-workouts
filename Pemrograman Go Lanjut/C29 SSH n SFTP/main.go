package main

import (
	"golang.org/x/crypto/ssh"
	"log"
	"os"
)

func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}

func main() {
	const SSH_ADDRESS = "0.0.0.0:22"
	const SSH_USERNAME = "user"
	const SSH_KEY = ""
	const SSH_PASSWORD = "password"

	sshConfig := &ssh.ClientConfig{
		User:            SSH_USERNAME,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Auth: []ssh.AuthMethod{
			ssh.Password(SSH_PASSWORD),
			// PublicKeyFile(SSH_KEY),
		},
	}

	client, err := ssh.Dial("tcp", SSH_ADDRESS, sshConfig)
	if client != nil {
		defer client.Close()
	}
	if err != nil {
		log.Fatal("Failed to dial. " + err.Error())
	}

	session, err := client.NewSession()
	if session != nil {
		defer session.Close()
	}
	if err != nil {
		log.Fatal("Failed to create session. " + err.Error())
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Run("ls -l ~/")
	if err != nil {
		log.Fatal("Command execution error. " + err.Error())
	}
}