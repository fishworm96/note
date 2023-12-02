package main

import (
	"encoding/xml"
	"fmt"
)

func main() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName) // XMLName: xml.Name{Space:"", Local:"Person"}
	fmt.Printf("Name: %q\n", v.Name) // Name: "Grace R. Emlin"
	fmt.Printf("Phone: %q\n", v.Phone) // Phone: "none"
	fmt.Printf("Email: %v\n", v.Email) // Email: [{home gre@example.com} {work gre@work.com}]
	fmt.Printf("Groups: %v\n", v.Groups) // Groups: [Friends Squash]
	fmt.Printf("Address: %v\n", v.Address) // Address: {Hanga Roa Easter Island}
}
