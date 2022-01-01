// Package fixedwidth provides encoding and decoding for fixed-width formatted Data.
package fixedwidth

type TextMarshalerFixedWidth interface {
	MarshalTextFixedWidth() (text []byte, err error)
}

type TextUnmarshalerFixedWidth interface {
	UnmarshalTextFixedWidth(text []byte) error
}
