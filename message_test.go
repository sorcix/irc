// Copyright 2014 Vic Demuzere
//
// Use of this source code is governed by the MIT license.

package irc

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleParseMessage() {
	message := ParseMessage("JOIN #help")

	fmt.Println(message.Params[0])

	// Output: #help
}

func ExampleMessage_String() {
	message := &Message{
		Prefix: &Prefix{
			Name: "sorcix",
			User: "sorcix",
			Host: "myhostname",
		},
		Command: "PRIVMSG",
		Params:  []string{"This is an example!"},
	}

	fmt.Println(message.String())

	// Output: :sorcix!sorcix@myhostname PRIVMSG :This is an example!
}

var messageTests = [...]*struct {
	parsed     *Message
	rawMessage string
	rawPrefix  string
	hostmask   bool // Is it very clear that the prefix is a hostname?
	server     bool // Is the prefix a servername?
}{
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "syrk",
				User: "kalt",
				Host: "millennium.stealth.net",
			},
			Command: "QUIT",
			Params:  []string{"Gone to have lunch"},
		},
		rawMessage: ":syrk!kalt@millennium.stealth.net QUIT :Gone to have lunch",
		rawPrefix:  "syrk!kalt@millennium.stealth.net",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "Trillian",
			},
			Command: "SQUIT",
			Params:  []string{"cm22.eng.umd.edu", "Server out of control"},
		},
		rawMessage: ":Trillian SQUIT cm22.eng.umd.edu :Server out of control",
		rawPrefix:  "Trillian",
		server:     true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "WiZ",
				User: "jto",
				Host: "tolsun.oulu.fi",
			},
			Command: "JOIN",
			Params:  []string{"#Twilight_zone"},
		},
		rawMessage: ":WiZ!jto@tolsun.oulu.fi JOIN #Twilight_zone",
		rawPrefix:  "WiZ!jto@tolsun.oulu.fi",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "WiZ",
				User: "jto",
				Host: "tolsun.oulu.fi",
			},
			Command: "PART",
			Params:  []string{"#playzone", "I lost"},
		},
		rawMessage: ":WiZ!jto@tolsun.oulu.fi PART #playzone :I lost",
		rawPrefix:  "WiZ!jto@tolsun.oulu.fi",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "WiZ",
				User: "jto",
				Host: "tolsun.oulu.fi",
			},
			Command: "MODE",
			Params:  []string{"#eu-opers", "-l"},
		},
		rawMessage: ":WiZ!jto@tolsun.oulu.fi MODE #eu-opers -l",
		rawPrefix:  "WiZ!jto@tolsun.oulu.fi",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Command: "MODE",
			Params:  []string{"&oulu", "+b", "*!*@*.edu", "+e", "*!*@*.bu.edu"},
		},
		rawMessage: "MODE &oulu +b *!*@*.edu +e *!*@*.bu.edu",
	},
	{
		parsed: &Message{
			Command: "PRIVMSG",
			Params:  []string{"#channel", "Message with :colons!"},
		},
		rawMessage: "PRIVMSG #channel :Message with :colons!",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "irc.vives.lan",
			},
			Command: "251",
			Params:  []string{"test", "There are 2 users and 0 services on 1 servers"},
		},
		rawMessage: ":irc.vives.lan 251 test :There are 2 users and 0 services on 1 servers",
		rawPrefix:  "irc.vives.lan",
		server:     true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "irc.vives.lan",
			},
			Command: "376",
			Params:  []string{"test", "End of MOTD command"},
		},
		rawMessage: ":irc.vives.lan 376 test :End of MOTD command",
		rawPrefix:  "irc.vives.lan",
		server:     true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "irc.vives.lan",
			},
			Command: "250",
			Params:  []string{"test", "Highest connection count: 1 (1 connections received)"},
		},
		rawMessage: ":irc.vives.lan 250 test :Highest connection count: 1 (1 connections received)",
		rawPrefix:  "irc.vives.lan",
		server:     true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "sorcix",
				User: "~sorcix",
				Host: "sorcix.users.quakenet.org",
			},
			Command: "PRIVMSG",
			Params:  []string{"#viveslan", "\001ACTION is testing CTCP messages!\001"},
		},
		rawMessage: ":sorcix!~sorcix@sorcix.users.quakenet.org PRIVMSG #viveslan :\001ACTION is testing CTCP messages!\001",
		rawPrefix:  "sorcix!~sorcix@sorcix.users.quakenet.org",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "sorcix",
				User: "~sorcix",
				Host: "sorcix.users.quakenet.org",
			},
			Command: "NOTICE",
			Params:  []string{"midnightfox", "\001PONG 1234567890\001"},
		},
		rawMessage: ":sorcix!~sorcix@sorcix.users.quakenet.org NOTICE midnightfox :\001PONG 1234567890\001",
		rawPrefix:  "sorcix!~sorcix@sorcix.users.quakenet.org",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "a",
				User: "b",
				Host: "c",
			},
			Command: "QUIT",
		},
		rawMessage: ":a!b@c QUIT",
		rawPrefix:  "a!b@c",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "a",
				User: "b",
			},
			Command: "PRIVMSG",
			Params:  []string{"message"},
		},
		rawMessage: ":a!b PRIVMSG message",
		rawPrefix:  "a!b",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "a",
				Host: "c",
			},
			Command: "NOTICE",
			Params:  []string{":::Hey!"},
		},
		rawMessage: ":a@c NOTICE ::::Hey!",
		rawPrefix:  "a@c",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "nick",
			},
			Command: "PRIVMSG",
			Params:  []string{"$@", "This message contains a\ttab!"},
		},
		rawMessage: ":nick PRIVMSG $@ :This message contains a\ttab!",
		rawPrefix:  "nick",
	},
	{
		parsed: &Message{
			Command: "TEST",
			Params:  []string{"$@", "", "param", "Trailing"},
		},
		rawMessage: "TEST $@  param Trailing",
	},
	{
		rawMessage: ": PRIVMSG test :Invalid message with empty prefix.",
		rawPrefix:  "",
	},
	{
		rawMessage: ":  PRIVMSG test :Invalid message with space prefix",
		rawPrefix:  " ",
	},
	{
		parsed: &Message{
			Command: "TOPIC",
			Params:  []string{"#foo", ""},
		},
		rawMessage: "TOPIC #foo :",
		rawPrefix:  "",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "name",
				User: "user",
				Host: "example.org",
			},
			Command: "PRIVMSG",
			Params:  []string{"#test", "Message with spaces at the end!  "},
		},
		rawMessage: ":name!user@example.org PRIVMSG #test :Message with spaces at the end!  ",
		rawPrefix:  "name!user@example.org",
		hostmask:   true,
	},
	{
		parsed: &Message{
			Command: "PASS",
			Params:  []string{"oauth:token_goes_here"},
		},
		rawMessage: "PASS oauth:token_goes_here",
		rawPrefix:  "",
	},
	{
		parsed: &Message{
			Command: "PRIVMSG",
			Params:  []string{"#some:channel", "http://example.com"},
		},
		rawMessage: "PRIVMSG #some:channel http://example.com",
		rawPrefix:  "",
	},
}

// -----
// PREFIX
// -----

func TestPrefix_IsHostmask(t *testing.T) {

	for i, test := range messageTests {

		// Skip tests that have no prefix
		if test.parsed == nil || test.parsed.Prefix == nil {
			continue
		}

		if test.hostmask && !test.parsed.Prefix.IsHostmask() {
			t.Errorf("Prefix %d should be recognized as a hostmask!", i)
		}

	}
}

func TestPrefix_IsServer(t *testing.T) {

	for i, test := range messageTests {

		// Skip tests that have no prefix
		if test.parsed == nil || test.parsed.Prefix == nil {
			continue
		}

		if test.server && !test.parsed.Prefix.IsServer() {
			t.Errorf("Prefix %d should be recognized as a server!", i)
		}

	}
}

func TestPrefix_String(t *testing.T) {
	var s string

	for i, test := range messageTests {

		// Skip tests that have no prefix
		if test.parsed == nil || test.parsed.Prefix == nil {
			continue
		}

		// Convert the prefix
		s = test.parsed.Prefix.String()

		// Result should be the same as the value in rawMessage.
		if s != test.rawPrefix {
			t.Errorf("Failed to stringify prefix %d:", i)
			t.Logf("Output: %s", s)
			t.Logf("Expected: %s", test.rawPrefix)
		}
	}
}

func TestPrefix_Len(t *testing.T) {
	var l int

	for i, test := range messageTests {

		// Skip tests that have no prefix
		if test.parsed == nil || test.parsed.Prefix == nil {
			continue
		}

		l = test.parsed.Prefix.Len()

		// Result should be the same as the value in rawMessage.
		if l != len(test.rawPrefix) {
			t.Errorf("Failed to calculate prefix length %d:", i)
			t.Logf("Output: %d", l)
			t.Logf("Expected: %d", len(test.rawPrefix))
		}
	}
}

func TestParsePrefix(t *testing.T) {
	var p *Prefix

	for i, test := range messageTests {

		// Skip tests that have no prefix
		if test.parsed == nil || test.parsed.Prefix == nil {
			continue
		}

		// Parse the prefix
		p = ParsePrefix(test.rawPrefix)

		// Result struct should be the same as the value in parsed.
		if *p != *test.parsed.Prefix {
			t.Errorf("Failed to parse prefix %d:", i)
			t.Logf("Output: %#v", p)
			t.Logf("Expected: %#v", test.parsed.Prefix)
		}
	}
}

// -----
// MESSAGE
// -----

func TestMessage_String(t *testing.T) {
	var s string

	for i, test := range messageTests {

		// Skip tests that have no valid struct
		if test.parsed == nil {
			continue
		}

		// Convert the prefix
		s = test.parsed.String()

		// Result should be the same as the value in rawMessage.
		if s != test.rawMessage {
			t.Errorf("Failed to stringify message %d:", i)
			t.Logf("Output: %s", s)
			t.Logf("Expected: %s", test.rawMessage)
		}
	}
}

func TestMessage_Len(t *testing.T) {
	var l int

	for i, test := range messageTests {

		// Skip tests that have no valid struct
		if test.parsed == nil {
			continue
		}

		l = test.parsed.Len()

		// Result should be the same as the value in rawMessage.
		if l != len(test.rawMessage) {
			t.Errorf("Failed to calculate message length %d:", i)
			t.Logf("Output: %d", l)
			t.Logf("Expected: %d", len(test.rawMessage))
		}
	}
}

func TestParseMessage(t *testing.T) {
	var p *Message

	for i, test := range messageTests {

		// Parse the prefix
		p = ParseMessage(test.rawMessage)

		// Result struct should be the same as the value in parsed.
		if !reflect.DeepEqual(p, test.parsed) {
			t.Errorf("Failed to parse message %d:", i)
			t.Logf("Output: %#v", p)
			t.Logf("Expected: %#v", test.parsed)
		}
	}
}

// -----
// MESSAGE DECODE -> ENCODE
// -----

func TestMessageDecodeEncode(t *testing.T) {
	var (
		p *Message
		s string
	)

	for i, test := range messageTests {

		// Skip invalid messages
		if test.parsed == nil {
			continue
		}

		// Decode the message, then encode it again.
		p = ParseMessage(test.rawMessage)
		s = p.String()

		// Result struct should be the same as the original.
		if s != test.rawMessage {
			t.Errorf("Message %d failed decode-encode sequence!", i)
		}
	}
}

// -----
// BENCHMARK
// -----

func BenchmarkPrefix_String_short(b *testing.B) {
	b.ReportAllocs()

	prefix := new(Prefix)
	prefix.Name = "Namename"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefix.String()
	}
}
func BenchmarkPrefix_String_long(b *testing.B) {
	b.ReportAllocs()

	prefix := new(Prefix)
	prefix.Name = "Namename"
	prefix.User = "Username"
	prefix.Host = "Hostname"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefix.String()
	}
}
func BenchmarkParsePrefix_short(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ParsePrefix("Namename")
	}
}
func BenchmarkParsePrefix_long(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ParsePrefix("Namename!Username@Hostname")
	}
}
func BenchmarkMessage_String(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		messageTests[0].parsed.String()
	}
}
func BenchmarkParseMessage_short(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ParseMessage("COMMAND arg1 :Message\r\n")
	}
}
func BenchmarkParseMessage_medium(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ParseMessage(":Namename COMMAND arg6 arg7 :Message message message\r\n")
	}
}
func BenchmarkParseMessage_long(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ParseMessage(":Namename!username@hostname COMMAND arg1 arg2 arg3 arg4 arg5 arg6 arg7 :Message message message message message\r\n")
	}
}
