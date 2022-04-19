package log

import (
	"encoding/json"

	"github.com/mailru/easyjson/jwriter"
)

var (
	_ *jwriter.Writer
)

func easyEncodeJSON(out *jwriter.Writer, in *log) {
	out.RawByte('{')
	{
		if in.error != "" {
			out.RawString("\"error\":")
			out.String(in.error)

			if in.message != "" {
				out.RawString(",\"message\":")
				out.String(in.message)
			}
		} else {
			out.RawString("\"message\":")
			out.String(in.message)
		}
	}
	{
		if len(in.callers) > 0 {
			out.RawString(",\"callers\":")
			out.RawByte('[')
			for index, caller := range in.callers {
				if index > 0 {
					out.RawByte(',')
				}
				out.String(caller)
			}
			out.RawByte(']')
		}
	}
	{
		out.RawString(",\"time\":")
		out.String(in.time)
	}
	{
		if len(in.data) > 0 {
			out.RawString(",\"data\":")
			out.RawByte('{')
			first := true
			for key, value := range in.data {
				if first {
					first = false
				} else {
					out.RawByte(',')
				}
				out.String(key)
				out.RawByte(':')
				out.Raw(json.Marshal(value))
			}
			out.RawByte('}')
		}
	}
	out.RawByte('}')
	out.RawByte('\n')
}

func (lg *log) toJSON() []byte {
	w := &jwriter.Writer{}
	easyEncodeJSON(w, lg)
	return w.Buffer.BuildBytes()
}
