package middlewares

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
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
		ctx := context.WithValue(r.Context(), "session", getMdn)
		next.ServeHTTP(w, r.WithContext(ctx))
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
	log.Println(responseData)
	return responseData
}

func GetSession(sessionId string) interface{} {
	data := make(map[string]string)
	data["session_id"] = sessionId
	response := Sfpas(hosts, data, "", "")
	log.Println("tes resulr middleware")
	log.Println(response["error"])
	if(response["error"].(float64) >= 400){
		return nil
	} else {
		return response["session"]
	}
}