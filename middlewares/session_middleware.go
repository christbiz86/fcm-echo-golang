package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"fcm-golang/services"
	"time"
	"net/http"
	"github.com/dgrijalva/jwt-go"
)

var hosts = "http://10.1.35.36:8353/sfpas/session/validate"
var responseData map[string]interface{}

type Session struct {
	Session_id string `json:session_id`
}

type jwtCustomClaims struct {
	SessionMdn  string `json:"sessionMdn"`
	jwt.StandardClaims
}

func SessionMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session_id Session
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&session_id)
		getMdn := GetSession(session_id.Session_id)
		q := getMdn.(map[string]interface{})
		if q["MDN"] != nil{
			ctx := context.WithValue(r.Context(), "sessionMdn", q["MDN"])
			next.ServeHTTP(w, r.WithContext(ctx))

			claims := &jwtCustomClaims{
				q["MDN"].(string),
				jwt.StandardClaims{
					ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
				},
			}

			// Create token with claims
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			// Generate encoded token and send it as response.
			token.SignedString([]byte(services.GoDotEnvVariable("JWT_KEY")))
			
			// save to session
			//store, _ := session.Start(context.Background(), w, r)
			//store.Set("foo", q["MDN"])
			//store.Save()
		}
	})
}

func Sfpas(sfpasUrl string, sessionData map[string]string, email string, passwd string) map[string]interface{} {
	postData := make(map[string]map[string]string)
	postData["session"] = sessionData
	if(email != "" && passwd != ""){
		postData["collection"]["id"] = email
		postData["collection"]["token"] = passwd
	}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(postData)
	res, _ := http.Post(sfpasUrl, "application/json; charset=utf-8", b)
	json.NewDecoder(res.Body).Decode(&responseData)
	return responseData
}

func GetSession(sessionId string) interface{} {
	data := make(map[string]string)
	data["session_id"] = sessionId
	response := Sfpas(hosts, data, "", "")
	if(response["error"].(float64) >= 400){
		return nil
	} else {
		return response["session"]
	}
}