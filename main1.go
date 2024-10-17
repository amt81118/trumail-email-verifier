package main

import (
	"fmt"

	emailverifier "github.com/AfterShip/email-verifier"
)

var (
	verifier = emailverifier.NewVerifier()
)

func main() {
	email := "amityadav81118@gmail.com"

	ret, err := verifier.Verify(email)
	if err != nil {
		fmt.Println("Verify email address failed, error is:", err)
		return
	}

	if !ret.Syntax.Valid {
		fmt.Println("Email address syntax is invalid")
		return
	}

	// Print formatted output
	fmt.Println("Email Validation Result:",ret)
	fmt.Printf("  Email: %s\n", ret.Email)
	fmt.Printf("  Disposable: %t\n", ret.Disposable)
	fmt.Printf("  Reachable: %s\n", ret.Reachable)
	fmt.Printf("  Role Account: %t\n", ret.RoleAccount)
	fmt.Printf("  Free Email: %t\n", ret.Free)
	fmt.Printf("  Syntax:\n")
	fmt.Printf("    Username: %s\n", ret.Syntax.Username)
	fmt.Printf("    Domain: %s\n", ret.Syntax.Domain)
	fmt.Printf("    Valid: %t\n", ret.Syntax.Valid)
	fmt.Printf("  Has MX Records: %t\n", ret.HasMxRecords)

}
