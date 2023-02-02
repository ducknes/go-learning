package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

//	------------------JSON------------------

// type User struct {
// 	Name      string `json:"name"`
// 	Age       int    `json:"age"`
// 	IsBlocked bool   `json:"is_blocked"`
// 	Books     []Book `json:"books"`
// }

// type Book struct {
// 	Name string `json:"name"`
// 	Year int    `json:"year"`
// }

// func main() {
// 	//Serialize()

// 	byt := []byte(`{"name":"Ilya","age":19,"is_blocked":false,"books":[{"name":"BN1","year":2003},{"name":"BN2","year":1950}]}`)
// 	var dat User //map[string]interface{}

// 	if err := json.Unmarshal(byt, &dat); err != nil {
// 		panic(err)
// 	}
// 	//fmt.Println(dat["books"].([]interface{})[0].(map[string]interface{})["name"])
// 	fmt.Println(dat.Books[0].Year)
// }

// func Serialize() {
// 	var books []Book
// 	book1 := Book{
// 		Name: "BN1",
// 		Year: 2003,
// 	}
// 	book2 := Book{
// 		Name: "BN2",
// 		Year: 1950,
// 	}
// 	books = append(books, book1, book2)

// 	sv := User{
// 		Name:      "Ilya",
// 		Age:       19,
// 		IsBlocked: false,
// 		Books:     books,
// 	}

// 	boolVar, _ := json.Marshal(sv)
// 	fmt.Println(string(boolVar))
// }

//	------------------GJSON, SJSON------------------

func main() {
	json := `
	{
		"name": {"first": "Tom", "last": "Anderson"},
		"age":37,
		"children": ["Sara","Alex","Jack"],
		"fav.movie": "Deer Hunter",
		"friends": [
		  {"first": "Dale", "last": "Murphy", "age": 44, "nets": ["ig", "fb", "tw"]},
		  {"first": "Roger", "last": "Craig", "age": 47, "nets": ["fb", "tw"]},
		  {"first": "Jane", "last": "Murphy", "age": 47, "nets": ["ig", "tw"]}
		]
	  }
	`

	// if !gjson.Valid(json) {
	// 	panic("JSON IS NOT VALID")
	// }

	// gjson.AddModifier("case", func(json, arg string) string {
	// 	if arg == "upper" {
	// 		return strings.ToUpper(json)
	// 	} else if arg == "lower" {
	// 		return strings.ToLower(json)
	// 	}
	// 	return json
	// })

	// value := gjson.Get(json, "friends.#(age>=44)#.last|@case:upper")
	// for _, last := range value.Array() {
	// 	fmt.Println(last.String())
	// }

	// result, ok := gjson.Parse(json).Value().(map[string]interface{})
	// if !ok {
	// 	panic("NOT OKAY PARSING TO MAP")
	// }
	// fmt.Println(result)
	
	value, _ := sjson.Set(json, "new_obj", map[string]interface{}{"hello":"world"})
	fmt.Println(value)

}

// Explain what is go.mod
