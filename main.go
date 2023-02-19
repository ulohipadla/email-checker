package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func checkDomain(domain string) {
	var hasMx, hasSpf, hasDMARC bool
	var spfRecord, dmarcwRecord string
	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Fatal(err)
	}

	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Fatal(err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSpf = true
			spfRecord = record
			break
		}
	}

	dmarcRecord, err2 := net.LookupTXT("_dmarc." + domain)

	if err != nil {
		log.Fatal(err2)
	}

	for _, record := range dmarcRecord {
		if strings.HasPrefix(record, "v=spf1") {
			hasDMARC = true
			dmarcwRecord = record
			break
		}
	}
	fmt.Println("hasMx, hasSpf, hasDMARC, spfRecord, dmarcwRecord")
	fmt.Println(hasMx, hasSpf, hasDMARC, spfRecord, dmarcwRecord)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from from input: ", err)
	}
}
