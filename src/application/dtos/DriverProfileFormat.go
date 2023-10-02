package dtos

import (
	"time"

	pb "github.com/maxmanovb/logistics-staff-service/src/application/protos/logistics_staff"
	"github.com/maxmanovb/logistics-staff-service/src/domain/models"
)


func NewDriverProfileFormat(driver *models.DriverProfile) *pb.DriverProfile {
	return &pb.DriverProfile{
		DriverId: driver.DriverID,
		Name: driver.Name,
		PhoneNumber: driver.PhoneNumber,
		ImageUrl: driver.ImageUrl,
		Active: driver.Active,
		CreatedAt: driver.CreatedAt.Format(time.RFC3339),
		UpdatedAt: driver.UpdatedAt.Format(time.RFC3339),
	}
}
