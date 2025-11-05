package util

import (
	"base_app/auth"
	"base_app/constants/api_constants"
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GetUserIdAndRequestIdFromContext(ctx context.Context) (string, string) {
	claims, _ := ctx.Value(auth.Claims).(jwt.MapClaims)
	var userId string
	if claims["sub"] != nil {
		userId = fmt.Sprintf("%v", claims["sub"])
	}

	var requestId string
	reqIDRaw := ctx.Value(api_constants.RequestId) // reqIDRaw at this point is of type 'interface{}'

	reqID, ok := reqIDRaw.(string)
	if !ok {
		requestId = uuid.New().String()
	} else {
		requestId = reqID
	}

	return userId, requestId
}
