package helpers

import (
	"database/sql"
	"github.com/twitchtv/twirp"
)

// NullStringFromUUID validates if the provided uuid is a valid uuid or an empty string otherwise return a Twirp error
func NullStringFromUUID(fieldName string, uuid string) (sql.NullString, error) {
	if !IsValidUUID(uuid) && uuid != "" {
		return sql.NullString{}, twirp.InvalidArgumentError(fieldName, "is an invalid uuid")
	}
	return sql.NullString{String: uuid, Valid: uuid != ""}, nil
}
