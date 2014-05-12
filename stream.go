package irc

import (
	"bufio"
	"io"
	"sync"
)

const delim byte = '\n'

// A decoder reads Message objects from an input stream.
type Decoder struct {
	reader *bufio.Reader
	line   string
	err    error
	mu     sync.Mutex
}

// NewDecoder returns a new Decoder that reads from r.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		reader: bufio.NewReader(r),
	}
}

// Decode attempts to read a single Message from the stream.
//
// Returns a non-nil error if the read failed.
func (dec *Decoder) Decode() (*Message, error) {

	dec.mu.Lock()
	defer dec.mu.Unlock()

	dec.line, dec.err = dec.reader.ReadString(delim)
	if dec.err != nil {
		return nil, dec.err
	}

	return ParseMessage(dec.line), nil
}

// An encoder writes Message objects to an output stream.
type Encoder struct {
	writer io.Writer
	err    error
	mu     sync.Mutex
}

// NewEncoder returns a new Encoder that writes to w.
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		writer: w,
	}
}

// Encode writes the IRC encoding of m to the stream.
//
// This method may be used from multiple goroutines.
//
// Returns an non-nil error if the write to the underlying stream stopped early.
func (enc *Encoder) Encode(m *Message) error {

	enc.mu.Lock()
	defer enc.mu.Unlock()

	if _, enc.err = enc.writer.Write(m.Bytes()); enc.err != nil {
		return enc.err
	}
	return nil
}
