package repository

import (
	"context"

	"github.com/maxmanovb/logistics-staff-service/src/domain/models"
)




type DriverRepository interface {
	// WithTx(ctx context.Context, f func(r DriverRepository) error) error
	SaveDriver(ctx context.Context, driver *models.Driver) error
	UpdateDriver(ctx context.Context, driver *models.Driver) error
	GetDriver(ctx context.Context, driverID string) (*models.Driver, error)
	GetDriverByPhone(ctx context.Context, phone string) (*models.Driver, error)
	SaveDriverProfile(ctx context.Context, profile *models.DriverProfile) error
	UpdateDriverProfile(ctx context.Context, profile *models.DriverProfile) error
	GetDriverProfile(ctx context.Context, driverID string) (*models.DriverProfile, error)
	// SaveDriverSmsCode(ctx context.Context, smsCode *models.DriverSmsCode) error
	GetDriverSmsCode(ctx context.Context, driverID, code string) (*models.DriverSmsCode, error)
	} 