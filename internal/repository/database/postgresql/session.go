package database

import (
	"context"
	"log"
	"social_network/internal/config/database"
	"social_network/internal/pkg/jwt"
)

var db = config.ConnectDB()

func CreateSessions(ctx context.Context, payload *jwt.PayloadJWT) error {
	insertSession := `INSERT INTO session (expiresrefresh,refreshuid,refreshtoken)
					VALUES($1,$2,$3)`

	_, err := db.Exec(ctx, insertSession, payload.ExpiresRefresh, payload.RefreshUID, payload.RefreshToken)

	if err != nil {
		log.Println(err, " :[ERROR] Insert JWT Payload")
		return err
	}

	return nil
}

func UpdateSessions(ctx context.Context, payload *jwt.PayloadJWT) error {
	insertSession := `UPDATE session (expiresrefresh,refreshuid,refreshtoken)
					VALUES($1,$2,$3)`

	_, err := db.Exec(ctx, insertSession, payload.ExpiresRefresh, payload.RefreshUID, payload.RefreshToken)

	if err != nil {
		log.Println(err, " :[ERROR] Update JWT Payload")
		return err
	}

	return nil
}
