//go:build !linux

package device

import (
	"github.com/Lordy82/wireguard-go/conn"
	"github.com/Lordy82/wireguard-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
