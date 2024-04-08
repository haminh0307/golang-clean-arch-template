package entity

import (
	"errors"
	"time"
	"unsafe"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ErrInvalidID = errors.New("invalid id")

// Base define base fields for all struct entity.
type Base struct {
	ID        ID        `json:"id" bson:"_id"`
	CreatedAt Timestamp `json:"createdAt" bson:"createdAt"`
	UpdatedAt Timestamp `json:"updatedAt" bson:"updatedAt"`
	IsDeleted bool      `json:"-" bson:"isDeleted"`
}

// BaseToCreate define base entity for input purpose.
//
// Note the json tag "-" to ignore those fields during JSON encoding and decoding.
type BaseToCreate struct {
	ID        ID      `json:"-" bson:"_id,omitempty"`
	CreatedAt TimeNow `json:"-" bson:"createdAt"`
	UpdatedAt TimeNow `json:"-" bson:"updatedAt"`
	IsDeleted bool    `json:"-" bson:"isDeleted"`
}

// ID is a custom ID that is string but support read from/write to MongoDB.
type ID string

const NilID = ID("")

func (id ID) String() string {
	return string(id)
}

func (id *ID) UnmarshalBSONValue(btype bsontype.Type, data []byte) error {
	return bson.UnmarshalValue(btype, data, (*string)(unsafe.Pointer(id)))
}

func (id ID) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if len(string(id)) == 0 {
		return bson.MarshalValue(primitive.NilObjectID)
	}

	oid, err := primitive.ObjectIDFromHex(string(id))
	switch {
	case err == nil:
		return bson.MarshalValue(oid)
	case errors.Is(err, primitive.ErrInvalidHex):
		return bson.TypeNull, nil, ErrInvalidID
	default:
		return bson.TypeNull, nil, err
	}
}

// TimeNow
type TimeNow struct {
	Timestamp
}

func (t TimeNow) MarshalBSONValue() (bsontype.Type, []byte, error) {
	return bson.MarshalValue(time.Now())
}
