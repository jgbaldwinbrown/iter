package iter

import (
	"testing"
	"strings"
	"encoding/json"
	"reflect"
)

const jsonExample = `1
"banana"
{"2": "earlobe"}
[{"hat": "cat"}, 3, 4]
`

var jsonExpect = []any {
	1,
	"banana",
	map[string]any{"2": "earlobe"},
	[]any{map[string]any{"hat": "cat"}, 3, 4},
}

func TestJsonDecoder(t *testing.T) {
	out, e := Collect[any](DecoderIter[any](json.NewDecoder(strings.NewReader(jsonExample))))
	if e != nil {
		panic(e)
	}
	if !reflect.DeepEqual(out, jsonExpect) {
		t.Errorf("out\n%#v\n!=\njsonExpect\n%#v", out, jsonExpect)
	}
}
