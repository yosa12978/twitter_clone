package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseModel struct {
	Id primitive.ObjectID `bson:"_id" json:"id"`
}

type User struct {
	BaseModel    `bson:",inline" json:",inline"`
	Username     string `json:"username" bson:"username"`
	PasswordHash string `json:"-" bson:"passwordHash"`
	Salt         string `json:"-" bson:"salt"`
	Email        string `json:"email" bson:"email"`
	Icon         string `json:"icon" bson:"icon"`
}

func (usr *User) ToDto() UserDto {
	return UserDto{
		Id:       usr.BaseModel.Id.Hex(),
		Username: usr.Username,
		Icon:     usr.Icon,
	}
}

type UserDto struct {
	Id       string `json:"id"`
	Username string `json:"username" bson:"username"`
	Icon     string `json:"icon" bson:"icon"`
}

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupDto struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func FromDto(dto SignupDto) User {
	return User{
		Username:     dto.Username,
		Email:        dto.Email,
		PasswordHash: dto.Password, // todo add hashing!!!
	}
}

type ArrRes struct {
	Result interface{} `json:"result"`
}

func NewArrRes(v interface{}) ArrRes {
	return ArrRes{Result: v}
}

type PageRes struct {
	Page      int         `json:"page"`
	PageCount int         `json:"pageCount"`
	ItemCount int         `json:"itemCount"`
	HasNext   bool        `json:"hasNext"`
	HasPrev   bool        `json:"hasPrev"`
	Links     []string    `json:"links"`
	Result    interface{} `json:"result"`
}
