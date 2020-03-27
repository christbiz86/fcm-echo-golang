package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

var hosts = "http://10.1.35.36:8353/sfpas/session/validate"
var responseData map[string]interface{}

type Session struct {
	Session_id string `json:session_id`
}

func SessionMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var session_id Session
		decoder := json.NewDecoder(r.Body)
		decoder.Decode(&session_id)
		getMdn := GetSession(session_id.Session_id)
		if getMdn != nil {
			q := getMdn.(map[string]interface{})
			ctx := context.WithValue(r.Context(), "sessionMdn", q["MDN"])
			next.ServeHTTP(w, r.WithContext(ctx))

			c := http.Cookie{
				Name:   "CookieData",
				Value:  q["MDN"].(string),
				Expires: time.Now().Add(5 * time.Minute),
			}
			http.SetCookie(w, &c)
		} else {
			http.Error(w, "Failed get MDN!!!", http.StatusUnauthorized)
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