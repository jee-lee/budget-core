package helpers

import "github.com/google/uuid"

func Pointer[T any](v T) *T {
	return &v
}

func GetUUID(s string) (*uuid.UUID, error) {
	if s == "" {
		return nil, nil
	}
	u, err := uuid.Parse(s)
	return &u, err
}
