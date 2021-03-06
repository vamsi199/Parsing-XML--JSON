package main

import (
	"encoding/xml"
	"fmt"

)

func main() {

	type Result struct {
		XMLName xml.Name `xml:"catalog"`
		Book_id    string      `xml:"book id"`
		Author   string  `xml:"author"`
		Title   string       `xml:"title"`
		Price  float32        `xml:"price"`
		Description string `xml:"description"`


	}

	var v Result


	data := `
		<?xml version="1.0"?>
<catalog>

      <author>Gambardella, Matthew</author>
      <title>XML Developer's Guide</title>
      <genre>Computer</genre>
      <price>44.95</price>
      <publish_date>2000-10-01</publish_date>
      <description>An in-depth look at creating applications
      with XML.</description>


</catalog> `

	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("id: %q\n", v.Book_id)
	fmt.Printf("price: %v\n", v.Price)
	fmt.Printf("author: %q\n", v.Author)
	fmt.Printf("description: %q\n", v.Description)
	fmt.Printf("title: %q\n", v.Title)
	fmt.Println(&v)



}

