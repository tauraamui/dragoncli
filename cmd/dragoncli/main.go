package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"os"

	"github.com/tacusci/logging/v2"
	"github.com/tauraamui/dragoncli/internal/common"
	"github.com/tauraamui/dragoncli/internal/gui"
)

func init() {
	rpc.Register(Session{})
}

type Session struct {
	Token      string
	CameraUUID string
}

func (s Session) GetToken(args string, resp *string) error {
	*resp = s.Token
	return nil
}

type dragonClient struct {
	app               *gui.Gui
	client            *rpc.Client
	rpcConnectionAddr string
	session           Session
}

func (d *dragonClient) Connect() error {
	client, err := rpc.DialHTTP("tcp", d.rpcConnectionAddr)
	if err != nil {
		return err
	}

	d.client = client
	return nil
}

func (d *dragonClient) Authenticate(username, password string) {
	// NOTE(tauraamui): obviously temp placeholder for real auth here
	if username == "remoteadmin" && password == "12345" {
		d.session.Token = "validtoken"
		d.app.Show(d.app.Connections())
	}
	usernameAndPassword := fmt.Sprintf("%s|%s", username, password)
	authToken := ""
	d.client.Call("MediaServer.Authenticate", &usernameAndPassword, &authToken)
}

func (d *dragonClient) ActiveConnections() ([]common.ConnectionData, error) {
	conns := []common.ConnectionData{}
	err := d.client.Call("MediaServer.ActiveConnections", &d.session, &conns)
	if err != nil {
		return nil, err
	}

	return conns, nil
}

func (d *dragonClient) RebootConnection(cameraUUID string) (bool, error) {
	ok := false
	d.session.CameraUUID = cameraUUID
	// ensure we don't accidentally leave this set no matter what
	defer func() { d.session.CameraUUID = "" }()

	err := d.client.Call("MediaServer.RebootConnection", &d.session, &ok)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (d *dragonClient) Shutdown() (bool, error) {
	ok := false
	err := d.client.Call("MediaServer.Shutdown", &d.session, &ok)
	if err != nil {
		return false, err
	}

	return ok, nil
}

type exitCode int

const (
	exitOK     exitCode = 0
	exitError  exitCode = 1
	exitCancel exitCode = 2
	exitAuth   exitCode = 4
)

func mainRun() exitCode {
	rpcConnectionAddrPtr := flag.String("rpcaddr", ":3121", "RPC server address")
	flag.Parse()

	dc := dragonClient{
		rpcConnectionAddr: *rpcConnectionAddrPtr,
	}

	err := dc.Connect()
	if err != nil {
		logging.Error("Unable to connect to daemon: %v...", err)
		return exitError
	}
	app := gui.NewGui(dc.ActiveConnections)

	dc.app = app

	app.Login().Callback(dc.Authenticate)
	app.Show(app.Login())
	app.SetFocusToPages()

	if err := app.Run(); err != nil {
		logging.Error("Error occurred during app execution: %v", err)
		return exitError
	}

	return exitOK
}

func main() {
	os.Exit(int(mainRun()))
}
