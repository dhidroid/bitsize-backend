package services

import (
	"bitslearn/v1/models"
	"bitslearn/v1/utils"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitUserService(db *mongo.Database) {
	userCollection = db.Collection("users")
}

// Register user
func RegisterWithEmail(user *models.User) (string, error) {
	// Check if email already exists
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.New("email already registered")
	}

	// 

	// Check if username already exists
	count, err = userCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	if err != nil {
		return "", err
	}
	if count > 0 {
		return "", errors.New("username already taken")
	}

	// default admin fasle
	user.IsAdmin = false

	// Hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPwd)

	// Generate UID (use UUID in real app)
	user.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	user.UpdatedAt = time.Now().Format("2006-01-02 15:04:05")

	// Insert into DB
	_, err = userCollection.InsertOne(ctx, user)
	if err != nil {
		return "", err
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.UID)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Login user
func LoginWithEmail(email, password string) (string, *models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.User
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return "", nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateJWT(user.UID)
	if err != nil {
		return "", nil, err
	}

	user.Password = ""
	return token, &user, nil
}

// get all users
func GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		user.Password = "" // hide password
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
