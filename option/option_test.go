package option

import (
	"encoding/json"
	"testing"
)

type testStruct struct {
	MyInt Option[int] `json:"MyInt,omitzero"`
}

func TestOptionJSON(t *testing.T) {

	s := testStruct{MyInt: Some(5)}
	testOptionJSONMarshalUnmarshal(t, s)

	s = testStruct{MyInt: None[int]()}
	testOptionJSONMarshalUnmarshal(t, s)

	s = testStruct{}
	if err := json.Unmarshal([]byte("{}"), &s); err != nil {
		t.Errorf("Unmarshal: %s", err)
	}

	z, ok := s.MyInt.Val()
	if !ok {
		t.Logf("Unmarshal: None")
	} else {
		t.Logf("Unmarshal: Some(%d)", z)
	}
}

func testOptionJSONMarshalUnmarshal(t *testing.T, s testStruct) {
	b, err := json.Marshal(s)
	if err != nil {
		t.Errorf("Marshal: %s", err)
	}

	t.Logf("Marshal: %s", string(b))

	var y testStruct
	if err := json.Unmarshal(b, &y); err != nil {
		t.Errorf("Unmarshal: %s", err)
	}

	z, ok := y.MyInt.Val()
	if !ok {
		t.Logf("Unmarshal: None")
	} else {
		t.Logf("Unmarshal: Some(%d)", z)
	}
}
