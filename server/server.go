package server

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

type ServerInfo struct {
	Name        string
	Host        string
	User        string
	Disk        string
	Password    string
	KeyFile     string `toml:"key_file"`
	KeyPassword string `toml:"key_password"`
}

func (serverInfo *ServerInfo) GetPrivateKeySigner() (ssh.Signer, error) {
	privateKey, err := ioutil.ReadFile(serverInfo.KeyFile)
	if err != nil {
		return nil, err
	}

	if serverInfo.KeyPassword == "" {
		return ssh.ParsePrivateKey(privateKey)
	}

	return ssh.ParsePrivateKeyWithPassphrase(privateKey, []byte(serverInfo.KeyPassword))
}

func (serverInfo *ServerInfo) HasNoKeyFile() bool {
	return serverInfo.KeyFile == ""
}
