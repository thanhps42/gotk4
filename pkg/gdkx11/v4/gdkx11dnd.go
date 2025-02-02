// Code generated by girgen. DO NOT EDIT.

package gdkx11

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
)

// #include <stdlib.h>
// #include <gdk/x11/gdkx.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeX11Drag = coreglib.Type(C.gdk_x11_drag_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeX11Drag, F: marshalX11Drag},
	})
}

type X11Drag struct {
	_ [0]func() // equal guard
	gdk.Drag
}

var (
	_ gdk.Dragger = (*X11Drag)(nil)
)

func wrapX11Drag(obj *coreglib.Object) *X11Drag {
	return &X11Drag{
		Drag: gdk.Drag{
			Object: obj,
		},
	}
}

func marshalX11Drag(p uintptr) (interface{}, error) {
	return wrapX11Drag(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}
