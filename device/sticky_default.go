//go:build !linux

package device

import (
	"github.com/nymtech/amneziawg-go/conn"
	"github.com/nymtech/amneziawg-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
