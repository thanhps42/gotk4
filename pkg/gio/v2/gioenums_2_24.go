// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"fmt"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
import "C"

// GType values.
var (
	GTypeConverterResult      = coreglib.Type(C.g_converter_result_get_type())
	GTypeZlibCompressorFormat = coreglib.Type(C.g_zlib_compressor_format_get_type())
	GTypeConverterFlags       = coreglib.Type(C.g_converter_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeConverterResult, F: marshalConverterResult},
		coreglib.TypeMarshaler{T: GTypeZlibCompressorFormat, F: marshalZlibCompressorFormat},
		coreglib.TypeMarshaler{T: GTypeConverterFlags, F: marshalConverterFlags},
	})
}

// ConverterResult results returned from g_converter_convert().
type ConverterResult C.gint

const (
	// ConverterError: there was an error during conversion.
	ConverterError ConverterResult = iota
	// ConverterConverted: some data was consumed or produced.
	ConverterConverted
	// ConverterFinished: conversion is finished.
	ConverterFinished
	// ConverterFlushed: flushing is finished.
	ConverterFlushed
)

func marshalConverterResult(p uintptr) (interface{}, error) {
	return ConverterResult(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ConverterResult.
func (c ConverterResult) String() string {
	switch c {
	case ConverterError:
		return "Error"
	case ConverterConverted:
		return "Converted"
	case ConverterFinished:
		return "Finished"
	case ConverterFlushed:
		return "Flushed"
	default:
		return fmt.Sprintf("ConverterResult(%d)", c)
	}
}

// ZlibCompressorFormat: used to select the type of data format to use for
// Decompressor and Compressor.
type ZlibCompressorFormat C.gint

const (
	// ZlibCompressorFormatZlib: deflate compression with zlib header.
	ZlibCompressorFormatZlib ZlibCompressorFormat = iota
	// ZlibCompressorFormatGzip: gzip file format.
	ZlibCompressorFormatGzip
	// ZlibCompressorFormatRaw: deflate compression with no header.
	ZlibCompressorFormatRaw
)

func marshalZlibCompressorFormat(p uintptr) (interface{}, error) {
	return ZlibCompressorFormat(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for ZlibCompressorFormat.
func (z ZlibCompressorFormat) String() string {
	switch z {
	case ZlibCompressorFormatZlib:
		return "Zlib"
	case ZlibCompressorFormatGzip:
		return "Gzip"
	case ZlibCompressorFormatRaw:
		return "Raw"
	default:
		return fmt.Sprintf("ZlibCompressorFormat(%d)", z)
	}
}

// ConverterFlags flags used when calling a g_converter_convert().
type ConverterFlags C.guint

const (
	// ConverterNoFlags: no flags.
	ConverterNoFlags ConverterFlags = 0b0
	// ConverterInputAtEnd: at end of input data.
	ConverterInputAtEnd ConverterFlags = 0b1
	// ConverterFlush: flush data.
	ConverterFlush ConverterFlags = 0b10
)

func marshalConverterFlags(p uintptr) (interface{}, error) {
	return ConverterFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for ConverterFlags.
func (c ConverterFlags) String() string {
	if c == 0 {
		return "ConverterFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(51)

	for c != 0 {
		next := c & (c - 1)
		bit := c - next

		switch bit {
		case ConverterNoFlags:
			builder.WriteString("None|")
		case ConverterInputAtEnd:
			builder.WriteString("InputAtEnd|")
		case ConverterFlush:
			builder.WriteString("Flush|")
		default:
			builder.WriteString(fmt.Sprintf("ConverterFlags(0b%b)|", bit))
		}

		c = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if c contains other.
func (c ConverterFlags) Has(other ConverterFlags) bool {
	return (c & other) == other
}
