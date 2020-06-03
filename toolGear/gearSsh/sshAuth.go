package gearSsh

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
)

type SshAuth struct {
	// SSH related.
	Username    string
	Password    string
	Host        string
	Port        string
	PublicKey   string
}
type SshAuthArgs SshAuth


func NewSshAuth(args ...SshAuth) *SshAuth {

	var _args SshAuth
	if len(args) > 0 {
		_args = args[0]
	}

	if _args.Username == "" {
		_args.Username = DefaultUsername
	}

	if _args.Password == "" {
		_args.Password = DefaultPassword
	}

	if _args.PublicKey == "" {
		_args.PublicKey = DefaultKeyFile
	}

	if _args.Host == "" {
		_args.Host = DefaultSshHost
	}

	if _args.Port == "" {
		_args.Port = DefaultSshPort
	}

	//sshAuth := &SshAuth{}
	//*sshAuth = SshAuth(_args)

	return &_args
}


func readPublicKeyFile(file string) (ssh.AuthMethod, error) {

	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		// fmt.Printf("# Error reading file '%s': %s\n", file, err)
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		// fmt.Printf("# Error parsing key '%s': %s\n", signer, err)
		return nil, err
	}

	return ssh.PublicKeys(signer), err
}
