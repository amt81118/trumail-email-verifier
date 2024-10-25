# Email Verification Tool

This is a simple email verification tool written in Go. The tool reads a list of email addresses from a CSV file, verifies each email using the Trumail API, and writes the verification results to another CSV file. The program also includes random pauses between email verifications to simulate a more realistic process.

## Features

- Reads email addresses from a CSV file.
- Uses the Trumail API to verify email addresses.
- Writes detailed verification results to a CSV file.
- Adds random delays between email verifications to avoid throttling or API limits.

-Make sure you have Go installed. Then, install the necessary dependencies:
    go get github.com/sdwolfe32/trumail/verifier

Run the program
    -go run fileName.go

files 
    for single email verify ==> singleVerify.go
    for bulk eamil verify using csv ====> bulkVerify.go
    another service to verify email ==> emailVerfier2.go 