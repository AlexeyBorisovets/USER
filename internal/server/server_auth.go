// Package server package server
package server

import (
	"context"
	"fmt"

	"github.com/AlexeyBorisovets/USER/internal/model"
	pb "github.com/AlexeyBorisovets/USER/proto"
	"golang.org/x/crypto/bcrypt"
)

// Authentication log-in
func (s *Server) Authentication(ctx context.Context, request *pb.AuthenticationRequest) (*pb.AuthenticationResponse, error) {
	idUser := request.GetId()
	password := request.GetPassword()
	accessToken, refreshToken, err := s.se.Authentication(ctx, idUser, password)
	if err != nil {
		return nil, err
	}
	return &pb.AuthenticationResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// Registration sign-up
func (s *Server) Registration(ctx context.Context, request *pb.RegistrRequest) (*pb.RegistrResponse, error) {
	hashPassword, err := hashingPassword(request.Password)
	if err != nil {
		return nil, fmt.Errorf("server: error while hashing password, %e", err)
	}
	request.Password = hashPassword
	p := model.User{
		Name:     request.Name,
		Balance:  uint(request.Balance),
		UserType: request.Usertype,
		Password: request.Password,
	}
	newID, err := s.se.CreateUser(ctx, &p)
	if err != nil {
		return nil, err
	}
	return &pb.RegistrResponse{Id: newID}, nil
}

// RefreshMyTokens refresh tokens
func (s *Server) RefreshMyTokens(ctx context.Context, refreshTokenString *pb.RefreshTokensRequest) (*pb.RefreshTokensResponse, error) { // refresh our tokens
	refreshTokenStr := refreshTokenString.GetRefreshToken()
	newRefreshToken, newAccessToken, err := s.se.RefreshTokens(ctx, refreshTokenStr)
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokensResponse{
		RefreshToken: newRefreshToken,
		AccessToken:  newAccessToken,
	}, nil
}

// Logout exit
func (s *Server) Logout(ctx context.Context, request *pb.LogoutRequest) (*pb.BufResponse, error) {
	err := s.se.Verify(request.AccessToken)
	if err != nil {
		return nil, err
	}
	idUser := request.GetId()
	err = s.se.UpdateUserAuth(ctx, idUser, "")
	if err != nil {
		return nil, err
	}
	return new(pb.BufResponse), nil
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
