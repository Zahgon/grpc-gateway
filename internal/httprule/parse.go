package httprule

// InvalidTemplateError indicates that the path template is not valid.
type InvalidTemplateError struct {
	tmpl string
	msg  string
}

func (e InvalidTemplateError) Error() string { _ = "STUB: not implemented"; return "" }

// Parse parses the string representation of path template
func Parse(tmpl string) (Compiler, error) { _ = "STUB: not implemented"; return *new(Compiler), nil }

func tokenize(path string) (tokens []string, verb string) {
	_ = "STUB: not implemented"
	return nil, ""
}

// See
// https://github.com/grpc-ecosystem/grpc-gateway/pull/1947#issuecomment-774523693 ;
// although normal and backwards-compat logic here is to use the last index
// of a colon, if the final segment is a variable followed by a colon, the
// part following the colon must be a verb. Hence if the previous token is
// an end var marker, we switch the index we're looking for to Index instead
// of LastIndex, so that we correctly grab the remaining part of the path as
// the verb.

// Not enough to be variable so skip this logic and don't result in an
// invalid index

// parser is a parser of the template syntax defined in github.com/googleapis/googleapis/google/api/http.proto.
type parser struct {
	tokens   []string
	accepted []string
}

// topLevelSegments is the target of this parser.
func (p *parser) topLevelSegments() ([]segment, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *parser) segments() ([]segment, error) { _ = "STUB: not implemented"; return nil, nil }

func (p *parser) segment() (segment, error) { _ = "STUB: not implemented"; return *new(segment), nil }

func (p *parser) literal() (segment, error) { _ = "STUB: not implemented"; return *new(segment), nil }

func (p *parser) variable() (segment, error) { _ = "STUB: not implemented"; return *new(segment), nil }

func (p *parser) fieldPath() (string, error) { _ = "STUB: not implemented"; return "", nil }

// A termType is a type of terminal symbols.
type termType string

// These constants define some of valid values of termType.
// They improve readability of parse functions.
//
// You can also use "/", "*", "**", "." or "=" as valid values.
const (
	typeIdent   = termType("ident")
	typeLiteral = termType("literal")
	typeEOF     = termType("$")
)

// eof is the terminal symbol which always appears at the end of token sequence.
const eof = "\u0000"

// accept tries to accept a token in "p".
// This function consumes a token and returns it if it matches to the specified "term".
// If it doesn't match, the function does not consume any tokens and return an error.
func (p *parser) accept(term termType) (string, error) { _ = "STUB: not implemented"; return "", nil }

// expectPChars determines if "t" consists of only pchars defined in RFC3986.
//
// https://www.ietf.org/rfc/rfc3986.txt, P.49
//
//	pchar         = unreserved / pct-encoded / sub-delims / ":" / "@"
//	unreserved    = ALPHA / DIGIT / "-" / "." / "_" / "~"
//	sub-delims    = "!" / "$" / "&" / "'" / "(" / ")"
//	              / "*" / "+" / "," / ";" / "="
//	pct-encoded   = "%" HEXDIG HEXDIG
func expectPChars(t string) error { _ = "STUB: not implemented"; return nil }

// unreserved

// unreserved

// sub-delims

// rest of pchar

// pct-encoded

// expectIdent determines if "ident" is a valid identifier in .proto schema ([[:alpha:]_][[:alphanum:]_]*).
func expectIdent(ident string) error { _ = "STUB: not implemented"; return nil }

func isHexDigit(r rune) bool { _ = "STUB: not implemented"; return false }
