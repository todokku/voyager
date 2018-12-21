// Code generated by go-swagger; DO NOT EDIT.

package store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/go-swagger/go-swagger/examples/contributed-templates/stratoscale/models"
)

// OrderGetReader is a Reader for the OrderGet structure.
type OrderGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OrderGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewOrderGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewOrderGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewOrderGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewOrderGetOK creates a OrderGetOK with default headers values
func NewOrderGetOK() *OrderGetOK {
	return &OrderGetOK{}
}

/*OrderGetOK handles this case with default header values.

successful operation
*/
type OrderGetOK struct {
	Payload *models.Order
}

func (o *OrderGetOK) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetOK  %+v", 200, o.Payload)
}

func (o *OrderGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Order)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOrderGetBadRequest creates a OrderGetBadRequest with default headers values
func NewOrderGetBadRequest() *OrderGetBadRequest {
	return &OrderGetBadRequest{}
}

/*OrderGetBadRequest handles this case with default header values.

Invalid ID supplied
*/
type OrderGetBadRequest struct {
}

func (o *OrderGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetBadRequest ", 400)
}

func (o *OrderGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewOrderGetNotFound creates a OrderGetNotFound with default headers values
func NewOrderGetNotFound() *OrderGetNotFound {
	return &OrderGetNotFound{}
}

/*OrderGetNotFound handles this case with default header values.

Order not found
*/
type OrderGetNotFound struct {
}

func (o *OrderGetNotFound) Error() string {
	return fmt.Sprintf("[GET /store/order/{orderId}][%d] orderGetNotFound ", 404)
}

func (o *OrderGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
