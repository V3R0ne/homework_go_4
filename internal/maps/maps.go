package maps

import (
	"math"
	"slices"
)

type User struct {
	ID     int
	Name   string
	City   string
	Active bool
	Tags   []string
}

type Product struct {
	SKU      string
	Quantity int
	Price    int
}

func CountWords(words []string) map[string]int {
	// TODO: посчитать, сколько раз встречается каждое слово.
	// Вернуть новую готовую к записи map. Регистр и пустая строка имеют значение.
	countWords := make(map[string]int)
	for _, word := range words {
		countWords[word]++
	}
	return countWords
}

func GetOrDefault(values map[string]int, key string, fallback int) int {
	// TODO: вернуть значение существующего ключа или fallback, если ключ отсутствует.
	// Нулевое значение по существующему ключу должно возвращаться как обычное значение.
	if value, ok := values[key]; ok {
		return value
	}
	return fallback
}

func DeleteAndReport(values map[string]int, key string) bool {
	// TODO: удалить ключ и сообщить, существовал ли он до вызова.
	// Для отсутствующего ключа, пустой и nil map вернуть false без panic.
	if _, ok := values[key]; ok {
		delete(values, key)
		return true
	}
	return false
}

func MergeCounters(left, right map[string]int) map[string]int {
	// TODO: объединить два счётчика в новой map, складывая значения общих ключей.
	// Исходные map после вызова должны остаться без изменений.
	newMap := make(map[string]int)
	for key, value := range left {
		newMap[key] = value
	}
	for key, value := range right {
		newMap[key] += value
	}
	return newMap
}

func KeysByValueAtLeast(values map[string]int, min int) []string {
	// TODO: вернуть ключи, значения которых не меньше min.
	// Результат должен иметь стабильный алфавитный порядок.
	result := make([]string, 0)
	for key, value := range values {
		if value >= min {
			result = append(result, key)
		}
	}
	slices.Sort(result)
	return result
}

func MaxKeyByValue(values map[string]int) (string, bool) {
	// TODO: найти ключ с максимальным значением.
	// Для пустого входа вернуть "", false; при равенстве выбрать меньший ключ по алфавиту.
	if len(values) == 0 {
		return "", false
	}

	keyWithMaxValue := ""
	maxValue := math.MinInt
	for key, value := range values {
		if value > maxValue {
			maxValue = value
			keyWithMaxValue = key
		} else if value == maxValue {
			keyWithMaxValue = min(keyWithMaxValue, key)
		}
	}
	return keyWithMaxValue, true
}

func BuildUserIndex(users []User) map[int]User {
	// TODO: построить индекс пользователей по ID.
	// Если один ID встречается несколько раз, в результате должен остаться последний пользователь.
	userIndex := make(map[int]User)
	for _, user := range users {
		userIndex[user.ID] = user
	}
	return userIndex
}

func GroupActiveUsersByCity(users []User) map[string][]User {
	// TODO: сгруппировать только активных пользователей по городу.
	// Порядок пользователей внутри каждого города должен совпадать с исходным слайсом.
	activeUsersByCity := make(map[string][]User)
	for _, user := range users {
		if user.Active {
			activeUsersByCity[user.City] = append(activeUsersByCity[user.City], user)
		}
	}
	return activeUsersByCity
}

func CountTags(users []User) map[string]int {
	// TODO: посчитать все появления тегов у всех пользователей.
	// Повторяющийся тег у одного пользователя также считается отдельным появлением.
	countTags := make(map[string]int)
	for _, user := range users {
		for _, tag := range user.Tags {
			countTags[tag]++
		}
	}
	return countTags
}

func BuildInventory(products []Product) map[string]Product {
	// TODO: построить склад по SKU.
	// При повторном SKU оставить последнюю запись из входного слайса.
	inventory := make(map[string]Product)
	for _, product := range products {
		inventory[product.SKU] = product
	}
	return inventory
}

func ReserveStock(inventory map[string]Product, sku string, count int) bool {
	// TODO: уменьшить остаток существующего товара на положительное count, если товара хватает.
	// При неуспехе вернуть false и не менять склад.
	if count <= 0 {
		return false
	}

	p, ok := inventory[sku]
	if !ok || p.Quantity < count {
		return false
	}

	p.Quantity -= count
	inventory[sku] = p
	return true
}

func Restock(inventory map[string]*Product, sku string, count int) bool {
	// TODO: увеличить остаток товара, хранящегося в map как указатель.
	// Ключ должен существовать, указатель не должен быть nil, count должен быть положительным.
	if count <= 0 {
		return false
	}

	p, ok := inventory[sku]
	if !ok || p == nil {
		return false
	}

	p.Quantity += count
	return true
}

func LowStockSKUs(inventory map[string]Product, limit int) []string {
	// TODO: вернуть SKU товаров с остатком не больше limit.
	// Результат должен иметь стабильный алфавитный порядок.
	productsWithQuantityLessLimit := make([]string, 0)
	for sku, product := range inventory {
		if product.Quantity <= limit {
			productsWithQuantityLessLimit = append(productsWithQuantityLessLimit, sku)
		}
	}

	slices.Sort(productsWithQuantityLessLimit)
	return productsWithQuantityLessLimit
}

func InventoryValue(inventory map[string]Product) int {
	// TODO: посчитать общую стоимость склада как сумму Quantity * Price.
	// Пустой и nil склад имеют стоимость 0.
	if len(inventory) == 0 {
		return 0
	}

	allSum := 0
	for _, product := range inventory {
		allSum += product.Quantity * product.Price
	}
	return allSum
}

func ApplyPriceUpdates(inventory map[string]Product, updates map[string]int) int {
	// TODO: применить положительные новые цены только к существующим товарам.
	// Вернуть число реально обновлённых товаров; неизвестные SKU и неположительные цены пропустить.
	countUpdates := 0
	for sku, price := range updates {
		if price <= 0 {
			continue
		}
		p, ok := inventory[sku]
		if !ok {
			continue
		}
		p.Price = price
		inventory[sku] = p
		countUpdates++
	}
	return countUpdates
}
