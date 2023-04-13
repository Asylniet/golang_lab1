package models

import (
	"lab1/db"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Store struct {
	Items []Item
}

// constructor
func InitializeStore() *Store {
	return &Store{
		Items: []Item{},
	}
}

func (store *Store) GetDataFromDB() {
	var items []Item
	db.DB.Find(&items)
	store.Items = items
}

// toString
func Print(items []Item) {
	for i := 0; i < len(items); i++ {
		items[i].Print(i)
	}
}

func (store *Store) Print() {
	Print(store.Items)
}

// Methods
func (store *Store) Add(item *Item) {
	db.DB.Create(item)
	store.Items = append(store.Items, *item)
}

func (store *Store) getItemByID(id int) (*Item, int) {
	for i := 0; i < len(store.Items); i++ {
		item := store.Items[i]

		if item.ID == id {
			return &item, i
		}
	}
	return nil, len(store.Items)
}

func (store *Store) Remove(id int) {
	_, index := store.getItemByID(id)
	items := store.Items
	if index < len(store.Items) {
		store.Items = append(items[:index], items[index+1:]...)
		db.DB.Delete(&Item{}, id)
	} else {
		panic("No item found to delete")
	}
}

func (store *Store) Search(text string) []Item {
	var items []Item
	for i := 0; i < len(store.Items); i++ {
		item := store.Items[i]

		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(text)) {
			items = append(items, item)
		}
	}
	return items
}

func SortItems(items []Item, filter string, order string) []Item {
	if filter != "rating" && filter != "price" {
		panic("No such filter\nTry again")
		return []Item{}
	}
	if order != "asc" && order != "desc" {
		panic("No such order\nTry again")
		return []Item{}
	}
	sort.Slice(items, func(i, j int) bool {
		if filter == "rating" {
			if order == "asc" {
				return items[i].Rating < items[j].Rating // ASC
			} else {
				return items[i].Rating > items[j].Rating // DESC
			}
		} else {
			if order == "asc" {
				return items[i].Price < items[j].Price // ASC
			} else {
				return items[i].Price > items[j].Price // DESC
			}
		}
	})

	return items
}

func FilterItems(items []Item, filterStr string) (result []Item) {
	if len(filterStr) == 0 {
		return items
	}

	filters := strings.Split(filterStr, " ")
	for i := 0; i < len(filters); i++ {
		fltr, compare, number := parseText(filters[i])
		for _, item := range items {
			if item.filterBy(fltr, compare, number) {
				result = append(result, item)
			}
		}
		items = result
		if i < len(filters)-1 {
			result = []Item{}
		}
	}
	return
}

func parseText(txt string) (string, string, string) {
	fltrRegex := regexp.MustCompile("[a-z]+")
	compareRegex := regexp.MustCompile("[<>=!]+")
	numRegex := regexp.MustCompile("[0-9]+(\\.[0-9]+)*")

	fltr := fltrRegex.FindStringSubmatch(txt)[0]
	compare := compareRegex.FindStringSubmatch(txt)[0]
	number := numRegex.FindStringSubmatch(txt)[0]
	return fltr, compare, number
}

func (item Item) filterBy(fltr string, compare string, number string) bool {
	if fltr != "rating" && fltr != "price" {
		panic("No such filter")
	}
	if fltr == "rating" {
		return filterRating(item, compare, number)
	} else if fltr == "price" {
		return filterPrice(item, compare, number)
	}
	return false
}

func filterRating(item Item, compare string, number string) bool {
	num, _ := strconv.ParseFloat(number, 32)
	switch compare {
	case ">":
		if item.Rating > float32(num) {
			return true
		}
		return false
	case "<":
		if item.Rating < float32(num) {
			return true
		}
		return false
	case ">=":
		if item.Rating >= float32(num) {
			return true
		}
		return false
	case "<=":
		if item.Rating <= float32(num) {
			return true
		}
		return false
	case "!=":
		if item.Rating != float32(num) {
			return true
		}
		return false
	}
	return false
}

func filterPrice(item Item, compare string, number string) bool {
	num, _ := strconv.ParseInt(number, 10, 32)
	switch compare {
	case ">":
		if item.Price > int(num) {
			return true
		}
		return false
	case "<":
		if item.Price < int(num) {
			return true
		}
		return false
	case ">=":
		if item.Price >= int(num) {
			return true
		}
		return false
	case "<=":
		if item.Price <= int(num) {
			return true
		}
		return false
	case "!=":
		if item.Price != int(num) {
			return true
		}
		return false
	}
	return false
}
