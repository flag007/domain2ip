package main

import (
	"fmt"
	"strings"
	"net"
	"os"
	"sync"
	"bufio"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	var domains []string
	domain_channel1 := make(chan string)
	var domainWG1 sync.WaitGroup

	for i := 0; i < 50; i++ {
		domainWG1.Add(1)

		go func() {
			for domain := range domain_channel1 {
				addr, err := net.ResolveIPAddr("ip", domain)
				if err != nil {
					continue
				}
				fmt.Fprintf(os.Stdout, "%s %s\n", domain, addr)
			}
		}()
		domainWG1.Done()
	}

	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		domain_channel1 <- domain
		domains = append(domains, domain)
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}

	domainWG1.Wait()
	close(domain_channel1)
}
