// Code generated by ogen, DO NOT EDIT.

package v1

import (
	"math/bits"
	"strconv"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"

	"github.com/ogen-go/ogen/validate"
)

// Encode implements json.Marshaler.
func (s *ErrorResponse) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *ErrorResponse) encodeFields(e *jx.Encoder) {
	{
		e.FieldStart("message")
		e.Str(s.Message)
	}
	{
		e.FieldStart("code")
		e.Int(s.Code)
	}
}

var jsonFieldsNameOfErrorResponse = [2]string{
	0: "message",
	1: "code",
}

// Decode decodes ErrorResponse from json.
func (s *ErrorResponse) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode ErrorResponse to nil")
	}
	var requiredBitSet [1]uint8

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "message":
			requiredBitSet[0] |= 1 << 0
			if err := func() error {
				v, err := d.Str()
				s.Message = string(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"message\"")
			}
		case "code":
			requiredBitSet[0] |= 1 << 1
			if err := func() error {
				v, err := d.Int()
				s.Code = int(v)
				if err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"code\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode ErrorResponse")
	}
	// Validate required fields.
	var failures []validate.FieldError
	for i, mask := range [1]uint8{
		0b00000011,
	} {
		if result := (requiredBitSet[i] & mask) ^ mask; result != 0 {
			// Mask only required fields and check equality to mask using XOR.
			//
			// If XOR result is not zero, result is not equal to expected, so some fields are missed.
			// Bits of fields which would be set are actually bits of missed fields.
			missed := bits.OnesCount8(result)
			for bitN := 0; bitN < missed; bitN++ {
				bitIdx := bits.TrailingZeros8(result)
				fieldIdx := i*8 + bitIdx
				var name string
				if fieldIdx < len(jsonFieldsNameOfErrorResponse) {
					name = jsonFieldsNameOfErrorResponse[fieldIdx]
				} else {
					name = strconv.Itoa(fieldIdx)
				}
				failures = append(failures, validate.FieldError{
					Name:  name,
					Error: validate.ErrFieldRequired,
				})
				// Reset bit.
				result &^= 1 << bitIdx
			}
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *ErrorResponse) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *ErrorResponse) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UploadStatusStatus as json.
func (o OptUploadStatusStatus) Encode(e *jx.Encoder) {
	if !o.Set {
		return
	}
	e.Str(string(o.Value))
}

// Decode decodes UploadStatusStatus from json.
func (o *OptUploadStatusStatus) Decode(d *jx.Decoder) error {
	if o == nil {
		return errors.New("invalid: unable to decode OptUploadStatusStatus to nil")
	}
	o.Set = true
	if err := o.Value.Decode(d); err != nil {
		return err
	}
	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s OptUploadStatusStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *OptUploadStatusStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode implements json.Marshaler.
func (s *UploadStatus) Encode(e *jx.Encoder) {
	e.ObjStart()
	s.encodeFields(e)
	e.ObjEnd()
}

// encodeFields encodes fields.
func (s *UploadStatus) encodeFields(e *jx.Encoder) {
	{
		if s.Status.Set {
			e.FieldStart("status")
			s.Status.Encode(e)
		}
	}
}

var jsonFieldsNameOfUploadStatus = [1]string{
	0: "status",
}

// Decode decodes UploadStatus from json.
func (s *UploadStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UploadStatus to nil")
	}

	if err := d.ObjBytes(func(d *jx.Decoder, k []byte) error {
		switch string(k) {
		case "status":
			if err := func() error {
				s.Status.Reset()
				if err := s.Status.Decode(d); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return errors.Wrap(err, "decode field \"status\"")
			}
		default:
			return d.Skip()
		}
		return nil
	}); err != nil {
		return errors.Wrap(err, "decode UploadStatus")
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s *UploadStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UploadStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}

// Encode encodes UploadStatusStatus as json.
func (s UploadStatusStatus) Encode(e *jx.Encoder) {
	e.Str(string(s))
}

// Decode decodes UploadStatusStatus from json.
func (s *UploadStatusStatus) Decode(d *jx.Decoder) error {
	if s == nil {
		return errors.New("invalid: unable to decode UploadStatusStatus to nil")
	}
	v, err := d.StrBytes()
	if err != nil {
		return err
	}
	// Try to use constant string.
	switch UploadStatusStatus(v) {
	case UploadStatusStatusReady:
		*s = UploadStatusStatusReady
	case UploadStatusStatusInProgress:
		*s = UploadStatusStatusInProgress
	case UploadStatusStatusFailed:
		*s = UploadStatusStatusFailed
	case UploadStatusStatusDeleted:
		*s = UploadStatusStatusDeleted
	default:
		*s = UploadStatusStatus(v)
	}

	return nil
}

// MarshalJSON implements stdjson.Marshaler.
func (s UploadStatusStatus) MarshalJSON() ([]byte, error) {
	e := jx.Encoder{}
	s.Encode(&e)
	return e.Bytes(), nil
}

// UnmarshalJSON implements stdjson.Unmarshaler.
func (s *UploadStatusStatus) UnmarshalJSON(data []byte) error {
	d := jx.DecodeBytes(data)
	return s.Decode(d)
}
