// Code generated by girgen. DO NOT EDIT.

package gdk

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gdk/gdk.h>
// #include <glib-object.h>
// extern void _gotk4_gdk3_DeviceManager_ConnectDeviceRemoved(gpointer, GdkDevice*, guintptr);
// extern void _gotk4_gdk3_DeviceManager_ConnectDeviceChanged(gpointer, GdkDevice*, guintptr);
// extern void _gotk4_gdk3_DeviceManager_ConnectDeviceAdded(gpointer, GdkDevice*, guintptr);
import "C"

// GType values.
var (
	GTypeDeviceManager = coreglib.Type(C.gdk_device_manager_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDeviceManager, F: marshalDeviceManager},
	})
}

// DeviceManager: in addition to a single pointer and keyboard for user
// interface input, GDK contains support for a variety of input devices,
// including graphics tablets, touchscreens and multiple pointers/keyboards
// interacting simultaneously with the user interface. Such input devices often
// have additional features, such as sub-pixel positioning information and
// additional device-dependent information.
//
// In order to query the device hierarchy and be aware of changes in the device
// hierarchy (such as virtual devices being created or removed, or physical
// devices being plugged or unplugged), GDK provides DeviceManager.
//
// By default, and if the platform supports it, GDK is aware of multiple
// keyboard/pointer pairs and multitouch devices. This behavior can be changed
// by calling gdk_disable_multidevice() before gdk_display_open(). There should
// rarely be a need to do that though, since GDK defaults to a compatibility
// mode in which it will emit just one enter/leave event pair for all devices on
// a window. To enable per-device enter/leave events and other multi-pointer
// interaction features, gdk_window_set_support_multidevice() must be called on
// Windows (or gtk_widget_set_support_multidevice() on widgets). window. See the
// gdk_window_set_support_multidevice() documentation for more information.
//
// On X11, multi-device support is implemented through XInput 2. Unless
// gdk_disable_multidevice() is called, the XInput 2 DeviceManager
// implementation will be used as the input source. Otherwise either the core or
// XInput 1 implementations will be used.
//
// For simple applications that don’t have any special interest in input
// devices, the so-called “client pointer” provides a reasonable approximation
// to a simple setup with a single pointer and keyboard. The device that has
// been set as the client pointer can be accessed via
// gdk_device_manager_get_client_pointer().
//
// Conceptually, in multidevice mode there are 2 device types. Virtual devices
// (or master devices) are represented by the pointer cursors and keyboard foci
// that are seen on the screen. Physical devices (or slave devices) represent
// the hardware that is controlling the virtual devices, and thus have no
// visible cursor on the screen.
//
// Virtual devices are always paired, so there is a keyboard device for every
// pointer device. Associations between devices may be inspected through
// gdk_device_get_associated_device().
//
// There may be several virtual devices, and several physical devices could be
// controlling each of these virtual devices. Physical devices may also be
// “floating”, which means they are not attached to any virtual device.
//
// Master and slave devices
//
//    carlossacarino:~$ xinput list
//    ⎡ Virtual core pointer                          id=2    [master pointer  (3)]
//    ⎜   ↳ Virtual core XTEST pointer                id=4    [slave  pointer  (2)]
//    ⎜   ↳ Wacom ISDv4 E6 Pen stylus                 id=10   [slave  pointer  (2)]
//    ⎜   ↳ Wacom ISDv4 E6 Finger touch               id=11   [slave  pointer  (2)]
//    ⎜   ↳ SynPS/2 Synaptics TouchPad                id=13   [slave  pointer  (2)]
//    ⎜   ↳ TPPS/2 IBM TrackPoint                     id=14   [slave  pointer  (2)]
//    ⎜   ↳ Wacom ISDv4 E6 Pen eraser                 id=16   [slave  pointer  (2)]
//    ⎣ Virtual core keyboard                         id=3    [master keyboard (2)]
//        ↳ Virtual core XTEST keyboard               id=5    [slave  keyboard (3)]
//        ↳ Power Button                              id=6    [slave  keyboard (3)]
//        ↳ Video Bus                                 id=7    [slave  keyboard (3)]
//        ↳ Sleep Button                              id=8    [slave  keyboard (3)]
//        ↳ Integrated Camera                         id=9    [slave  keyboard (3)]
//        ↳ AT Translated Set 2 keyboard              id=12   [slave  keyboard (3)]
//        ↳ ThinkPad Extra Buttons                    id=15   [slave  keyboard (3)]
//
// By default, GDK will automatically listen for events coming from all master
// devices, setting the Device for all events coming from input devices. Events
// containing device information are K_MOTION_NOTIFY, K_BUTTON_PRESS,
// K_2BUTTON_PRESS, K_3BUTTON_PRESS, K_BUTTON_RELEASE, K_SCROLL, K_KEY_PRESS,
// K_KEY_RELEASE, K_ENTER_NOTIFY, K_LEAVE_NOTIFY, K_FOCUS_CHANGE,
// K_PROXIMITY_IN, K_PROXIMITY_OUT, K_DRAG_ENTER, K_DRAG_LEAVE, K_DRAG_MOTION,
// K_DRAG_STATUS, K_DROP_START, K_DROP_FINISHED and K_GRAB_BROKEN. When dealing
// with an event on a master device, it is possible to get the source (slave)
// device that the event originated from via gdk_event_get_source_device().
//
// On a standard session, all physical devices are connected by default to the
// "Virtual Core Pointer/Keyboard" master devices, hence routing all events
// through these. This behavior is only modified by device grabs, where the
// slave device is temporarily detached for as long as the grab is held, and
// more permanently by user modifications to the device hierarchy.
//
// On certain application specific setups, it may make sense to detach a
// physical device from its master pointer, and mapping it to an specific
// window. This can be achieved by the combination of gdk_device_grab() and
// gdk_device_set_mode().
//
// In order to listen for events coming from devices other than a virtual
// device, gdk_window_set_device_events() must be called. Generally, this
// function can be used to modify the event mask for any given device.
//
// Input devices may also provide additional information besides X/Y. For
// example, graphics tablets may also provide pressure and X/Y tilt information.
// This information is device-dependent, and may be queried through
// gdk_device_get_axis(). In multidevice mode, virtual devices will change axes
// in order to always represent the physical device that is routing events
// through it. Whenever the physical device changes, the Device:n-axes property
// will be notified, and gdk_device_list_axes() will return the new device axes.
//
// Devices may also have associated “keys” or macro buttons. Such keys can be
// globally set to map into normal X keyboard events. The mapping is set using
// gdk_device_set_key().
//
// In GTK+ 3.20, a new Seat object has been introduced that supersedes
// DeviceManager and should be preferred in newly written code.
type DeviceManager struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DeviceManager)(nil)
)

// DeviceManagerer describes types inherited from class DeviceManager.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type DeviceManagerer interface {
	coreglib.Objector
	baseDeviceManager() *DeviceManager
}

var _ DeviceManagerer = (*DeviceManager)(nil)

func wrapDeviceManager(obj *coreglib.Object) *DeviceManager {
	return &DeviceManager{
		Object: obj,
	}
}

func marshalDeviceManager(p uintptr) (interface{}, error) {
	return wrapDeviceManager(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (deviceManager *DeviceManager) baseDeviceManager() *DeviceManager {
	return deviceManager
}

// BaseDeviceManager returns the underlying base object.
func BaseDeviceManager(obj DeviceManagerer) *DeviceManager {
	return obj.baseDeviceManager()
}

// ConnectDeviceAdded signal is emitted either when a new master pointer is
// created, or when a slave (Hardware) input device is plugged in.
func (deviceManager *DeviceManager) ConnectDeviceAdded(f func(device Devicer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(deviceManager, "device-added", false, unsafe.Pointer(C._gotk4_gdk3_DeviceManager_ConnectDeviceAdded), f)
}

// ConnectDeviceChanged signal is emitted whenever a device has changed in the
// hierarchy, either slave devices being disconnected from their master device
// or connected to another one, or master devices being added or removed a slave
// device.
//
// If a slave device is detached from all master devices
// (gdk_device_get_associated_device() returns NULL), its DeviceType will change
// to GDK_DEVICE_TYPE_FLOATING, if it's attached, it will change to
// GDK_DEVICE_TYPE_SLAVE.
func (deviceManager *DeviceManager) ConnectDeviceChanged(f func(device Devicer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(deviceManager, "device-changed", false, unsafe.Pointer(C._gotk4_gdk3_DeviceManager_ConnectDeviceChanged), f)
}

// ConnectDeviceRemoved signal is emitted either when a master pointer is
// removed, or when a slave (Hardware) input device is unplugged.
func (deviceManager *DeviceManager) ConnectDeviceRemoved(f func(device Devicer)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(deviceManager, "device-removed", false, unsafe.Pointer(C._gotk4_gdk3_DeviceManager_ConnectDeviceRemoved), f)
}

// ClientPointer returns the client pointer, that is, the master pointer that
// acts as the core pointer for this application. In X11, window managers may
// change this depending on the interaction pattern under the presence of
// several pointers.
//
// You should use this function seldomly, only in code that isn’t triggered by a
// Event and there aren’t other means to get a meaningful Device to operate on.
//
// Deprecated: Use gdk_seat_get_pointer() instead.
//
// The function returns the following values:
//
//    - device: client pointer. This memory is owned by GDK and must not be freed
//      or unreferenced.
//
func (deviceManager *DeviceManager) ClientPointer() Devicer {
	var _arg0 *C.GdkDeviceManager // out
	var _cret *C.GdkDevice        // in

	_arg0 = (*C.GdkDeviceManager)(unsafe.Pointer(coreglib.InternObject(deviceManager).Native()))

	_cret = C.gdk_device_manager_get_client_pointer(_arg0)
	runtime.KeepAlive(deviceManager)

	var _device Devicer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gdk.Devicer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Devicer)
			return ok
		})
		rv, ok := casted.(Devicer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
		}
		_device = rv
	}

	return _device
}

// Display gets the Display associated to device_manager.
//
// The function returns the following values:
//
//    - display (optional) to which device_manager is associated to, or NULL.
//      This memory is owned by GDK and must not be freed or unreferenced.
//
func (deviceManager *DeviceManager) Display() *Display {
	var _arg0 *C.GdkDeviceManager // out
	var _cret *C.GdkDisplay       // in

	_arg0 = (*C.GdkDeviceManager)(unsafe.Pointer(coreglib.InternObject(deviceManager).Native()))

	_cret = C.gdk_device_manager_get_display(_arg0)
	runtime.KeepAlive(deviceManager)

	var _display *Display // out

	if _cret != nil {
		_display = wrapDisplay(coreglib.Take(unsafe.Pointer(_cret)))
	}

	return _display
}

// ListDevices returns the list of devices of type type currently attached to
// device_manager.
//
// Deprecated: , use gdk_seat_get_pointer(), gdk_seat_get_keyboard() and
// gdk_seat_get_slaves() instead.
//
// The function takes the following parameters:
//
//    - typ: device type to get.
//
// The function returns the following values:
//
//    - list of Devices. The returned list must be freed with g_list_free (). The
//      list elements are owned by GTK+ and must not be freed or unreffed.
//
func (deviceManager *DeviceManager) ListDevices(typ DeviceType) []Devicer {
	var _arg0 *C.GdkDeviceManager // out
	var _arg1 C.GdkDeviceType     // out
	var _cret *C.GList            // in

	_arg0 = (*C.GdkDeviceManager)(unsafe.Pointer(coreglib.InternObject(deviceManager).Native()))
	_arg1 = C.GdkDeviceType(typ)

	_cret = C.gdk_device_manager_list_devices(_arg0, _arg1)
	runtime.KeepAlive(deviceManager)
	runtime.KeepAlive(typ)

	var _list []Devicer // out

	_list = make([]Devicer, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GdkDevice)(v)
		var dst Devicer // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gdk.Devicer is nil")
			}

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Devicer)
				return ok
			})
			rv, ok := casted.(Devicer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gdk.Devicer")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})

	return _list
}
