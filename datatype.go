package datatype

/*
*Primitive data type : numeric, string, bool
* numeric :
* - int (int8, int16, int32, int64, uint8, uint16, uint32, uint64)
* - float (float32, float64)
* - complex(complex64, complex128)
*check memori : unsafe.Sizeof(var)
*check type : reflect.TypeOf(var)

*byte is an alias from int8, byte is used to represent of ASCII character

* string
* - for loops char used string(var)
* - for loops ASCII character from string byte(var) or print directly

* bool has true or false value

* NON PRIMITIVE DATA TYPE NEXT SESSION
 */

import (
	"fmt"
	"math/cmplx"
	"reflect"
	"strconv"
	"unsafe"
)

func DataType() {
	// DATA PRIMITIVE : INTEGER, FLOAT, STRING, BOOL

	fmt.Printf("DATA TYPE\n")
	//integer
	var intVar int = 23
	fmt.Printf("Integer = %d\n", intVar)

	//float
	var intFloat float32 = 25.13
	fmt.Printf("Float = %v\n", intFloat)

	//string
	name := "kagerou"
	fmt.Printf("String = %v\n", name)
	fmt.Println("Len string is ", len(name))
	for _, val := range name {
		fmt.Println(val) // !this will print ASCII
	}
	for _, val := range name {
		fmt.Println(string(val)) //print character from name
	}

	//boolean
	var myStatus bool
	fmt.Printf("Boolean = %v\n", myStatus)

	//unsigned integer -- non negativ value
	var unsigInteger uint16 = 34
	fmt.Printf("Unsigned Integer = %v\n", unsigInteger)

	//complex
	var z complex128 = cmplx.Sqrt(-5 + 12i)
	fmt.Printf("Complex128 = %v and type = %T\n", z, z)

	// *CONVERT DATA TYPE
	fmt.Println("CONVERT DATA")
	// convert NUMERIC
	newFloat := float32(intVar)
	fmt.Printf("Integer to Float = %.2f\n", newFloat)

	newInt := int(intFloat)
	fmt.Printf("Float to Integer = %v\n", newInt)
	fmt.Printf("Type = %T\n", newInt)

	//convert numeric to string
	//integer to string
	var intToString = strconv.Itoa(intVar) // !common conversion
	fmt.Printf("Numeric to String = %v\n", intToString)
	fmt.Printf("Type = %T\n", intToString)

	//string to numeric (int)
	var newString = "25"
	newString2 := "26"

	if stringToInteger, err := strconv.Atoi(newString); err == nil { // !common conversion
		fmt.Printf("String to integer = %v and type = %T\n", stringToInteger, stringToInteger)
	}

	if s, err := strconv.ParseInt(newString2, 10, 32); err == nil { // !specific conversion
		fmt.Printf("String to integer = %v, and type = %T", s, s)
	}

	binStr := "001"
	strToInt, er := strconv.ParseInt(binStr, 2, 64)
	if er != nil {
		fmt.Println("\nError during conversion")
		return
	}
	fmt.Printf("\n\nBinary String to Integer = %v and type %T\n", strToInt, strToInt)

	// STRING TO FLOAT
	strFloat := "3.14"
	floatStr, errf := strconv.ParseFloat(strFloat, 64)
	if errf != nil {
		fmt.Println("\nError during conversion, your data is not matching with aim type")
		return
	}
	fmt.Printf("\nString Float to Float = %v and type %T\n", floatStr, floatStr)

	// STRING TO BOOL
	strBool := "false"
	boolStr, errb := strconv.ParseBool(strBool)
	if errb != nil {
		fmt.Println("\nError during conversion")
	}
	fmt.Printf("\nString Bool to Boolean = %v and type %T\n", boolStr, boolStr)

	// NUMERIC TO STRING

	// INTEGER TO STRING
	myage := 25
	myAgeToString := strconv.FormatInt(int64(myage), 26)
	fmt.Printf("\nInteger to String = %v and type %T\n", myAgeToString, myAgeToString)

	// FLOAT TO STRING
	myFloat := 25.13
	myFloat2 := 24e+2

	floatToString := strconv.FormatFloat(float64(myFloat), 'f', 2, 64)
	floatToString2 := strconv.FormatFloat(float64(myFloat2), 'e', 2, 64)

	fmt.Printf("\nFloat to String = %v and type %T\n", floatToString, floatToString)
	fmt.Printf("\nFloat to String = %v and type %T\n", floatToString2, floatToString2)

	// BOOLEAN TO STRING
	myBoolean := true
	booleanToString := strconv.FormatBool(myBoolean)
	fmt.Printf("\nBoolean to String %v and type %T\n", booleanToString, booleanToString)

	//COMPLEX
	//initial
	comp := complex(3, 1)
	fmt.Println(comp)
	fmt.Printf("Comp is %d bytes\n", unsafe.Sizeof(comp))
	fmt.Printf("comp type's is %s\n", reflect.TypeOf(comp))

	/*
	*Byte is an alias for uint8 in Go
	*byte is used to represent the ASCII character
	 */
	var rbyte byte = 'a'
	fmt.Printf("Var : %c\n", rbyte)
	fmt.Printf("Size : %d byte\n", unsafe.Sizeof(rbyte)) // check size from variable
	fmt.Printf("Type : %s\n", reflect.TypeOf(rbyte))     //check type from variable

	s := "abc"
	s1 := "ABC"
	fmt.Println([]byte(s))
	fmt.Println([]byte(s1))

}
