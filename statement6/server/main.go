package main

import (
	"context"
	"log"
	"net"

	pb "github.com/CynthiaAntony/CynthiaInt1/statement6/proto"
	"github.com/CynthiaAntony/CynthiaInt1/statement6/trie"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedTrieServiceServer
	t *trie.Trie
}

func (s *server) Add(ctx context.Context, req *pb.WordRequest) (*pb.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}
	a := req.GetWord()
	s.t.Add(a)
	log.Printf("Added word: %s", a)
	return &pb.Empty{}, nil
}
func (s *server) Check(ctx context.Context, req *pb.WordRequest) (*pb.BoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}
	a := req.GetWord()
	exists := s.t.Check(a)
	log.Printf("Checked word: %s, exists: %t", a, exists)
	return &pb.BoolResponse{Exists: exists}, nil
}
func (s *server) Remove(ctx context.Context, req *pb.WordRequest) (*pb.BoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}
	a := req.GetWord()
	removed := s.t.Remove(a)
	log.Printf("Removed word: %s, removed: %t", a, removed)
	return &pb.BoolResponse{Exists: removed}, nil
}
func (s *server) List(ctx context.Context, req *pb.Empty) (*pb.ListResponse, error) {
	words := s.t.List()
	log.Printf("Listed words: %v", words)
	return &pb.ListResponse{Words: words}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTrieServiceServer(s, &server{t: trie.InitTrie()})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
