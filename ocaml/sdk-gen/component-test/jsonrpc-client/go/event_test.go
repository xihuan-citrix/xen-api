package componenttest

import (
	"reflect"
	"testing"
	"time"
	"xenapi"
)

func TestEventFrom(t *testing.T) {
	session, err := GetSession("xapi-24/event_from_11")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	var classes = []string{
		"vm",
	}
	token := "token0"
	var timeout float64 = 600.0
	var snapshot xenapi.RecordInterface = map[string]interface{}{
		"NameLabel": "snapshot0",
	}
	var eventRecord0 = xenapi.EventRecord{
		Snapshot:  snapshot,
		ID:        10,
		Timestamp: time.Date(2024, time.January, 1, 13, 20, 11, 0, time.UTC),
		Class:     "vm",
		Operation: "add",
		Ref:       "OpaqueRef:5event01-7b62-f571-d38b-754341a973e5",
		ObjUUID:   "cevnet08-3e97-53bb-4f97-bd039ca8c217",
	}
	var expectedResult = xenapi.EventBatch{
		Token: "token0",
		ValidRefCounts: map[string]int{
			"OpaqueRef:5event01-7b62-f571-d38b-754341a973e5": 6,
		},
		Events: []xenapi.EventRecord{
			eventRecord0,
		},
	}
	result, err := xenapi.Event.From(session, classes, token, timeout)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	if reflect.DeepEqual(result, expectedResult) {
		t.Fail()
		return
	}
}
