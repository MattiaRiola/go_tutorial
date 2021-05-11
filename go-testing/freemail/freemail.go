package freemail

import "strings"

var freemails = []string{"gmail.com", "yahoo.com", "outlook.com"}

// return true if the email contains a free email provider that I've in this local "database"
func IsFreemail(email string) bool {
	for _, provider := range freemails {
		if strings.Contains(email, provider) {
			return true
		}
	}
	return false
}
