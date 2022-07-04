package main

import (
	"log"
	"sync"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
	"serversdf/config"
	"serversdf/output"
	"serversdf/server"
)

func main() {
	cfg, err := config.ReadConfigurationFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}

	outputIsToTerminal, err := output.IsToTerminal()
	if err != nil {
		log.Fatal(err)
	}

	var outputContext output.OutputContext
	if outputIsToTerminal {
		numberOfLinesToWrite := cfg.GetNumberOfServers()
		outputContext = output.NewTerminalOutput(numberOfLinesToWrite, cfg.GetLongestServerName())
	} else {
		outputContext = output.NewPipeOutput(cfg.GetLongestServerName())
	}

	// HostkeyCallback is needed to add server to knownhosts when connecting to a server
	hostKeyCallback, err := knownhosts.New(cfg.KnownHosts)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	for i, serverInfo := range cfg.Servers {
		// We only want to write a progress status if we're running in the terminal
		if outputIsToTerminal {
			outputContext.WriteInfo(serverInfo.Name, "fetching disk usage", i)
		}

		wg.Add(1)
		go func(serverInfo server.ServerInfo, hostKeyCallback ssh.HostKeyCallback, index int) {
			defer wg.Done()
			serverConnection, err := server.NewServerConnection(serverInfo, hostKeyCallback)
			if err != nil {
				outputContext.WriteError(serverInfo.Name, err.Error(), index)
				return
			}

			usage, err := serverConnection.GetDiskUsage()
			if err != nil {
				outputContext.WriteError(serverInfo.Name, err.Error(), index)
				return
			}

			outputContext.WriteSuccess(serverInfo.Name, usage+" used", index)
		}(serverInfo, hostKeyCallback, i)
	}

	wg.Wait()
}
