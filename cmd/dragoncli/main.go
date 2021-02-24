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
	Token string
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
		logging.Fatal("Unable to fetch active connections: %v", err)
	}

	if len(conns) > 0 {
		logging.Info("CONNECTION: %v", conns[0])
	}
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
