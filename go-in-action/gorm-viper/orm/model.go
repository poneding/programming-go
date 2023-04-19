package orm

import (
	"gorm-viper/util"
	"time"

	"gorm.io/gorm"
)

type PureModel struct {
	Id string `gorm:"type:char(36);primaryKey" json:"id"`
}

type Model struct {
	Id        string    `gorm:"type:char(36);primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type IntegerModel struct {
	Id        uint64    `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (m *PureModel) BeforeCreate(tx *gorm.DB) error {
	m.Id = util.GeneratePureUUID()
	return nil
}

func (m *Model) BeforeCreate(tx *gorm.DB) error {
	m.Id = util.GeneratePureUUID()
	return nil
}
