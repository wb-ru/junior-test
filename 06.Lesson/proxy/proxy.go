package proxy

import (
	"regexp"
	"strings"
)

// a simple proxy that proxies regexp
// simple alhanumeric queris are done without regexes
// regexes that are already existant are cached
type RegProxy struct {
	cache map[string]*regexp.Regexp
}

// Checks if all symbols in s are alhanumeric
func (p *RegProxy) isAlphaNum(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
			return false
		}
	}
	return true
}

// MatchString reports whether the string s contains any match of the regular expression query q
func (p *RegProxy) MatchString(s string, q string) bool {
	if p.cache == nil {
		p.cache = make(map[string]*regexp.Regexp)
	}
	// check whether this regex was cached
	r, ok := p.cache[q]
	if ok {
		return r.MatchString(s)
	}
	// check if query is short and simple
	if len(q) < 256 && p.isAlphaNum(q) {
		return strings.Contains(s, q)
	}
	r, _ = regexp.Compile(q)
	// cache it!
	p.cache[q] = r

	return r.MatchString(s)
}
