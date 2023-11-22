package openapi

import "gorm.io/gorm"

type CartItem struct {
	gorm.Model

	MenuItem   MenuItem
	MenuItemID int32
}
