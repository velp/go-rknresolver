package main

import (
	"net"
	"fmt"
	"time"
	"io/ioutil"
	"encoding/xml"
	"encoding/json"
)

const countWorkers int = 20
const partSize int = 1000

type XMLrkn struct {
	XMLName  xml.Name     `xml:"register"`
	Content  []XMLContent `xml:"content"`
}

type XMLContent struct {
	AttrIncludeTime   string       `xml:"includeTime,attr"`
	AttrEntryType     string       `xml:"entryType,attr"`
	IP                string       `xml:"ip"`
	Domain            string       `xml:"domain"`
}

type RKNBlocks struct {
	Cont []RKNBlock `json:"content"`
}

type RKNBlock struct {
	IP       []string  `json:"ips"`
	Domain   string    `json:"domain"`
	Date     string    `json:"date"`
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func worker(id int, jobs <-chan *[]XMLContent, results chan<- *[]RKNBlock) {
	for j := range jobs {
		fmt.Println("Worker", id, "started job")
		var rknBlocks []RKNBlock
		for _, cont := range *j {
			ips, err := net.LookupHost(cont.Domain)
			if err != nil {
				ips = []string{"None"}
			}
			rknBlocks = append(rknBlocks, RKNBlock{ips, cont.Domain, cont.AttrIncludeTime})
		}
		fmt.Println("Worker", id, "finished job")
		results <- &rknBlocks
	}
}

func main() {
	startTime := time.Now()

	data, err := ioutil.ReadFile("dump.xml")
	if err != nil {
		panic(err)
	}

	var reestr XMLrkn
	err = xml.Unmarshal(data, &reestr)
	if err != nil {
		panic(err)
	}

	// Create workers
	jobs := make(chan *[]XMLContent, 100)
	results := make(chan *[]RKNBlock, 100)
	for w := 1; w <= countWorkers; w++ {
		go worker(w, jobs, results)
	}
	fmt.Println("Create", countWorkers, "workers")

	// Send jobs
	contSize := len(reestr.Content)
	countParts := 0
	for i := 0; i < contSize; i += partSize {
		partConts := reestr.Content[i:min(i+partSize, contSize)]
		jobs <- &partConts
		countParts += 1
	}
	fmt.Println("Send all jobs")
	close(jobs)

	// Catch results
	var rknBlocks RKNBlocks
	for r:=0; r<countParts; r++ {
		res := <- results
		rknBlocks.Cont = append(rknBlocks.Cont, *res...)
		fmt.Printf("Progress (%s) %d/%d\n", time.Since(startTime).String(), len(rknBlocks.Cont), contSize)
	}

	// Gen json
	jsonRKNBlocks, err := json.MarshalIndent(&rknBlocks, "", "  ")
	err = ioutil.WriteFile("dump.json", jsonRKNBlocks, 0644)
	if err != nil {
		panic(err)
	}
}