package main

import (
	"net/rpc"

	"github.com/tacusci/logging/v2"
)

func init() {
	rpc.Register(Session{})
	rpc.Register(ConnectionData{})
}

type Session struct {
	Token      string
	CameraUUID string
}

func (s Session) GetToken(args string, resp *string) error {
	*resp = s.Token
	return nil
}

type ConnectionData struct {
	UUID, Title string
}

func (c ConnectionData) GetUUID(args string, dst *string) error {
	*dst = c.UUID
	return nil
}

func (c ConnectionData) GetTitle(args string, dst *string) error {
	*dst = c.Title
	return nil
}

type dragonClient struct {
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
	}
}

func (d *dragonClient) ActiveConnections() ([]ConnectionData, error) {
	conns := []ConnectionData{}
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
		logging.Fatal("Unable to connect to dragon daemon: %v...", err)
	}

	conns, err := dc.ActiveConnections()
	if err != nil {
		logging.Fatal("Unable to fetch active connections: %v...", err)
	}

	if len(conns) > 0 {
		logging.Info("REBOOTING CONNECTION: %v", conns[0])
		_, err := dc.RebootConnection(conns[0].UUID)

		if err != nil {
			logging.Error("Unable to reboot connection of UUID %s: %v...", conns[0].UUID, err)
		}
	}

	// dc.Shutdown()

	// app := gui.NewGui()
	// // go func() {
	// // 	time.Sleep(time.Second * 3)
	// // 	app.Close()
	// // }()
	// app.Login().Callback(dc.Authenticate)
	// app.Show(app.Login())
	// app.SetFocusToPages()
	// app.Run()
}
