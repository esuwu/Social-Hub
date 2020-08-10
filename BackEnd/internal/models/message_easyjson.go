// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson4086215fDecodeMainInternalModels(in *jlexer.Lexer, out *Message) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "chatId":
			out.ChatId = string(in.String())
		case "chatPhoto":
			out.ChatPhoto = string(in.String())
		case "chatName":
			out.ChatName = string(in.String())
		case "authorName":
			out.AuthorName = string(in.String())
		case "authorSurname":
			out.AuthorSurname = string(in.String())
		case "authorUrl":
			out.AuthorUrl = string(in.String())
		case "authorPhoto":
			out.AuthorPhoto = string(in.String())
		case "text":
			out.Text = string(in.String())
		case "time":
			out.Time = string(in.String())
		case "sticker":
			out.Sticker = string(in.String())
		case "isMe":
			out.IsMe = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson4086215fEncodeMainInternalModels(out *jwriter.Writer, in Message) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"chatId\":"
		out.RawString(prefix[1:])
		out.String(string(in.ChatId))
	}
	{
		const prefix string = ",\"chatPhoto\":"
		out.RawString(prefix)
		out.String(string(in.ChatPhoto))
	}
	{
		const prefix string = ",\"chatName\":"
		out.RawString(prefix)
		out.String(string(in.ChatName))
	}
	{
		const prefix string = ",\"authorName\":"
		out.RawString(prefix)
		out.String(string(in.AuthorName))
	}
	{
		const prefix string = ",\"authorSurname\":"
		out.RawString(prefix)
		out.String(string(in.AuthorSurname))
	}
	{
		const prefix string = ",\"authorUrl\":"
		out.RawString(prefix)
		out.String(string(in.AuthorUrl))
	}
	{
		const prefix string = ",\"authorPhoto\":"
		out.RawString(prefix)
		out.String(string(in.AuthorPhoto))
	}
	{
		const prefix string = ",\"text\":"
		out.RawString(prefix)
		out.String(string(in.Text))
	}
	{
		const prefix string = ",\"time\":"
		out.RawString(prefix)
		out.String(string(in.Time))
	}
	if in.Sticker != "" {
		const prefix string = ",\"sticker\":"
		out.RawString(prefix)
		out.String(string(in.Sticker))
	}
	{
		const prefix string = ",\"isMe\":"
		out.RawString(prefix)
		out.Bool(bool(in.IsMe))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Message) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson4086215fEncodeMainInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Message) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson4086215fEncodeMainInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Message) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson4086215fDecodeMainInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Message) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson4086215fDecodeMainInternalModels(l, v)
}
