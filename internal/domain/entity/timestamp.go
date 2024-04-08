package entity

import (
	"encoding/json"
	"strconv"
	"time"
	"unsafe"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// Timestamp is a custom type that support both time.Time and Unix timestamp.
type Timestamp time.Time

// MarshalJSON returns the JSON encoding of Timestamp.
func (t Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

// UnmarshalJSON parses the JSON-encoded data and stores the result in the Timestamp receiver.
func (t *Timestamp) UnmarshalJSON(b []byte) error {
	p := (*time.Time)(unsafe.Pointer(t))

	timestamp, err := strconv.ParseInt(string(b), 10, 64)
	if err != nil {
		return json.Unmarshal(b, p)
	}

	*p = time.Unix(timestamp, 0)

	return nil
}

// MarshalBSONValue returns the BSON encoding of Timestamp.
func (t Timestamp) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(time.Time(t))
}

// UnmarshalBSONValue parses the BSON-encoded data and stores the result in the receiver.
func (t *Timestamp) UnmarshalBSONValue(btype bsontype.Type, data []byte) error {
	return bson.UnmarshalValue(btype, data, (*time.Time)(unsafe.Pointer(t)))
}
