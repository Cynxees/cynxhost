package entity

import (
	"cynxhost/internal/constant/types"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type TblServerTemplate struct {
	Id          int       `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate time.Time `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate time.Time `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name        string    `gorm:"size:255;not null" json:"name" visibility:"1"`
	Description *string   `gorm:"type:text" json:"description" visibility:"2"`
	MinimumRam  int       `gorm:"not null" json:"minimum_ram" visibility:"1"`
	MinimumCpu  int       `gorm:"not null" json:"minimum_cpu" visibility:"1"`
	MinimumDisk int       `gorm:"not null" json:"minimum_disk" visibility:"1"`
	ImageUrl    string    `gorm:"size:255" json:"image_url" visibility:"1"`
	ScriptId    int       `gorm:"not null" json:"script_id" visibility:"1"`
	Script      TblScript `gorm:"foreignKey:ScriptId" json:"script" visibility:"1"`
}

type TblServerTemplateCategory struct {
	Id               int                `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate      time.Time          `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate      time.Time          `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name             string             `gorm:"size:255;not null" json:"name" visibility:"1"`
	Description      *string            `gorm:"type:text" json:"description" visibility:"2"`
	ParentId         int                `gorm:"default:null" json:"parent_id"`
	ImagePath        *string            `gorm:"size:255" json:"image_path" visibility:"10"`
	ServerTemplateId *int               `gorm:"default:null" json:"server_template_id"`
	ServerTemplate   *TblServerTemplate `gorm:"foreignKey:ServerTemplateId" json:"server_template"`
}

type TblScript struct {
	Id             int       `gorm:"primaryKey" json:"id" visibility:"1"`
	CreatedDate    time.Time `gorm:"autoCreateTime" json:"created_date" visibility:"1"`
	UpdatedDate    time.Time `gorm:"autoUpdateTime" json:"updated_date" visibility:"1"`
	Name           string    `gorm:"size:255;not null" json:"name" visibility:"1"`
	Variables      Variables `gorm:"type:json" json:"variables" visibility:"2"` // Custom type for JSON handling
	SetupScript    string    `gorm:"type:text;not null" json:"setup_script" visibility:"2"`
	StartScript    string    `gorm:"type:text;not null" json:"start_script" visibility:"2"`
	StopScript     string    `gorm:"type:text;not null" json:"stop_script" visibility:"2"`
	ShutdownScript string    `gorm:"type:text;not null" json:"shutdown_script" visibility:"2"`
	ConfigPath     string    `gorm:"size:255;not null" json:"config_path" visibility:"2"`
}

type ScriptVariable struct {
	Name    string                   `json:"name"`
	Type    types.ScriptVariableType `json:"type"`
	Content []ScriptVariableContent  `json:"content"`
}

type ScriptVariableContent struct {
	Name  string                 `json:"name"`
	Value map[string]interface{} `json:"value"`
}

// Custom type for handling JSON serialization
type Variables []ScriptVariable

// Value implements the `driver.Valuer` interface to convert the struct to JSON.
func (v Variables) Value() (driver.Value, error) {
	return json.Marshal(v)
}

// Scan implements the `sql.Scanner` interface to convert JSON to the struct.
func (v *Variables) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, v)
}
