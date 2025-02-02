// Code generated by girgen. DO NOT EDIT.

package gio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gio/gio.h>
// #include <glib-object.h>
// extern GTlsCertificateFlags _gotk4_gio2_TlsCertificateClass_verify(GTlsCertificate*, GSocketConnectable*, GTlsCertificate*);
// GTlsCertificateFlags _gotk4_gio2_TLSCertificate_virtual_verify(void* fnptr, GTlsCertificate* arg0, GSocketConnectable* arg1, GTlsCertificate* arg2) {
//   return ((GTlsCertificateFlags (*)(GTlsCertificate*, GSocketConnectable*, GTlsCertificate*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeTLSCertificate = coreglib.Type(C.g_tls_certificate_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTLSCertificate, F: marshalTLSCertificate},
	})
}

// TLSCertificateOverrides contains methods that are overridable.
type TLSCertificateOverrides struct {
	// Verify: this verifies cert and returns a set of CertificateFlags
	// indicating any problems found with it. This can be used to verify a
	// certificate outside the context of making a connection, or to check a
	// certificate against a CA that is not part of the system CA database.
	//
	// If identity is not NULL, cert's name(s) will be compared against it, and
	// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value if it does
	// not match. If identity is NULL, that bit will never be set in the return
	// value.
	//
	// If trusted_ca is not NULL, then cert (or one of the certificates in its
	// chain) must be signed by it, or else G_TLS_CERTIFICATE_UNKNOWN_CA will be
	// set in the return value. If trusted_ca is NULL, that bit will never be
	// set in the return value.
	//
	// (All other CertificateFlags values will always be set or unset as
	// appropriate.).
	//
	// The function takes the following parameters:
	//
	//    - identity (optional): expected peer identity.
	//    - trustedCa (optional): certificate of a trusted authority.
	//
	// The function returns the following values:
	//
	//    - tlsCertificateFlags: appropriate CertificateFlags.
	//
	Verify func(identity SocketConnectabler, trustedCa TLSCertificater) TLSCertificateFlags
}

func defaultTLSCertificateOverrides(v *TLSCertificate) TLSCertificateOverrides {
	return TLSCertificateOverrides{
		Verify: v.verify,
	}
}

// TLSCertificate: certificate used for TLS authentication and encryption. This
// can represent either a certificate only (eg, the certificate received by a
// client from a server), or the combination of a certificate and a private key
// (which is needed when acting as a ServerConnection).
type TLSCertificate struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*TLSCertificate)(nil)
)

// TLSCertificater describes types inherited from class TLSCertificate.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type TLSCertificater interface {
	coreglib.Objector
	baseTLSCertificate() *TLSCertificate
}

var _ TLSCertificater = (*TLSCertificate)(nil)

func init() {
	coreglib.RegisterClassInfo[*TLSCertificate, *TLSCertificateClass, TLSCertificateOverrides](
		GTypeTLSCertificate,
		initTLSCertificateClass,
		wrapTLSCertificate,
		defaultTLSCertificateOverrides,
	)
}

func initTLSCertificateClass(gclass unsafe.Pointer, overrides TLSCertificateOverrides, classInitFunc func(*TLSCertificateClass)) {
	pclass := (*C.GTlsCertificateClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeTLSCertificate))))

	if overrides.Verify != nil {
		pclass.verify = (*[0]byte)(C._gotk4_gio2_TlsCertificateClass_verify)
	}

	if classInitFunc != nil {
		class := (*TLSCertificateClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapTLSCertificate(obj *coreglib.Object) *TLSCertificate {
	return &TLSCertificate{
		Object: obj,
	}
}

func marshalTLSCertificate(p uintptr) (interface{}, error) {
	return wrapTLSCertificate(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (cert *TLSCertificate) baseTLSCertificate() *TLSCertificate {
	return cert
}

// BaseTLSCertificate returns the underlying base object.
func BaseTLSCertificate(obj TLSCertificater) *TLSCertificate {
	return obj.baseTLSCertificate()
}

// NewTLSCertificateFromFile creates a Certificate from the PEM-encoded data in
// file. The returned certificate will be the first certificate found in file.
// As of GLib 2.44, if file contains more certificates it will try to load a
// certificate chain. All certificates will be verified in the order found
// (top-level certificate should be the last one in the file) and the
// Certificate:issuer property of each certificate will be set accordingly if
// the verification succeeds. If any certificate in the chain cannot be
// verified, the first certificate in the file will still be returned.
//
// If file cannot be read or parsed, the function will return NULL and set
// error. Otherwise, this behaves like g_tls_certificate_new_from_pem().
//
// The function takes the following parameters:
//
//    - file containing a PEM-encoded certificate to import.
//
// The function returns the following values:
//
//    - tlsCertificate: new certificate, or NULL on error.
//
func NewTLSCertificateFromFile(file string) (*TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(file)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_tls_certificate_new_from_file(_arg1, &_cerr)
	runtime.KeepAlive(file)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = wrapTLSCertificate(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// NewTLSCertificateFromFiles creates a Certificate from the PEM-encoded data in
// cert_file and key_file. The returned certificate will be the first
// certificate found in cert_file. As of GLib 2.44, if cert_file contains more
// certificates it will try to load a certificate chain. All certificates will
// be verified in the order found (top-level certificate should be the last one
// in the file) and the Certificate:issuer property of each certificate will be
// set accordingly if the verification succeeds. If any certificate in the chain
// cannot be verified, the first certificate in the file will still be returned.
//
// If either file cannot be read or parsed, the function will return NULL and
// set error. Otherwise, this behaves like g_tls_certificate_new_from_pem().
//
// The function takes the following parameters:
//
//    - certFile: file containing one or more PEM-encoded certificates to import.
//    - keyFile: file containing a PEM-encoded private key to import.
//
// The function returns the following values:
//
//    - tlsCertificate: new certificate, or NULL on error.
//
func NewTLSCertificateFromFiles(certFile, keyFile string) (*TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(certFile)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(keyFile)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.g_tls_certificate_new_from_files(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(certFile)
	runtime.KeepAlive(keyFile)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = wrapTLSCertificate(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// NewTLSCertificateFromPem creates a Certificate from the PEM-encoded data in
// data. If data includes both a certificate and a private key, then the
// returned certificate will include the private key data as well. (See the
// Certificate:private-key-pem property for information about supported
// formats.)
//
// The returned certificate will be the first certificate found in data. As of
// GLib 2.44, if data contains more certificates it will try to load a
// certificate chain. All certificates will be verified in the order found
// (top-level certificate should be the last one in the file) and the
// Certificate:issuer property of each certificate will be set accordingly if
// the verification succeeds. If any certificate in the chain cannot be
// verified, the first certificate in the file will still be returned.
//
// The function takes the following parameters:
//
//    - data: PEM-encoded certificate data.
//    - length of data, or -1 if it's 0-terminated.
//
// The function returns the following values:
//
//    - tlsCertificate: new certificate, or NULL if data is invalid.
//
func NewTLSCertificateFromPem(data string, length int) (*TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 C.gssize           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(data)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.gssize(length)

	_cret = C.g_tls_certificate_new_from_pem(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(data)
	runtime.KeepAlive(length)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = wrapTLSCertificate(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// NewTLSCertificateFromPKCS11URIs creates a Certificate from a PKCS \#11 URI.
//
// An example pkcs11_uri would be
// pkcs11:model=Model;manufacturer=Manufacture;serial=1;token=My20Client20Certificate;id=01
//
// Where the token’s layout is:
//
//    Object 0:
//      URL: pkcs11:model=Model;manufacturer=Manufacture;serial=1;token=My20Client20Certificate;id=01;object=private20key;type=private
//      Type: Private key (RSA-2048)
//      ID: 01
//
//    Object 1:
//      URL: pkcs11:model=Model;manufacturer=Manufacture;serial=1;token=My20Client20Certificate;id=01;object=Certificate20for20Authentication;type=cert
//      Type: X.509 Certificate (RSA-2048)
//      ID: 01
//
//
// In this case the certificate and private key would both be detected and used
// as expected. pkcs_uri may also just reference an X.509 certificate object and
// then optionally private_key_pkcs11_uri allows using a private key exposed
// under a different URI.
//
// Note that the private key is not accessed until usage and may fail or require
// a PIN later.
//
// The function takes the following parameters:
//
//    - pkcs11Uri: PKCS \#11 URI.
//    - privateKeyPkcs11Uri (optional): PKCS \#11 URI.
//
// The function returns the following values:
//
//    - tlsCertificate: new certificate, or NULL on error.
//
func NewTLSCertificateFromPKCS11URIs(pkcs11Uri, privateKeyPkcs11Uri string) (*TLSCertificate, error) {
	var _arg1 *C.gchar           // out
	var _arg2 *C.gchar           // out
	var _cret *C.GTlsCertificate // in
	var _cerr *C.GError          // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(pkcs11Uri)))
	defer C.free(unsafe.Pointer(_arg1))
	if privateKeyPkcs11Uri != "" {
		_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(privateKeyPkcs11Uri)))
		defer C.free(unsafe.Pointer(_arg2))
	}

	_cret = C.g_tls_certificate_new_from_pkcs11_uris(_arg1, _arg2, &_cerr)
	runtime.KeepAlive(pkcs11Uri)
	runtime.KeepAlive(privateKeyPkcs11Uri)

	var _tlsCertificate *TLSCertificate // out
	var _goerr error                    // out

	_tlsCertificate = wrapTLSCertificate(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _tlsCertificate, _goerr
}

// Issuer gets the Certificate representing cert's issuer, if known.
//
// The function returns the following values:
//
//    - tlsCertificate (optional): certificate of cert's issuer, or NULL if cert
//      is self-signed or signed with an unknown certificate.
//
func (cert *TLSCertificate) Issuer() TLSCertificater {
	var _arg0 *C.GTlsCertificate // out
	var _cret *C.GTlsCertificate // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(cert).Native()))

	_cret = C.g_tls_certificate_get_issuer(_arg0)
	runtime.KeepAlive(cert)

	var _tlsCertificate TLSCertificater // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(TLSCertificater)
				return ok
			})
			rv, ok := casted.(TLSCertificater)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSCertificater")
			}
			_tlsCertificate = rv
		}
	}

	return _tlsCertificate
}

// IsSame: check if two Certificate objects represent the same certificate. The
// raw DER byte data of the two certificates are checked for equality. This has
// the effect that two certificates may compare equal even if their
// Certificate:issuer, Certificate:private-key, or Certificate:private-key-pem
// properties differ.
//
// The function takes the following parameters:
//
//    - certTwo: second certificate to compare.
//
// The function returns the following values:
//
//    - ok: whether the same or not.
//
func (certOne *TLSCertificate) IsSame(certTwo TLSCertificater) bool {
	var _arg0 *C.GTlsCertificate // out
	var _arg1 *C.GTlsCertificate // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(certOne).Native()))
	_arg1 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(certTwo).Native()))

	_cret = C.g_tls_certificate_is_same(_arg0, _arg1)
	runtime.KeepAlive(certOne)
	runtime.KeepAlive(certTwo)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Verify: this verifies cert and returns a set of CertificateFlags indicating
// any problems found with it. This can be used to verify a certificate outside
// the context of making a connection, or to check a certificate against a CA
// that is not part of the system CA database.
//
// If identity is not NULL, cert's name(s) will be compared against it, and
// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value if it does not
// match. If identity is NULL, that bit will never be set in the return value.
//
// If trusted_ca is not NULL, then cert (or one of the certificates in its
// chain) must be signed by it, or else G_TLS_CERTIFICATE_UNKNOWN_CA will be set
// in the return value. If trusted_ca is NULL, that bit will never be set in the
// return value.
//
// (All other CertificateFlags values will always be set or unset as
// appropriate.).
//
// The function takes the following parameters:
//
//    - identity (optional): expected peer identity.
//    - trustedCa (optional): certificate of a trusted authority.
//
// The function returns the following values:
//
//    - tlsCertificateFlags: appropriate CertificateFlags.
//
func (cert *TLSCertificate) Verify(identity SocketConnectabler, trustedCa TLSCertificater) TLSCertificateFlags {
	var _arg0 *C.GTlsCertificate     // out
	var _arg1 *C.GSocketConnectable  // out
	var _arg2 *C.GTlsCertificate     // out
	var _cret C.GTlsCertificateFlags // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(cert).Native()))
	if identity != nil {
		_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(identity).Native()))
	}
	if trustedCa != nil {
		_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(trustedCa).Native()))
	}

	_cret = C.g_tls_certificate_verify(_arg0, _arg1, _arg2)
	runtime.KeepAlive(cert)
	runtime.KeepAlive(identity)
	runtime.KeepAlive(trustedCa)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}

// Verify: this verifies cert and returns a set of CertificateFlags indicating
// any problems found with it. This can be used to verify a certificate outside
// the context of making a connection, or to check a certificate against a CA
// that is not part of the system CA database.
//
// If identity is not NULL, cert's name(s) will be compared against it, and
// G_TLS_CERTIFICATE_BAD_IDENTITY will be set in the return value if it does not
// match. If identity is NULL, that bit will never be set in the return value.
//
// If trusted_ca is not NULL, then cert (or one of the certificates in its
// chain) must be signed by it, or else G_TLS_CERTIFICATE_UNKNOWN_CA will be set
// in the return value. If trusted_ca is NULL, that bit will never be set in the
// return value.
//
// (All other CertificateFlags values will always be set or unset as
// appropriate.).
//
// The function takes the following parameters:
//
//    - identity (optional): expected peer identity.
//    - trustedCa (optional): certificate of a trusted authority.
//
// The function returns the following values:
//
//    - tlsCertificateFlags: appropriate CertificateFlags.
//
func (cert *TLSCertificate) verify(identity SocketConnectabler, trustedCa TLSCertificater) TLSCertificateFlags {
	gclass := (*C.GTlsCertificateClass)(coreglib.PeekParentClass(cert))
	fnarg := gclass.verify

	var _arg0 *C.GTlsCertificate     // out
	var _arg1 *C.GSocketConnectable  // out
	var _arg2 *C.GTlsCertificate     // out
	var _cret C.GTlsCertificateFlags // in

	_arg0 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(cert).Native()))
	if identity != nil {
		_arg1 = (*C.GSocketConnectable)(unsafe.Pointer(coreglib.InternObject(identity).Native()))
	}
	if trustedCa != nil {
		_arg2 = (*C.GTlsCertificate)(unsafe.Pointer(coreglib.InternObject(trustedCa).Native()))
	}

	_cret = C._gotk4_gio2_TLSCertificate_virtual_verify(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(cert)
	runtime.KeepAlive(identity)
	runtime.KeepAlive(trustedCa)

	var _tlsCertificateFlags TLSCertificateFlags // out

	_tlsCertificateFlags = TLSCertificateFlags(_cret)

	return _tlsCertificateFlags
}

// TLSCertificateListNewFromFile creates one or more Certificates from the
// PEM-encoded data in file. If file cannot be read or parsed, the function will
// return NULL and set error. If file does not contain any PEM-encoded
// certificates, this will return an empty list and not set error.
//
// The function takes the following parameters:
//
//    - file containing PEM-encoded certificates to import.
//
// The function returns the following values:
//
//    - list: a #GList containing Certificate objects. You must free the list and
//      its contents when you are done with it.
//
func TLSCertificateListNewFromFile(file string) ([]TLSCertificater, error) {
	var _arg1 *C.gchar  // out
	var _cret *C.GList  // in
	var _cerr *C.GError // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(file)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.g_tls_certificate_list_new_from_file(_arg1, &_cerr)
	runtime.KeepAlive(file)

	var _list []TLSCertificater // out
	var _goerr error            // out

	_list = make([]TLSCertificater, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GTlsCertificate)(v)
		var dst TLSCertificater // out
		{
			objptr := unsafe.Pointer(src)
			if objptr == nil {
				panic("object of type gio.TLSCertificater is nil")
			}

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(TLSCertificater)
				return ok
			})
			rv, ok := casted.(TLSCertificater)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.TLSCertificater")
			}
			dst = rv
		}
		_list = append(_list, dst)
	})
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _list, _goerr
}
