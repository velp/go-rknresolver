package rknresolver

import (
	"fmt"
	"log"
	"time"
	"github.com/miekg/dns"
)

type Resolver struct {
	workers   int
	partSize  int
	ns        string
	logger    *log.Logger
}

func NewResolver(workers int, partSize int, ns string, logger *log.Logger) *Resolver {
	return &Resolver{
		workers,
		partSize,
		ns,
		logger,
	}
}

func (r *Resolver) ActualizeReg(reg *Register) *Register {
	startTime := time.Now()
	// Create workers
	jobs := make(chan *[]Content, 100)
	results := make(chan *[]Content, 100)
	for w := 1; w <= r.workers; w++ {
		go r.worker(w, jobs, results)
	}
	r.logger.Println("Create", r.workers, "workers")

	// Send jobs
	contSize := len(reg.Content)
	countParts := 0
	for i := 0; i < contSize; i += r.partSize {
		partConts := reg.Content[i:min(i+r.partSize, contSize)]
		jobs <- &partConts
		countParts += 1
	}
	r.logger.Println("Send all", countParts, "jobs")
	close(jobs)

	// Catch results
	var dump Register
	for c:=0; c<countParts; c++ {
		res := <- results
		dump.Content = append(dump.Content, *res...)
		r.logger.Printf("Progress (%s) %d/%d\n", time.Since(startTime).String(), len(dump.Content), contSize)
	}
	return &dump
}

func (r *Resolver) worker(id int, jobs <-chan *[]Content, results chan<- *[]Content) {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.RecursionDesired = true
	for j := range jobs {
		var contents []Content
		for _, cont := range *j {
			if len(cont.Domain) > 0 {
				for _, domain := range cont.Domain {
					m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
					ips, err := r.lookupHost(c, m)
					if err == nil {
						cont.IP = merge(cont.IP, ips)
					}
				}
			}
			contents = append(contents, cont)
		}
		results <- &contents
	}
}

func (r *Resolver) lookupHost(c *dns.Client, m *dns.Msg) ([]string, error) {
	rr, _, err := c.Exchange(m, r.ns)
	if rr == nil || err != nil {
		return nil, err
	}
	if rr.Rcode != dns.RcodeSuccess {
		return nil, fmt.Errorf("invalid answer name")
	}
	var ips []string
	for _, ans := range rr.Answer {
		if rec, ok := ans.(*dns.A); ok {
			ips = append(ips, rec.A.String())
		}
	}
	return ips, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func merge(sl1 []string, sl2 []string) []string {
	for _, el := range sl2 {
		if !exist(sl1, el) {
			sl1 = append(sl1, el)
		}
	}
	return sl1
}

func exist(slice []string, element string) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}
