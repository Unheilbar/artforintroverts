package entities

type User struct {
	ID       string `json:"id" bson:"id,omitempty"`
	Nickname string `json:"nickname" bson:"nickname,omitempty"`
	Age      uint   `json:"age" bson:"age,omitempty"`
}
