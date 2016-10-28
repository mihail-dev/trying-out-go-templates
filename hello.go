package main

import (
	"fmt"
	"os"
	"text/template"
)

       
type Inventory struct {
        Material string
        Count    uint
}


func main() {
	fmt.Printf("Hello, world.\n")
	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}\n")
	if err != nil { panic(err) }
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil { panic(err) }

}
