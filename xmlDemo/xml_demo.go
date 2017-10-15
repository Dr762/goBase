package xmlDemo

import (
	"io/ioutil"
	"log"
	"strings"
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName Name `xml:"person"`
	Name Name `xml:"name"`
	Email []Email `xml:"email"`
}

type Name struct {
	Family string `xml:"family"`
	Personal string `xml:"personal"`
}

type Email struct {
	Type string `xml:"type,attr"`
	Address string `xml:",chardata"`
}

func ParseXML(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	reader := strings.NewReader(string(bytes))
	parser := xml.NewDecoder(reader)
	depth := 0

	for {
		token, err := parser.Token()
		if err != nil {
			break
		}

		switch t := token.(type) {

		case xml.StartElement:
			elem := xml.StartElement(t)
			name := elem.Name.Local
			printElem(name, depth)
			depth++

		case xml.EndElement:
			depth--
			elem := xml.EndElement(t)
			name := elem.Name.Local
			printElem(name, depth)

		case xml.CharData:
			bytes := xml.CharData(t)
			printElem("\""+string([]byte(bytes))+"\"",depth)

		case xml.Comment:
			printElem("Comment", depth)

		case xml.ProcInst:
			printElem("ProcInst", depth)

		case xml.Directive:
			printElem("Directive", depth)
		default:
			fmt.Println("Unknown")

		}
	}
}

func UnmarshalXML(filename string){
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var person Person
	if xml.Unmarshal(bytes,&person) != nil {
		log.Fatal(err)
	}

	fmt.Println("Family name: \"" + person.Name.Family + "\"")
	fmt.Println("Second email address: \"" + person.Email[1].Address + "\"")
}


func printElem(str string,depth int){
	for n:=0;n<depth;n++{
		fmt.Print(" ")
	}
	fmt.Println(str)
}
