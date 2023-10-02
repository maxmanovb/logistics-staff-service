package driver

import (
	"context"

	"github.com/maxmanovb/logistics-staff-service/src/domain/models"
	repositories "github.com/maxmanovb/logistics-staff-service/src/domain/repository"

	"gorm.io/gorm"
)

const(

  driverTable = "support.drivers"
  driverProfileTable = "support.driver_profile"
  driverSmsTable = "support.driver_sms_code"
)

type driverrepoImpl struct {
	db *gorm.DB
}
 
func NewDriverRepository(db *gorm.DB) repositories.DriverRepository{
	return &driverrepoImpl{
		db:db,
	}
}
func (r *driverrepoImpl) GetDriverByPhone(ctx context.Context, phone string) (*models.Driver, error) {

	var driver models.Driver
	result := r.db.WithContext(ctx).Table(driverTable).First(&driver, "phone_number = ?", phone)
	if result.Error != nil {
		return nil, result.Error
	}
	return &driver, nil
}


func (r *driverrepoImpl)SaveDriver(ctx context.Context, driver *models.Driver) error{
	result := r.db.WithContext(ctx).Table(driverTable).Create(&driver)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}


func (r *driverrepoImpl)SaveDriverProfile(ctx context.Context, driver *models.DriverProfile) error{
	result := r.db.WithContext(ctx).Table(driverProfileTable).Create(&driver)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}

func (r *driverrepoImpl)UpdateDriver(ctx context.Context, driver *models.Driver) error{
	result := r.db.WithContext(ctx).Table(driverTable).Save(&driver)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}

func (r *driverrepoImpl)UpdateDriverProfile(ctx context.Context, driver *models.DriverProfile) error{
	result := r.db.WithContext(ctx).Table(driverProfileTable).Save(&driver)

	if result.Error != nil{
		return result.Error
	}

	return nil;
}
func (r *driverrepoImpl)GetDriver(ctx context.Context, driverID string)(*models.Driver, error){
	var driver *models.Driver
	result := r.db.WithContext(ctx).Table(driverTable).First(&driver, "id = ?", driverID)

	if result.Error != nil {
		return nil, result.Error
	}

	return driver, nil;
}

func (r *driverrepoImpl)GetDriverProfile(ctx context.Context, driverID string)(*models.DriverProfile, error){
	var driver *models.DriverProfile
	result := r.db.WithContext(ctx).Table(driverSmsTable).First(&driver, "driver_id = ?", driverID)

	if result.Error != nil {
		return nil, result.Error
	}

	return driver, nil;
}



func (r *driverrepoImpl) GetDriverSmsCode(ctx context.Context,PhoneNumber, code string) (*models.DriverSmsCode, error) {

	var driver models.DriverSmsCode
	result := r.db.WithContext(ctx).Table(driverSmsTable).First(&driver, "code = ? AND phone_number = ?", code, PhoneNumber)
	if result.Error != nil {
		return nil, result.Error
	}
	return &driver, nil
}
