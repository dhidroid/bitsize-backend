package models

type User struct {
	UID            string   `json:"uid" bson:"uid"`
	Username       string   `json:"username" bson:"username"`
	Name           string   `json:"name" bson:"name"`
	Email          string   `json:"email" bson:"email"`
	Password       string   `json:"password,omitempty" bson:"password"`
	ProfilePic     string   `json:"profilePic" bson:"profilePic"`
	AreaOfInterest []string `json:"areaOfInterest" bson:"areaOfInterest"`
	IsCreatorMode  bool     `json:"isCreatorMode" bson:"isCreatorMode"`
	CreatedAt      string   `json:"createdAt" bson:"createdAt"`
	UpdatedAt      string   `json:"updatedAt" bson:"updatedAt"`
	IsAdmin        bool     `json:"admin" bson:"admin"`
}
