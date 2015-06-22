// Copyright 2014 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

package irc

import (
	"bytes"
	"io"
	"reflect"
	"strings"
	"testing"
)

var normalizedStream = "PING :port80a.se.quakenet.org\r\n:port80a.se.quakenet.org PONG port80a.se.quakenet.org :port80a.se.quakenet.org\r\nPING :chat.freenode.net\r\n:wilhelm.freenode.net PONG wilhelm.freenode.net :chat.freenode.net\r\n"
var stream = "PING port80a.se.quakenet.org\r\n:port80a.se.quakenet.org PONG port80a.se.quakenet.org :port80a.se.quakenet.org\r\nPING chat.freenode.net\r\n:wilhelm.freenode.net PONG wilhelm.freenode.net :chat.freenode.net\r\n"

var result = [...]*Message{
	{
		Command: PING,
		Params:  []string{"port80a.se.quakenet.org"},
	},
	{
		Prefix: &Prefix{
			Name: "port80a.se.quakenet.org",
		},
		Command: PONG,
		Params:  []string{"port80a.se.quakenet.org", "port80a.se.quakenet.org"},
	},
	{
		Command: PING,
		Params:  []string{"chat.freenode.net"},
	},
	{
		Prefix: &Prefix{
			Name: "wilhelm.freenode.net",
		},
		Command: PONG,
		Params:  []string{"wilhelm.freenode.net", "chat.freenode.net"},
	},
}

func TestDecoder_Decode(t *testing.T) {

	reader := strings.NewReader(stream)
	dec := NewDecoder(reader)

	for i, test := range result {
		if message, err := dec.Decode(); err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		} else {
			if !reflect.DeepEqual(message, test) {
				t.Fatalf("Decoded message looks wrong! (%d)", i)
			}
		}
	}

	if _, err := dec.Decode(); err != io.EOF {
		t.Fatal("Decode should return an EOF error!")
	}
}

func TestEncoder_Encode(t *testing.T) {

	buffer := new(bytes.Buffer)
	enc := NewEncoder(buffer)

	for _, test := range result {
		if err := enc.Encode(test); err != nil {
			t.Fatalf("Unexpected error: %s", err.Error())
		}
	}

	if buffer.String() != normalizedStream {
		t.Fatalf("Encoded stream looks wrong!")
	}

}
