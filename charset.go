package validpass

import (
    "regexp"
    "strings"
)

var (
    regexpUpper = regexp.MustCompile("[[:upper:]]")
    regexpLower = regexp.MustCompile("[[:lower:]]")
    regexpDigit = regexp.MustCompilePOSIX("[[:digit:]]")
    regexpSpecial = regexp.MustCompile("[!\"#$%&'()*+,\\-./:;<=>?@[\\\\\\]^_`{|}~]")
    regexpInvalid = regexp.MustCompile(`(\n|\t|\v|\r|\a|\f)`)
)

// HasUpper returns whether the string has an uppercase character in it
func HasUpper(pass string) bool {
    return regexpUpper.MatchString(pass)
}

// HasLower returns whether the string has a lowercase character in it
func HasLower(pass string) bool {
    return regexpLower.MatchString(pass)
}

// HasDigit returns whether the string has a digit (0-9) in it
func HasDigit(pass string) bool {
    return regexpDigit.MatchString(pass)
}

// HasSpace returns whether the string has a space (" ") in it
func HasSpace(pass string) bool {
    return strings.Contains(pass, " ")
}

// HasSpecial returns whether the string has a special character in it.
// Special characters are defined as one of !"#$%&'()*+,\-./:;<=>?@[]^_`{|}~
func HasSpecial(pass string) bool {
    return regexpSpecial.MatchString(pass)
}

// HasInvalid returns whether the string contains any invalid characters, i.e.
// anything that could not be entered into a form field.
func HasInvalid(pass string) bool {
    return regexpInvalid.MatchString(pass)
}