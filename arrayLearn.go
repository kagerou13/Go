package array

import (
	"fmt"
	"unsafe"
)

func ArrayLearn() {
	// declare array -> set of same data value
	// format :
	// var nameArray = [len]type{}
	// nameArray := [len]type{}
	// nameArray := [len]type{val1, val2, val3}
	// array must have set len first

	fmt.Println("===== NON PRIMITIVE DATA TYPE =====")
	fmt.Println("Array")

	var arr1 = [3]int{}
	arr2 := [4]string{}
	arr3 := [5]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)

	//update value
	arr2[0] = "ninja shoyo"
	arr2[1] = "kageyama"
	arr2[2] = "miya"

	fmt.Print("Array after update : ")
	fmt.Println(arr2)
	fmt.Printf("Length = %d\n", len(arr2))
	fmt.Printf("Size = %d byte\n", unsafe.Sizeof(arr2))

	// if we have many of we can declare with this format
	var arr4 = [5]string{
		"haikyuu",
		"bleach",
		"naruto",
		"one piece",
		"dragon ball",
	}

	fmt.Println(arr4)

	// without len we can declare array, example
	// nameArray := [...]type{}
	arr5 := [...]int{
		1,
		2,
		3,
	}
	fmt.Println(arr5)
}

func SliceLearn() {
	// slice is part of array
	// length from slice is changeable or unconstantly, we can change it
	// have 3 data are pointer, length, and capacity
	// how to make slice
	// 1. sliceName := arrayName[]
	// 2. var sliceName []type = arrayname[]

	fmt.Println("Slice")
	anime := [...]string{
		"naruto",
		"bleach",
		"luffy",
		"haikyuu",
		"ashito",
		"sasuke",
		"aquamarine",
	}
	fmt.Print("Array = ")
	fmt.Println(anime)
	fmt.Println("slice it : >>")
	//from above array we can make slice with 4 format:

	/*
	* format : slice = array[low:high] -> get array from low index until high -1 index
	 */
	slice1 := anime[2:4] // -> only get 2 value, index 2 and index 3
	fmt.Print("Slice1 from 2 to 4 = ")
	fmt.Println(slice1)
	// pointer = 2
	// len of slice = 2
	// capacity = len of array - pointer = 7 - 2 = 5
	// length must be less than capacity

	/*
	* format : slice = array[low:]  -> get array from low index until end of array
	 */
	slice2 := anime[2:]
	fmt.Print("Slice2 from 2 to end = ")
	fmt.Println(slice2)

	/*
	* format : slice = array[:high]  -> get array from begin of array until high -1 index
	 */
	var slice3 []string = anime[:5]
	fmt.Print("Slice3 from begin to 5 = ")
	fmt.Println(slice3)

	/*
	* format : slice = array[:]  -> get all of array
	 */
	var slice4 []string = anime[:]
	fmt.Print("Slice4 all of array = ")
	fmt.Println(slice4)

	// FUNCTIONS ON SLICE
	// len(slice) -> to get length of slice
	// cap(slice) -> to get capacity
	// append(slice, data) -> make new slice and add new data after end index of old slice
	// make([]typedata, length, capacity) -> make new slice
	// copy(destination, source)  -> copy slice from source to destination

	//len(slice)
	fmt.Print("Len slice1 = ")
	fmt.Println(len(slice1))

	//cap(slice)
	fmt.Print("Capacity slice1 = ")
	fmt.Println(cap(slice1))

	//append(slice, data)
	slice5 := append(slice1, "rimuru")
	fmt.Print("Slice1 after append = ")
	fmt.Println(slice5)

	// update value
	slice5[0] = "villain"
	fmt.Print("Slice5 update = ")
	fmt.Println(slice5)

	// make([]typedata, length, capacity)
	slice6 := make([]string, 2, 4) // we can append new value
	slice6[0] = "tabun"
	slice6[1] = "idol"
	fmt.Print("Slice from make func = ")
	fmt.Println(slice6)

	// append
	slice6 = append(slice6, "romance")
	fmt.Print("After append = ")
	fmt.Println(slice6)
	fmt.Printf("Capacity = %d\n", cap(slice6))

	slice7 := append(slice6, "gunjuo")
	fmt.Print("After append = ")
	fmt.Println(slice7)
	fmt.Printf("Len = %d\n", len(slice7))
	fmt.Printf("Capacity = %d\n", cap(slice7))
	slice7[0] = "zero"
	fmt.Printf("Slice6 = %v\n", slice6) // change because still same array

	slice7 = append(slice7, "comet")
	fmt.Print("After append = ")
	fmt.Println(slice7)
	fmt.Printf("Len = %d\n", len(slice7))
	fmt.Printf("Capacity = %d\n", cap(slice7))

	slice7 = append(slice7, "monster")
	fmt.Print("After append = ")
	fmt.Println(slice7)
	fmt.Printf("Len = %d\n", len(slice7))
	fmt.Printf("Capacity = %d\n", cap(slice7))

	slice7[0] = "haruka"
	fmt.Printf("Slice6 = %v\n", slice6) // not change because is not same array
	fmt.Printf("Slice7 = %v\n", slice7)

	// copy(destination,source)
	fromThis := anime[:]
	toThis := make([]string, len(fromThis), cap(fromThis))
	fmt.Println("Copy array : ")
	copy(toThis, fromThis)
	fmt.Println(fromThis)
	fmt.Println(toThis)

	var arrayVar = [...]int{1, 2, 3, 4} // this is array
	var sliceVar = []int{1, 2, 3, 4}    // this is slice

	fmt.Printf("Capacity array = %d\n", cap(arrayVar))
	fmt.Printf("Capacity Slice = %d\n", cap(sliceVar))

	sliceVar = append(sliceVar, 5)
	fmt.Printf("Capacity Slice = %d\n", cap(sliceVar))
}

func MapsLearn() {

	// Map is similar with array and slice, but index change with key
	// key is unique and must different
	// format : var nameMaps := map[keyType]valueType{}
	// format : nameMaps := map[keyType]valueType{}

	// make maps
	person := map[string]string{
		"name": "kagerou",
		"age":  "24",
		"job":  "technical support",
	}

	// access maps
	// nameMaps[key]
	fmt.Printf("Maps : %s\n", person)
	fmt.Println(person["name"])
	fmt.Println(person["idol"]) // will return default value

	// functions on maps
	// len(mapsVar) 		-> check length of map
	// map[key]   			-> to get data from map
	// map[key] = value		-> to update or change value with key
	// make(map[typeKey]TypeValue) -> make new map
	// delete(map, key)  	-> to delete data on map with key

	// len
	fmt.Printf("Length of Map = %d\n", len(person))
	// access
	fmt.Printf("Data = %s\n", person["job"])
	// update value
	person["job"] = "developer"
	fmt.Printf("After update = %s\n", person["job"])
	// make map
	nameAnime := make(map[int]string)
	nameAnime[0] = "naruto"
	nameAnime[1] = "bleach"
	fmt.Printf("New Map = %v\n", nameAnime)

	// delete data on map with key
	delete(nameAnime, 1)
	fmt.Printf("New Map = %v\n", nameAnime)

}
