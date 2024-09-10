package cache

import "fmt"

type Cacheable interface {
	Set(string, interface{})
	Get(string)
	Delete(string)
}

type Cache map[string]interface{}

func New() Cache {
	return map[string]interface{}{}
}

func (c *Cache) Set(key string, val interface{}) {
	(*c)[key] = val
}

func (c *Cache) Get(key string) interface{} {
	if c.check(key) {
		return (*c)[key]
	}

	return nil
}

func (c *Cache) Delete(key string) {
	if c.check(key) {
		delete(*c, key)
	}
}

func (c *Cache) check(key string) bool {
	_, ok := (*c)[key]

	if ok {
		return true
	}

	defer HandlePanic()

	panic(fmt.Sprint("Ошибка, нет такого юзера: ", key, " !"))

}

func HandlePanic() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
