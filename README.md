mixedcase
=========
[![GoDoc](https://godoc.org/github.com/mohae/mixedcase?status.svg)](https://godoc.org/github.com/mohae/mixedcase)[![Build Status](https://travis-ci.org/mohae/mixedcase.png)](https://travis-ci.org/mohae/mixedcase)

MixedCase converts snake_case identifiers to MixedCase or mixedCase identifiers. It will also ensure that the identifiers do not start with invalid characters; converting numeric characters, [0-9], to their word equivalent and discarding invalid start characters. Common initialisms are Upper Cased unless they begin a mixedCase identifier, which are lower cased.

The resulting string is a valid Go identifier.

## Initialisms
Currently, initialisms do not cover all possible initialisms; just the ones present in Go's linter.
