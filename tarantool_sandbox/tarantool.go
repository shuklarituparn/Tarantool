package main

import (
	"fmt"
	"github.com/tarantool/go-tarantool"
	"log"
)

func main() {
	conn, err := tarantool.Connect("127.0.0.1:3301", tarantool.Opts{
		User: "admin",
		Pass: "pass",
	})

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	_, err = conn.Insert("tester", []interface{}{4, "ABBA", 1972})
	resp2, err := conn.Select("tester", "primary", 0, 1, tarantool.IterEq, []interface{}{4})

	/*


		This selects a tuple by the primary key with offset = 0 and limit = 1 from a space named tester (in our example,
		this is the index named primary, based on the id field of each tuple).
	*/

	resp2, err = conn.Select("tester", "secondary", 0, 1, tarantool.IterEq, []interface{}{"ABBA"})

	// As we updated our database with secondary key, it works perfectly fine

	resp2, err = conn.Update("tester", "primary", []interface{}{4}, []interface{}{[]interface{}{"+", 2, 3}})

	/*
		This increases by 3 the value of field 2 in the tuple with id = 4. If a tuple with this id doesn’t exist,
		Tarantool will return an error.
	*/

	resp2, err = conn.Replace("tester", []interface{}{4, "New band", 2011}) //ABBA and 1975 will change

	resp2, err = conn.Upsert("tester", []interface{}{4, "Another band", 2000}, []interface{}{[]interface{}{"+", 2, 5}})
	if err != nil {
		log.Fatal(err)
	}

	/*	code := resp.Code
		data := resp.Data
	*/
	code2 := resp2.Code
	data2 := resp2.Data

	fmt.Printf("Response Code: %d\n", code2)
	fmt.Printf("Response Data: %+v\n", data2)
}

/*

Response code: 0
Response Data: [[4 ABBA 1972]]


*/

/*
reason to use the interface in go:

1. Allows us to pass the data of different type to the function
2.]interface{}{"+", 2, 3}. This structure includes a string ("+"), followed by two numbers. Using interface{}
allows you to represent this mixed data structure within a slice.
3.The use of interface{} allows you to pass different numbers of parameters to the function.
*/
