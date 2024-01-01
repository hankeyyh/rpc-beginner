package proto

import (
	"encoding/gob"
	"io"
)

type Content struct {
	Msg string
	Seq int
}

func NewContent(msg string, seq int) *Content {
	return &Content{Msg: msg, Seq: seq}
}

type ContentGobCodec struct {
	enc *gob.Encoder
	dec *gob.Decoder
}

func NewContentGobCodec(rwc io.ReadWriteCloser) *ContentGobCodec {
	return &ContentGobCodec{
		enc: gob.NewEncoder(rwc),
		dec: gob.NewDecoder(rwc),
	}
}

func (codec *ContentGobCodec) Encode(v *Content) error {
	return codec.enc.Encode(v)
}

func (codec *ContentGobCodec) Decode(v *Content) error {
	return codec.dec.Decode(v)
}