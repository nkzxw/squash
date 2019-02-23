package local

import (
	"os/exec"
)

type GdbInterface struct{}

func (g *GdbInterface) GetRemoteConnectionCmd(plankName, plankNamespace, podName, podNamespace string, localPort, remotePort int) *exec.Cmd {
	// proxy through the debug container
	return GetPortForwardCmd(plankName, plankNamespace, localPort, remotePort)
}

func (d *GdbInterface) GetDebugCmd(localPort int) *exec.Cmd {
	// TODO
	return nil
}

func (d *GdbInterface) ExpectRunningPlank() bool {
	// TODO
	return false
}