package middlewares

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MicahParks/keyfunc/v2"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Jwt() gin.HandlerFunc {
	// Get the JWKS URL.
	//
	// This is a sample JWKS service. Visit https://jwks-service.appspot.com/ and grab a token to test this example.
	jwksURL := os.Getenv("AUTH0_JWKS_URI")

	// Create the JWKS from the resource at the given URL.
	// jwks, err := keyfunc.Get(jwksURL, options)

	jwks, err := keyfunc.Get(jwksURL, keyfunc.Options{
		RefreshInterval:  time.Hour,
		RefreshRateLimit: time.Minute * 5,
		RefreshTimeout:   time.Second * 10,
	})

	if err != nil {
		log.Fatalf("Failed to create JWKS from resource at the given URL.\nError: %s", err.Error())
	}

	return func(c *gin.Context) {

		// Retrieve JWT from the "Authorization" header
		authHeader := c.GetHeader("Authorization")
		splitToken := strings.Split(authHeader, "Bearer ")

		if len(splitToken) != 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		tokenString := splitToken[1]

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			c.Abort()
			return
		}

		// Get a JWT to parse.
		jwtB64 := "eyJraWQiOiJlZThkNjI2ZCIsInR5cCI6IkpXVCIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJXZWlkb25nIiwiYXVkIjoiVGFzaHVhbiIsImlzcyI6Imp3a3Mtc2VydmljZS5hcHBzcG90LmNvbSIsImlhdCI6MTYzMTM2OTk1NSwianRpIjoiNDY2M2E5MTAtZWU2MC00NzcwLTgxNjktY2I3NDdiMDljZjU0In0.LwD65d5h6U_2Xco81EClMa_1WIW4xXZl8o4b7WzY_7OgPD2tNlByxvGDzP7bKYA9Gj--1mi4Q4li4CAnKJkaHRYB17baC0H5P9lKMPuA6AnChTzLafY6yf-YadA7DmakCtIl7FNcFQQL2DXmh6gS9J6TluFoCIXj83MqETbDWpL28o3XAD_05UP8VLQzH2XzyqWKi97mOuvz-GsDp9mhBYQUgN3csNXt2v2l-bUPWe19SftNej0cxddyGu06tXUtaS6K0oe0TTbaqc3hmfEiu5G0J8U6ztTUMwXkBvaknE640NPgMQJqBaey0E4u0txYgyvMvvxfwtcOrDRYqYPBnA"

		// Parse the JWT.
		token, err := jwt.Parse(jwtB64, jwks.Keyfunc)
		if err != nil {
			// retorne o erro "Failed to parse the JWT.\nError: %s"
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Failed to parse the JWT: " + err.Error()})
			c.Abort()
			return
		}

		// Check if the token is valid.
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "The token is not valid."})
			c.Abort()
			return
		}
		// // End the background refresh goroutine when it's no longer needed.
		// cancel()

		// // This will be ineffectual because the line above this canceled the parent context.Context.
		// // This method call is idempotent similar to context.CancelFunc.
		// jwks.EndBackground()
	}
}
