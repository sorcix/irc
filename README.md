# Go **irc** package

[Package Documentation][Documentation] @ godoc.org

[![Build Status](https://drone.io/github.com/sorcix/irc/status.png)](https://drone.io/github.com/sorcix/irc/latest)

## Features
Package irc allows your application to speak the IRC protocol.

 - **Limited scope**, does one thing and does it well.
 - Focus on simplicity and **speed**.
 - **Stable API**: updates shouldn't break existing software.
 - Well [documented][Documentation] code.

*This package does not manage your entire IRC connection. It only translates the protocol to easy to use Go types. It is meant as a single component in a larger IRC library, or for basic IRC bots for which a large IRC package would be overkill.*

### Message
The [Message][] and [Prefix][] types provide translation to and from IRC message format.

    // Parse the IRC-encoded data and stores the result in a new struct.
    message := irc.ParseMessage(raw)

    // Returns the IRC encoding of the message.
    raw = message.String()

### Encoder & Decoder
The [Encoder][] and [Decoder][] types allow working with IRC message streams.

    // Create a decoder that reads from given io.Reader
    dec := irc.NewDecoder(reader)

    // Decode the next IRC message
    message, err := dec.Decode()

    // Create an encoder that writes to given io.Writer
    enc := irc.NewEncoder(writer)

    // Send a message to the writer.
    enc.Encode(message)

### Conn
The [Conn][] type combines an [Encoder][] and [Decoder][] for a duplex connection.

    c, err := irc.Dial("irc.server.net:6667")

    // Methods from both Encoder and Decoder are available
    message, err := c.Decode()

## Future plans

 - Basic event-based client, in a separate package.
 - Support for IRCv3 message tags.
 - Example code


[Documentation]: https://godoc.org/github.com/sorcix/irc "Package documentation by Godoc.org"
[Message]: http://godoc.org/github.com/sorcix/irc#Message "Message type documentation"
[Prefix]: http://godoc.org/github.com/sorcix/irc#Prefix "Prefix type documentation"
[Encoder]: http://godoc.org/github.com/sorcix/irc#Encoder "Encoder type documentation"
[Decoder]: http://godoc.org/github.com/sorcix/irc#Decoder "Decoder type documentation"
[Conn]: http://godoc.org/github.com/sorcix/irc#Conn "Conn type documentation"
[RFC1459]: http://tools.ietf.org/html/rfc1459.html "RFC 1459"
