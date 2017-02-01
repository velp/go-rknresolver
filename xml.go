package rknresolver

import (
	"time"
	"bytes"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"

	"golang.org/x/net/html/charset"
)

// XML scheme: http://vigruzki.rkn.gov.ru/docs/description_for_operators_actual.pdf
type Register struct {
	Content  []Content `xml:"content" json:"content"`
}

type Content struct {
	ID            int       `xml:"id,attr" json:"id"`
	IP            []string  `xml:"ip" json:"ip"`
	Subnet        []string  `xml:"ipSubnet" json:"ipSubnet"`
	Domain        []string  `xml:"domain" json:"domain"`
	URL           []string  `xml:"url" json:"url"`
	IncludeTime   RegTime   `xml:"includeTime,attr" json:"includeTime"`
	EntryType     int       `xml:"entryType,attr" json:"entryType"`
	BlockType     string    `xml:"blockType,attr" json:"blockType"`
	UrgencyType   int       `xml:"urgencyType,attr" json:"urgencyType"`
	Hash          string    `xml:"hash,attr" json:"hash"`
	Decision      Decision  `xml:"decision" json:"decision"`
}

type Decision struct {
	Date    RegTime  `xml:"date,attr" json:"date"`
	Number  string   `xml:"number,attr" json:"number"`
	Org     string   `xml:"org,attr" json:"org"`
}

type RegTime struct {
	time.Time
}

func (rt *RegTime) UnmarshalXMLAttr(attr xml.Attr) error {
	var shortForm string
	if attr.Name.Local == "date" {
		shortForm = "2006-01-02" // yyyymmdd date format
	} else {
		shortForm = "2006-01-02T15:04:05"
	}
	
	parse, err := time.Parse(shortForm, attr.Value)
	if err != nil {
		return err
	}
	*rt = RegTime{parse}
	return nil
}

func Parse(filepath string) (*Register, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var dump Register
	reader := bytes.NewReader(data)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&dump)
	if err != nil {
		return nil, err
	}
	return &dump, nil
}

func (r *Register) WriteJSONFile(filepath string) error {
	jsonContent, err := json.MarshalIndent(r, "", "  ")
	err = ioutil.WriteFile(filepath, jsonContent, 0644)
	if err != nil {
		return err
	}
	return nil
}



