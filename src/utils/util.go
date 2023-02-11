package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/bytedance/sonic"
	"github.com/valyala/fastjson"
)

var Logger *log.Logger

func HandleError(err error) {
	fmt.Println("Error:", err)
}

func ValidateRequest(r *http.Request) error {
	if r.Method != http.MethodPost {
		return fmt.Errorf("Invalid request method")
	}
	if r.Header.Get("Content-Type") != "application/json" {
		return fmt.Errorf("Invalid Content-Type")
	}
	return nil
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	WriteJSON(w, statusCode, fastjson.MustParse(fmt.Sprintf(`{"error": "%s"}`, message)))
}

func WriteJSON(w http.ResponseWriter, statusCode int, data *fastjson.Value) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	_, err := w.Write(data.MarshalTo(nil))
	if err != nil {
		HandleError(err)
	}
}

func LogRequest(r *http.Request, start time.Time) {
	Logger.Printf("%s %s %v\n", r.Method, r.URL.Path, time.Since(start))
}

func InitLogger() {
	Logger = log.New(os.Stdout, "", log.LstdFlags)
}

func WriteJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(data)
	if err != nil {
		HandleError(err)
	}
}

func EncodeJSON(v interface{}) (*fastjson.Value, error) {
	var p fastjson.Parser
	b, err := sonic.Marshal(v)
	if err != nil {
		return nil, err
	}
	return p.ParseBytes(b)
}

func DecodeJSON(r io.ReadCloser, v interface{}) error {
	var p fastjson.Parser
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	value, err := p.ParseBytes(b)
	if err != nil {
		return err
	}
	return sonic.Unmarshal(value.MarshalTo(nil), v)
}
