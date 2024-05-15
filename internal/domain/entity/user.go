package entity

type User struct {
	Base `bson:"inline"`

	Username     string `json:"username" bson:"username"`
	PasswordHash []byte `json:"-" bson:"passwordHash"`
	FullName     string `json:"fullName" bson:"fullName"`
}

type UserToCreate struct {
	BaseToCreate `bson:"inline"`

	Username     string `json:"username" bson:"username" binding:"required"`
	PasswordHash []byte `json:"-" bson:"passwordHash"`
	Password     string `json:"password" bson:"-" binding:"required"`
	FullName     string `json:"fullName" bson:"fullName" binding:"required"`
}

type UserToUpdate struct {
	UpdatedAt TimeNow `json:"-" bson:"updated_at"`
	Username  string  `json:"username" bson:"username,omitempty" binding:"required_without=FullName"`
	FullName  string  `json:"fullName" bson:"fullName,omitempty" binding:"required_without=Username"`
}
