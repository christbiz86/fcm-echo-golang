package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http"
)

var hosts = "http://10.1.35.36:8353"
var userAgent = "MySmartfren"
var sfpasUrl = "/sfpas/session/validate"
var responseData map[string]interface{}

func session_middleware(w http.ResponseWriter, r http.Request) error{

	

	//$msg = $request->getAttribute('input');
	//
	//if (!isset($msg['session_id'])) {
	//throw new Exception('sessionId missing', 1);
	//}
	//
	//$session = $this->getSession($msg['session_id']);
	//
	//$request = $request->withAttribute('session', $session);
	//
	//$response = $next($request, $response);
	//return $response
}

func sfpas(sfpasUrl string, sessionData map[string]string, email string, passwd string) map[string]interface{} {
	postData := make(map[string]map[string]string)
	postData["session"] = sessionData
	if(email != "" && passwd != ""){
		postData["collection"]["id"] = email
		postData["collection"]["token"] = passwd
	}
	requestData, _ := json.Marshal(postData)
	body := bytes.NewReader(requestData)
	req, err := http.NewRequest("POST", sfpasUrl, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", string(len(requestData)))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		json.NewDecoder(resp.Body).Decode(&responseData)
	}
	defer resp.Body.Close()
	return responseData
}

func getSession(sessionId string) interface{} {
	data := make(map[string]string)
	data["session_id"] = sessionId
	response := sfpas(sfpasUrl, data, "", "")
	if(response["http_data"].(int) >= 400){
		return nil
	}
	q := response["payload"].(map[string]interface{})
	return q["session"]
}