Go IRC library
===============

[![Build Status](https://drone.io/github.com/sorcix/irc/status.png)](https://drone.io/github.com/sorcix/irc/latest)

[Message]: http://godoc.org/github.com/sorcix/irc#Message "Message type documentation"
[Encoder]: http://godoc.org/github.com/sorcix/irc#Encoder "Encoder type documentation"
[Decoder]: http://godoc.org/github.com/sorcix/irc#Decoder "Decoder type documentation"
[Prefix]: http://godoc.org/github.com/sorcix/irc#Prefix "Prefix type documentation"
[RFC1459]: http://tools.ietf.org/html/rfc1459.html "RFC 1459"
[JSON]: http://golang.org/pkg/encoding/json/ "Package encoding/json in the standard library"

This minimalistic Go IRC library provides helpers for IRC connectivity.

**This library does not abstract away the IRC protocol from your application, knowledge of the IRC protocol is required!** Most of the time, I need to implement only parts of the IRC protocol, and most other IRC libraries are overkill. The aim of this project is to provide a package similar to the official [net/http](http://golang.org/pkg/net/http/ "HTTP") package, providing you with the basics *required for every IRC client and server*.

If you want something up and running very fast, this isn't for you. If you prefer minimal code and fast execution times, **keep reading**!

[Documentation by Godoc.org](http://godoc.org/github.com/sorcix/irc "Documentation")

    go get github.com/sorcix/irc

Message
--------
The [Message][] and [Prefix][] types provide translation to and from raw IRC messages, as described in [RFC1459][] (section 2.3.1).

    // Parse the IRC-encoded data and stores the result in a new struct.
    message := irc.ParseMessage(raw)

    // Returns the IRC encoding of the message.
    raw = message.String()

Encoder & Decoder
-----------------
Similar to the types found in [encoding/json][JSON], the [Encoder][] type translates a [Message][] and writes it to a stream. The [Decoder][] type allows you to read [Message][]s from a stream, one by one.

    // Create a decoder that reads from given io.Reader
    dec := irc.NewDecoder(reader)

    // Decode the next IRC message
    message, err := dec.Decode()

    // Create an encoder that writes to given io.Writer
    enc := irc.NewEncoder(writer)

    // Send a message to the writer.
    enc.Encode(message)
