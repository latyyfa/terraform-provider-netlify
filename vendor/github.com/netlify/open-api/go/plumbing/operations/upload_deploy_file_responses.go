// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/netlify/open-api/go/models"
)

// UploadDeployFileReader is a Reader for the UploadDeployFile structure.
type UploadDeployFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadDeployFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadDeployFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUploadDeployFileDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUploadDeployFileOK creates a UploadDeployFileOK with default headers values
func NewUploadDeployFileOK() *UploadDeployFileOK {
	return &UploadDeployFileOK{}
}

/*UploadDeployFileOK handles this case with default header values.

OK
*/
type UploadDeployFileOK struct {
	Payload *models.File
}

func (o *UploadDeployFileOK) Error() string {
	return fmt.Sprintf("[PUT /deploys/{deploy_id}/files/{path}][%d] uploadDeployFileOK  %+v", 200, o.Payload)
}

func (o *UploadDeployFileOK) GetPayload() *models.File {
	return o.Payload
}

func (o *UploadDeployFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadDeployFileDefault creates a UploadDeployFileDefault with default headers values
func NewUploadDeployFileDefault(code int) *UploadDeployFileDefault {
	return &UploadDeployFileDefault{
		_statusCode: code,
	}
}

/*UploadDeployFileDefault handles this case with default header values.

error
*/
type UploadDeployFileDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the upload deploy file default response
func (o *UploadDeployFileDefault) Code() int {
	return o._statusCode
}

func (o *UploadDeployFileDefault) Error() string {
	return fmt.Sprintf("[PUT /deploys/{deploy_id}/files/{path}][%d] uploadDeployFile default  %+v", o._statusCode, o.Payload)
}

func (o *UploadDeployFileDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *UploadDeployFileDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
