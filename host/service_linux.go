package host

import (
	"github.com/nextdns/nextdns/host/service"
	"github.com/nextdns/nextdns/host/service/ddwrt"
	"github.com/nextdns/nextdns/host/service/edgeos"
	"github.com/nextdns/nextdns/host/service/entware"
	"github.com/nextdns/nextdns/host/service/merlin"
	"github.com/nextdns/nextdns/host/service/procd"
	"github.com/nextdns/nextdns/host/service/synology"
	"github.com/nextdns/nextdns/host/service/systemd"
	"github.com/nextdns/nextdns/host/service/sysv"
	"github.com/nextdns/nextdns/host/service/upstart"
)

func NewService(c service.Config) (service.Service, error) {
	if s, err := systemd.New(c); err == nil {
		return s, nil
	}
	if s, err := procd.New(c); err == nil {
		return s, nil
	}
	if s, err := merlin.New(c); err == nil {
		return s, nil
	}
	if s, err := edgeos.New(c); err == nil {
		return s, nil
	}
	if s, err := ddwrt.New(c); err == nil {
		return s, nil
	}
	if s, err := synology.New(c); err == nil {
		return s, nil
	}
	if s, err := entware.New(c); err == nil {
		return s, nil
	}
	if s, err := upstart.New(c); err == nil {
		return s, nil
	}
	if s, err := sysv.New(c); err == nil {
		return s, nil
	}
	return nil, service.ErrNotSuported
}
