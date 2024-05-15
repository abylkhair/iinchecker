package dtos

import (
	"encoding/json"
	"net/http"
	"reflect"
)

const (
	OkStatus   = "true"
	FailStatus = "false"
)

type ResponseOptions struct {
}

type ResponseDto struct {
	ResponseOptions

	Data    []interface{} `json:"data"`
	Success string        `json:"success"`
	Errors  string        `json:"errors,omitempty"`

	w http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *ResponseDto {
	return &ResponseDto{
		Success: OkStatus,
		Data:    make([]interface{}, 0),
		w:       w,
	}
}

func (r *ResponseDto) SetError(msg string) *ResponseDto {
	r.Errors = msg
	return r
}

func (r *ResponseDto) SetData(data interface{}) {
	r.Data = r.reflectData(data)
}

func (r *ResponseDto) SetHeaders(headers map[string]string) *ResponseDto {
	for k, v := range headers {
		r.w.Header().Set(k, v)
	}
	return r
}

func (r *ResponseDto) SetStatus(status int) *ResponseDto {
	r.w.WriteHeader(status)
	return r
}

func (r *ResponseDto) JSON() error {
	r.w.Header().Set("Content-Type", "application/json; charset=utf-8")

	data, err := json.Marshal(r)
	if err != nil {
		r.Errors = "error marshal" // FIXME
	}

	_, err = r.w.Write(data)
	if err != nil {
		r.Errors = "error write" // FIXME
	}

	if len(r.Errors) != 0 {
		r.Success = FailStatus
	}

	return nil
}

func (r *ResponseDto) reflectData(in interface{}) []interface{} {
	sType := reflect.ValueOf(in)

	if sType.Kind() != reflect.Slice {
		return []interface{}{}
	}

	ret := make([]interface{}, sType.Len())

	for i := 0; i < sType.Len(); i++ {
		ret[i] = sType.Index(i).Interface()
	}

	return ret
}
