package repository

import (
	"time"

	"github.com/flanksource/confighub/db/models"
	"gorm.io/gorm"
)

// DBRepo should satisfy the database repository interface
type DBRepo struct {
	db *gorm.DB
}

// NewRepo is the factory function for the database repo instance
func NewRepo(db *gorm.DB) Database {
	return &DBRepo{
		db: db,
	}
}

// GetOneConfigItem returns a single config item result
func (d *DBRepo) GetOneConfigItem(extID string) (*models.ConfigItem, error) {

	ci := models.ConfigItem{}
	if err := d.db.First(&ci, "external_id = ?", extID).Error; err != nil {
		return nil, err
	}

	return &ci, nil
}

// CreateConfigItem inserts a new config item row in the db
func (d *DBRepo) CreateConfigItem(ci *models.ConfigItem) error {

	ci.CreatedAt = time.Now().UTC()

	if err := d.db.Create(ci).Error; err != nil {
		return err
	}

	return nil
}

// UpdateAllFieldsConfigItem updates all the fields of a given config item row
func (d *DBRepo) UpdateAllFieldsConfigItem(ci *models.ConfigItem) error {

	ci.UpdatedAt = time.Now().UTC()

	if err := d.db.Save(ci).Error; err != nil {
		return err
	}

	return nil
}

// CreateConfigChange inserts a new config change row in the db
func (d *DBRepo) CreateConfigChange(cc *models.ConfigChange) error {

	cc.CreatedAt = time.Now().UTC()

	if err := d.db.Create(cc).Error; err != nil {
		return err
	}

	return nil
}