package headers

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

const crlf = "\r\n"

type Headers map[string]string

func NewHeaders() Headers {
	return map[string]string{}
}

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 {
		return 0, false, nil
	}
	if idx == 0 {
		// the empty line
		// headers are done, consume the CRLF
		return 2, true, nil
	}

	parts := bytes.SplitN(data[:idx], []byte(":"), 2)
	key := string(parts[0])

	if key != strings.TrimRight(key, " ") {
		return 0, false, fmt.Errorf("invalid header name: %s", key)
	}

	value := bytes.TrimSpace(parts[1])
	key = strings.TrimSpace(key)
	key = strings.ToLower(key)

	for _, c := range key {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) && !isAllowedChar(c) {
			return 0, false, fmt.Errorf("invalid char in header: %c", c)
		}
	}

	_, ok := h[key]
	if ok {
		fmt.Println("Before conact: ", h[key])
		h[key] = h[key] + ", " + string(value)
	} else {
		h.Set(key, string(value))
	}

	return idx + 2, false, nil
}

func (h Headers) Set(key, value string) {
	h[key] = value
}

func isAllowedChar(r rune) bool {
	allowed := "!#$%&'*+-.^_`|~"

	for _, c := range allowed {
		if r == c {
			return true
		}
	}

	return false
}
