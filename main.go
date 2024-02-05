package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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

}
