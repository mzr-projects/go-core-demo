package shopping

import "com.mt/go-demo/shopping/db"

func NameCheck(itemId int) (string, bool) {
	item := db.LoadItem(itemId)
	if item != nil {
		return "", false
	}
	return item.Name, true
}
