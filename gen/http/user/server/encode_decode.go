// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// user HTTP server encoders and decoders
//
// Command:
// $ goa gen
// github.com/akm/goa_v2_post_payload_including_snake_case_fields/design

package server

import (
	"context"
	"io"
	"net/http"

	userviews "github.com/akm/goa_v2_post_payload_including_snake_case_fields/gen/user/views"
	goa "goa.design/goa"
	goahttp "goa.design/goa/http"
)

// EncodeCreateResponse returns an encoder for responses returned by the user
// create endpoint.
func EncodeCreateResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*userviews.User)
		enc := encoder(ctx, w)
		body := NewCreateResponseBody(res.Projected)
		w.WriteHeader(http.StatusCreated)
		return enc.Encode(body)
	}
}

// DecodeCreateRequest returns a decoder for requests sent to the user create
// endpoint.
func DecodeCreateRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			body struct {
				Firstname *string
				Lastname  *string
			}
			err error
		)
		err = decoder(r).Decode(&body)
		if err != nil {
			if err == io.EOF {
				return nil, goa.MissingPayloadError()
			}
			return nil, goa.DecodePayloadError(err.Error())
		}
		if body.Firstname == nil {
			err = goa.MergeErrors(err, goa.MissingFieldError("firstname", "body"))
		}
		if err != nil {
			return nil, err
		}
		payload := NewCreateUserPayload(body)

		return payload, nil
	}
}
