// Code generated by ogen, DO NOT EDIT.

package v1

// SetFake set fake values.
func (s *ErrorResponse) SetFake() {
	{
		{
			s.Message = "string"
		}
	}
	{
		{
			s.Code = int(0)
		}
	}
}

// SetFake set fake values.
func (s *OptUploadStatusStatus) SetFake() {
	var elem UploadStatusStatus
	{
		elem.SetFake()
	}
	s.SetTo(elem)
}

// SetFake set fake values.
func (s *UploadStatus) SetFake() {
	{
		{
			s.Status.SetFake()
		}
	}
}

// SetFake set fake values.
func (s *UploadStatusStatus) SetFake() {
	*s = UploadStatusStatusReady
}
