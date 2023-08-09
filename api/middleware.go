package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/georgejawoods/Bank-app/token"

	"github.com/gin-gonic/gin"
)

const (
	autorizationHeaderKey   = "autorization"
	authorizationTypeBearer = "bearer"
	autorizationPayloadKey  = "autorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		autorizationHeader := ctx.GetHeader(autorizationHeaderKey)
		if len(autorizationHeader) == 0 {
			err := errors.New("autorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		fields := strings.Fields(autorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid autorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		autorizationType := strings.ToLower(fields[0])
		if autorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported autorization type %s", autorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		ctx.Set(autorizationPayloadKey, payload)
		ctx.Next()
	}
}
