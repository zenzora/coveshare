package auth

import "time"

// login via email - sends login link, with optional "redirect" to encrypted message

// CreateLoginCode (from email)
func CreateLoginCode(email string, expiration time.Time) {

}

// Validate login code
