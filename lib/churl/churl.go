package churl

// This file is a copy of `net/url/parse.go`, with some modifications to make it work for ClickHouse HA URLs.
// Most private fuctions and constants are copied as-is from Go 1.25.7.
// The reason is that the origin `Parse()` func does not support multiple hosts in the Host part after Go 1.26.
// See the original issue to more details: https://github.com/golang/go/issues/75859

import (
	neturl "net/url"
)

type encoding int

const (
	encodePath encoding = 1 + iota
	encodePathSegment
	encodeHost
	encodeZone
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

const upperhex = "0123456789ABCDEF"

func Parse(rawURL string) (*neturl.URL, error) {
	_ = "STUB: not implemented"
	// Cut off #frag
	return nil, nil
}

func parse(rawURL string, viaRequest bool) (*neturl.URL, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Split off possible leading "http:", "mailto:", etc.
// Cannot contain escaped characters.

// We consider rootless paths per RFC 3986 as opaque.

// Avoid confusion with malformed schemes, like cache_object:foo/bar.
// See golang.org/issue/16822.
//
// RFC 3986, §3.3:
// In addition, a URI reference (Section 4.1) may be a relative-path reference,
// in which case the first path segment cannot contain a colon (":") character.

// First path segment has colon. Not allowed in relative URL.

// OmitHost is set to true when rawURL has an empty host (authority).
// See golang.org/issue/46059.

// Set Path and, optionally, RawPath.
// RawPath is a hint of the encoding of Path. We don't want to set it if
// the default escaping of Path is equivalent, to help make sure that people
// don't rely on it in general.

func getScheme(rawURL string) (scheme, path string, err error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

// do nothing

// we have encountered an invalid character,
// so there is no valid scheme

func parseAuthority(authority string) (user *neturl.Userinfo, host string, err error) {
	_ = "STUB: not implemented"
	return nil, "", nil
}

func parseHost(host string) (string, error) { _ = "STUB: not implemented"; return "", nil }

// Parse an IP-Literal in RFC 3986 and RFC 6874.
// E.g., "[fe80::1]", "[fe80::1%25en0]", "[fe80::1]:80".

// RFC 6874 defines that %25 (%-encoded percent) introduces
// the zone identifier, and the zone identifier can use basically
// any %-encoding it likes. That's different from the host, which
// can only %-encode non-ASCII bytes.
// We do impose some restrictions on the zone, to avoid stupidity
// like newlines.

// Per RFC 3986, only a host identified by a valid
// IPv6 address can be enclosed by square brackets.
// This excludes any IPv4, but notably not IPv4-mapped addresses.

func validUserinfo(s string) bool { _ = "STUB: not implemented"; return false }

// `RFC 3986 section 3.2.1` does not allow '@' in userinfo.
// It is a delimiter between userinfo and host.
// However, URLs are diverse, and in some cases,
// the userinfo may contain an '@' character,
// for example, in "http://username:p@ssword@google.com",
// the string "username:p@ssword" should be treated as valid userinfo.
// Ref:
//   https://go.dev/issue/3439
//   https://go.dev/issue/22655

func validOptionalPort(port string) bool { _ = "STUB: not implemented"; return false }

func unescape(s string, mode encoding) (string, error) {
	_ = "STUB: not implemented"
	// Count %, check that they're well-formed.
	return "", nil
}

// Per https://tools.ietf.org/html/rfc3986#page-21
// in the host component %-encoding can only be used
// for non-ASCII bytes.
// But https://tools.ietf.org/html/rfc6874#section-2
// introduces %25 being allowed to escape a percent sign
// in IPv6 scoped-address literals. Yay.

// RFC 6874 says basically "anything goes" for zone identifiers
// and that even non-ASCII can be redundantly escaped,
// but it seems prudent to restrict %-escaped bytes here to those
// that are valid host name bytes in their unescaped form.
// That is, you can use escaping in the zone identifier but not
// to introduce bytes you couldn't just write directly.
// But Windows puts spaces here! Yay.

func escape(s string, mode encoding) string { _ = "STUB: not implemented"; return "" }

func shouldEscape(c byte, mode encoding) bool {
	_ = "STUB: not implemented"
	// §2.3 Unreserved characters (alphanum)
	return false
}

// §3.2.2 Host allows
//	sub-delims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
// as part of reg-name.
// We add : because we include :port as part of host.
// We add [ ] because we include [ipv6]:port as part of host.
// We add < > because they're the only characters left that
// we could possibly allow, and Parse will reject them if we
// escape them (because hosts can't use %-encoding for
// ASCII bytes).

// §2.3 Unreserved characters (mark)

// §2.2 Reserved characters (reserved)
// Different sections of the URL allow a few of
// the reserved characters to appear unescaped.

// §3.3
// The RFC allows : @ & = + $ but saves / ; , for assigning
// meaning to individual path segments. This package
// only manipulates the path as a whole, so we allow those
// last three as well. That leaves only ? to escape.

// §3.3
// The RFC allows : @ & = + $ but saves / ; , for assigning
// meaning to individual path segments.

// §3.2.1
// The RFC allows ';', ':', '&', '=', '+', '$', and ',' in
// userinfo, so we must escape only '@', '/', and '?'.
// The parsing of userinfo treats ':' as special so we must escape
// that too.

// §3.4
// The RFC reserves (so we must escape) everything.

// §4.1
// The RFC text is silent but the grammar allows
// everything, so escape nothing.

// RFC 3986 §2.2 allows not escaping sub-delims. A subset of sub-delims are
// included in reserved from RFC 2396 §2.2. The remaining sub-delims do not
// need to be escaped. To minimize potential breakage, we apply two restrictions:
// (1) we always escape sub-delims outside of the fragment, and (2) we always
// escape single quote to avoid breaking callers that had previously assumed that
// single quotes would be escaped. See issue #19917.

// Everything else must be escaped.

func ishex(c byte) bool { _ = "STUB: not implemented"; return false }

func unhex(c byte) byte { _ = "STUB: not implemented"; return 0 }

func stringContainsCTLByte(s string) bool { _ = "STUB: not implemented"; return false }

func setFragment(u *neturl.URL, f string) error { _ = "STUB: not implemented"; return nil }

// Default encoding is fine.

func setPath(u *neturl.URL, p string) error { _ = "STUB: not implemented"; return nil }

// Default encoding is fine.
