/*
 * Continous Food Delievery
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddCartItem - Add a menu item a cart
func AddCartItem(c *gin.Context) {
	var menuItem MenuItem

	if err := c.BindJSON(&menuItem); err != nil {
		return
	}

	cartItem := CartItem{MenuItem: menuItem}
	DB.Create(&cartItem)

	DB.Debug().AutoMigrate(&CartItem{})

	c.JSON(http.StatusOK, gin.H{"data": cartItem})
}

// DeleteCartItem - Remove item from cart
func DeleteCartItem(c *gin.Context) {
	id := c.Param("itemId")
	result, _ := strconv.Atoi(id)

	var cartItems []CartItem
	DB.Find(&cartItems)

	for _, ci := range cartItems {
		x := int(ci.MenuItem.ID)
		if x == result {
			findresult := DB.Find(&ci.MenuItem.ID).Where("ImageId = ?", id).Delete(&ci)
			fmt.Println(findresult)
			DB.Save(&ci)
			c.JSON(http.StatusOK, "Cart item Deleted")
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "cart item not found"})
}

// ListCart - List all cart items
func ListCart(c *gin.Context) {
	var cartItems []CartItem
	findresult := DB.Find(&cartItems)
	if findresult == nil {
		c.JSON(http.StatusNoContent, cartItems)
	}

	c.JSON(http.StatusOK, cartItems)
}

func GetMenuItem(menuitem int32) (cartitems []CartItem, err error) {
	return cartitems, DB.Where("menuitem_id = ?", menuitem).Set("gorm:auto_preload", true).Find(&cartitems).Error
}
