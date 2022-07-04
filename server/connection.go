package server

import (
	"bytes"
	"fmt"
	"strings"

	"golang.org/x/crypto/ssh"
)

type ServerConnection struct {
	ServerInfo   ServerInfo
	ClientConfig *ssh.ClientConfig
}

func NewServerConnection(serverInfo ServerInfo, hostKeyCallback ssh.HostKeyCallback) (*ServerConnection, error) {
	if serverInfo.HasNoKeyFile() {
		return &ServerConnection{
			serverInfo,
			clientConfigWithPasswordAuth(serverInfo, hostKeyCallback),
		}, nil
	}

	config, err := clientConfigWithPublicKeyAuth(serverInfo, hostKeyCallback)
	if err != nil {
		return nil, err
	}

	return &ServerConnection{serverInfo, config}, nil
}

func (sc *ServerConnection) GetDiskUsage() (string, error) {
	var buff bytes.Buffer
	var port string = "22"

	conn, err := ssh.Dial("tcp", sc.ServerInfo.Host+":"+port, sc.ClientConfig)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	session.Stdout = &buff
	command := freeDiskSpaceCommand(sc.ServerInfo.Disk)
	if err := session.Run(command); err != nil {
		return "", nil
	}

	return strings.TrimSuffix(buff.String(), "\n"), nil
}

func freeDiskSpaceCommand(device string) string {
	regex := fmt.Sprintf("^%s\\b", device)
	awk := awkCmdToFormatDFOutput()
	return `df -h | grep "` + regex + `" | ` + awk + ` | head -n 1`
}

func awkCmdToFormatDFOutput() string {
	// $2 is total disk space
	// $3 is used disk space
	// $5 is percentage of disk space used
	return `awk '{print $3 " / " $2 " (" $5 ")"}'`
}

func clientConfigWithPasswordAuth(serverInfo ServerInfo, hostKeyCallback ssh.HostKeyCallback) *ssh.ClientConfig {
	return &ssh.ClientConfig{
		User: serverInfo.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(serverInfo.Password),
		},
		HostKeyCallback: hostKeyCallback,
	}
}

func clientConfigWithPublicKeyAuth(serverInfo ServerInfo, hostKeyCallback ssh.HostKeyCallback) (*ssh.ClientConfig, error) {
	signer, err := serverInfo.GetPrivateKeySigner()
	if err != nil {
		return nil, err
	}

	return &ssh.ClientConfig{
		User: serverInfo.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
	}, nil
}
