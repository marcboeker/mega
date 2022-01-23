package uuidgql

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTimestamp(ts time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		_, _ = io.WriteString(w, strconv.Quote(ts.Format(time.RFC3339)))
	})
}

func UnmarshalTimestamp(v interface{}) (ts time.Time, err error) {
	s, ok := v.(string)
	if !ok {
		return time.Time{}, fmt.Errorf("invalid type %T, expect string", v)
	}
	return time.Parse(s, time.RFC3339)
}
