// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
import "C"

// DragFindWindowForScreen finds the destination window and DND protocol to use
// at the given pointer position.
//
// This function is called by the drag source to obtain the dest_window and
// protocol parameters for gdk_drag_motion().
//
// The function takes the following parameters:
//
//    - context: DragContext.
//    - dragWindow: window which may be at the pointer position, but should be
//      ignored, since it is put up by the drag source as an icon.
//    - screen where the destination window is sought.
//    - xRoot: x position of the pointer in root coordinates.
//    - yRoot: y position of the pointer in root coordinates.
//
// The function returns the following values:
//
//    - destWindow: location to store the destination window in.
//    - protocol: location to store the DND protocol in.
//
func DragFindWindowForScreen(context *DragContext, dragWindow Windower, screen *Screen, xRoot, yRoot int) (Windower, DragProtocol) {
	var _arg1 *C.GdkDragContext // out
	var _arg2 *C.GdkWindow      // out
	var _arg3 *C.GdkScreen      // out
	var _arg4 C.gint            // out
	var _arg5 C.gint            // out
	var _arg6 *C.GdkWindow      // in
	var _arg7 C.GdkDragProtocol // in

	_arg1 = (*C.GdkDragContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GdkWindow)(unsafe.Pointer(coreglib.InternObject(dragWindow).Native()))
	_arg3 = (*C.GdkScreen)(unsafe.Pointer(coreglib.InternObject(screen).Native()))
	_arg4 = C.gint(xRoot)
	_arg5 = C.gint(yRoot)

	C.gdk_drag_find_window_for_screen(_arg1, _arg2, _arg3, _arg4, _arg5, &_arg6, &_arg7)
	runtime.KeepAlive(context)
	runtime.KeepAlive(dragWindow)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(xRoot)
	runtime.KeepAlive(yRoot)

	var _destWindow Windower   // out
	var _protocol DragProtocol // out

	{
		objptr := unsafe.Pointer(_arg6)
		if objptr == nil {
			panic("object of type gdk.Windower is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Windower)
			return ok
		})
		rv, ok := casted.(Windower)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Windower")
		}
		_destWindow = rv
	}
	_protocol = DragProtocol(_arg7)

	return _destWindow, _protocol
}
