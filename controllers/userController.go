// package controllers

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/RND2002/goChatApp/db"
// 	"github.com/RND2002/goChatApp/models"
// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt/v4"
// )

// type RegistrationData struct {
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// // Register @Summary Register a new user
// // @Description Register a new user with username and password
// // @Tags auth
// // @Accept json
// // @Produce json
// // @Param user body models.User true "User data"
// // @Success 200 {object} models.User
// // @Failure 400
// // @Router /auth/register [post]
// func Register(c *gin.Context) {
// 	var formData RegistrationData
// 	err := c.ShouldBindJSON(&formData)

// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid form data"})
// 		return
// 	}

// 	if formData.Username == "" || formData.Email == "" || formData.Password == "" {
// 		c.JSON(400, gin.H{"error": "Empty fields not allowed"})
// 		return
// 	}

// 	var user models.User
// 	user.Username = formData.Username
// 	user.Email = formData.Email
// 	user.Password = HashPassword(formData.Password)

// 	err = db.DB.Create(&user).Error
// 	if err != nil {
// 		c.JSON(500, gin.H{"message": "Error registering user", "error": err})
// 		return
// 	}

// 	c.JSON(200, gin.H{"message": "User saved successfully", "user": user})
// }

// type LoginData struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// // Login @Summary Log in a user
// // @Description Log in with username and password
// // @Tags auth
// // @Accept json
// // @Produce json
// // @Param user body LoginData true "Login data"
// // @Success 200 {object} models.User
// // @Failure 400 {object} gin.H{"error": "Invalid form data"}
// // @Failure 404 {object} gin.H{"error": "User not found"}
// // @Failure 401 {object} gin.H{"error": "Invalid password"}
// // @Failure 500 {object} gin.H{"error": "Error generating token Or Token Expired"}
// // @Router /auth/login [post]
// func Login(c *gin.Context) {
// 	var loginData LoginData

// 	err := c.ShouldBindJSON(&loginData)
// 	if err != nil {
// 		c.JSON(400, gin.H{"error": "Invalid form data"})
// 		return
// 	}

// 	var user models.User
// 	err = db.DB.Where("username = ?", loginData.Username).First(&user).Error
// 	if err != nil {
// 		c.JSON(404, gin.H{"error": "User not found"})
// 		return
// 	}
// 	if !CompareHashedPassword(loginData.Password, user.Password) {
// 		c.JSON(401, gin.H{"error": "Invalid password"})
// 		return
// 	}

// 	token, err := GenerateToken(user)
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Error generating token Or Token Expired"})
// 		return
// 	}

// 	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

// 	c.JSON(200, gin.H{"message": "User logged in successfully", "user": user, "token": token})

// }

// // DeleteUser @Summary Delete a user
// // @Description Delete a user by ID
// // @Tags auth
// // @Accept json
// // @Produce json
// // @Param id path string true "User ID"
// // @Success 200 {object} gin.H{"message": "User deleted successfully", "user": models.User}
// // @Failure 404 {object} gin.H{"error": "User not found"}
// // @Router /auth/delete/{id} [delete]
// func DeleteUser(c *gin.Context) {
// 	id := c.Param("id")
// 	var user models.User
// 	err := db.DB.Where("id = ?", id).First(&user).Error
// 	if err != nil {
// 		c.JSON(404, gin.H{"error": "User not found"})
// 		return
// 	}
// 	db.DB.Delete(&user)
// 	c.JSON(200, gin.H{"message": "User deleted successfully", "user": user})
// }

// // GetUsers @Summary Get all users
// // @Description Fetch all users
// // @Tags auth
// // @Accept json
// // @Produce json
// // @Success 200 {object} []models.User
// // @Failure 500 {object} gin.H{"error": "Error fetching users"}
// // @Router /auth/users [get]
// func GetUsers(c *gin.Context) {
// 	var users []models.User
// 	err := db.DB.Find(&users).Error
// 	if err != nil {
// 		c.JSON(500, gin.H{"error": "Error fetching users"})
// 		return
// 	}
// 	c.JSON(200, gin.H{"users": users})
// }

// // Logout @Summary Log out a user
// // @Description Log out the currently logged-in user
// // @Tags auth
// // @Accept json
// // @Produce json
// // @Success 200 {object} gin.H{"message": "User logged out successfully"}
// // @Router /auth/logout [post]
// func Logout(c *gin.Context) {
// 	c.SetCookie("jwt", "", -1, "", "", false, true)
// 	c.JSON(200, gin.H{"message": "User logged out successfully"})
// }

// // User @Summary Get currently logged-in user
// // @Description Fetch details of the currently logged-in user
// // @Tags auth
// // @Accept json
// // @Produce json
// // @Success 200 {object} models.User
// // @Failure 401 {object} gin.H{"message": "No JWT token provided Or No user Logged in"}
// // @Failure 400 {object} gin.H{"message": "Error parsing token or token is invalid"}
// // @Failure 404 {object} gin.H{"message": "User not found"}
// // @Router /auth/user [get]
// func User(c *gin.Context) {
// 	fmt.Println("Request to Get User")

// 	// Get the "jwt" cookie by name
// 	cookie, err := c.Cookie("jwt")
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"message": "No JWT token provided Or No user Logged in"})
// 		return
// 	}

// 	// Parse the JWT token
// 	token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secret), nil
// 	})
// 	if err != nil || !token.Valid {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing token or token is invalid"})
// 		return
// 	}

// 	// Extract and process claims as before
// 	claims, ok := token.Claims.(*CustomClaims)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse claims"})
// 		return
// 	}

// 	id := claims.UserID

// 	// Fetch the user from the database
// 	user := models.User{ID: id}
// 	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{"message": "This was the user logged in", "user": user})
//	}
package controllers

import (
	"fmt"
	"net/http"

	"github.com/RND2002/goChatApp/db"
	"github.com/RND2002/goChatApp/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type RegistrationData struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register @Summary Register a new user
// @Description Register a new user with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User data"
// @Success 200 {object} models.User
// @Failure 400
// @Router /auth/register [post]
func Register(c *gin.Context) {
	var formData RegistrationData
	err := c.ShouldBindJSON(&formData)

	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid form data"})
		return
	}

	if formData.Username == "" || formData.Email == "" || formData.Password == "" {
		c.JSON(400, gin.H{"error": "Empty fields not allowed"})
		return
	}

	var user models.User
	user.Username = formData.Username
	user.Email = formData.Email
	user.Password, _ = HashPassword(formData.Password)

	err = db.DB.Create(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"message": "Error registering user", "error": err})
		return
	}

	c.JSON(200, gin.H{"message": "User saved successfully", "user": user})
}

type LoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login @Summary Log in a user

func Login(c *gin.Context) {
	var loginData LoginData

	err := c.ShouldBindJSON(&loginData)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid form data"})
		return
	}

	var user models.User
	err = db.DB.Where("username = ?", loginData.Username).First(&user).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	if !CompareHashedPassword(loginData.Password, user.Password) {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	token, err := GenerateToken(user)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token Or Token Expired"})
		return
	}

	c.SetCookie("jwt", token, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "User logged in successfully", "user": user, "token": token})

}

// DeleteUser @Summary Delete a user
// @Description Delete a user by ID
// @Tags auth
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} gin.H{"message": "User deleted successfully", "user": models.User}
// @Failure 404 {object} gin.H{"error": "User not found"}
// @Router /auth/delete/{id} [delete]
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	err := db.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}
	db.DB.Delete(&user)
	c.JSON(200, gin.H{"message": "User deleted successfully", "user": user})
}

// GetUsers @Summary Get all users
// @Description Fetch all users
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Failure 500 {object} gin.H{"error": "Error fetching users"}
// @Router /auth/users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	err := db.DB.Find(&users).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Error fetching users"})
		return
	}
	c.JSON(200, gin.H{"users": users})
}

// Logout @Summary Log out a user
// @Description Log out the currently logged-in user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} gin.H{"message": "User logged out successfully"}
// @Router /auth/logout [post]
func Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(200, gin.H{"message": "User logged out successfully"})
}

// User @Summary Get currently logged-in user
// @Description Fetch details of the currently logged-in user
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} models.User
// @Failure 401 {object} gin.H{"message": "No JWT token provided Or No user Logged in"}
// @Failure 400 {object} gin.H{"message": "Error parsing token or token is invalid"}
// @Failure 404 {object} gin.H{"message": "User not found"}
// @Router /auth/user [get]
func User(c *gin.Context) {
	fmt.Println("Request to Get User")

	// Get the "jwt" cookie by name
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "No JWT token provided Or No user Logged in"})
		return
	}

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(cookie, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Error parsing token or token is invalid"})
		return
	}

	// Extract and process claims as before
	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse claims"})
		return
	}

	id := claims.UserID

	// Fetch the user from the database
	user := models.User{ID: id}
	if err := db.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "This was the user logged in", "user": user})
}
