package main

import (
	"fmt"
	"sort"
)

// Визначення структури з одним полем
type Item struct {
	ID int
}

// Функція для пошуку унікальних елементів та їх сортування
func uniqueAndSorted(items []Item) []Item {
	uniqueMap := make(map[int]bool)
	var uniqueItems []Item

	// Знаходимо унікальні елементи
	for _, item := range items {
		if _, exists := uniqueMap[item.ID]; !exists {
			uniqueMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}

	// Сортуємо унікальні елементи
	sort.Slice(uniqueItems, func(i, j int) bool {
		return uniqueItems[i].ID < uniqueItems[j].ID
	})

	return uniqueItems
}

func main() {
	// Приклад використання функції
	items := []Item{{3}, {2}, {1}, {2}}
	fmt.Println("Початковий слайс:", items)

	result := uniqueAndSorted(items)
	fmt.Println("Унікальні та відсортовані елементи:", result)
}
