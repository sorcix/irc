package irc

import (
	"testing"
)

// MessageTest
type messageTest struct {
	parsed     *Message
	rawMessage string
	rawPrefix  string
	paramLen   int
}

var messageTests = [10]*messageTest{
	&messageTest{
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
	&messageTest{
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
	&messageTest{
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
	&messageTest{
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
	&messageTest{
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
	&messageTest{
		parsed: &Message{
			Command: "MODE",
			Params:  []string{"&oulu", "+b", "*!*@*.edu", "+e", "*!*@*.bu.edu"},
		},
		rawMessage: "MODE &oulu +b *!*@*.edu +e *!*@*.bu.edu\r\n",
	},
	&messageTest{
		parsed: &Message{
			Command:  "PRIVMSG",
			Params:   []string{"#channel"},
			Trailing: "Message with :colons!",
		},
		rawMessage: "PRIVMSG #channel :Message with :colons!\r\n",
	},
	&messageTest{
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
	&messageTest{
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
	&messageTest{
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
}

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
		if !p.equals(test.parsed) {
			t.Errorf("Failed to parse prefix %d:", i)
			t.Logf("Output: %#v", p)
			t.Logf("Expected: %#v", test.parsed)
		}
	}
}

func BenchmarkPrefix_String_short(b *testing.B) {
	prefix := new(Prefix)
	prefix.Name = "Namename"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		prefix.String()
	}
}
func BenchmarkPrefix_String_long(b *testing.B) {
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
	for i := 0; i < b.N; i++ {
		ParsePrefix("Namename")
	}
}
func BenchmarkParsePrefix_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParsePrefix("Namename!Username@Hostname")
	}
}
func BenchmarkMessage_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		messageTests[0].parsed.String()
	}
}
func BenchmarkParseMessage_short(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage("COMMAND arg1 :Message\r\n")
	}
}
func BenchmarkParseMessage_medium(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(":Namename COMMAND arg6 arg7 :Message message message\r\n")
	}
}
func BenchmarkParseMessage_long(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseMessage(":Namename!username@hostname COMMAND arg1 arg2 arg3 arg4 arg5 arg6 arg7 :Message message message message message\r\n")
	}
}

func (m *Message) equals(o *Message) bool {

	if (m.Prefix != nil && *m.Prefix != *o.Prefix) || m.Trailing != o.Trailing || m.Command != o.Command || len(m.Params) != len(o.Params) {
		return false
	}

	for i, param := range m.Params {
		if param != o.Params[i] {
			return false
		}
	}

	return true
}
