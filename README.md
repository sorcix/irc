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

## Examples
Check these other projects for an example on how to use the package:

Clients:

 - https://github.com/nickvanw/ircx (great simple example)
 - https://github.com/FSX/jun
 - https://github.com/jnwhiteh/wallops
 - https://github.com/Alligator/gomero
 - https://github.com/msparks/iq
 - https://github.com/TheCreeper/HackBot

Servers:

 - https://github.com/nightexcessive/excessiveircd


[Documentation]: https://godoc.org/github.com/sorcix/irc "Package documentation by Godoc.org"
[Message]: https://godoc.org/github.com/sorcix/irc#Message "Message type documentation"
[Prefix]: https://godoc.org/github.com/sorcix/irc#Prefix "Prefix type documentation"
[Encoder]: https://godoc.org/github.com/sorcix/irc#Encoder "Encoder type documentation"
[Decoder]: https://godoc.org/github.com/sorcix/irc#Decoder "Decoder type documentation"
[Conn]: https://godoc.org/github.com/sorcix/irc#Conn "Conn type documentation"
[RFC1459]: https://tools.ietf.org/html/rfc1459.html "RFC 1459"
