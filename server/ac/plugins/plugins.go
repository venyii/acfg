package plugins

import (
	"os"
	"os/exec"

	"github.com/venyii/acfg/server/ac/portalloc"
)

type Plugins []Plugin

type Plugin interface {
	Pid() int
	Cmd() *exec.Cmd
	LogFile() *os.File

	Name() string
	RecvPorts(pm portalloc.PortsMap)
	PortReqs() portalloc.PortReqs
	CreateTmpConfig(tmpDir string) error
	Start(logsDir string) error
	Stop() error
}
