package main

import (
	"fmt"
	//"github.com/golang/protobuf/ptypes/any"
)

// slice, map , chan --> make to instantiate it
// map is a reference type
// key of a map can be any type; on which ever the type == operator is used.
// if map is not instantiated it is nil
func main() {

	// key can be of any thing where you can directly use == operator
	// string , int, bool ,float64
	var mymap map[string]any // only declaration of map not instatiated.

	if mymap == nil {
		fmt.Println("mymap is nil")
	}

	mymap = make(map[string]any)

	mymap["trivandrum-1"] = 590034
	mymap["bangalore-1"] = 560086
	mymap["hyderabad-2"] = 543234
	mymap["bangalore-2"] = 560032
	//["bangalore-1"] = 560076

	for key, val := range mymap {
		fmt.Println("Key:", key, "Value:", val)
	}

	val, ok := mymap["bangalore-1"]
	//val1 := mymap["bangalore-2"]
	if ok {
		fmt.Println("Key already there,so the value", val)
	}

	val, ok = mymap["bangalore-3"]
	//val1 := mymap["bangalore-2"]
	if ok {
		fmt.Println("Key already there,so the value", val)
	} else {
		fmt.Println("Key not there")
	}

	delete(mymap, "bangalore-1")

	//var mymap2 map[string]any
	var mymap2 = make(map[string]any)
	fmt.Printf("print mymap2 before deep copy %v \n", mymap2)
	err := deep_copy(mymap2, mymap)
	if err != nil {
		fmt.Printf("error %v in deep copy", err)
	}
	fmt.Printf("print mymap after deep copy %v \n", mymap)
	fmt.Printf("print mymap2 after deep copy %v \n", mymap2)

	// mymap1 := mymap

}

func deep_copy(m1 map[string]any, m2 map[string]any) error {

	if m2 == nil {
		fmt.Println("source map is nil")
	}

	for key, val := range m2 {
		m1[key] = val
	}
	return nil
}

// perform deep copy of a map
// update an element from the map, return whether it has updated or not. There should
// be an error parameter to tell that if the key does not exist shoudl return "unable to update due to no key"

// delete an element from the map , it should tell whether deleted or not.
// check nil conditions for all functions.. if nil then return as an error
