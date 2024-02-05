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
	fmt.Println("domain, haxMx, hasSPF, sprRecord, hasDMARC, demarcRecord \n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error: Could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, demarcRecord string

	maxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error :%v\n", err)
	}

	if len(maxRecords) > 0 {
		hasMx = true

	}
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error:%v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	demarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error%v\n", err)
	}
	for _, record := range demarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			demarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v ", domain, hasMx, hasSPF, spfRecord, hasDMARC, demarcRecord)

}
