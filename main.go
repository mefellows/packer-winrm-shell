package main

import (
	"github.com/mefellows/packer-winrm-shell/provisioner/winrm-shell"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {

	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(winrmshell.Provisioner))
	server.Serve()
}
