package utils

import "github.com/google/uuid"

func IsUUIDValid(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
