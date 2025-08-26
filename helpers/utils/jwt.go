package utils

import (
	"inventory-service/model"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

type ClaimsToken struct {
	Id                        string
	Name                      string
	Email                     string
	Role                      string
	UserCompaniesReferencesId string
	CompaniesRefencesId       string
	EmployeeReferencesId      string
	Expire                    float64
}

func GenerateToken(user *model.VerifyCredentialRes, role string, exp int64) (string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["role"] = role
	claims["exp"] = exp

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte(viper.GetString("JWT_SECRET")))

	return t, err
}

func UseToken(token *jwt.Token) ClaimsToken {
	var result ClaimsToken

	claims := token.Claims.(jwt.MapClaims)

	if claims["user_id"] != nil {
		result.Id = claims["user_id"].(string)
		result.Name = claims["name"].(string)
		result.Role = claims["role"].(string)
	}

	return result

}

func GetApiCredential(c echo.Context) (*model.AccessApiCredential, bool) {
	apiId := c.Request().Header["Api-Secret-Id"]
	apiKey := c.Request().Header["Api-Secret-Key"]

	if len(apiId) == 0 || len(apiKey) == 0 {
		return nil, false
	}

	credential := &model.AccessApiCredential{
		ApiSecretId:  apiId[0],
		ApiSecretKey: apiKey[0],
	}

	return credential, true
}
