package grpc

import (
	"context"

	pb "github.com/maxmanovb/logistics-staff-service/src/application/protos/logistics_staff"
	 "github.com/maxmanovb/logistics-staff-service/src/application/services"
)

type Server struct {
    pb.StaffServiceServer
    driverApp service.DriverApplicationService
}

func NewServer(driverApp service.DriverApplicationService) *Server {
    return &Server{
        driverApp: driverApp,
    }
}

func (s *Server) RegisterDriver(ctx context.Context, r *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {
    response, err := s.driverApp.RegisterDriver(ctx, r)
    if err != nil {
        return nil, err
    }
    return response, nil
}

// func (s *Server) ChangeDriverPassword(ctx context.Context, r *pb.ChangeDriverPasswordRequest) (*pb.ChangeDriverPasswordResponse, error) {
//     err := s.driverApp.ChangeDriverPassword(ctx, r)
//     if err != nil {
//         return nil, err
//     }
//     return &pb.ChangeDriverPasswordResponse{}, nil
// }

func (s *Server) LoginDriver(ctx context.Context, r *pb.LoginDriverRequest) (*pb.LoginDriverResponse, error) {
    response, err := s.driverApp.LoginDriver(ctx, r)
    if err != nil {
        return nil, err
    }
    return response, nil
}

func (s *Server) ConfirmSMSCode(ctx context.Context, r *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error) {
    response, err := s.driverApp.ConfirmSMSCode(ctx, r)
    if err != nil {
        return nil, err
    }
    return response, nil
}

func (s *Server) GetDriverProfile(ctx context.Context, r *pb.GetDriverProfileRequest) (*pb.GetDriverProfileResponse, error) {
    response, err := s.driverApp.GetDriverProfile(ctx, r)
    if err != nil {
        return nil, err
    }
    return response, nil
}

func (s *Server) UpdateDriverProfile(ctx context.Context, r *pb.UpdateDriverProfileRequest) (*pb.UpdateDriverProfileResponse, error) {
    response, err := s.driverApp.UpdateDriverProfile(ctx, r)
    if err != nil {
        return nil, err
    }
    return response, nil
}
