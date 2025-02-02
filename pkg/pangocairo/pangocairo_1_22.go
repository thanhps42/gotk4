// Code generated by girgen. DO NOT EDIT.

package pangocairo

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/cairo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/pango"
)

// #include <stdlib.h>
// #include <pango/pangocairo.h>
import "C"

// CreateContext creates a context object set up to match the current
// transformation and target surface of the Cairo context.
//
// This context can then be used to create a layout using pango.Layout.New.
//
// This function is a convenience function that creates a context using the
// default font map, then updates it to cr. If you just need to create a layout
// for use with cr and do not need to access PangoContext directly, you can use
// create_layout instead.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//
// The function returns the following values:
//
//    - context: newly created PangoContext. Free with g_object_unref().
//
func CreateContext(cr *cairo.Context) *pango.Context {
	var _arg1 *C.cairo_t      // out
	var _cret *C.PangoContext // in

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))

	_cret = C.pango_cairo_create_context(_arg1)
	runtime.KeepAlive(cr)

	var _context *pango.Context // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_context = &pango.Context{
			Object: obj,
		}
	}

	return _context
}

// ShowGlyphItem draws the glyphs in glyph_item in the specified cairo context,
//
// embedding the text associated with the glyphs in the output if the output
// format supports it (PDF for example), otherwise it acts similar to
// show_glyph_string.
//
// The origin of the glyphs (the left edge of the baseline) will be drawn at the
// current point of the cairo context.
//
// Note that text is the start of the text for layout, which is then indexed by
// glyph_item->item->offset.
//
// The function takes the following parameters:
//
//    - cr: cairo context.
//    - text: UTF-8 text that glyph_item refers to.
//    - glyphItem: PangoGlyphItem.
//
func ShowGlyphItem(cr *cairo.Context, text string, glyphItem *pango.GlyphItem) {
	var _arg1 *C.cairo_t        // out
	var _arg2 *C.char           // out
	var _arg3 *C.PangoGlyphItem // out

	_arg1 = (*C.cairo_t)(unsafe.Pointer(cr.Native()))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(text)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.PangoGlyphItem)(gextras.StructNative(unsafe.Pointer(glyphItem)))

	C.pango_cairo_show_glyph_item(_arg1, _arg2, _arg3)
	runtime.KeepAlive(cr)
	runtime.KeepAlive(text)
	runtime.KeepAlive(glyphItem)
}
