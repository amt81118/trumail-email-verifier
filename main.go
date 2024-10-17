package main

import (
	"fmt"
	"log"

	trumail "github.com/sdwolfe32/trumail/verifier"
)

func main() {
	v := trumail.NewVerifier("localhost", "amit.yadav@code-b.dev")

	// Validate a single email address
	result, err := v.Verify("vinita@paisabazaar.com")
	if err != nil {
		log.Fatalf("Error verifying email: %v", err)
	}

	// Print the result in a formatted way
	// log.Printf("Result: %+v\n", result)
	fmt.Println("Verification Result:")
	fmt.Printf("  Address: %s\n", result.Address.Address)
	fmt.Printf("  Username: %s\n", result.Address.Username)
	fmt.Printf("  Domain: %s\n", result.Address.Domain)
	fmt.Printf("  MD5 Hash: %s\n", result.Address.MD5Hash)
	fmt.Printf("  Valid Format: %t\n", result.ValidFormat)
	fmt.Printf("  Deliverable: %t\n", result.Deliverable)
	fmt.Printf("  Full Inbox: %t\n", result.FullInbox)
	fmt.Printf("  Host Exists: %t\n", result.HostExists)
	fmt.Printf("  Catch-All: %t\n", result.CatchAll)
}
