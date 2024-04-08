package entity

type User struct {
	Base `bson:"inline"`

	Username     string `json:"username" bson:"username"`
	PasswordHash []byte `json:"-" bson:"passwordHash"`
	FullName     string `json:"fullName" bson:"fullName"`
}

type UserToCreate struct {
	BaseToCreate `bson:"inline"`

	Username     string `json:"username" bson:"username"`
	PasswordHash []byte `json:"-" bson:"passwordHash"`
	Password     string `json:"password" bson:"-"`
	FullName     string `json:"fullName" bson:"fullName"`
}

type UserToUpdate struct {
	UpdatedAt TimeNow `json:"-" bson:"updated_at"`
	Username  string  `json:"username" bson:"username,omitempty"`
	FullName  string  `json:"fullName" bson:"fullName,omitempty"`
}
