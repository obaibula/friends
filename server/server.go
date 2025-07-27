package main

import (
	"context"
	"fmt"

	pb "github.com/obaibula/friends/proto"
)

type FriendServer struct {
	pb.UnimplementedFriendServiceServer
	Graph *Graph
}

func NewFriendServer(g *Graph) *FriendServer {
	return &FriendServer{Graph: g}
}

func (s *FriendServer) AddFriend(ctx context.Context, req *pb.AddFriendRequest) (*pb.AddFriendResponse, error) {
	err := s.Graph.AddFriend(ctx, req.GetFrom(), req.GetTo())
	if err != nil {
		return nil, err
	}
	return &pb.AddFriendResponse{Message: fmt.Sprintf("%q added %q to friends", req.From, req.To)}, nil
}

func (s *FriendServer) GetMutualFriends(ctx context.Context, req *pb.MutualFriendsRequest) (*pb.MutualFriendsResponse, error) {
	mutualFriends, err := s.Graph.GetMutualFriends(ctx, req.GetUser1(), req.GetUser2())
	if err != nil {
		return nil, err
	}
	return &pb.MutualFriendsResponse{MutualFriends: mutualFriends}, nil
}
