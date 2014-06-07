// Copyright 2014 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

package ctcp

import (
	"testing"
)

func TestDecode(t *testing.T) {
	if _, _, ok := Decode("\x01\x01"); ok {
		t.Error("Message is invalid, but ok is true.")
	}
	if _, _, ok := Decode("\x01"); ok {
		t.Error("Message is invalid, but ok is true.")
	}
	if _, _, ok := Decode("\x01VERSION"); ok {
		t.Error("Message is invalid, but ok is true.")
	}
	if tag, message, ok := Decode("\x01VERSION\x01"); tag != "VERSION" || len(message) > 0 || !ok {
		t.Error("Message contains only a tag, wrong results.")
	}
	if tag, message, ok := Decode("\x01PING 123456789\x01"); tag != "PING" || message != "123456789" || !ok {
		t.Error("Message contains tag and a message, wrong results.")
	}
	if tag, message, ok := Decode("\x01CLIENTINFO A B C\x01"); tag != "CLIENTINFO" || message != "A B C" || !ok {
		t.Error("Message contains tag and a message with spaces, wrong results.")
	}
}

func TestEncode(t *testing.T) {
	if text := Encode("", "INVALID"); len(text) > 0 {
		t.Error("Message is invalid, but returns a non-empty string.")
	}
	if text := Encode("VERSION", ""); text != "\x01VERSION\x01" {
		t.Error("Message contains only a tag, wrong results.")
	}
	if text := Encode("PING", "123456789"); text != "\x01PING 123456789\x01" {
		t.Error("Message contains tag and a message, wrong results.")
	}
	if text := Encode("CLIENTINFO", "A B C"); text != "\x01CLIENTINFO A B C\x01" {
		t.Error("Message contains tag and a message with spaces, wrong results.")
	}
}

func TestAction(t *testing.T) {
	if text := Action("A B C"); text != "\x01ACTION A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestPing(t *testing.T) {
	if text := Ping("A B C"); text != "\x01PING A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestPong(t *testing.T) {
	if text := Pong("A B C"); text != "\x01PONG A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestVersion(t *testing.T) {
	if text := Version("A B C"); text != "\x01VERSION A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestUserInfo(t *testing.T) {
	if text := UserInfo("A B C"); text != "\x01USERINFO A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestClientInfo(t *testing.T) {
	if text := ClientInfo("A B C"); text != "\x01CLIENTINFO A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestFinger(t *testing.T) {
	if text := Finger("A B C"); text != "\x01FINGER A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestSource(t *testing.T) {
	if text := Source("A B C"); text != "\x01SOURCE A B C\x01" {
		t.Error("Wrong result!")
	}
}

func TestTime(t *testing.T) {
	if text := Time("A B C"); text != "\x01TIME A B C\x01" {
		t.Error("Wrong result!")
	}
}
