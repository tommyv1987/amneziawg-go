//go:build !linux

package device

import (
	"github.com/tommyv1987/amneziawg-go/conn"
	"github.com/tommyv1987/amneziawg-go/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
