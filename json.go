package jsonmap

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func printBool(m bool, k string, format string) {
	if format == JSON_FORMAT {
		log.Printf("\n{\n\t%v : %v\n}", k, m)
	} else if format == VALUE_FORMAT {
		log.Printf("Key: %v", k)
		log.Printf("Value: %v", m)
	}

}

func printInt(m int, k string, format string) {
	if format == JSON_FORMAT {
		log.Printf("\n{\n\t%v : %v\n}", k, m)
	} else if format == VALUE_FORMAT {
		log.Printf("Key: %v", k)
		log.Printf("Value: %v", m)
	}
}

func printString(m string, k string, format string) {
	if format == JSON_FORMAT {
		log.Printf("\n{\n\t%v : %v\n}", k, m)
	} else if format == VALUE_FORMAT {
		log.Printf("Key: %v", k)
		log.Printf("Value: %v", m)
	}
}

func printSlice(m []string, k string, format string) {

	if format == JSON_FORMAT {

		log.Printf("\n{\n\t%v : [\n", k)

		for index, value := range m {
			if index != len(m)-1 {

				fmt.Printf("\t\t%v,\n", value)
			} else {
				fmt.Printf("\t\t%v\n\t]\n}\n", value)
			}
		}

	} else if format == VALUE_FORMAT {

		log.Printf("Key: %v", k)
		log.Printf("Loop")

		for index, value := range m {
			log.Printf("Value %v: %v", index, value)
		}
	}

}

func (m *Map) PrintSortOrder() {
	var keys []string
	for k := range m.m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Printf("%T", m.m[k])
		fmt.Println("")
		// in case of map value, call recursive print function
		// for each value on the map
		fmt.Println("Key:", k, "Value:", m.m[k])
	}
}

func (m *Map) PrintReverseOrder() {
	var keys []string
	for k := range m.m {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(keys)))

	for _, k := range keys {
		fmt.Printf("%T", m.m[k])
		fmt.Println("")
		// in case of map value, call recursive print function
		// for each value on the map
		fmt.Println("Key:", k, "Value:", m.m[k])
	}
}

// Print Json contents
func (m *Map) loop(key string) {
	if key == "" {
		log.Printf("All map's values are printed.")
	} else {
		log.Printf("Key: %v", key)
		log.Printf("Values:\n")
	}

	for i, k := range m.m {
		log.Printf("Tipo %v", reflect.TypeOf(k).Kind())
		if reflect.TypeOf(k).Kind() == reflect.Map {
			log.Printf("Key: %v", i)
			log.Printf("Loop")
			temp := CreateJson(Maps{})
			temp.m = k.(Maps)
			temp.loop(key)
		} else if reflect.TypeOf(k).Kind() == reflect.Slice {
			printSlice(k.([]string), i, VALUE_FORMAT)
		} else if reflect.TypeOf(k).Kind() == reflect.Bool {
			printBool(k.(bool), i, VALUE_FORMAT)
		} else if reflect.TypeOf(k).Kind() == reflect.String {
			printString(k.(string), i, VALUE_FORMAT)
		} else if reflect.TypeOf(k).Kind() == reflect.Int {
			printInt(k.(int), i, VALUE_FORMAT)
		}
	}
}

func (m *Map) IsMap(key string) bool {
	if reflect.TypeOf(m.Get(key)).Kind() == reflect.Map &&
		len(reflect.ValueOf(m.Get(key)).MapKeys()) > 0 {
		return true
	}

	return false
}

func (m *Map) IsSlice(key string) bool {
	if reflect.TypeOf(m.Get(key)).Kind() == reflect.Slice &&
		reflect.ValueOf(m.Get(key)).Len() > 0 {
		return true
	}

	return false
}

func (m *Map) IsInt(key string) bool {
	if reflect.TypeOf(m.Get(key)).Kind() == reflect.Int &&
		reflect.ValueOf(m.Get(key)).Int() > 0 {
		return true
	}

	return false
}

func (m *Map) IsString(key string) bool {
	if reflect.TypeOf(m.Get(key)).Kind() == reflect.String &&
		reflect.ValueOf(m.Get(key)).String() > " " {
		return true
	}

	return false
}

func (m *Map) Exist(key string) bool {

	if m.IsMap(key) {
		return true
	} else if m.IsSlice(key) {
		return true
	} else if m.IsInt(key) {
		return true
	} else if m.IsString(key) {
		return true
	}

	return false
}

func (m *Map) Search(key string, format string) {

	if m.IsMap(key) {

		mu := CreateJson(Maps{})
		mu.m = m.Get(key).(Maps)
		log.Printf("Key to search: %v\n", key)

		if format == JSON_FORMAT {
			mu.IndentJsonPrint()
		} else if format == VALUE_FORMAT {
			mu.loop(key)
		}

	} else if m.IsSlice(key) {
		mu := m.Get(key).([]string)

		printSlice(mu, key, format)
	} else if m.IsInt(key) {
		mu := m.Get(key).(int)

		printInt(mu, key, format)
	} else if m.IsString(key) {
		mu := m.Get(key).(string)

		printString(mu, key, format)
	}
}

func CreateJson(elem Maps) Map {
	mappazza := Map{}
	mappazza.m = make(Maps)
	mappazza.m = Maps{}

	if reflect.ValueOf(elem).Len() > 0 {
		mappazza.m = elem
	}

	return mappazza
}

func (m *Map) Get(key string) interface{} {
	n := CreateJson(Maps{})

	if key == "" {
		return m.m
	}

	if strings.Contains(key, " ") {
		s := strings.Split(key, " ")

		for k, _ := range m.m {
			if k == s[0] {
				if reflect.TypeOf(m.m[k]).Kind() == reflect.Map {
					q := strings.Join(s[1:len(s)], " ")
					n.m = m.m[k].(Maps)
					return n.Get(q)
				}
			}
		}
	} else {
		for k, v := range m.m {
			if k == key {
				return v
			}
		}
	}
	return n.m
}

func (m *Map) assign(key string, values ...interface{}) {
	if len(values) > 1 {
		m.m[key] = values
	} else {
		m.m[key] = values[0]
	}
}

func (m *Map) Set(key string, values ...interface{}) {

	if !strings.Contains(key, " ") {

		m.assign(key, values[0])
	} else {

		s := strings.Split(key, " ")
		stringa := strings.Join(s, " ")
		for i, k := range s {
			if k == "" {
				err := strings.Join([]string{"Value of index", strconv.FormatInt(int64(i), 16), "of key [", stringa, "] ,not valid"}, " ")
				panic(err)
			}
		}

		mappa10 := CreateJson(Maps{})

		for i := len(s) - 1; i >= 0; i-- {
			mappa11 := CreateJson(Maps{})

			if m.Exist(strings.Join(s[0:i], " ")) {
				for k, v := range m.Get(strings.Join(s[0:i], " ")).(Maps) {
					mappa11.Set(k, v)
				}
			}

			if i != len(s)-1 {
				mappa11.assign(s[i], mappa10.m)

			} else {
				mappa11.assign(s[len(s)-1], values[0])
			}

			mappa10.m = mappa11.m
		}

		m.m = mappa10.m
	}
}

func (m *Map) Delete(key string) {
	if !strings.Contains(key, " ") {
		delete(m.m, key)
	} else {
		s := strings.Split(key, " ")

		mappa1 := CreateJson(Maps{})

		if m.Exist(s[len(s)-2]) {
			if m.IsMap(s[len(s)-2]) {
				mappa1.m = m.Get(s[len(s)-2]).(Maps)

				delete(mappa1.m, s[len(s)-1])

				m.Set(s[len(s)-2], Maps{})

				for k, v := range mappa1.m {

					q := strings.Join(s[0:len(s)-1], " ")

					if len(q) > 0 {
						q = strings.Join([]string{q, k}, " ")
					} else {
						q = k
					}

					m.Set(q, v)
				}
			}

		}
	}
}

func (m *Map) IndentJsonPrint() {
	d, err := json.MarshalIndent(m.m, " ", "\t")

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("map: %v", string(d))
}
