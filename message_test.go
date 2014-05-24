package irc

import (
	"reflect"
	"testing"
)

type messageTest struct {
	parsed     *Message
	rawMessage string
	rawPrefix  string
	paramLen   int
}

var messageTests = [16]*messageTest{
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "syrk",
				User: "kalt",
				Host: "millennium.stealth.net",
			},
			Command:  "QUIT",
			Trailing: "Gone to have lunch",
		},
		rawMessage: ":syrk!kalt@millennium.stealth.net QUIT :Gone to have lunch\r\n",
		rawPrefix:  "syrk!kalt@millennium.stealth.net",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "Trillian",
			},
			Command:  "SQUIT",
			Params:   []string{"cm22.eng.umd.edu"},
			Trailing: "Server out of control",
		},
		rawMessage: ":Trillian SQUIT cm22.eng.umd.edu :Server out of control\r\n",
		rawPrefix:  "Trillian",
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
		rawMessage: ":WiZ!jto@tolsun.oulu.fi JOIN #Twilight_zone\r\n",
		rawPrefix:  "WiZ!jto@tolsun.oulu.fi",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "WiZ",
				User: "jto",
				Host: "tolsun.oulu.fi",
			},
			Command:  "PART",
			Params:   []string{"#playzone"},
			Trailing: "I lost",
		},
		rawMessage: ":WiZ!jto@tolsun.oulu.fi PART #playzone :I lost\r\n",
		rawPrefix:  "WiZ!jto@tolsun.oulu.fi",
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
		rawMessage: ":WiZ!jto@tolsun.oulu.fi MODE #eu-opers -l\r\n",
		rawPrefix:  "WiZ!jto@tolsun.oulu.fi",
	},
	{
		parsed: &Message{
			Command: "MODE",
			Params:  []string{"&oulu", "+b", "*!*@*.edu", "+e", "*!*@*.bu.edu"},
		},
		rawMessage: "MODE &oulu +b *!*@*.edu +e *!*@*.bu.edu\r\n",
	},
	{
		parsed: &Message{
			Command:  "PRIVMSG",
			Params:   []string{"#channel"},
			Trailing: "Message with :colons!",
		},
		rawMessage: "PRIVMSG #channel :Message with :colons!\r\n",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "irc.vives.lan",
			},
			Command:  "251",
			Params:   []string{"test"},
			Trailing: "There are 2 users and 0 services on 1 servers",
		},
		rawMessage: ":irc.vives.lan 251 test :There are 2 users and 0 services on 1 servers\r\n",
		rawPrefix:  "irc.vives.lan",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "irc.vives.lan",
			},
			Command:  "376",
			Params:   []string{"test"},
			Trailing: "End of MOTD command",
		},
		rawMessage: ":irc.vives.lan 376 test :End of MOTD command\r\n",
		rawPrefix:  "irc.vives.lan",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "irc.vives.lan",
			},
			Command:  "250",
			Params:   []string{"test"},
			Trailing: "Highest connection count: 1 (1 connections received)",
		},
		rawMessage: ":irc.vives.lan 250 test :Highest connection count: 1 (1 connections received)\r\n",
		rawPrefix:  "irc.vives.lan",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "sorcix",
				User: "~sorcix",
				Host: "sorcix.users.quakenet.org",
			},
			Command:  "PRIVMSG",
			Params:   []string{"#viveslan"},
			Trailing: "\001ACTION is testing CTCP messages!\001",
		},
		rawMessage: ":sorcix!~sorcix@sorcix.users.quakenet.org PRIVMSG #viveslan :\001ACTION is testing CTCP messages!\001\r\n",
		rawPrefix:  "sorcix!~sorcix@sorcix.users.quakenet.org",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "sorcix",
				User: "~sorcix",
				Host: "sorcix.users.quakenet.org",
			},
			Command:  "NOTICE",
			Params:   []string{"midnightfox"},
			Trailing: "\001PONG 1234567890\001",
		},
		rawMessage: ":sorcix!~sorcix@sorcix.users.quakenet.org NOTICE midnightfox :\001PONG 1234567890\001\r\n",
		rawPrefix:  "sorcix!~sorcix@sorcix.users.quakenet.org",
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
		rawMessage: ":a!b@c QUIT\r\n",
		rawPrefix:  "a!b@c",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "a",
				User: "b",
			},
			Command:  "PRIVMSG",
			Trailing: "message",
		},
		rawMessage: ":a!b PRIVMSG :message\r\n",
		rawPrefix:  "a!b",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "a",
				Host: "c",
			},
			Command:  "NOTICE",
			Trailing: ":::Hey!",
		},
		rawMessage: ":a@c NOTICE ::::Hey!\r\n",
		rawPrefix:  "a@c",
	},
	{
		parsed: &Message{
			Prefix: &Prefix{
				Name: "nick",
			},
			Command:  "PRIVMSG",
			Params:   []string{"$@"},
			Trailing: "This message contains a\ttab!",
		},
		rawMessage: ":nick PRIVMSG $@ :This message contains a\ttab!\r\n",
		rawPrefix:  "nick",
	},
}

// -----
// PREFIX
// -----

func TestPrefix_String(t *testing.T) {
	var s string

	for i, test := range messageTests {

		// Skip tests that have no prefix
		if test.parsed.Prefix == nil {
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
		if test.parsed.Prefix == nil {
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
		if test.parsed.Prefix == nil {
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
			t.Errorf("Failed to parse prefix %d:", i)
			t.Logf("Output: %#v", p)
			t.Logf("Expected: %#v", test.parsed)
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
