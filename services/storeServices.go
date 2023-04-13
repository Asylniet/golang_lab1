package services

import (
	"gorm.io/gorm"
	"lab1/db"
	"lab1/models"
)

func GetItems() ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := db.DB.Find(&items)
	return items, result
}

func GetItem(id string) (models.Item, *gorm.DB) {
	var item models.Item
	result := db.DB.First(&item, id)
	return item, result
}

func AddItem(name string, price int) *gorm.DB {
	item := models.NewItem(name, price, 0)
	result := db.DB.Create(&item)
	return result
}

func RemoveItem(id string) *gorm.DB {
	result := db.DB.Delete(&models.Item{}, id)
	return result
}

func UpdateItem(itemQuery models.Item, id string) (*gorm.DB, string) {
	var item models.Item
	result := db.DB.First(&item, id)

	if item == (models.Item{}) {
		return result, "No item found"
	}

	item.Name = itemQuery.Name
	item.Price = itemQuery.Price
	result = db.DB.Save(&item)

	return result, ""
}

func SearchItems(name, filter, sort string) ([]models.Item, *gorm.DB) {
	var items []models.Item
	result := db.DB.Where("LOWER(name) LIKE ?", "%"+name+"%").Order(sort).Find(&items)
	items = models.FilterItems(items, filter)
	return items, result
}
