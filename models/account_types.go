package models

import (
	"time"

	"github.com/jgreat/fog-wallet/api"
	"gorm.io/gorm"
)

// Recreate gorm.Model type without the ID. ID will be defined by openapi model.
type base struct {
	// ID        uint `gorm:"primarykey"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

/**
Account - An account in the wallet.
An Account is associated with one AccountKey, containing a View keypair and a Spend keypair.
**/
type Account struct {
	// base fields
	base

	// Inherit types from generated openapi.
	api.Account
}
