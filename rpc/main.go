package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/sharma-ayush1999/rpc/proto/device"
)

// deviceServer implements pb.DeviceServiceServer
type deviceServer struct {
	pb.UnimplementedDeviceServiceServer // embed for forward compatibility
    // In production: inject a repository interface here
    // repo DeviceRepository
}

// GetDeviceStatus — Unary RPC
func (s *deviceServer) GetDeviceStatus(ctx context.Context, req *pb.GetDeviceStatusRequest) (*pb.GetDeviceStatusResponse, error) {
	//validate input
	if req.DeviceId == ""{
		return nil, status.Error(codes.InvalidArgument, "device_id is required")
	}

    // In production: s.repo.FindByID(ctx, req.DeviceId)
    // Simulated lookup:
	if req.DeviceId == "unknown" {
		return nil, status.Errorf(codes.NotFound, "device %s not found", req.DeviceId)
	}

	return &pb.GetDeviceStatusResponse {
		Device: &pb.Device {
			Id: req.DeviceId,
			Name: "Gateway " + req.DeviceId,
			Status: "online",
			LastSeen: timestamppb.Now(),
		},
	}, nil
}

// StreamDeviceEvents — Server Streaming RPC
func (s *deviceServer) StreamDeviceEvents(req *pb.StreamDeviceEventsRequest, stream pb.DeviceService_StreamDeviceEventsServer) error {
	if req.DeviceId == "" {
		return status.Error(codes.InvalidArgument, "device_id is required")
	}

	// Simulate streaming 5 events to client
	events := []string{"config_applied", "link_up", "ospf_neighbor_ip", "tunnel_established", "heartbeat"}
	for _, evt := range events {
        // Check if client disconnected
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		default:
		}

		if err := stream.Send(&pb.DeviceEvent {
			DeviceId: req.DeviceId,
			EventType: evt,
			Message: fmt.Sprintf("Event: %s on device %s", evt, req.DeviceId),
			OccurredAt: timestamppb.Now(),
		}); err != nil {
			return err
		}
		time.Sleep(500 * time.Millisecond)
	}
	return nil //stream closed gracefully
}


func main(){
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Add interceptors for logging and auth
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(loggingInterceptor, authInterceptor),
	)
	pb.RegisterDeviceServiceServer(grpcServer, &deviceServer{})

	log.Printf("grpc server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// ── Interceptors ─────────────────────────────────────────────────────────
func loggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	start := time.Now()
	resp, err := handler(ctx, req)
	log.Printf("RPC: %s | Duration: %v | Error: %v", info.FullMethod, time.Since(start), err)

	return resp, err
}

func authInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error){
	// Extract token from gRPC metadata
    // md, _ := metadata.FromIncomingContext(ctx)
    // token := md.Get("authorization")
    // validate token...
	return handler(ctx, req)
}