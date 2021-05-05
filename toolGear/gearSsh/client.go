package gearSsh

import (
	"fmt"
	"github.com/gearboxworks/scribeHelpers/ux"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"time"
)


func NewSshClient(args ...SshClientArgs) *Ssh {

	var _args SshClientArgs
	if len(args) > 0 {
		_args = args[0]
	}

	_args.ClientAuth = NewSshAuth(*_args.ClientAuth)

	if _args.StatusLine.UpdateDelay == 0 {
		_args.StatusLine.UpdateDelay = DefaultStatusLineUpdateDelay
	}

	sshClient := &Ssh{}
	*sshClient = Ssh(_args)

	return sshClient
}


func (s *Ssh) Connect() error {
	var err error
	if state := s.IsNil(); state.IsError() {
		return state.GetError()
	}

	for range onlyOnce {
		sshConfig := &ssh.ClientConfig{}

		var auth []ssh.AuthMethod

		// Try SSH key file first.
		var keyfile ssh.AuthMethod
		keyfile, err = readPublicKeyFile(s.ClientAuth.PublicKey)

		if err == nil && keyfile != nil {
			// Authenticate using SSH key.
			auth = []ssh.AuthMethod{keyfile}
		} else {
			// Authenticate using password
			auth = []ssh.AuthMethod{ssh.Password(s.ClientAuth.Password)}
		}

		sshConfig = &ssh.ClientConfig {
			User: s.ClientAuth.Username,
			Auth: auth,
			// HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil },
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         time.Second * 10,
		}

		s.ClientInstance, err = ssh.Dial("tcp", s.ClientAuth.GetHost(), sshConfig)
		if err != nil {
			break
		}

		s.ClientSession, err = s.ClientInstance.NewSession()
		//noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer s.ClientSession.Close()
		//noinspection GoDeferInLoop,GoUnhandledErrorResult
		defer s.ClientInstance.Close()
		if err != nil {
			break
		}

		// Set IO
		s.ClientSession.Stdout = os.Stdout
		s.ClientSession.Stderr = os.Stderr
		s.ClientSession.Stdin = os.Stdin


		// Trap signals.
		//sigc := make(chan os.Signal, 1)
		//signal.Notify(sigc,
		//    //syscall.SIGHUP,
		//    //syscall.SIGINT,
		//    //syscall.SIGTERM,
		//    //syscall.SIGQUIT,
		//    )
		//go func() {
		//	for {
		//		sig := <-sigc
		//		if sig == syscall.SIGURG {
		//			continue
		//		}
		//		fmt.Printf("SIGNAL: '%s'\n", sig.String())
		//		err = s.ClientSession.Signal(ssh.Signal(sig.String()))
		//		fmt.Printf("err: '%v'\n", err)
		//	}
		//}()

		// Set up terminal modes
		modes := ssh.TerminalModes {
			ssh.ECHO:          1,
			ssh.TTY_OP_ISPEED: 19200,
			ssh.TTY_OP_OSPEED: 19200,
			ssh.ISIG: 1,
		}


		// Request pseudo terminal
		fileDescriptor := int(os.Stdin.Fd())
		if terminal.IsTerminal(fileDescriptor) {
			originalState, err := terminal.MakeRaw(fileDescriptor)
			if err != nil {
				break
			}
			//noinspection GoDeferInLoop,GoUnhandledErrorResult
			defer terminal.Restore(fileDescriptor, originalState)

			s.StatusLine.TermWidth, s.StatusLine.TermHeight, err = terminal.GetSize(fileDescriptor)
			if err != nil {
				//break	- IGNORE for now.
				s.StatusLine.TermWidth = 80
				s.StatusLine.TermHeight = 25
			}

			// xterm-256color
			err = s.ClientSession.RequestPty("xterm-256color", s.StatusLine.TermHeight, s.StatusLine.TermWidth, modes)
			if err != nil {
				break
			}
		}

		if s.StatusLine.Enable {
			go s.StatusLineUpdate()
			go s.statusLineWorker()
		}

		if s.FsMount != "" {
			s.SshFsTunnel()
		}


		// Determine if the current shell is a pipe or terminal.
		//var shLvl string
		//shLvl, err = strconv.Atoi(os.Getenv("SHLVL"))
		// SHLVL is really a reliable way. So we're going to set a new env variable.
		// Any containers built prior to 2021/05/04 (May) will need to be rebuilt to avoid the shell level warning.
		fi, _ := os.Stdin.Stat()
		if (fi.Mode() & os.ModeCharDevice) == 0 {
			s.Env["GB_PIPE"] = "YES"
			delete(s.Env, "GB_TERMINAL")
			//_ = os.Setenv("GB_PIPE", "YES")
			//_ = os.Unsetenv("GB_TERMINAL")
		} else {
			delete(s.Env, "GB_PIPE")
			s.Env["GB_TERMINAL"] = "YES"
			//_ = os.Unsetenv("GB_PIPE")
			//_ = os.Setenv("GB_TERMINAL", "YES")
		}


		// Assign all shell variables to outgoing SSH.
		for k, v := range s.Env {
			if s.Debug {
				fmt.Printf("DEBUG: ENV %s => %s\r\n", k, v)
			}
			err = s.ClientSession.Setenv(k, v)
			if err != nil {
				break
			}
		}


		// Start remote shell
		if len(s.CmdArgs) == 0 {
			err = s.ClientSession.Shell()
			if err != nil {
				break
			}

			err = s.ClientSession.Wait()
			if err != nil {
				break
			}

		} else {
			cmd := ""
			if s.Debug {
				fmt.Printf("DEBUG: s.CmdArgs == %s\r\n", strings.Join(s.CmdArgs, " "))
			}
			for i, v := range s.CmdArgs {
				if s.Debug {
					fmt.Printf("DEBUG: Command arg[%d] %s\r\n", i, v)
				}

				if strings.Contains(v, " ") {
					v = `'` + v + `'`
				}
				cmd = fmt.Sprintf("%s %s", cmd, v)
			}

			if s.Debug {
				fmt.Printf("DEBUG: Command %s\r\n", cmd)
			}

			err = s.ClientSession.Run(cmd)
			if err != nil {
				break
			}
		}


		//if len(s.CmdArgs) == 0 {
		//	// Set up terminal modes
		//	modes := ssh.TerminalModes {
		//		ssh.ECHO:          1,
		//		ssh.TTY_OP_ISPEED: 19200,
		//		ssh.TTY_OP_OSPEED: 19200,
		//	}
		//
		//	// Request pseudo terminal
		//	fileDescriptor := int(os.Stdin.Fd())
		//	if terminal.IsTerminal(fileDescriptor) {
		//		originalState, err := terminal.MakeRaw(fileDescriptor)
		//		if err != nil {
		//			break
		//		}
		//		//noinspection GoDeferInLoop,GoUnhandledErrorResult
		//		defer terminal.Restore(fileDescriptor, originalState)
		//
		//		s.StatusLine.TermWidth, s.StatusLine.TermHeight, err = terminal.GetSize(fileDescriptor)
		//		if err != nil {
		//			break
		//		}
		//
		//		// xterm-256color
		//		err = s.ClientSession.RequestPty("xterm-256color", s.StatusLine.TermHeight, s.StatusLine.TermWidth, modes)
		//		if err != nil {
		//			break
		//		}
		//	}
		//
		//	go s.StatusLineUpdate()
		//	go s.statusLineWorker()
		//
		//	// Start remote shell
		//	err = s.ClientSession.Shell()
		//	if err != nil {
		//		break
		//	}
		//
		//	err = s.ClientSession.Wait()
		//	if err != nil {
		//		break
		//	}
		//
		//} else {
		//	cmd := ""
		//	if len(s.CmdArgs) > 0 {
		//		for _, v := range s.CmdArgs {
		//			cmd = fmt.Sprintf("%s %s", cmd, v)
		//		}
		//	}
		//
		//	err = s.ClientSession.Run(cmd)
		//	if err != nil {
		//		break
		//	}
		//}

		s.resetView()
	}

	return err
}


func (s *Ssh) GetEnv() *ux.State {
	if state := s.IsNil(); state.IsError() {
		return state
	}

	for range onlyOnce {
		s.Env = make(Environment)
		for _, item := range os.Environ() {
			if strings.HasPrefix(item, "TMPDIR=") {
				continue
			}

			sa := strings.SplitN(item, "=", 2)
			s.Env[sa[0]] = sa[1]
		}
	}

	return s.State
}
