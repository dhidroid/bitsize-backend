package models

type User struct {
	UID            string   `json:"uid" firestore:"uid"`
	Username       string   `json:"username" firestore:"username"`
	Name           string   `json:"name" firestore:"name"`
	Email          string   `json:"email" firestore:"email"`
	Password       string   `json:"password,omitempty" firestore:"password,omitempty"`
	ProfilePic     string   `json:"profilePic" firestore:"profilePic"`
	Create         bool     `json:"create" firestore:"create"`
	AreaOfInterest []string `json:"areaOfInterest" firestore:"areaOfInterest"`
}
