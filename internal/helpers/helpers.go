package helpers

import (
	"fmt"
	"github.com/jackc/pgtype"
)

func UUIDToString(uuid pgtype.UUID) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid.Bytes[0:4], uuid.Bytes[4:6], uuid.Bytes[6:8], uuid.Bytes[8:10], uuid.Bytes[10:16])
}

func StringToUUID(uuid string) pgtype.UUID {
	var result pgtype.UUID
	_ = (&result).Set(uuid)
	return result
}
