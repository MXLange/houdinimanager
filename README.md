This is a go lib to manage your go routines.

Here is an example to test:

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/MXLange/houdinimanager"
)

func main() {

	houdiniError1, err := houdinimanager.NewHoudiniManager(true, false, 0)
	if err != nil {
		fmt.Printf("Error creating routine manager houdiniError1: %v\n", err.Error())
	} else {
		defer houdiniError1.Close()
	}

	houdiniError2, err := houdinimanager.NewHoudiniManager(false, true, 2)
	if err != nil {
		fmt.Printf("Error creating routine manager houdiniError2: %v\n", err.Error())
	} else {
		defer houdiniError2.Close()
	}

	houdini, err := houdinimanager.NewHoudiniManager(true, true, 2)
	if err != nil {
		log.Fatalf("Error creating routine manager: %v", err.Error())
	}
	defer houdini.Close()

	for i := 1; i <= 5; i++ {
		houdini.Execute(func() {
			teste(i, "true, true, 5")
		})
	}
	houdini.Wait()
	fmt.Println("The first case completed because the program waited for the houdinis to finish,\nand the max number of routines was set")
	time.Sleep(time.Duration(3) * time.Second)

	houdini2, err := houdinimanager.NewHoudiniManager(false, true, 0)
	if err != nil {
		log.Fatalf("Error creating routine manager %v", err.Error())
	}
	defer houdini2.Close()

	for i := 1; i <= 10; i++ {
		houdini2.Execute(func() {
			teste(i, "false, false, 0")
		})
	}
	houdini2.Wait() //this will make the program wait for the houdinis to finish
	fmt.Println("The second case completed because the program waited for the houdinis to finish,\nand all the routines were executed concurrently")
	time.Sleep(time.Duration(3) * time.Second)

	houdini3, err := houdinimanager.NewHoudiniManager(false, false, 0)
	if err != nil {
		log.Fatalf("Error creating routine manager %v", err.Error())
	}
	defer houdini3.Close()

	for i := 1; i <= 50; i++ {
		houdini3.Execute(func() {
			teste(i, "false, false, 0")
		})
	}
	houdini3.Wait() //this will not make the program wait for the houdinis to finish

	fmt.Println("The third case didn't complete because the program didn't wait for the houdinis to finish")
}

func teste(i int, config string) {
	fmt.Printf("Executando rotina %v, config: (%v)\n", i, config)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("Rotina %v finalizada, config: (%v)\n", i, config)
}

