package handler

import (
	"crypto/md5"
	"database/sql"
	"depogunabangunan/apps/interfaces"
	"depogunabangunan/apps/model"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	db             *sql.DB
	userRepository interfaces.IUserService
}

func NewLoginHandler(db *sql.DB, userRepository interfaces.IUserService) *LoginHandler {
	return &LoginHandler{
		db:             db,
		userRepository: userRepository,
	}
}

// LoginHandler handles the login endpoint
func (h *LoginHandler) Login(c *gin.Context) {
	var userLogin model.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Retrieve the user from the database based on the provided username
	user, err := getUserByUsernameAndPassword(h.db, userLogin.Username, userLogin.Pass)
	if err != nil {
		recordLoginHistory(h.db, userLogin.Username, false)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate and return the JWT token
	token, err := generateJWTToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		recordLoginHistory(h.db, user.Username, false)
		return
	}
	recordLoginHistory(h.db, user.Username, true)
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// getUserByUsernameAndPassword retrieves a user from the database by username and password
func getUserByUsernameAndPassword(db *sql.DB, username, password string) (*model.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = $1"
	row := db.QueryRow(query, username)

	user := &model.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	// Encrypt the provided password using MD5
	hashedPassword := encryptMD5(password)

	// Compare the hashed password with the password from the database
	if user.Password != hashedPassword {
		return nil, fmt.Errorf("invalid password")
	}

	return user, nil
}

// encryptMD5 encrypts the password using MD5
func encryptMD5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

// generateJWTToken generates a JWT token for the user
func generateJWTToken(user *model.User) (string, error) {
	// Create the claims
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (1 day)
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with a secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// recordLoginHistory records the login history in the database
func recordLoginHistory(db *sql.DB, username string, success bool) error {
	query := "INSERT INTO login_history (username, login_time, login_success) VALUES ($1, $2, $3)"
	_, err := db.Exec(query, username, time.Now(), success)
	if err != nil {
		return err
	}

	return nil
}
