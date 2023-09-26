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
	// Create a scanner to read user input from the command line.
	scanner := bufio.NewScanner(os.Stdin)

	// Print a header for the output.
	fmt.Printf("domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord\n")

	fmt.Printf("Enter the domain - ")

	// Read input lines until EOF (Ctrl+D in most terminals).
	for scanner.Scan() {
		// Call the checkDomain function for each domain provided by the user.
		checkDomain(scanner.Text())
	}

	// Check for errors when reading from input.
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error: could not read from input: %v\n", err)
	}
}

func checkDomain(domain string) {
	// Declare variables to store information about domain records.
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	// Look up MX (Mail Exchange) records for the given domain.
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Error looking up MX records for %s: %v\n", domain, err)
	}

	// Check if there are MX records for the domain.
	if len(mxRecords) > 0 {
		hasMx = true
	}

	// Look up TXT (text) records for the given domain.
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error looking up TXT records for %s: %v\n", domain, err)
	}

	// Iterate through TXT records to find an SPF (Sender Policy Framework) record.
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// Look up DMARC (Domain-based Message Authentication, Reporting, and Conformance) records for the domain.
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error looking up DMARC records for %s: %v\n", domain, err)
	}

	// Iterate through DMARC records to find a DMARC record.
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	// Print the results for the domain.
	fmt.Printf("%s, %v, %v, %s, %v, %s\n", domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
