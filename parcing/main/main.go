package main

import (
	"encoding/xml"
	"fmt"
	"encoding/json"
	"os"
)

func main() {

	type Result struct {
		XMLName xml.Name `xml:"book"`
		Id    string      `xml:"id"`
		Author   []string  `xml:"author"`
		Title   string       `xml:"title"`
		Price  float32        `xml:"price"`
		Description string `xml:"description"`


	}
	var v Result

	data := `
		<book>
		<id>bk</id>
      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>

      <description>An in-depth look at creating applications
      with XML.</description>
   </book> `

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("id: %q\n", v.Id)
	fmt.Printf("price: %v\n", v.Price)
	fmt.Printf("author: %q\n", v.Author)
	fmt.Printf("description: %q\n", v.Description)
	fmt.Printf("title: %q\n", v.Title)
	fmt.Println(v)
	fmt.Println("------------------------------------------------------------------------------------------")

	b, err := json.Marshal(v)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)


}

