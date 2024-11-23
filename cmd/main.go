package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	pb "github.com/liju-github/CentralisedFoodbuddyMicroserviceProto/Admin"
	"google.golang.org/grpc"
)

type adminServer struct {
	pb.UnimplementedAdminServiceServer
}

func (s *adminServer) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	// Load credentials from environment
	adminUsername := os.Getenv("ADMINUSERNAME")
	adminPassword := os.Getenv("ADMINPASSWORD")

	// Check credentials
	if req.Username == adminUsername && req.Password == adminPassword {
		return &pb.AdminLoginResponse{
			Success: true,
			Message: "Login successful",
		}, nil
	}

	return &pb.AdminLoginResponse{
		Success: false,
		Message: "Invalid credentials",
	}, nil
}

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Get port from environment
	port := os.Getenv("ADMINGRPCPORT")
	if port == "" {
		port = "50053"
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterAdminServiceServer(s, &adminServer{})

	log.Printf("Admin service starting on port %s", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
