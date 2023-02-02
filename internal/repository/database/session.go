package database


import(
	"context"
	"social_network/internal/pkg/jwt"
	"social_network/internal/config/database"
	"log"

)

func CreateSessions(ctx context.Context, payload *jwt.PayloadJWT) error {
	db := config.ConnectDB()
	insertSession := `INSERT INTO session (expiresrefresh,refreshuid,refreshtoken)
					VALUES($1,$2,$3)`

	_, err := db.Exec(ctx, insertSession, payload.ExpiresRefresh, payload.RefreshUID, payload.RefreshToken)

	if err != nil {
		log.Println(err, " :[ERROR] Insert JWT Payload")
		return err
	}

	return nil
}
