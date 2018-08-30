package validpass

import (
    "github.com/GeertJohan/go.rice"
    "strings"
)

var (
    top207 map[string]struct{}
    top1575 map[string]struct{}
    top12k map[string]struct{}
    top304k map[string]struct{}
)

func parseContents(file string) map[string]struct{} {
    m := make(map[string]struct{})
    strs := strings.Split(file, "\n")
    for _, str := range strs {
        m[strings.ToLower(str)] = struct{}{}
    }
    return m
}

func init() {
    box := rice.MustFindBox("lists")
    top207 = parseContents(box.MustString("top-207.txt"))
    top1575 = parseContents(box.MustString("top-1575.txt"))
    top12k = parseContents(box.MustString("top-12k.txt"))
    top304k = parseContents(box.MustString("top-304k.txt"))
}

// IsInTop207 returns if a given password is in the top 207 most common
// passwords
func IsInTop207(pass string) bool {
    _, ok := top207[strings.ToLower(pass)]
    return ok
}

// IsInTop1575 returns if a given password is in the top 1575 most common
// passwords
func IsInTop1575(pass string) bool {
    _, ok := top1575[strings.ToLower(pass)]
    return ok
}

// IsInTop12k returns if a given password is in the top 12,000 most common
// passwords
func IsInTop12k(pass string) bool {
    _, ok := top12k[strings.ToLower(pass)]
    return ok
}

// IsInTop304k returns if a given password is in the top 304,000 most common
// passwords
func IsInTop304k(pass string) bool {
    _, ok := top304k[strings.ToLower(pass)]
    return ok
}