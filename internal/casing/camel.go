package casing

// Camel returns the CamelCased name.
//
// This was moved from the now deprecated github.com/golang/protobuf/protoc-gen-go/generator package
//
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercases names, it's extremely unlikely to have two fields
// with different capitalizations.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func Camel(s string) string { _ = "STUB: not implemented"; return "" }

// Need a capital letter; drop the '_'.

// Invariant: if the next letter is lower case, it must be converted
// to upper case.
// That is, we process a word at a time, where words are marked by _ or
// upper case letter. Digits are treated as words.

// Skip the underscore in s.

// Assume we have a letter now - if not, it's a bogus identifier.
// The next word is a sequence of characters that must start upper case.

// Make it a capital letter.

// Guaranteed not lower case.
// Accept lower case sequence that follows.

// CamelIdentifier returns the CamelCased identifier without affecting the package name/path if any.
func CamelIdentifier(s string) string { _ = "STUB: not implemented"; return "" }

// JSONCamelCase converts a snake_case identifier to a camelCase identifier,
// according to the protobuf JSON specification.
func JSONCamelCase(s string) string { _ = "STUB: not implemented"; return "" }

// proto identifiers are always ASCII

// convert to uppercase

// And now lots of helper functions.

// Is c an ASCII lower-case letter?
func isASCIILower(c byte) bool { _ = "STUB: not implemented"; return false }

// Is c an ASCII digit?
func isASCIIDigit(c byte) bool { _ = "STUB: not implemented"; return false }
