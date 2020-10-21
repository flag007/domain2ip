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
	domain_channel := make(chan string)
	var domainWG sync.WaitGroup

	for i := 0; i < 20; i++ {
		domainWG.Add(1)

		go func() {
			for domain := range domain_channel {
				ips, err := net.LookupHost(domain)
				//addr, err := net.ResolveIPAddr("ip", domain)
				if err != nil {
					continue
				}

				for _, ip := range ips {
					fmt.Fprintf(os.Stdout, "%30-s%s\n", domain, ip)
				}
			}
			domainWG.Done()
		}()
	}

	for sc.Scan() {
		domain := strings.ToLower(sc.Text())
		domain_channel <- domain
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to read input: %s\n", err)
	}
	close(domain_channel)
	domainWG.Wait()

}
