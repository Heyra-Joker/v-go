/*
@File    :   util/exists.go
@Time    :   2021/12/10 11:42:03
@Author  :   Joker
@Contact :   jokermonster030@gmail.com
@Desc    :   None
🤡
code is far away from bugs with the god animal protecting
I love animals. They taste delicious.
🤡
*/

package util

import (
	"fmt"
	"os"
	"reflect"
)

// ItemContainer check item whether in items
func ItemContainer(item, items interface{}) bool {
	kind := reflect.TypeOf(item).Kind()
	itemsList := reflect.ValueOf(items)

	// set if it's Ptr
	if itemsList.Kind() == reflect.Ptr {
		itemsList = itemsList.Elem()
	}

	for i := 0; i < itemsList.Len(); i++ {
		subValue := itemsList.Index(i)
		sKind := subValue.Kind()
		if sKind != kind {
			panic(fmt.Sprintf("item kind: %#+v not equal items sub value: %#+v", kind, sKind))
		}

		if reflect.DeepEqual(item, subValue.Interface()) {
			return true
		}
	}

	return false
}

// PathExists check target whether exists
func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	return true
}
