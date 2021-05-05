package gearSsh

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)


// From https://sosedoff.com/2015/05/25/ssh-port-forwarding-with-go.html
// Handle local client connections and tunnel data to the remote server
// Will use io.Copy - http://golang.org/pkg/io/#Copy
func (s *Ssh) handleClient(client net.Conn, remote net.Conn) {
	//goland:noinspection GoUnhandledErrorResult
	defer client.Close()
	chDone := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(client, remote)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy remote->local: %s", err))
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy local->remote: %s", err))
		}
		chDone <- true
	}()

	<-chDone
}


func (s *Ssh) SshFsTunnel() {

	for range onlyOnce {
		if !s.FsRemote {
			break
		}
		var err error

		s.FsAuth = NewSshAuth()
		s.FsAuth.Host = "localhost"
		s.FsAuth.Port = "0"
		s.FsAuth.Username = s.ServerAuth.Username
		s.FsAuth.Password = s.ServerAuth.Password
		s.FsAuth.PublicKey = s.ServerAuth.PublicKey

		// Listen on remote server port
		s.FsListener, err = s.ClientInstance.Listen("tcp", s.FsAuth.GetHost())
		if err != nil {
			s.State.SetError("Listen open port ON remote server error: %s", err)
			break
		}
		s.FsAuth.SetUrl(s.FsListener.Addr())

		err = os.Setenv("GEARBOX_MOUNT_HOST", s.FsAuth.Host)
		err = os.Setenv("GEARBOX_MOUNT_PORT", s.FsAuth.Port)
		err = os.Setenv("GEARBOX_MOUNT_USER", s.FsAuth.Username)
		err = os.Setenv("GEARBOX_MOUNT_PASSWORD", s.FsAuth.Password)
		//err = os.Setenv("GEARBOX_MOUNT_PATH", s.FsMount)

		s.State = s.GetEnv()
		if err != nil {
			break
		}

		//s.PrintInfo()

		go s.SshFsTunnelListener()
	}
}


func (s *Ssh) SshFsTunnelListener() {

	for range onlyOnce {
		// handle incoming connections on reverse forwarded tunnel
		for {
			//s.PrintInfo()

			// Open a (local) connection to localEndpoint whose content will be forwarded so serverEndpoint
			local, err := net.Dial("tcp", "localhost:" + s.ServerAuth.Port) // s.ServerAuth.Host + ":" + s.ServerAuth.Port)
			if err != nil {
				log.Fatalln(fmt.Printf("Dial INTO local service error: %s", err))
			}

			client, err := s.FsListener.Accept()
			if err != nil {
				log.Fatalln(err)
			}

			s.handleClient(client, local)
		}
	}
}


func (s *Ssh) PrintInfo() {
	for range onlyOnce {
		fmt.Printf("ClientAuth: %s\r\n",
			s.ClientAuth.GetHost(),
		)

		fmt.Printf("FsRemote: %v\r\n",
			s.FsRemote,
		)
		fmt.Printf("\r\n")

		fmt.Printf("FsAuth: %s\r\n",
			s.FsAuth.GetHost(),
		)
		if s.FsListener != nil {
			fmt.Printf("FsListener.Addr: %v\r\n",
				s.FsListener.Addr(),
			)
		}
		fmt.Printf("\r\n")

		fmt.Printf("ServerAuth: %s\r\n",
			s.ServerAuth.GetHost(),
		)
		if s.ServerListener != nil {
			fmt.Printf("ServerListener.Addr: %v\r\n",
				s.ServerListener.Addr(),
			)
		}
		fmt.Printf("\r\n")

		if s.ServerConnection != nil {
			fmt.Printf("ServerConnection LocalAddr: %v -> RemoteAddr: %v\r\n",
				s.ServerConnection.LocalAddr(),
				s.ServerConnection.RemoteAddr(),
			)
		}
		fmt.Printf("\r\n")

		if s.ClientInstance != nil {
			fmt.Printf("ClientInstance LocalAddr: %v -> RemoteAddr: %v\r\n",
				s.ClientInstance.LocalAddr(),
				s.ClientInstance.RemoteAddr(),
			)
		}
		fmt.Printf("\r\n")

		for k, v := range s.Env {
			if strings.Contains(k, "_MOUNT_") {
				fmt.Printf("%s => %s\r\n", k, v)
			}
		}
		fmt.Printf("\r\n")
	}
}
