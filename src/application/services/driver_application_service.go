package service

import (
	"context"

	"github.com/maxmanovb/logistics-staff-service/src/application/dtos"
	pb "github.com/maxmanovb/logistics-staff-service/src/application/protos/logistics_staff"
	"github.com/maxmanovb/logistics-staff-service/src/domain/services"
)

type DriverApplicationService interface {
    RegisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error)
    // ChangeDriverPassword(ctx context.Context, req *pb.Dr) (*pb.ChangeDriverPasswordResponse, error)
    LoginDriver(ctx context.Context, req *pb.LoginDriverRequest) (*pb.LoginDriverResponse, error)
    ConfirmSMSCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error)
    GetDriverProfile(ctx context.Context, req *pb.GetDriverProfileRequest) (*pb.GetDriverProfileResponse, error)
    UpdateDriverProfile(ctx context.Context, req *pb.UpdateDriverProfileRequest) (*pb.UpdateDriverProfileResponse, error)
}

type driverAppSvcImpl struct {
    driverSvc services.DriverService
}

func NewDriverApplicationService(driverSvc services.DriverService) DriverApplicationService {
    return &driverAppSvcImpl{
        driverSvc: driverSvc,
    }
}

func (s *driverAppSvcImpl) RegisterDriver(ctx context.Context, req *pb.RegisterDriverRequest) (*pb.RegisterDriverResponse, error) {

	driverID,err := s.driverSvc.RegisterDriver(ctx, req.PhoneNumber, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterDriverResponse{
		DriverId: driverID,
	}, nil
}



func (s *driverAppSvcImpl) LoginDriver(ctx context.Context, req *pb.LoginDriverRequest) (*pb.LoginDriverResponse, error) {
	driverID,err := s.driverSvc.LoginDriver(ctx, req.PhoneNumber, req.Password)
	if err != nil {
		return nil, err
	}

	return &pb.LoginDriverResponse{
		DriverId: driverID,
	}, nil
}

func (s *driverAppSvcImpl) ConfirmSMSCode(ctx context.Context, req *pb.ConfirmSmsCodeRequest) (*pb.ConfirmSmsCodeResponse, error) {
	driverProfile,err := s.driverSvc.ConfirmSMSCode(ctx, req.DriverId, req.SmsCode)
	if err != nil {
		return nil, err
	}

	return &pb.ConfirmSmsCodeResponse{
		Profile: dtos.NewDriverProfileFormat(driverProfile),
	}, nil}

func (s *driverAppSvcImpl) GetDriverProfile(ctx context.Context, req *pb.GetDriverProfileRequest) (*pb.GetDriverProfileResponse, error) {
	driverProfile,err := s.driverSvc.GetDriverProfile(ctx, req.DriverId)
	if err != nil {
		return nil, err
	}

	return &pb.GetDriverProfileResponse{
		Profile: dtos.NewDriverProfileFormat(driverProfile),
	}, nil

}

func (s *driverAppSvcImpl) UpdateDriverProfile(ctx context.Context, req *pb.UpdateDriverProfileRequest) (*pb.UpdateDriverProfileResponse, error) {
	driverProfile,err := s.driverSvc.UpdateDriverProfile(ctx, req.DriverId, req.Name, req.ImageUrl)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateDriverProfileResponse{
		Profile: dtos.NewDriverProfileFormat(driverProfile),
	}, nil
}