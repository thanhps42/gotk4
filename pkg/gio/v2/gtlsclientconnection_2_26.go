// Code generated by girgen. DO NOT EDIT.

package gio

// #include <stdlib.h>
// #include <gio/gio.h>
import "C"

// TLSClientConnectionInterface: vtable for a ClientConnection implementation.
//
// An instance of this type is always passed by reference.
type TLSClientConnectionInterface struct {
	*tlsClientConnectionInterface
}

// tlsClientConnectionInterface is the struct that's finalized.
type tlsClientConnectionInterface struct {
	native *C.GTlsClientConnectionInterface
}
