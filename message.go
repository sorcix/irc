// Copyright 2014 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

package irc

import (
	"bytes"
	"strings"
)

// Various constants used for formatting IRC messages.
const (
	prefix     byte = 0x3A // Prefix or last argument
	prefixUser byte = 0x21 // Username
	prefixHost byte = 0x40 // Hostname
	space      byte = 0x20 // Separator

	maxLength = 510 // Maximum length is 512 - 2 for the line endings.
)

func cutsetFunc(r rune) bool {
	// Characters to trim from prefixes/messages.
	return r == '\r' || r == '\n'
}

// Objects implementing the Sender interface are able to send messages to an IRC server.
//
// As there might be a message queue, it is possible that Send returns a nil
// error, but the message is not sent (yet). The error value is only used when
// it is certain that sending the message is impossible.
//
// This interface is not used inside this package, and shouldn't have been
// defined here in the first place. For backwards compatibility only.
type Sender interface {
	Send(*Message) error
}

// Prefix represents the prefix (sender) of an IRC message.
// See RFC1459 section 2.3.1.
//
//    <servername> | <nick> [ '!' <user> ] [ '@' <host> ]
//
type Prefix struct {
	Name string // Nick- or servername
	User string // Username
	Host string // Hostname
}

// ParsePrefix takes a string and attempts to create a Prefix struct.
func ParsePrefix(raw string) (p *Prefix) {

	p = new(Prefix)

	user := indexByte(raw, prefixUser)
	host := indexByte(raw, prefixHost)

	switch {

	case user > 0 && host > user:
		p.Name = raw[:user]
		p.User = raw[user+1 : host]
		p.Host = raw[host+1:]

	case user > 0:
		p.Name = raw[:user]
		p.User = raw[user+1:]

	case host > 0:
		p.Name = raw[:host]
		p.Host = raw[host+1:]

	default:
		p.Name = raw

	}

	return p
}

// Len calculates the length of the string representation of this prefix.
func (p *Prefix) Len() (length int) {
	length = len(p.Name)
	if len(p.User) > 0 {
		length = length + len(p.User) + 1
	}
	if len(p.Host) > 0 {
		length = length + len(p.Host) + 1
	}
	return
}

// Bytes returns a []byte representation of this prefix.
func (p *Prefix) Bytes() []byte {
	buffer := new(bytes.Buffer)
	p.writeTo(buffer)
	return buffer.Bytes()
}

// String returns a string representation of this prefix.
func (p *Prefix) String() (s string) {
	// Benchmarks revealed that in this case simple string concatenation
	// is actually faster than using a ByteBuffer as in (*Message).String()
	s = p.Name
	if len(p.User) > 0 {
		s = s + string(prefixUser) + p.User
	}
	if len(p.Host) > 0 {
		s = s + string(prefixHost) + p.Host
	}
	return
}

// IsHostmask returns true if this prefix looks like a user hostmask.
func (p *Prefix) IsHostmask() bool {
	return len(p.User) > 0 && len(p.Host) > 0
}

// IsServer returns true if this prefix looks like a server name.
func (p *Prefix) IsServer() bool {
	return len(p.User) <= 0 && len(p.Host) <= 0 // && indexByte(p.Name, '.') > 0
}

// writeTo is an utility function to write the prefix to the bytes.Buffer in Message.String().
func (p *Prefix) writeTo(buffer *bytes.Buffer) {
	buffer.WriteString(p.Name)
	if len(p.User) > 0 {
		buffer.WriteByte(prefixUser)
		buffer.WriteString(p.User)
	}
	if len(p.Host) > 0 {
		buffer.WriteByte(prefixHost)
		buffer.WriteString(p.Host)
	}
	return
}

// Message represents an IRC protocol message.
// See RFC1459 section 2.3.1.
//
//    <message>  ::= [':' <prefix> <SPACE> ] <command> <params> <crlf>
//    <prefix>   ::= <servername> | <nick> [ '!' <user> ] [ '@' <host> ]
//    <command>  ::= <letter> { <letter> } | <number> <number> <number>
//    <SPACE>    ::= ' ' { ' ' }
//    <params>   ::= <SPACE> [ ':' <trailing> | <middle> <params> ]
//
//    <middle>   ::= <Any *non-empty* sequence of octets not including SPACE
//                   or NUL or CR or LF, the first of which may not be ':'>
//    <trailing> ::= <Any, possibly *empty*, sequence of octets not including
//                   NUL or CR or LF>
//
//    <crlf>     ::= CR LF
type Message struct {
	*Prefix
	Command string
	Params  []string
}

// ParseMessage takes a string and attempts to create a Message struct.
// Returns nil if the Message is invalid.
func ParseMessage(raw string) (m *Message) {

	// Ignore empty messages.
	if raw = strings.TrimFunc(raw, cutsetFunc); len(raw) < 2 {
		return nil
	}

	m = new(Message)

	if raw[0] == prefix {

		// Split the string on space, skipping the first character
		split := strings.SplitN(raw[1:], " ", 2)
		if len(split) < 2 || len(split[0]) < 1 {
			return nil
		}

		// Parse the first part of the split as the prefix
		m.Prefix = ParsePrefix(split[0])

		// We can continue to parse the remainder of the message
		raw = split[1]

	}

	// We split out trailing now. This way, using Fields to split the rest of
	// the parameters will remove duplicate spaces
	split := strings.SplitN(raw, " :", 2)

	// Because params can be delimited by one or more spaces, we have to use
	// strings.FieldsFunc, rather than strings.Split
	m.Params = strings.FieldsFunc(split[0], func(r rune) bool {
		return r == ' '
	})

	// The first param is the command, so we bail if we can't find it
	if len(m.Params) < 1 {
		return nil
	}

	m.Command = m.Params[0]
	m.Params = m.Params[1:]

	// Append the trailing argument onto the params if we had one
	if len(split) == 2 {
		m.Params = append(m.Params, split[1])
	}

	// If there were no params, set it to nil to be consistent
	if len(m.Params) == 0 {
		m.Params = nil
	}

	return m

}

//Trailing returns the last param of a message, or an empty string if there
// were no params.
func (m *Message) Trailing() string {

	if len(m.Params) == 0 {
		return ""
	}

	return m.Params[len(m.Params)-1]

}

// Len calculates the length of the string representation of this message.
func (m *Message) Len() (length int) {

	if m.Prefix != nil {
		length = m.Prefix.Len() + 2 // Include prefix and trailing space
	}

	length = length + len(m.Command)

	if len(m.Params) > 0 {
		// len(m.Params) is the number of spaces. We add one for the : in the
		// trailing argument.
		length = length + len(m.Params) + 1
		for _, param := range m.Params {
			length = length + len(param)
		}
	}

	return
}

// Bytes returns a []byte representation of this message.
//
// As noted in rfc2812 section 2.3, messages should not exceed 512 characters
// in length. This method forces that limit by discarding any characters
// exceeding the length limit.
func (m *Message) Bytes() []byte {

	buffer := new(bytes.Buffer)

	// Message prefix
	if m.Prefix != nil {
		buffer.WriteByte(prefix)
		m.Prefix.writeTo(buffer)
		buffer.WriteByte(space)
	}

	// Command is required
	buffer.WriteString(m.Command)

	// Space separated list of arguments
	if len(m.Params) > 0 {
		if len(m.Params) > 1 {
			buffer.WriteByte(space)
			buffer.WriteString(strings.Join(m.Params[:len(m.Params)-1], string(space)))
		}
		buffer.WriteByte(space)
		buffer.WriteByte(prefix)
		buffer.WriteString(m.Params[len(m.Params)-1])
	}

	// We need the limit the buffer length.
	if buffer.Len() > (maxLength) {
		buffer.Truncate(maxLength)
	}

	return buffer.Bytes()
}

// String returns a string representation of this message.
//
// As noted in rfc2812 section 2.3, messages should not exceed 512 characters
// in length. This method forces that limit by discarding any characters
// exceeding the length limit.
func (m *Message) String() string {
	return string(m.Bytes())
}
