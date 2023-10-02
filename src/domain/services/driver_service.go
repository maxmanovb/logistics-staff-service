package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/maxmanovb/logistics-staff-service/src/domain/models"
	"github.com/maxmanovb/logistics-staff-service/src/domain/repository"
	"github.com/maxmanovb/logistics-staff-service/src/infrastructure/crypto"
	"github.com/maxmanovb/logistics-staff-service/src/infrastructure/rand"
)

type DriverService interface {
	RegisterDriver(ctx context.Context, phoneNumber, password string) (string, error)
	ChangeDriverPassword(ctx context.Context, driverID, currentPassword, newPassword string) error
	LoginDriver(ctx context.Context, phoneNumber, password string) (string, error)
	ConfirmSMSCode(ctx context.Context, driverID, smsCode string) (*models.DriverProfile, error)
	GetDriverProfile(ctx context.Context, driverID string) (*models.DriverProfile, error)
	UpdateDriverProfile(ctx context.Context, driverID, name, imageUrl string) (*models.DriverProfile, error)
	}
 
type driverSvcImpl struct {
	driverRepo repository.DriverRepository
}

func NewDriverSvc(driverRepo repository.DriverRepository) DriverService {
	return &driverSvcImpl{
		driverRepo: driverRepo,
	}
}

func (s *driverSvcImpl) RegisterDriver(ctx context.Context, phoneNumber, password string) (string, error) {

	driver, err := s.driverRepo.GetDriverByPhone(ctx, phoneNumber)
	if err == nil {
		return driver.ID, fmt.Errorf("driver with this phone already exists: %s", phoneNumber)
	}
	var (
		driverID   = rand.UUID()
		salt        = crypto.GenerateSalt()
		saltedPass  = crypto.Combine(salt, password)
		passHash    = crypto.HashPassword(saltedPass)
		now         = time.Now().UTC()
	)
	driver = &models.Driver{
		ID:           driverID,
		PhoneNumber:  phoneNumber,
		PasswordHash: passHash,
		PasswordSalt: salt,
		CreatedAt:    now,
		UpdatedAt:    now,
	}



		if err := s.driverRepo.SaveDriver(ctx, driver); err != nil {
			return "driver with this phone", err
		}

		return driverID ,nil

}

func (s *driverSvcImpl) ChangeDriverPassword(ctx context.Context, driverID, currentPassword, newPassword string) error {

	driver, err := s.driverRepo.GetDriver(ctx, driverID)
	if err != nil {
		return err
	}
	if !crypto.PasswordMatch(currentPassword, driver.PasswordSalt, driver.PasswordHash) {
		return errors.New("password doesn't match")
	}

	var (
		salt       = crypto.GenerateSalt()
		saltedPass = crypto.Combine(salt, newPassword)
		passHash   = crypto.HashPassword(saltedPass)
	)

	if driver.PasswordHash != passHash {
		return errors.New("invalid or missing currentPassword")
	}
	updatedDriver := &models.Driver{
        ID:           driver.ID,
        PasswordSalt: salt,
        PasswordHash: passHash,
    }
	err = s.driverRepo.UpdateDriver(ctx, updatedDriver)

	if err != nil {
		return err
	}

	return nil
}

func (s *driverSvcImpl) LoginDriver(ctx context.Context,  phoneNumber, password string) (string, error) {

	driver, err := s.driverRepo.GetDriverByPhone(ctx, phoneNumber)

	if err != nil {
		return "err", err
	}
	if !crypto.PasswordMatch(password, driver.PasswordSalt, driver.PasswordHash) {
		return "rr", errors.New("password doesn't match")
	}



	return driver.ID, nil
}

func (s *driverSvcImpl) ConfirmSMSCode(ctx context.Context, driverID, code string) (*models.DriverProfile, error) {
	// smsCode, err := s.driverRepo.GetDriverSmsCode(ctx, driverID, code)
	// if err != nil {
	// 	return nil, err
	// }
    // currentTime := time.Now()
	// if currentTime.After(smsCode.ExpiresIn) {
    //     return true // Code has expired
    // }
	// if smsCode.ExpiresIn() {
	// 	return nil, errors.New("code is expired")
	// }

	driverProfile, err := s.driverRepo.GetDriverProfile(ctx, driverID)
	if err != nil {
		return nil, err
	}

	return driverProfile, nil
}

func (s *driverSvcImpl) GetDriverProfile(ctx context.Context, driverId string) (*models.DriverProfile, error) {

	driver, err := s.driverRepo.GetDriverProfile(ctx, driverId)

	if err != nil {
		return nil, err
	}




	return driver, nil
}


func (s *driverSvcImpl) UpdateDriverProfile(ctx context.Context, driverID, name, imageUrl string) (*models.DriverProfile, error) {

	driver, err := s.driverRepo.GetDriverProfile(ctx, driverID)
	if err != nil {
		return nil,err
	}
	driver.Name = name
	driver.ImageUrl = imageUrl
	driver.UpdatedAt = time.Now().UTC()


	err = s.driverRepo.UpdateDriverProfile(ctx, driver)

	if err != nil {
		return nil,err
	}

	return driver, nil
}