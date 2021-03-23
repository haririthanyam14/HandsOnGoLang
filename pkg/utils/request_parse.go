package utils

import (
	"HandsOnGoLang/pkg/liberror"
	"encoding/json"
	"errors"
	"net/http"
)

func ParseRequest(req *http.Request, data interface{}) (err error) {
	if req == nil {
		return liberror.Builder().SetOperation("ParseRequest").SetKind(liberror.ValidationError).SetCause(errors.New("request is nil")).Build()
	}

	if req.Body == nil {
		return liberror.Builder().SetOperation("ParseRequest").SetKind(liberror.ValidationError).SetCause(errors.New("request body is nil")).Build()
	}

	decoder := json.NewDecoder(req.Body)

	err = decoder.Decode(&data)
	if err != nil {
		return liberror.Builder().SetOperation("ParseRequest.decoder.Decode").SetKind(liberror.ValidationError).SetCause(err).Build()
	}

	return
}
