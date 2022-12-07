// Package server p
package server

import (
	"github.com/AlexeyBorisovets/USER/internal/model"
	"github.com/AlexeyBorisovets/USER/internal/service"
	pb "github.com/AlexeyBorisovets/USER/proto"

	"context"
)

// Server struct
type Server struct {
	pb.UnimplementedUSERServer
	se *service.Service
}

// NewServer create new server connection
func NewServer(serv *service.Service) *Server {
	return &Server{se: serv}
}

// GetUser get user by id from db
func (s *Server) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	accessToken := request.GetAccessToken()
	if err := s.se.Verify(accessToken); err != nil {
		return nil, err
	}
	userID := request.GetId()
	userDB, err := s.se.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	userProto := &pb.GetUserResponse{
		User: &pb.User{
			Id:       userDB.ID,
			Name:     userDB.Name,
			Balance:  uint64(userDB.Balance),
			Usertype: userDB.UserType,
			Password: userDB.Password,
		},
	}
	return userProto, nil
}

// GetAllUsers get all users from db
func (s *Server) GetAllUsers(ctx context.Context, _ *pb.GetAllUsersRequest) (*pb.GetAllUsersResponse, error) {
	users, err := s.se.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	var list []*pb.User
	for _, user := range users {
		userProto := new(pb.User)
		userProto.Id = user.ID
		userProto.Name = user.Name
		userProto.Balance = uint64(user.Balance)
		userProto.Usertype = user.UserType
		list = append(list, userProto)
	}
	return &pb.GetAllUsersResponse{Users: list}, nil
}

// DeleteUser delete user by id
func (s *Server) DeleteUser(ctx context.Context, request *pb.DeleteUserRequest) (*pb.BufResponse, error) {
	idUser := request.GetId()
	err := s.se.DeleteUser(ctx, idUser)
	if err != nil {
		return nil, err
	}
	return new(pb.BufResponse), nil
}

// UpdateUser update user with new parameters
func (s *Server) UpdateUser(ctx context.Context, request *pb.UpdateUserRequest) (*pb.BufResponse, error) {
	accessToken := request.GetAccessToken()
	if err := s.se.Verify(accessToken); err != nil {
		return nil, err
	}
	user := &model.User{
		Name:     request.User.Name,
		UserType: request.User.Usertype,
		Balance:  uint(request.User.Balance),
	}
	idUser := request.GetId()
	err := s.se.UpdateUser(ctx, idUser, user)
	if err != nil {
		return nil, err
	}
	return new(pb.BufResponse), nil
}

func (s *Server) GetUserByUserType(ctx context.Context, request *pb.GetUserByUserTypeRequest) (*pb.GetUserByUserTypeResponse, error) {
	accessToken := request.GetAccessToken()
	if err := s.se.Verify(accessToken); err != nil {
		return nil, err
	}
	usertype := request.GetUsertype()
	users, err := s.se.GetUserByUserType(ctx, usertype)
	if err != nil {
		return nil, err
	}
	var list []*pb.User
	for _, user := range users {
		userProto := new(pb.User)
		userProto.Id = user.ID
		userProto.Name = user.Name
		userProto.Balance = uint64(user.Balance)
		userProto.Usertype = user.UserType
		list = append(list, userProto)
	}
	return &pb.GetUserByUserTypeResponse{Users: list}, nil
}

func (s *Server) GetBalanceByID(ctx context.Context, request *pb.GetBalanceByIDRequest) (*pb.GetBalanceByIDResponse, error) {
	userID := request.GetId()
	BalanceDB, err := s.se.GetBalanceByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	BalanceProto := &pb.GetBalanceByIDResponse{Balance: uint64(BalanceDB)}
	return BalanceProto, nil
}

func (s *Server) UpdateBalance(ctx context.Context, request *pb.UpdateBalanceRequest) (*pb.BufResponse, error) {
	userID := request.GetId()
	newBalance := request.GetBalance()

	err := s.se.UpdateBalance(ctx, userID, uint(newBalance))
	if err != nil {
		return nil, err
	}
	return new(pb.BufResponse), nil
}
