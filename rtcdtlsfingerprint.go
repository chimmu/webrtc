package webrtc

import "crypto/sha256"

type digestFunc func([]byte) []byte

var fingerprintAlgoritms = map[string]digestFunc{
	// "md2":     nil, // [RFC3279]
	// "md5":     nil, // [RFC3279]
	// "sha-1":   nil, // [RFC3279]
	// "sha-224": nil, // [RFC4055]
	"sha-256": sha256Digest, // [RFC4055]
	// "sha-384": nil, // [RFC4055]
	// "sha-512": nil, // [RFC4055]
}

func sha256Digest(b []byte) []byte {
	hash := sha256.Sum256(b)
	return hash[:]
}

// RTCDtlsFingerprint specifies the hash function algorithm and certificate
// fingerprint as described in https://tools.ietf.org/html/rfc4572.
type RTCDtlsFingerprint struct {
	// Algorithm specifies one of the the hash function algorithms defined in
	// the 'Hash function Textual Names' registry.
	Algorithm string

	// Value specifies the value of the certificate fingerprint in lowercase
	// hex string as expressed utilizing the syntax of 'fingerprint' in
	// https://tools.ietf.org/html/rfc4572#section-5.
	Value string
}
