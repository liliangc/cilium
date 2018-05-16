// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetHelloReader is a Reader for the GetHello structure.
type GetHelloReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHelloReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHelloOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHelloOK creates a GetHelloOK with default headers values
func NewGetHelloOK() *GetHelloOK {
	return &GetHelloOK{}
}

/*GetHelloOK handles this case with default header values.

Success
*/
type GetHelloOK struct {
}

func (o *GetHelloOK) Error() string {
	return fmt.Sprintf("[GET /hello][%d] getHelloOK ", 200)
}

func (o *GetHelloOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
