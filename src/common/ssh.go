package common

// go ssh 远程登录主机
import (
	"bytes"
	"configcenter/src/common/blog"
	"fmt"
	"net"
)

func SSHConnect( user, password, host string, port int ) ( *ssh.Session, error ) {
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
		return nil
	}

	clientConfig = &ssh.ClientConfig{
		User:               user,
		Auth:               auth,
		// Timeout:             30 * time.Second,
		HostKeyCallback:    hostKeyCallbk,
	}

	// connet to ssh
	addr = fmt.Sprintf( "%s:%d", host, port )

	if client, err = ssh.Dial( "tcp", addr, clientConfig ); err != nil {
		return nil, err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}

func RunSsh(user, password, host string, port int) (s string, err error){

	var stdOut, stdErr bytes.Buffer

	session, err := SSHConnect( user, password, host, port)
	if err != nil {
		blog.Errorf("[ssh ] SSHConnect failed, err: %v",err)
	}
	defer session.Close()

	session.Stdout = &stdOut
	session.Stderr = &stdErr
	minionId := host
	runErr := session.Run("curl -s 192.168.31.210/yumrepo/conf/install-saltminion.sh | bash -s "+minionId)
	if runErr != nil {
		panic(runErr)
	}
	return stdOut.String(),runErr

}

