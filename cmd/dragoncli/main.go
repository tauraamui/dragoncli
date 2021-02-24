package main

import (
	"net/rpc"

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
	rpcConnectionPort string
	session           Session
}

func (d *dragonClient) Connect() error {
	client, err := rpc.DialHTTP("tcp", d.rpcConnectionPort)
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

func main() {
	dc := dragonClient{
		rpcConnectionPort: ":3121",
	}
	err := dc.Connect()
	if err != nil {
		logging.Fatal("Unable to connect to daemon: %v...", err)
	}
	app := gui.NewGui(dc.ActiveConnections)

	dc.app = app

	app.Login().Callback(dc.Authenticate)
	app.Show(app.Login())
	app.SetFocusToPages()

	// err := dc.Connect()
	// if err != nil {
	// 	logging.Fatal("Unable to connect to dragon daemon: %v...", err)
	// }

	// conns, err := dc.ActiveConnections()
	// if err != nil {
	// 	logging.Fatal("Unable to fetch active connections: %v...", err)
	// }

	// if len(conns) > 0 {
	// 	logging.Info("REBOOTING CONNECTION: %v", conns[0])
	// 	_, err := dc.RebootConnection(conns[0].UUID)

	// 	if err != nil {
	// 		logging.Error("Unable to reboot connection of UUID %s: %v...", conns[0].UUID, err)
	// 	}
	// }

	// dc.Shutdown()

	app.Run()
}
