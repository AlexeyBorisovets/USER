package service

import (
	"context"
	"fmt"
	"time"

	"github.com/AlexeyBorisovets/USER/internal/model"
	"github.com/AlexeyBorisovets/USER/internal/repository"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

// AccessTokenWorkTime RefreshTokenWorkTime duration time
var (
	AccessTokenWorkTime  = time.Now().Add(time.Minute * 10).Unix()
	RefreshTokenWorkTime = time.Now().Add(time.Hour * 3).Unix()
)

// Authentication log-in
func (se *Service) Authentication(ctx context.Context, id, password string) (accessToken, refreshToken string, err error) {
	authUser, err := se.repos.GetUserByID(ctx, id)
	if err != nil {
		return "", "", err
	}
	incoming := []byte(password)
	existing := []byte(authUser.Password)
	_ = bcrypt.CompareHashAndPassword(existing, incoming) // check passwords
	// if err != nil {
	// 	return "", "", err //Problem with comparing
	// }
	authUser.Password = password
	accessToken, refreshToken, err = se.CreateJWT(ctx, se.repos, authUser)
	if err != nil {
		return "", "", err
	}
	return
}

// CreateUser create new user, add him to db
func (se *Service) CreateUser(ctx context.Context, u *model.User) (string, error) {
	hashedPassword, err := hashingPassword(u.Password)
	if err != nil {
		return "", err
	}
	u.Password = hashedPassword
	return se.repos.CreateUser(ctx, u)
}

// RefreshTokens refresh tokens
func (se *Service) RefreshTokens(ctx context.Context, refreshTokenStr string) (newRefreshToken, newAccessToken string, err error) { // refresh our tokens
	refreshToken, err := jwt.Parse(refreshTokenStr, func(t *jwt.Token) (interface{}, error) {
		return se.jwtKey, nil
	}) // parse it into string format
	if err != nil {
		log.Errorf("service: can't parse refresh token - %e", err)
		return "", "", err
	}
	if !refreshToken.Valid {
		return "", "", fmt.Errorf("service: invalid refresh token")
	}
	claims := refreshToken.Claims.(jwt.MapClaims)
	userUUID := claims["jti"]
	if userUUID == "" {
		return "", "", fmt.Errorf("service: error while parsing claims, ID couldnt be empty")
	}
	user, err := se.repos.SelectByIDAuth(ctx, userUUID.(string))
	if err != nil {
		return "", "", fmt.Errorf("service: token refresh failed - %e", err)
	}
	if refreshTokenStr != user.RefreshToken {
		return "", "", fmt.Errorf("service: invalid refresh token")
	}

	return se.CreateJWT(ctx, se.repos, &user)
}

// CreateJWT create jwt tokens
func (se *Service) CreateJWT(ctx context.Context, rps repository.Repository, user *model.User) (accessTokenStr, refreshTokenStr string, err error) {
	accessToken := jwt.New(jwt.SigningMethodHS256)            // encrypt access token by SigningMethodHS256 method
	claimsA := accessToken.Claims.(jwt.MapClaims)             // fill access-token`s claims
	claimsA["exp"] = AccessTokenWorkTime                      // work time
	claimsA["username"] = user.Name                           // payload
	accessTokenStr, err = accessToken.SignedString(se.jwtKey) // convert token to string format
	if err != nil {
		log.Errorf("service: can't generate access token - %v", err)
		return "", "", err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claimsR := refreshToken.Claims.(jwt.MapClaims)
	claimsR["username"] = user.Name
	claimsR["exp"] = RefreshTokenWorkTime
	claimsR["jti"] = user.ID
	refreshTokenStr, err = refreshToken.SignedString(se.jwtKey)
	if err != nil {
		log.Errorf("service: can't generate access token - %v", err)
		return "", "", err
	}
	err = rps.UpdateAuth(ctx, user.ID, refreshTokenStr) // add into user refresh token
	if err != nil {
		log.Errorf("service: can't generate access token - %v", err)
		return "", "", err
	}
	return
}

// UpdateUserAuth update auth user, add token
func (se *Service) UpdateUserAuth(ctx context.Context, id, refreshToken string) error {
	return se.repos.UpdateAuth(ctx, id, refreshToken)
}

// Verify verify access token
func (se *Service) Verify(accessTokenString string) error {
	accessToken, err := jwt.Parse(accessTokenString, func(t *jwt.Token) (interface{}, error) {
		return se.jwtKey, nil
	})
	if err != nil {
		log.Errorf("service: can't parse accessToken - ", err)
		return err
	}
	if !accessToken.Valid {
		//return fmt.Errorf("service: expired accessToken")
	}
	return nil
}

// hashingPassword _
func hashingPassword(password string) (string, error) {
	bytesPassword := []byte(password)
	hashedBytesPassword, err := bcrypt.GenerateFromPassword(bytesPassword, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	hashPassword := string(hashedBytesPassword)
	return hashPassword, nil
}
