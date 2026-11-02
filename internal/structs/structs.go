package structs

import (
	"fmt"
	"slices"
	"unsafe"
)

type User struct {
	ID         int
	Name       string
	Active     bool
	LoginCount int
	Tags       []string
}

func UpdateUserName(users map[int]User, id int, name string) bool {
	// TODO: изменить имя существующего пользователя, хранящегося в map как значение.
	// Если пользователя нет, вернуть false и не добавлять новый ключ.
	u, ok := users[id]
	if !ok {
		return false
	}
	u.Name = name
	users[id] = u
	return true
}

func ActivateUser(users map[int]User, id int) bool {
	// TODO: сделать существующего пользователя активным.
	// Если пользователя нет, вернуть false.
	u, ok := users[id]
	if !ok {
		return false
	}
	u.Active = true
	users[id] = u
	return true
}

func IncrementLoginCount(users map[int]User, id int) bool {
	// TODO: увеличить LoginCount существующего пользователя на единицу.
	// Если пользователя нет, вернуть false.
	u, ok := users[id]
	if !ok {
		return false
	}
	u.LoginCount++
	users[id] = u
	return true
}

func CopyUsers(users map[int]User) map[int]User {
	// TODO: вернуть независимую копию map и вложенных слайсов Tags.
	// Изменения любой части результата не должны менять исходные данные.
	result := make(map[int]User, len(users))
	for id, u := range users {
		u.Tags = slices.Clone(u.Tags)
		result[id] = u
	}
	return result
}

func BuildPointerIndex(users []User) map[int]*User {
	// TODO: построить индекс ID -> *User из независимых копий входных пользователей.
	// Указатели разных ID не должны вести к одному объекту или менять исходный слайс.
	index := make(map[int]*User, len(users))
	for _, user := range users {
		u := user
		u.Tags = slices.Clone(user.Tags)
		index[u.ID] = &u
	}
	return index
}

func RenameThroughPointer(users map[int]*User, id int, name string) bool {
	// TODO: изменить имя объекта по существующему ненулевому указателю.
	// Для отсутствующего ключа или nil-значения вернуть false.
	u, ok := users[id]
	if !ok || u == nil {
		return false
	}
	u.Name = name
	return true
}

func (u User) DisplayName() string {
	// TODO: вернуть строку "ID:Name".
	return fmt.Sprintf("%d:%s", u.ID, u.Name)
}

type Admin struct {
	User
	Permissions []string
}

func (a Admin) DisplayName() string {
	// TODO: вернуть строку "admin:ID:Name", перекрывая одноимённый метод встроенного User.
	return "admin:" + a.User.DisplayName()
}

func (a Admin) HasPermission(permission string) bool {
	// TODO: проверить точное наличие permission в списке администратора.
	// Регистр, пробелы и повторения не нормализуются.
	return slices.Contains(a.Permissions, permission)
}

func RenameEmbeddedUser(admin *Admin, name string) {
	// TODO: изменить Name у встроенного User.
	// Для nil admin функция должна безопасно завершиться.
	if admin == nil {
		return
	}
	admin.User.Name = name
}

type Account struct {
	ID      int
	Balance int
}

type PremiumAccount struct {
	Account
	Bonus int
}

func (p PremiumAccount) Total() int {
	// TODO: вернуть сумму основного баланса и бонуса.
	return p.Balance + p.Bonus
}

func PremiumLabel(p PremiumAccount) string {
	// TODO: вернуть строку "ID:Balance total=Total".
	return fmt.Sprintf("%d:%d total=%d", p.ID, p.Balance, p.Total())
}

type BadLayout struct {
	Active   bool
	Amount   int64
	Verified bool
	Retries  int32
}

type GoodLayout struct {
	Amount   int64
	Retries  int32
	Active   bool
	Verified bool
}

func UserValueSize(user User) uintptr {
	// TODO: вернуть фактический размер переданного значения User в байтах.
	return unsafe.Sizeof(user)
}

func UserPointerSize(user *User) uintptr {
	// TODO: вернуть фактический размер переменной-указателя на User в байтах.
	// Результат не зависит от того, nil указатель или нет.
	return unsafe.Sizeof(user)
}

func LayoutSizes() (bad, good, saved uintptr) {
	// TODO: вернуть размеры BadLayout и GoodLayout, затем число сэкономленных байт.
	bad = unsafe.Sizeof(BadLayout{})
	good = unsafe.Sizeof(GoodLayout{})
	saved = bad - good
	return bad, good, saved
}

var _ uintptr = unsafe.Sizeof(User{})
