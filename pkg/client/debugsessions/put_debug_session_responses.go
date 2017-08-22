package debugsessions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/solo-io/squash/pkg/models"
)

// PutDebugSessionReader is a Reader for the PutDebugSession structure.
type PutDebugSessionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutDebugSessionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPutDebugSessionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutDebugSessionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutDebugSessionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 412:
		result := NewPutDebugSessionPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPutDebugSessionUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutDebugSessionCreated creates a PutDebugSessionCreated with default headers values
func NewPutDebugSessionCreated() *PutDebugSessionCreated {
	return &PutDebugSessionCreated{}
}

/*PutDebugSessionCreated handles this case with default header values.

Debug session created
*/
type PutDebugSessionCreated struct {
	Payload *models.DebugSession
}

func (o *PutDebugSessionCreated) Error() string {
	return fmt.Sprintf("[PUT /debugconfig/{debugConfigId}/session][%d] putDebugSessionCreated  %+v", 201, o.Payload)
}

func (o *PutDebugSessionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DebugSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutDebugSessionBadRequest creates a PutDebugSessionBadRequest with default headers values
func NewPutDebugSessionBadRequest() *PutDebugSessionBadRequest {
	return &PutDebugSessionBadRequest{}
}

/*PutDebugSessionBadRequest handles this case with default header values.

Invalid ID supplied
*/
type PutDebugSessionBadRequest struct {
}

func (o *PutDebugSessionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /debugconfig/{debugConfigId}/session][%d] putDebugSessionBadRequest ", 400)
}

func (o *PutDebugSessionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDebugSessionNotFound creates a PutDebugSessionNotFound with default headers values
func NewPutDebugSessionNotFound() *PutDebugSessionNotFound {
	return &PutDebugSessionNotFound{}
}

/*PutDebugSessionNotFound handles this case with default header values.

debug config not found
*/
type PutDebugSessionNotFound struct {
}

func (o *PutDebugSessionNotFound) Error() string {
	return fmt.Sprintf("[PUT /debugconfig/{debugConfigId}/session][%d] putDebugSessionNotFound ", 404)
}

func (o *PutDebugSessionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDebugSessionPreconditionFailed creates a PutDebugSessionPreconditionFailed with default headers values
func NewPutDebugSessionPreconditionFailed() *PutDebugSessionPreconditionFailed {
	return &PutDebugSessionPreconditionFailed{}
}

/*PutDebugSessionPreconditionFailed handles this case with default header values.

Debug state currently exists
*/
type PutDebugSessionPreconditionFailed struct {
}

func (o *PutDebugSessionPreconditionFailed) Error() string {
	return fmt.Sprintf("[PUT /debugconfig/{debugConfigId}/session][%d] putDebugSessionPreconditionFailed ", 412)
}

func (o *PutDebugSessionPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutDebugSessionUnprocessableEntity creates a PutDebugSessionUnprocessableEntity with default headers values
func NewPutDebugSessionUnprocessableEntity() *PutDebugSessionUnprocessableEntity {
	return &PutDebugSessionUnprocessableEntity{}
}

/*PutDebugSessionUnprocessableEntity handles this case with default header values.

Invalid input
*/
type PutDebugSessionUnprocessableEntity struct {
}

func (o *PutDebugSessionUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /debugconfig/{debugConfigId}/session][%d] putDebugSessionUnprocessableEntity ", 422)
}

func (o *PutDebugSessionUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}