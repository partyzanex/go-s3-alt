// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"io"
	"net/http"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/ogen-go/ogen/conv"
	ht "github.com/ogen-go/ogen/http"
	"github.com/ogen-go/ogen/uri"
)

func encodeDownloadFileResponse(response DownloadFileRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *DownloadFileOKHeaders:
		w.Header().Set("Content-Type", "application/octet-stream")
		// Encoding response headers.
		{
			h := uri.NewHeaderEncoder(w.Header())
			// Encode "Content-Disposition" header.
			{
				cfg := uri.HeaderParameterEncodingConfig{
					Name:    "Content-Disposition",
					Explode: false,
				}
				if err := h.EncodeParam(cfg, func(e uri.Encoder) error {
					if val, ok := response.ContentDisposition.Get(); ok {
						return e.EncodeValue(conv.StringToString(val))
					}
					return nil
				}); err != nil {
					return errors.Wrap(err, "encode Content-Disposition header")
				}
			}
		}
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		writer := w
		if _, err := io.Copy(writer, response.Response); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ErrorResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeGetFileStatusResponse(response GetFileStatusRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UploadStatus:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ErrorResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(404)
		span.SetStatus(codes.Error, http.StatusText(404))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeUploadFileResponse(response UploadFileRes, w http.ResponseWriter, span trace.Span) error {
	switch response := response.(type) {
	case *UploadStatus:
		if err := func() error {
			if err := response.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			return errors.Wrap(err, "validate")
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		span.SetStatus(codes.Ok, http.StatusText(200))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	case *ErrorResponse:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(400)
		span.SetStatus(codes.Error, http.StatusText(400))

		e := new(jx.Encoder)
		response.Encode(e)
		if _, err := e.WriteTo(w); err != nil {
			return errors.Wrap(err, "write")
		}

		return nil

	default:
		return errors.Errorf("unexpected response type: %T", response)
	}
}

func encodeErrorResponse(response *ErrorResponseStatusCode, w http.ResponseWriter, span trace.Span) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	code := response.StatusCode
	if code == 0 {
		// Set default status code.
		code = http.StatusOK
	}
	w.WriteHeader(code)
	if st := http.StatusText(code); code >= http.StatusBadRequest {
		span.SetStatus(codes.Error, st)
	} else {
		span.SetStatus(codes.Ok, st)
	}

	e := new(jx.Encoder)
	response.Response.Encode(e)
	if _, err := e.WriteTo(w); err != nil {
		return errors.Wrap(err, "write")
	}

	if code >= http.StatusInternalServerError {
		return errors.Wrapf(ht.ErrInternalServerErrorResponse, "code: %d, message: %s", code, http.StatusText(code))
	}
	return nil

}
