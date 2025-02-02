// Code generated by girgen. DO NOT EDIT.

package atk

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <atk/atk.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeScrollType = coreglib.Type(C.atk_scroll_type_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeScrollType, F: marshalScrollType},
	})
}

// ScrollType specifies where an object should be placed on the screen when
// using scroll_to.
type ScrollType C.gint

const (
	// ScrollTopLeft: scroll the object vertically and horizontally to bring its
	// top left corner to the top left corner of the window.
	ScrollTopLeft ScrollType = iota
	// ScrollBottomRight: scroll the object vertically and horizontally to bring
	// its bottom right corner to the bottom right corner of the window.
	ScrollBottomRight
	// ScrollTopEdge: scroll the object vertically to bring its top edge to the
	// top edge of the window.
	ScrollTopEdge
	// ScrollBottomEdge: scroll the object vertically to bring its bottom edge
	// to the bottom edge of the window.
	ScrollBottomEdge
	// ScrollLeftEdge: scroll the object vertically and horizontally to bring
	// its left edge to the left edge of the window.
	ScrollLeftEdge
	// ScrollRightEdge: scroll the object vertically and horizontally to bring
	// its right edge to the right edge of the window.
	ScrollRightEdge
	// ScrollAnywhere: scroll the object vertically and horizontally so that as
	// much as possible of the object becomes visible. The exact placement is
	// determined by the application.
	ScrollAnywhere
)

func marshalScrollType(p uintptr) (interface{}, error) {
	return ScrollType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ScrollType.
func (s ScrollType) String() string {
	switch s {
	case ScrollTopLeft:
		return "TopLeft"
	case ScrollBottomRight:
		return "BottomRight"
	case ScrollTopEdge:
		return "TopEdge"
	case ScrollBottomEdge:
		return "BottomEdge"
	case ScrollLeftEdge:
		return "LeftEdge"
	case ScrollRightEdge:
		return "RightEdge"
	case ScrollAnywhere:
		return "Anywhere"
	default:
		return fmt.Sprintf("ScrollType(%d)", s)
	}
}
