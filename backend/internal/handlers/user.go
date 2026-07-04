package handlers

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"silirek/internal/config"
	"silirek/internal/models"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nfnt/resize"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func UsersHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	// ambil role dari context (set di middleware)
	roleIfc, exists := c.Get("user_role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "missing role in context"})
		return
	}
	userRole := roleIfc.(string)

	var users []models.User
	dbQuery := h.db.Preload("Organization").Model(&models.User{})

	switch userRole {
	case string(models.UserRoleSuperadmin):
		// superadmin: no filter
	case string(models.UserRoleAdmin):
		// Admin: exclude both superadmin and admin roles
		dbQuery = dbQuery.Where("role NOT IN ?", []string{
			string(models.UserRoleSuperadmin),
			string(models.UserRoleAdmin),
		})
	default:
		// user biasa: hanya boleh lihat diri sendiri
		userID := c.GetUint("user_id")
		dbQuery = dbQuery.Where("id = ?", userID)
	}

	if err := dbQuery.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set password default
	rawPassword := "12345"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := h.db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id":    user.ID,
		"email": user.Email,
	})
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.db.Preload("Organization").First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid multipart form"})
		return
	}

	// Update data dari form field
	if fullName := form.Value["full_name"]; len(fullName) > 0 {
		user.FullName = fullName[0]
	}
	if email := form.Value["email"]; len(email) > 0 {
		user.Email = email[0]
	}
	if nip := form.Value["nip"]; len(nip) > 0 {
		user.Nip = &nip[0]
	}

	// Tambahkan pengecekan untuk organization_id
	if organizationID := form.Value["organization_id"]; len(organizationID) > 0 && organizationID[0] != "" {
		if idVal, err := strconv.ParseUint(organizationID[0], 10, 64); err == nil {
			orgID := uint(idVal)
			user.OrganizationID = &orgID
		}
	}

	// Optional: handle password
	if password := form.Value["password"]; len(password) > 0 && password[0] != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password[0]), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
			return
		}
		user.Password = string(hashedPassword)
	}

	// Optional: handle photo
	if files := form.File["photo"]; len(files) > 0 {
		// Hapus foto lama jika ada
		if user.Photo != nil {
			oldPath := "uploads/photos/" + *user.Photo
			if err := os.Remove(oldPath); err != nil && !os.IsNotExist(err) {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete old photo"})
				return
			}
		}

		photoFile := files[0]
		filename, err := saveResizedPhoto(photoFile)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		user.Photo = &filename
	}

	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := h.db.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	if err := h.db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user deleted"})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := h.db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}
	if !user.IsActive {
		c.JSON(http.StatusForbidden, gin.H{"error": "akun tidak aktif"})
		return
	}

	// Compare hash
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    string(user.Role),
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(config.JWTKey())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}

func (h *UserHandler) Me(c *gin.Context) {
	claims := c.MustGet("claims").(jwt.MapClaims)
	userID := uint(claims["user_id"].(float64))

	var user models.User
	if err := h.db.Preload("Organization").First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"photo": user.Photo,
		"role":  user.Role,
	})
}

// saveResizedPhoto decodes, resizes (if needed) to max 500px wide, and saves the image.
// It returns the saved filename or an error.
func saveResizedPhoto(header *multipart.FileHeader) (string, error) {
	f, err := header.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open photo file")
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return "", fmt.Errorf("invalid image format")
	}

	// Only resize if the image is wider than 500px
	var processed image.Image
	if img.Bounds().Dx() > 500 {
		processed = resize.Resize(500, 0, img, resize.Lanczos3)
	} else {
		processed = img
	}

	filename := time.Now().Format("20060102150405") + "_" + header.Filename
	savePath := "uploads/photos/" + filename

	out, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("failed to create photo file")
	}
	defer out.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext == ".png" {
		err = png.Encode(out, processed)
	} else {
		// default to jpeg with quality 80
		err = jpeg.Encode(out, processed, &jpeg.Options{Quality: 80})
	}
	if err != nil {
		return "", fmt.Errorf("failed to encode and save image")
	}

	return filename, nil
}
