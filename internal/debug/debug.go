package debug

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
)

func Inspect(v interface{}, w io.Writer) {
	s, _ := json.MarshalIndent(v, "", "  ")
	fmt.Fprintf(w, "\n%v -> %v\n", reflect.TypeOf(v), string(s))
}
