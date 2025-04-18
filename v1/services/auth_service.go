package services

import (
	"bitslearn/config"
	"bitslearn/v1/models"
	"context"
	"errors"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// reg user 
func RegisterWithEmail(user *models.User) (string, error) {
	fsClient, _ := config.App.Firestore(context.Background())
	defer fsClient.Close()

	authClient, _ := config.App.Auth(context.Background())

	// Check if email exists
	iter := fsClient.Collection("users").Where("email", "==", user.Email).Documents(context.Background())
	if docs, _ := iter.GetAll(); len(docs) > 0 {
		return "", errors.New("email already registered")
	}

	// Hash password
	hashedPwd, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPwd)
	user.Create = true

	// Create user in Firestore
	docRef := fsClient.Collection("users").NewDoc()
	user.UID = docRef.ID
	if _, err := docRef.Set(context.Background(), user); err != nil {
		return "", err
	}

	// Create custom token
	token, err := authClient.CustomToken(context.Background(), user.UID)
	if err != nil {
		return "", err
	}

	return token, nil
}



// login user 
func LoginWithEmail(email, password string) (string, *models.User, error) {
	fsClient, _ := config.App.Firestore(context.Background())
	defer fsClient.Close()

	authClient, _ := config.App.Auth(context.Background())

	// Find user by email
	iter := fsClient.Collection("users").Where("email", "==", email).Documents(context.Background())
	docs, _ := iter.GetAll()
	if len(docs) == 0 {
		return "", nil, errors.New("user not found")
	}

	var user models.User
	_ = docs[0].DataTo(&user)

	// Check password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Create token
	token, err := authClient.CustomToken(context.Background(), user.UID)
	if err != nil {
		return "", nil, err
	}

	return token, &user, nil
}



func VerifyAndStoreUser(idToken string) (*models.User, error) {
	authClient, err := config.App.Auth(context.Background())
	if err != nil {
		return nil, err
	}

	token, err := authClient.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}

	uid := token.UID
	userInfo, err := authClient.GetUser(context.Background(), uid)
	if err != nil {
		return nil, err
	}

	firestoreClient, err := config.App.Firestore(context.Background())
	if err != nil {
		return nil, err
	}
	defer firestoreClient.Close()

	docRef := firestoreClient.Collection("users").Doc(uid)
	doc, err := docRef.Get(context.Background())

	var user models.User
	if err != nil {
		// New User
		user = models.User{
			UID:            uid,
			Username:       strings.Split(userInfo.Email, "@")[0],
			Name:           userInfo.DisplayName,
			Email:          userInfo.Email,
			ProfilePic:     userInfo.PhotoURL,
			Create:         true,
			AreaOfInterest: []string{},
		}
		_, err = docRef.Set(context.Background(), user)
		if err != nil {
			return nil, err
		}
	} else {
		err = doc.DataTo(&user)
		if err != nil {
			return nil, errors.New("error reading user data")
		}
	}

	return &user, nil
}
