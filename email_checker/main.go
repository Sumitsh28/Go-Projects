package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main(){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("DOMIAN | hasMX | hasSPF | spfRecord | hasDMARC | dmarcRecord")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	err := scanner.Err()

	if err != nil {
		log.Fatal("Error couldn't read from input %v\n",err)
	}
}

func checkDomain(domain string) {
	
	var hasDMARC, hasSPF, hasMX bool
	var dmarcRecord, spfRecord string

	mxRecord, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("ErrorL%v\n", err)
	}

	for len(mxRecord) > 0 {

		hasMX = true;

	}

	txtRecord, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("ErrorL%v\n", err)
	}

	for _, record := range txtRecord {

		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("ErrorL%v\n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)


}