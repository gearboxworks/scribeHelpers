package gearSsh

import (
	"github.com/newclarity/scribeHelpers/ux"
	"golang.org/x/crypto/ssh"
	//"io"
	"net"

	//"syscall"
	"time"
)

const onlyOnce = "1"


type Ssh struct {
	// SSH client
	ClientInstance *ssh.Client
	ClientSession  *ssh.Session
	ClientAuth     *SshAuth

	// SSH server for SSHFS
	ServerConfig      *ssh.ServerConfig
	ServerListener    net.Listener
	ServerConnection  net.Conn
	ServerAuth        *SshAuth

	FsAuth            *SshAuth
	FsListener        net.Listener
	FsRemote          bool
	FsReadOnly        bool
	FsMount           string

	// Status line related.
	StatusLine  StatusLine
	GearName    string
	GearVersion string

	// Shell related.
	Shell      bool
	Env        Environment
	CmdArgs    []string

	State      *ux.State
	Debug      bool
}
func (s *Ssh) IsNil() *ux.State {
	return ux.IfNilReturnError(s)
}

type SshClientArgs Ssh

type Environment map[string]string

const DefaultUsername = "gearbox"
const DefaultPassword = "box"
const DefaultKeyFile = "./keyfile.pub"
const DefaultSshHost = "localhost"
const DefaultSshPort = "22"
const DefaultStatusLineUpdateDelay = time.Second * 2


func (s *Ssh) IsValid() *ux.State {
	if state := ux.IfNilReturnError(s); state.IsError() {
		return state
	}

	for range onlyOnce {
		s.State = s.State.EnsureNotNil()

		if s.GearName == "" {
			s.State.SetError("name is nil")
			break
		}

		if s.GearVersion == "" {
			s.State.SetError("version is nil")
			break
		}
	}

	return s.State
}
