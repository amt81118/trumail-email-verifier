package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	trumail "github.com/sdwolfe32/trumail/verifier"
)

func main() {
	//? Seed the random number generator ->
	rand.Seed(time.Now().UnixNano())

	//? Open the CSV file for reading
	file, err := os.Open("pharma.csv")
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer file.Close()

	//? Initialize the CSV reader
	reader := csv.NewReader(file)

	//? Read all the records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	//? Open the result CSV file in append mode
	resultFile, err := os.OpenFile("pharmaVer1.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening result CSV file: %v", err)
	}
	defer resultFile.Close()

	//? Initialize the CSV writer
	writer := csv.NewWriter(resultFile)
	defer writer.Flush()

	//? Check if the file is empty, and if so, write the header
	fileInfo, err := resultFile.Stat()
	if err != nil {
		log.Fatalf("Error getting file info: %v", err)
	}
	if fileInfo.Size() == 0 {
		//? Write the header row to the CSV file if it's empty
		header := []string{
			"Email Address", "Username", "Domain", "MD5 Hash", "Valid Format",
			"Deliverable", "Full Inbox", "Host Exists", "Catch-All", "Error Message",
		}
		if err := writer.Write(header); err != nil {
			log.Fatalf("Error writing header to CSV file: %v", err)
		}
	}

	//? Initialize the email verifier
	v := trumail.NewVerifier("localhost", "amit.yadav@code-b.dev")

	for _, record := range records {
		if len(record) > 0 {
			email := record[0]

			// Random timeout between 1 to 5 minutes
			// randomMinutes := rand.Intn(5) + 1
			// fmt.Printf("Pausing for %d minute(s) before verifying email: %s\n", randomMinutes, email)
			// time.Sleep(time.Duration(randomMinutes) * time.Minute)

			randomMilliseconds := rand.Intn(10000) + 5001
			fmt.Printf("Pausing for %d millisecond(s) before verifying email: %s\n", randomMilliseconds, email)
			time.Sleep(time.Duration(randomMilliseconds) * time.Millisecond)

			result, err := v.Verify(email)
			fmt.Printf("Verification result for %s: %+v\n", email, result)

			// Prepare row data
			var row []string

			if err != nil {
				log.Printf("Error verifying email %s: %v", email, err)

				row = []string{
					email,     
					"",        
					"",         
					"",         
					"false",    
					"false",   
					"false",   
					"false",    
					"false",    
					err.Error(),
				}
			} else {
				row = []string{
					result.Address.Address,
					result.Address.Username,
					result.Address.Domain,
					result.Address.MD5Hash,
					fmt.Sprintf("%t", result.ValidFormat),
					fmt.Sprintf("%t", result.Deliverable),
					fmt.Sprintf("%t", result.FullInbox),
					fmt.Sprintf("%t", result.HostExists),
					fmt.Sprintf("%t", result.CatchAll),
					"", // No error message, so leave it empty
				}
			}

			//? Write the row to the CSV file immediately after processing
			if err := writer.Write(row); err != nil {
				log.Fatalf("Error writing row to CSV file: %v", err)
			}
			//? Flush after each write to ensure data is written to the file
			writer.Flush() 
		}
	}
}
