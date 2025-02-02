// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"context"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gcancel"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern void callbackDelete(gpointer);
// extern void _gotk4_gio2_VfsClass_local_file_removed(GVfs*, char*);
// extern void _gotk4_gio2_VfsClass_local_file_moved(GVfs*, char*, char*);
// extern void _gotk4_gio2_VfsClass_add_writable_namespaces(GVfs*, GFileAttributeInfoList*);
// extern gchar** _gotk4_gio2_VfsClass_get_supported_uri_schemes(GVfs*);
// extern gboolean _gotk4_gio2_VfsClass_local_file_set_attributes(GVfs*, char*, GFileInfo*, GFileQueryInfoFlags, GCancellable*, GError**);
// extern gboolean _gotk4_gio2_VfsClass_is_active(GVfs*);
// extern GFile* _gotk4_gio2_VfsClass_parse_name(GVfs*, char*);
// extern GFile* _gotk4_gio2_VfsClass_get_file_for_uri(GVfs*, char*);
// extern GFile* _gotk4_gio2_VfsClass_get_file_for_path(GVfs*, char*);
// extern GFile* _gotk4_gio2_VFSFileLookupFunc(GVfs*, char*, gpointer);
// GFile* _gotk4_gio2_VFS_virtual_get_file_for_path(void* fnptr, GVfs* arg0, char* arg1) {
//   return ((GFile* (*)(GVfs*, char*))(fnptr))(arg0, arg1);
// };
// GFile* _gotk4_gio2_VFS_virtual_get_file_for_uri(void* fnptr, GVfs* arg0, char* arg1) {
//   return ((GFile* (*)(GVfs*, char*))(fnptr))(arg0, arg1);
// };
// GFile* _gotk4_gio2_VFS_virtual_parse_name(void* fnptr, GVfs* arg0, char* arg1) {
//   return ((GFile* (*)(GVfs*, char*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gio2_VFS_virtual_is_active(void* fnptr, GVfs* arg0) {
//   return ((gboolean (*)(GVfs*))(fnptr))(arg0);
// };
// gboolean _gotk4_gio2_VFS_virtual_local_file_set_attributes(void* fnptr, GVfs* arg0, char* arg1, GFileInfo* arg2, GFileQueryInfoFlags arg3, GCancellable* arg4, GError** arg5) {
//   return ((gboolean (*)(GVfs*, char*, GFileInfo*, GFileQueryInfoFlags, GCancellable*, GError**))(fnptr))(arg0, arg1, arg2, arg3, arg4, arg5);
// };
// gchar** _gotk4_gio2_VFS_virtual_get_supported_uri_schemes(void* fnptr, GVfs* arg0) {
//   return ((gchar** (*)(GVfs*))(fnptr))(arg0);
// };
// void _gotk4_gio2_VFS_virtual_add_writable_namespaces(void* fnptr, GVfs* arg0, GFileAttributeInfoList* arg1) {
//   ((void (*)(GVfs*, GFileAttributeInfoList*))(fnptr))(arg0, arg1);
// };
// void _gotk4_gio2_VFS_virtual_local_file_moved(void* fnptr, GVfs* arg0, char* arg1, char* arg2) {
//   ((void (*)(GVfs*, char*, char*))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gio2_VFS_virtual_local_file_removed(void* fnptr, GVfs* arg0, char* arg1) {
//   ((void (*)(GVfs*, char*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeVFS = coreglib.Type(C.g_vfs_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVFS, F: marshalVFS},
	})
}

// VFS_EXTENSION_POINT_NAME: extension point for #GVfs functionality. See
// [Extending GIO][extending-gio].
const VFS_EXTENSION_POINT_NAME = "gio-vfs"

// VFSOverrides contains methods that are overridable.
type VFSOverrides struct {
	// The function takes the following parameters:
	//
	AddWritableNamespaces func(list *FileAttributeInfoList)
	// FileForPath gets a #GFile for path.
	//
	// The function takes the following parameters:
	//
	//    - path: string containing a VFS path.
	//
	// The function returns the following values:
	//
	//    - file: #GFile. Free the returned object with g_object_unref().
	//
	FileForPath func(path string) *File
	// FileForURI gets a #GFile for uri.
	//
	// This operation never fails, but the returned object might not support any
	// I/O operation if the URI is malformed or if the URI scheme is not
	// supported.
	//
	// The function takes the following parameters:
	//
	//    - uri: string containing a URI.
	//
	// The function returns the following values:
	//
	//    - file: #GFile. Free the returned object with g_object_unref().
	//
	FileForURI func(uri string) *File
	// SupportedURISchemes gets a list of URI schemes supported by vfs.
	//
	// The function returns the following values:
	//
	//    - utf8s: NULL-terminated array of strings. The returned array belongs
	//      to GIO and must not be freed or modified.
	//
	SupportedURISchemes func() []string
	// IsActive checks if the VFS is active.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if construction of the vfs was successful and it is now
	//      active.
	//
	IsActive func() bool
	// The function takes the following parameters:
	//
	//    - source
	//    - dest
	//
	LocalFileMoved func(source, dest string)
	// The function takes the following parameters:
	//
	LocalFileRemoved func(filename string)
	// The function takes the following parameters:
	//
	//    - ctx (optional)
	//    - filename
	//    - info
	//    - flags
	//
	LocalFileSetAttributes func(ctx context.Context, filename string, info *FileInfo, flags FileQueryInfoFlags) error
	// ParseName: this operation never fails, but the returned object might not
	// support any I/O operations if the parse_name cannot be parsed by the
	// #GVfs module.
	//
	// The function takes the following parameters:
	//
	//    - parseName: string to be parsed by the VFS module.
	//
	// The function returns the following values:
	//
	//    - file for the given parse_name. Free the returned object with
	//      g_object_unref().
	//
	ParseName func(parseName string) *File
}

func defaultVFSOverrides(v *VFS) VFSOverrides {
	return VFSOverrides{
		AddWritableNamespaces:  v.addWritableNamespaces,
		FileForPath:            v.fileForPath,
		FileForURI:             v.fileForURI,
		SupportedURISchemes:    v.supportedURISchemes,
		IsActive:               v.isActive,
		LocalFileMoved:         v.localFileMoved,
		LocalFileRemoved:       v.localFileRemoved,
		LocalFileSetAttributes: v.localFileSetAttributes,
		ParseName:              v.parseName,
	}
}

// VFS: entry point for using GIO functionality.
type VFS struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*VFS)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*VFS, *VFSClass, VFSOverrides](
		GTypeVFS,
		initVFSClass,
		wrapVFS,
		defaultVFSOverrides,
	)
}

func initVFSClass(gclass unsafe.Pointer, overrides VFSOverrides, classInitFunc func(*VFSClass)) {
	pclass := (*C.GVfsClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeVFS))))

	if overrides.AddWritableNamespaces != nil {
		pclass.add_writable_namespaces = (*[0]byte)(C._gotk4_gio2_VfsClass_add_writable_namespaces)
	}

	if overrides.FileForPath != nil {
		pclass.get_file_for_path = (*[0]byte)(C._gotk4_gio2_VfsClass_get_file_for_path)
	}

	if overrides.FileForURI != nil {
		pclass.get_file_for_uri = (*[0]byte)(C._gotk4_gio2_VfsClass_get_file_for_uri)
	}

	if overrides.SupportedURISchemes != nil {
		pclass.get_supported_uri_schemes = (*[0]byte)(C._gotk4_gio2_VfsClass_get_supported_uri_schemes)
	}

	if overrides.IsActive != nil {
		pclass.is_active = (*[0]byte)(C._gotk4_gio2_VfsClass_is_active)
	}

	if overrides.LocalFileMoved != nil {
		pclass.local_file_moved = (*[0]byte)(C._gotk4_gio2_VfsClass_local_file_moved)
	}

	if overrides.LocalFileRemoved != nil {
		pclass.local_file_removed = (*[0]byte)(C._gotk4_gio2_VfsClass_local_file_removed)
	}

	if overrides.LocalFileSetAttributes != nil {
		pclass.local_file_set_attributes = (*[0]byte)(C._gotk4_gio2_VfsClass_local_file_set_attributes)
	}

	if overrides.ParseName != nil {
		pclass.parse_name = (*[0]byte)(C._gotk4_gio2_VfsClass_parse_name)
	}

	if classInitFunc != nil {
		class := (*VFSClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapVFS(obj *coreglib.Object) *VFS {
	return &VFS{
		Object: obj,
	}
}

func marshalVFS(p uintptr) (interface{}, error) {
	return wrapVFS(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// FileForPath gets a #GFile for path.
//
// The function takes the following parameters:
//
//    - path: string containing a VFS path.
//
// The function returns the following values:
//
//    - file: #GFile. Free the returned object with g_object_unref().
//
func (vfs *VFS) FileForPath(path string) *File {
	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_get_file_for_path(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(path)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// FileForURI gets a #GFile for uri.
//
// This operation never fails, but the returned object might not support any I/O
// operation if the URI is malformed or if the URI scheme is not supported.
//
// The function takes the following parameters:
//
//    - uri: string containing a URI.
//
// The function returns the following values:
//
//    - file: #GFile. Free the returned object with g_object_unref().
//
func (vfs *VFS) FileForURI(uri string) *File {
	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_get_file_for_uri(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(uri)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// SupportedURISchemes gets a list of URI schemes supported by vfs.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of strings. The returned array belongs to
//      GIO and must not be freed or modified.
//
func (vfs *VFS) SupportedURISchemes() []string {
	var _arg0 *C.GVfs   // out
	var _cret **C.gchar // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))

	_cret = C.g_vfs_get_supported_uri_schemes(_arg0)
	runtime.KeepAlive(vfs)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// IsActive checks if the VFS is active.
//
// The function returns the following values:
//
//    - ok: TRUE if construction of the vfs was successful and it is now active.
//
func (vfs *VFS) IsActive() bool {
	var _arg0 *C.GVfs    // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))

	_cret = C.g_vfs_is_active(_arg0)
	runtime.KeepAlive(vfs)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ParseName: this operation never fails, but the returned object might not
// support any I/O operations if the parse_name cannot be parsed by the #GVfs
// module.
//
// The function takes the following parameters:
//
//    - parseName: string to be parsed by the VFS module.
//
// The function returns the following values:
//
//    - file for the given parse_name. Free the returned object with
//      g_object_unref().
//
func (vfs *VFS) ParseName(parseName string) *File {
	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(parseName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_parse_name(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(parseName)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// RegisterURIScheme registers uri_func and parse_name_func as the #GFile URI
// and parse name lookup functions for URIs with a scheme matching scheme. Note
// that scheme is registered only within the running application, as opposed to
// desktop-wide as it happens with GVfs backends.
//
// When a #GFile is requested with an URI containing scheme (e.g. through
// g_file_new_for_uri()), uri_func will be called to allow a custom constructor.
// The implementation of uri_func should not be blocking, and must not call
// g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
//
// When g_file_parse_name() is called with a parse name obtained from such file,
// parse_name_func will be called to allow the #GFile to be created again. In
// that case, it's responsibility of parse_name_func to make sure the parse name
// matches what the custom #GFile implementation returned when
// g_file_get_parse_name() was previously called. The implementation of
// parse_name_func should not be blocking, and must not call
// g_vfs_register_uri_scheme() or g_vfs_unregister_uri_scheme().
//
// It's an error to call this function twice with the same scheme. To unregister
// a custom URI scheme, use g_vfs_unregister_uri_scheme().
//
// The function takes the following parameters:
//
//    - scheme: URI scheme, e.g. "http".
//    - uriFunc (optional): FileLookupFunc.
//    - parseNameFunc (optional): FileLookupFunc.
//
// The function returns the following values:
//
//    - ok: TRUE if scheme was successfully registered, or FALSE if a handler for
//      scheme already exists.
//
func (vfs *VFS) RegisterURIScheme(scheme string, uriFunc, parseNameFunc VFSFileLookupFunc) bool {
	var _arg0 *C.GVfs              // out
	var _arg1 *C.char              // out
	var _arg2 C.GVfsFileLookupFunc // out
	var _arg3 C.gpointer
	var _arg4 C.GDestroyNotify
	var _arg5 C.GVfsFileLookupFunc // out
	var _arg6 C.gpointer
	var _arg7 C.GDestroyNotify
	var _cret C.gboolean // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(_arg1))
	if uriFunc != nil {
		_arg2 = (*[0]byte)(C._gotk4_gio2_VFSFileLookupFunc)
		_arg3 = C.gpointer(gbox.Assign(uriFunc))
		_arg4 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}
	if parseNameFunc != nil {
		_arg5 = (*[0]byte)(C._gotk4_gio2_VFSFileLookupFunc)
		_arg6 = C.gpointer(gbox.Assign(parseNameFunc))
		_arg7 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.g_vfs_register_uri_scheme(_arg0, _arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(scheme)
	runtime.KeepAlive(uriFunc)
	runtime.KeepAlive(parseNameFunc)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// UnregisterURIScheme unregisters the URI handler for scheme previously
// registered with g_vfs_register_uri_scheme().
//
// The function takes the following parameters:
//
//    - scheme: URI scheme, e.g. "http".
//
// The function returns the following values:
//
//    - ok: TRUE if scheme was successfully unregistered, or FALSE if a handler
//      for scheme does not exist.
//
func (vfs *VFS) UnregisterURIScheme(scheme string) bool {
	var _arg0 *C.GVfs    // out
	var _arg1 *C.char    // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(scheme)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_vfs_unregister_uri_scheme(_arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(scheme)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
func (vfs *VFS) addWritableNamespaces(list *FileAttributeInfoList) {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.add_writable_namespaces

	var _arg0 *C.GVfs                   // out
	var _arg1 *C.GFileAttributeInfoList // out

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.GFileAttributeInfoList)(gextras.StructNative(unsafe.Pointer(list)))

	C._gotk4_gio2_VFS_virtual_add_writable_namespaces(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(list)
}

// fileForPath gets a #GFile for path.
//
// The function takes the following parameters:
//
//    - path: string containing a VFS path.
//
// The function returns the following values:
//
//    - file: #GFile. Free the returned object with g_object_unref().
//
func (vfs *VFS) fileForPath(path string) *File {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.get_file_for_path

	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(path)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_VFS_virtual_get_file_for_path(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(path)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// fileForURI gets a #GFile for uri.
//
// This operation never fails, but the returned object might not support any I/O
// operation if the URI is malformed or if the URI scheme is not supported.
//
// The function takes the following parameters:
//
//    - uri: string containing a URI.
//
// The function returns the following values:
//
//    - file: #GFile. Free the returned object with g_object_unref().
//
func (vfs *VFS) fileForURI(uri string) *File {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.get_file_for_uri

	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_VFS_virtual_get_file_for_uri(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(uri)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// supportedURISchemes gets a list of URI schemes supported by vfs.
//
// The function returns the following values:
//
//    - utf8s: NULL-terminated array of strings. The returned array belongs to
//      GIO and must not be freed or modified.
//
func (vfs *VFS) supportedURISchemes() []string {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.get_supported_uri_schemes

	var _arg0 *C.GVfs   // out
	var _cret **C.gchar // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))

	_cret = C._gotk4_gio2_VFS_virtual_get_supported_uri_schemes(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(vfs)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}

// isActive checks if the VFS is active.
//
// The function returns the following values:
//
//    - ok: TRUE if construction of the vfs was successful and it is now active.
//
func (vfs *VFS) isActive() bool {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.is_active

	var _arg0 *C.GVfs    // out
	var _cret C.gboolean // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))

	_cret = C._gotk4_gio2_VFS_virtual_is_active(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(vfs)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - source
//    - dest
//
func (vfs *VFS) localFileMoved(source, dest string) {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.local_file_moved

	var _arg0 *C.GVfs // out
	var _arg1 *C.char // out
	var _arg2 *C.char // out

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(source)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(dest)))
	defer C.free(unsafe.Pointer(_arg2))

	C._gotk4_gio2_VFS_virtual_local_file_moved(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(source)
	runtime.KeepAlive(dest)
}

// The function takes the following parameters:
//
func (vfs *VFS) localFileRemoved(filename string) {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.local_file_removed

	var _arg0 *C.GVfs // out
	var _arg1 *C.char // out

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))

	C._gotk4_gio2_VFS_virtual_local_file_removed(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(filename)
}

// The function takes the following parameters:
//
//    - ctx (optional)
//    - filename
//    - info
//    - flags
//
func (vfs *VFS) localFileSetAttributes(ctx context.Context, filename string, info *FileInfo, flags FileQueryInfoFlags) error {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.local_file_set_attributes

	var _arg0 *C.GVfs               // out
	var _arg4 *C.GCancellable       // out
	var _arg1 *C.char               // out
	var _arg2 *C.GFileInfo          // out
	var _arg3 C.GFileQueryInfoFlags // out
	var _cerr *C.GError             // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	{
		cancellable := gcancel.GCancellableFromContext(ctx)
		defer runtime.KeepAlive(cancellable)
		_arg4 = (*C.GCancellable)(unsafe.Pointer(cancellable.Native()))
	}
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(filename)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GFileInfo)(unsafe.Pointer(coreglib.InternObject(info).Native()))
	_arg3 = C.GFileQueryInfoFlags(flags)

	C._gotk4_gio2_VFS_virtual_local_file_set_attributes(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4, &_cerr)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(info)
	runtime.KeepAlive(flags)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// parseName: this operation never fails, but the returned object might not
// support any I/O operations if the parse_name cannot be parsed by the #GVfs
// module.
//
// The function takes the following parameters:
//
//    - parseName: string to be parsed by the VFS module.
//
// The function returns the following values:
//
//    - file for the given parse_name. Free the returned object with
//      g_object_unref().
//
func (vfs *VFS) parseName(parseName string) *File {
	gclass := (*C.GVfsClass)(coreglib.PeekParentClass(vfs))
	fnarg := gclass.parse_name

	var _arg0 *C.GVfs  // out
	var _arg1 *C.char  // out
	var _cret *C.GFile // in

	_arg0 = (*C.GVfs)(unsafe.Pointer(coreglib.InternObject(vfs).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(parseName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gio2_VFS_virtual_parse_name(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(vfs)
	runtime.KeepAlive(parseName)

	var _file *File // out

	_file = wrapFile(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _file
}

// VFSGetDefault gets the default #GVfs for the system.
//
// The function returns the following values:
//
//    - vfs which will be the local file system #GVfs if no other implementation
//      is available.
//
func VFSGetDefault() *VFS {
	var _cret *C.GVfs // in

	_cret = C.g_vfs_get_default()

	var _vfs *VFS // out

	_vfs = wrapVFS(coreglib.Take(unsafe.Pointer(_cret)))

	return _vfs
}

// VFSGetLocal gets the local #GVfs for the system.
//
// The function returns the following values:
//
//    - vfs: #GVfs.
//
func VFSGetLocal() *VFS {
	var _cret *C.GVfs // in

	_cret = C.g_vfs_get_local()

	var _vfs *VFS // out

	_vfs = wrapVFS(coreglib.Take(unsafe.Pointer(_cret)))

	return _vfs
}

// VFSClass: instance of this type is always passed by reference.
type VFSClass struct {
	*vfsClass
}

// vfsClass is the struct that's finalized.
type vfsClass struct {
	native *C.GVfsClass
}
