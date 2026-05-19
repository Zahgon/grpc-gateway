package genopenapi

import (
	"regexp"
)

// LookupNamingStrategy looks up the given naming strategy and returns the naming
// strategy function for it. The naming strategy function takes in the list of all
// fully-qualified proto message names, and returns a mapping from fully-qualified
// name to OpenAPI name.
func LookupNamingStrategy(strategyName string) func([]string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// resolveNamesFQN uses the fully-qualified proto message name as the
// OpenAPI name, stripping the leading dot.
func resolveNamesFQN(messages []string) map[string]string { _ = "STUB: not implemented"; return nil }

// strip leading dot from proto fqn

// resolveNamesLegacy takes the names of all proto messages and generates unique references by
// applying the legacy heuristics for deriving unique names: starting from the bottom of the name hierarchy, it
// determines the minimum number of components necessary to yield a unique name, adds one
// to that number, and then concatenates those last components with no separator in between
// to form a unique name.
//
// E.g., if the fully qualified name is `.a.b.C.D`, and there are other messages with fully
// qualified names ending in `.D` but not in `.C.D`, it assigns the unique name `bCD`.
func resolveNamesLegacy(messages []string) map[string]string { _ = "STUB: not implemented"; return nil }

// resolveNamesSimple takes the names of all proto messages and generates unique references by using a simple
// heuristic: starting from the bottom of the name hierarchy, it determines the minimum
// number of components necessary to yield a unique name, and then concatenates those last
// components with a "." separator in between to form a unique name.
//
// E.g., if the fully qualified name is `.a.b.C.D`, and there are other messages with
// fully qualified names ending in `.D` but not in `.C.D`, it assigns the unique name `C.D`.
func resolveNamesSimple(messages []string) map[string]string { _ = "STUB: not implemented"; return nil }

// resolveNamesPackage takes the names of all proto messages and generates unique references by
// starting with the package-scoped name (with nested message types qualified by their containing
// "parent" types), and then following the "simple" heuristic above to add package name components
// until each message has a unique name with a "." between each component.
//
// E.g., if the fully qualified name is `.a.b.C.D`, the name is `C.D` unless there is another
// package-scoped name ending in "C.D", in  which case it would be `b.C.D` (unless that also
// conflicted, in which case the name would be the fully-qualified `a.b.C`).
func resolveNamesPackage(messages []string) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// For the "package" naming strategy, we rely on the convention that package names are lowercase
// but message names are capitalized.
var pkgEndRegexp = regexp.MustCompile(`\.[A-Z]`)

// Take the names of every proto message and generates a unique reference by:
// first, separating each message name into its components by splitting at dots. Then,
// take the shortest suffix slice from each components slice that is unique among all
// messages, and convert it into a component name by taking extraContext additional
// components into consideration and joining all components with componentSeparator.
func resolveNamesUniqueWithContext(messages []string, extraContext int, componentSeparator string, qualifyNestedMessages bool) map[string]string {
	_ = "STUB: not implemented"
	return nil
}

// Fall back to non-qualified behavior if search based on convention fails.

// Return each package component as an element, followed by the full message name
// (potentially qualified, if nested) as a single element.

// depth + extraContext > 0 ensures that we only break for values of depth when the
// resulting slice of name components is non-empty. Otherwise, we would return the
// empty string as the concise unique name is len(messages) == 1 (which is
// technically correct).

// When using empty separator (legacy mode), apply camelCase by title-casing
// intermediate lowercase package components. E.g., "google.rpc.Status" -> "googleRpcStatus"
// We only title-case components that are entirely lowercase (package names),
// not message names which are already PascalCase.
// Skip the first non-empty component (keep it lowercase for camelCase).
