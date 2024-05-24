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

	houdini, err := houdinimanager.NewHoudiniManager(true, true, 2)
	if err != nil {
		log.Fatalf("Error creating routine manager: %v", err.Error())
	}

	for i := 1; i <= 5; i++ {
		houdini.Execute(func() {
			teste(i, "true, true, 5")
		})
	}
	houdini.Wait()

	houdini2, err := houdinimanager.NewHoudiniManager(false, false, 0)
	if err != nil {
		log.Fatalf("Error creating routine manager: %v", err.Error())
	}

	for i := 1; i <= 50; i++ {
		houdini2.Execute(func() {
			teste(i, "false, false, 0")
		})
	}
	houdini2.Wait() //this will not make the program wait for the houdinis to finish

	fmt.Println("The second case didn't complete because the program didn't wait for the houdinis to finish")
}

func teste(i int, config string) {
	fmt.Printf("Executando rotina %v, config: (%v\n)", i, config)
	time.Sleep(time.Duration(i) * time.Second)
	fmt.Printf("Rotina %v finalizada, config: (%v\n)", i, config)
}
