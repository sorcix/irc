// Copyright 2014 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

package irc

import (
	"bytes"
	"strings"
)

const (
	prefixTag      byte = 0x40
	prefixTagValue byte = 0x3D
	tagSeparator   byte = 0x2C

	empty = ""
)

// Tags represents the key-value pairs in IRCv3 message tags.
//
// The values of this map contain the exact strings as sent over the connection,
// use the Get and Set function to decode.
type Tags map[string]string

func ParseTags(raw string) (t Tags) {

	parts := strings.Split(raw, string(tagSeparator))

	var value int

	t = make(Tags)

	for i := range parts {
		value = strings.IndexByte(parts[i], prefixTagValue)

		if value < 0 {

			// Malformed message tags may have an equals sign but no value.
			if len(parts[i]) < value+1 {
				t[parts[i]] = empty
				continue
			}

			t[parts[i][:value]] = parts[i][value+1:]
			continue
		}

		// This tag has no value, add an empty string.
		t[parts[i]] = empty
	}

	return t
}

// Len calculates the length of the string representation of this taglist.
func (t Tags) Len() (length int) {
	for key, value := range t {
		if len(value) < 1 {
			length = length + len(key) + 1
			continue
		}
		length = length + len(key) + len(value) + 2
	}
	length--

	return
}

// Bytes returns a []byte representation of this taglist.
func (t Tags) Bytes() []byte {
	buffer := new(bytes.Buffer)
	t.writeTo(buffer)
	return buffer.Bytes()
}

// String returns a string representation of this taglist.
func (t Tags) String() string {
	return string(t.Bytes())
}

// writeTo is an utility function to write message tags to the bytes.Buffer in Message.String().
func (t Tags) writeTo(buffer *bytes.Buffer) {

	var (
		max     int = len(t)
		current int
	)

	for key, value := range t {
		buffer.WriteString(key)
		if len(value) > 0 {
			buffer.WriteByte(prefixTagValue)
			buffer.WriteString(value)
		}
		if current < max {
			buffer.WriteByte(tagSeparator)
		}
		current++
	}
}

// Get returns the unescaped value of given tag key.
//
// If you need to get the escaped value as found in the IRC message, access the map directly.
func (t Tags) Get(key string) (string, bool) {
	//TODO(sorcix): Unescape
	if value, ok := t[key]; ok {
		return value, true
	}
	return empty, false
}

// Set escapes given value and saves it as the value for given key.
//
// If you want to save an already escaped value or a value that does not need escaping, access the map directly.
func (t Tags) Set(key, value string) {
	//TODO(sorcix): Escape
	t[key] = value
}
