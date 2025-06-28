package utils

import (
	"errors"
	"mime/multipart"
	"regexp"
)

// ValidateFile checks if the uploaded file is valid based on its size and type.
func ValidateFile(file *multipart.FileHeader, maxSize int64, allowedTypes []string) error {
	if file.Size > maxSize {
		return errors.New("file size exceeds the maximum limit")
	}

	fileType := file.Header.Get("Content-Type")
	for _, allowedType := range allowedTypes {
		if fileType == allowedType {
			return nil
		}
	}

	return errors.New("file type is not allowed")
}

// ValidateText checks if the provided text input is valid.
func ValidateText(input string, minLength int) error {
	if len(input) < minLength {
		return errors.New("input text is too short")
	}

	return nil
}

// ValidateEmail checks if the provided email is valid.
func ValidateEmail(email string) error {
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}

	return nil
}