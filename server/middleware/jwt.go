package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
	"server/model"
	"server/utilities"
	"strings"
	"time"
)


func CreateToken(user *model.User) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "btlweb2020") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["UserCode"] = user.UserCode
	atClaims["UserName"] = user.UserName
	atClaims["RoleCode"] = user.RoleCode
	atClaims["FullName"] = user.FullName
	atClaims["PassWord"] = user.PassWord
	atClaims["Email"] = user.Email
	atClaims["exp"] = time.Now().Add(time.Minute * 150).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return "bearToken " + token, nil
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	utilities.InfoLog.Printf("Token split: %v", strArr )
	if len(strArr) == 2 {
		utilities.InfoLog.Printf("len: %v\n", len(strArr) )
		return strArr[1]
	}
	return ""
}

func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := ExtractToken(r)
	utilities.InfoLog.Printf("VerifyToken => tokenString: %s\n",tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		utilities.InfoLog.Printf("ACCESS_SECRET => tokenString: %s\n",os.Getenv("ACCESS_SECRET"))
		return []byte("btlweb2020"), nil
	})
	if err != nil {
		return nil, err
	}
	utilities.InfoLog.Printf("VerifyToken => tokenString: %s\n",tokenString)
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*model.UserInfo, error) {
	token, err := VerifyToken(r)
	utilities.InfoLog.Printf("ExtractTokenMetadata => token: %s\n",token)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	utilities.InfoLog.Printf("MapClaims: %v\n", claims)
	if ok && token.Valid {
		utilities.InfoLog.Printf("Token is valid\n")
		return &model.UserInfo{
			UserName:  claims["UserName"].(string),
			RoleCode:  claims["RoleCode"].(string),
			FullName:  claims["FullName"].(string),
			PassWord:  claims["PassWord"].(string),
			Email:     claims["Email"].(string),
		}, nil
	}
	return nil, err
}