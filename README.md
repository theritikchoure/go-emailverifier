# Domain Records Checker

The Domain Records Checker is a simple GoLang application that checks various DNS records for a list of domains. It uses the Go standard library and the `net` package to perform DNS lookups and report whether certain DNS records exist for each domain.

## Table of Contents
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Usage](#usage)
- [Project Flow](#project-flow)

## Features

- Check the existence of MX (Mail Exchange) records for domains.
- Check for the presence of SPF (Sender Policy Framework) records and report the first SPF record found.
- Check for the existence of DMARC (Domain-based Message Authentication, Reporting, and Conformance) records and report the first DMARC record found.
- Read a list of domains from standard input or a text file and report the results in a CSV-like format.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Go (Golang): You should have Go installed on your system.
- Internet Connection: The application requires an internet connection to perform DNS lookups.

## Getting Started

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/theritikchoure/go-emailverifier.git
   cd go-emailverifier
    ```

2. **Run the application**:
    ```bash
    go run main.go
    ```

3. **Enter the domain in the terminal**

4. **View Results:**
    The application will output the results in a CSV-like format, indicating the presence or absence of MX, SPF, and DMARC records for each domain.

## Usage

- To check multiple domains, list them one per line in a text file and provide the file as input to the application.

- The application will display the results in the following format:
    ```
    domain, hasMx, hasSPF, spfRecord, hasDMARC, dmarcRecord
    ```

- `hasMx`, `hasSPF`, and `hasDMARC` indicate whether the respective records were found (true/false).

- `spfRecord` and `dmarcRecord` display the SPF and DMARC records if found, otherwise, they are empty.

## Project Flow

The Domain Records Checker follows these main steps during execution:

1. Initialization:
    - The program initializes by importing required packages and setting up variables.

2. User Input:
    - The program prompts the user to provide a list of domains. Users can either enter domains interactively or pipe a text file containing domains into the application.

3. Domain Checking Loop:
    - For each domain provided by the user, the program calls the checkDomain function.

4. DNS Lookups:
    - Inside the checkDomain function:
        - The program performs DNS lookups for MX, SPF, and DMARC records for the given domain using the net package.
        - It checks for errors during each lookup operation and logs any errors that occur.

5. Results Reporting:
    - The program formats the results for each domain in a CSV-like format.
    - It indicates whether MX, SPF, and DMARC records were found (true/false).
    - It includes the SPF and DMARC records if found, otherwise, they are empty.

6. Output:
    - The program prints the results to the console for each domain.