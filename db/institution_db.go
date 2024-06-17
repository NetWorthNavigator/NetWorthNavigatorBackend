package db

import (
	"log"

	"github.com/NetWorthNavigator/NetWorthNavigatorBackend/models"
	"gorm.io/gorm"
)

type InstitutionDB struct {
    DBClient *gorm.DB
}

func NewInstitutionDB(DBClient *gorm.DB) *InstitutionDB {
    return &InstitutionDB{DBClient: DBClient}
}

// CreateInstitution creates a new Institution record.
func (idb *InstitutionDB) CreateInstitution(institution models.Institution) error {
    // Check if the institution already exists with the same InstitutionID
    var existingInstitution models.Institution
    if err := idb.DBClient.Where("institution_id = ?", institution.InstitutionID).First(&existingInstitution).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            // Institution does not exist, proceed with creation
            result := idb.DBClient.Create(&institution)
            if result.Error != nil {
                log.Printf("Error creating Institution: %v", result.Error)
                return result.Error
            }
            return nil
        }
        // Other error occurred
        return err
    }
    // Institution with the same InstitutionID already exists, return nil or a custom error
    log.Printf("Institution with InstitutionID %s already exists", institution.InstitutionID)
    return nil // or return an error indicating the institution already exists
}

// GetInstitutions retrieves all institutions.
func (idb *InstitutionDB) GetInstitutions() ([]models.Institution, error) {
    var institutions []models.Institution
    result := idb.DBClient.Find(&institutions)
    if result.Error != nil {
        log.Printf("Error retrieving institutions: %v", result.Error)
        return nil, result.Error
    }
    return institutions, nil
}

// GetInstitution retrieves a single institution based on InstitutionID.
func (idb *InstitutionDB) GetInstitution(institutionID string) (*models.Institution, error) {
    var institution models.Institution
    result := idb.DBClient.Where("institution_id = ?", institutionID).First(&institution)
    if result.Error != nil {
        log.Printf("Error retrieving institution with InstitutionID %s: %v", institutionID, result.Error)
        return nil, result.Error
    }
    return &institution, nil
}

// DeleteInstitution deletes an institution based on InstitutionID.
func (idb *InstitutionDB) DeleteInstitution(institutionID string) error {
    result := idb.DBClient.Delete(&models.Institution{}, "institution_id = ?", institutionID)
    if result.Error != nil {
        log.Printf("Error deleting institution with InstitutionID %s: %v", institutionID, result.Error)
        return result.Error
    }
    return nil
}

// UpdateInstitution updates an existing institution's details.
func (idb *InstitutionDB) UpdateInstitution(institution models.Institution) error {
    result := idb.DBClient.Save(&institution)
    if result.Error != nil {
        log.Printf("Error updating institution with InstitutionID %s: %v", institution.InstitutionID, result.Error)
        return result.Error
    }
    return nil
}