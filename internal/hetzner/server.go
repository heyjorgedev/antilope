package hetzner

import (
	"context"
	"errors"

	"github.com/hetznercloud/hcloud-go/hcloud"
	"github.com/jorgemurta/antilope/internal/server"
)

func EnsureServerExists(ctx context.Context, client *hcloud.Client, serverName string) (server.Server, error) {

	srv, _, err := client.Server.GetByName(ctx, serverName)
	if err != nil {
		return server.Server{}, err
	}

	if srv != nil {
		return server.Server{
			Name: serverName,
			IpAddress: srv.PublicNet.IPv4.IP.String(),
		}, nil;
	}

	newServer, _, _ := client.Server.Create(ctx, hcloud.ServerCreateOpts{
		Name: serverName,
		ServerType: &hcloud.ServerType{Name: "cx11"},
		Image: &hcloud.Image{Name: "ubuntu-20.04"},
		Location: &hcloud.Location{Name: "nbg1"},
	})

	return server.Server {
		Name: serverName,
		IpAddress: newServer.Server.PublicNet.IPv4.IP.String(),
	}, nil
}

func DeleteServer(ctx context.Context, client *hcloud.Client, serverName string) error {

	srv, _, err := client.Server.GetByName(ctx, serverName)

	if err != nil {
		return err
	}

	if srv == nil {
		return errors.New("server not found")
	}

	_, _, err = client.Server.DeleteWithResult(ctx, srv)

	return err
}