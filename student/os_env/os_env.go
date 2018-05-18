package main

import "fmt"
import "os"

func main() {
	var JAVAHOME string
	JAVAHOME = os.Getenv("JAVA_HOME")
	fmt.Println(JAVAHOME)
}
