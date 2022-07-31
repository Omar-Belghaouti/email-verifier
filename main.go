package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	domain = strings.TrimSpace(domain)

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Fatal(err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Fatal(err)
	}
	for _, txtRecord := range txtRecords {
		if strings.HasPrefix(txtRecord, "v=spf1") {
			hasSPF = true
			spfRecord = txtRecord
		}
		if strings.HasPrefix(txtRecord, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = txtRecord
		}
	}

	fmt.Printf("%s, %t, %t, %s, %t, %s\n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
