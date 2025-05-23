// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudPvminstancesVolumesPutReader is a Reader for the PcloudPvminstancesVolumesPut structure.
type PcloudPvminstancesVolumesPutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudPvminstancesVolumesPutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudPvminstancesVolumesPutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudPvminstancesVolumesPutBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudPvminstancesVolumesPutUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudPvminstancesVolumesPutForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudPvminstancesVolumesPutNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPcloudPvminstancesVolumesPutConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudPvminstancesVolumesPutInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}] pcloud.pvminstances.volumes.put", response, response.Code())
	}
}

// NewPcloudPvminstancesVolumesPutOK creates a PcloudPvminstancesVolumesPutOK with default headers values
func NewPcloudPvminstancesVolumesPutOK() *PcloudPvminstancesVolumesPutOK {
	return &PcloudPvminstancesVolumesPutOK{}
}

/*
PcloudPvminstancesVolumesPutOK describes a response with status code 200, with default header values.

OK
*/
type PcloudPvminstancesVolumesPutOK struct {
	Payload models.Object
}

// IsSuccess returns true when this pcloud pvminstances volumes put o k response has a 2xx status code
func (o *PcloudPvminstancesVolumesPutOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud pvminstances volumes put o k response has a 3xx status code
func (o *PcloudPvminstancesVolumesPutOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes put o k response has a 4xx status code
func (o *PcloudPvminstancesVolumesPutOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes put o k response has a 5xx status code
func (o *PcloudPvminstancesVolumesPutOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes put o k response a status code equal to that given
func (o *PcloudPvminstancesVolumesPutOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pcloud pvminstances volumes put o k response
func (o *PcloudPvminstancesVolumesPutOK) Code() int {
	return 200
}

func (o *PcloudPvminstancesVolumesPutOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutOK %s", 200, payload)
}

func (o *PcloudPvminstancesVolumesPutOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutOK %s", 200, payload)
}

func (o *PcloudPvminstancesVolumesPutOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesPutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesPutBadRequest creates a PcloudPvminstancesVolumesPutBadRequest with default headers values
func NewPcloudPvminstancesVolumesPutBadRequest() *PcloudPvminstancesVolumesPutBadRequest {
	return &PcloudPvminstancesVolumesPutBadRequest{}
}

/*
PcloudPvminstancesVolumesPutBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudPvminstancesVolumesPutBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes put bad request response has a 2xx status code
func (o *PcloudPvminstancesVolumesPutBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes put bad request response has a 3xx status code
func (o *PcloudPvminstancesVolumesPutBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes put bad request response has a 4xx status code
func (o *PcloudPvminstancesVolumesPutBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes put bad request response has a 5xx status code
func (o *PcloudPvminstancesVolumesPutBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes put bad request response a status code equal to that given
func (o *PcloudPvminstancesVolumesPutBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the pcloud pvminstances volumes put bad request response
func (o *PcloudPvminstancesVolumesPutBadRequest) Code() int {
	return 400
}

func (o *PcloudPvminstancesVolumesPutBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVolumesPutBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutBadRequest %s", 400, payload)
}

func (o *PcloudPvminstancesVolumesPutBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesPutBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesPutUnauthorized creates a PcloudPvminstancesVolumesPutUnauthorized with default headers values
func NewPcloudPvminstancesVolumesPutUnauthorized() *PcloudPvminstancesVolumesPutUnauthorized {
	return &PcloudPvminstancesVolumesPutUnauthorized{}
}

/*
PcloudPvminstancesVolumesPutUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudPvminstancesVolumesPutUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes put unauthorized response has a 2xx status code
func (o *PcloudPvminstancesVolumesPutUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes put unauthorized response has a 3xx status code
func (o *PcloudPvminstancesVolumesPutUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes put unauthorized response has a 4xx status code
func (o *PcloudPvminstancesVolumesPutUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes put unauthorized response has a 5xx status code
func (o *PcloudPvminstancesVolumesPutUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes put unauthorized response a status code equal to that given
func (o *PcloudPvminstancesVolumesPutUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the pcloud pvminstances volumes put unauthorized response
func (o *PcloudPvminstancesVolumesPutUnauthorized) Code() int {
	return 401
}

func (o *PcloudPvminstancesVolumesPutUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVolumesPutUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutUnauthorized %s", 401, payload)
}

func (o *PcloudPvminstancesVolumesPutUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesPutUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesPutForbidden creates a PcloudPvminstancesVolumesPutForbidden with default headers values
func NewPcloudPvminstancesVolumesPutForbidden() *PcloudPvminstancesVolumesPutForbidden {
	return &PcloudPvminstancesVolumesPutForbidden{}
}

/*
PcloudPvminstancesVolumesPutForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudPvminstancesVolumesPutForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes put forbidden response has a 2xx status code
func (o *PcloudPvminstancesVolumesPutForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes put forbidden response has a 3xx status code
func (o *PcloudPvminstancesVolumesPutForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes put forbidden response has a 4xx status code
func (o *PcloudPvminstancesVolumesPutForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes put forbidden response has a 5xx status code
func (o *PcloudPvminstancesVolumesPutForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes put forbidden response a status code equal to that given
func (o *PcloudPvminstancesVolumesPutForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the pcloud pvminstances volumes put forbidden response
func (o *PcloudPvminstancesVolumesPutForbidden) Code() int {
	return 403
}

func (o *PcloudPvminstancesVolumesPutForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVolumesPutForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutForbidden %s", 403, payload)
}

func (o *PcloudPvminstancesVolumesPutForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesPutForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesPutNotFound creates a PcloudPvminstancesVolumesPutNotFound with default headers values
func NewPcloudPvminstancesVolumesPutNotFound() *PcloudPvminstancesVolumesPutNotFound {
	return &PcloudPvminstancesVolumesPutNotFound{}
}

/*
PcloudPvminstancesVolumesPutNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudPvminstancesVolumesPutNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes put not found response has a 2xx status code
func (o *PcloudPvminstancesVolumesPutNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes put not found response has a 3xx status code
func (o *PcloudPvminstancesVolumesPutNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes put not found response has a 4xx status code
func (o *PcloudPvminstancesVolumesPutNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes put not found response has a 5xx status code
func (o *PcloudPvminstancesVolumesPutNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes put not found response a status code equal to that given
func (o *PcloudPvminstancesVolumesPutNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the pcloud pvminstances volumes put not found response
func (o *PcloudPvminstancesVolumesPutNotFound) Code() int {
	return 404
}

func (o *PcloudPvminstancesVolumesPutNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVolumesPutNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutNotFound %s", 404, payload)
}

func (o *PcloudPvminstancesVolumesPutNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesPutNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesPutConflict creates a PcloudPvminstancesVolumesPutConflict with default headers values
func NewPcloudPvminstancesVolumesPutConflict() *PcloudPvminstancesVolumesPutConflict {
	return &PcloudPvminstancesVolumesPutConflict{}
}

/*
PcloudPvminstancesVolumesPutConflict describes a response with status code 409, with default header values.

Conflict
*/
type PcloudPvminstancesVolumesPutConflict struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes put conflict response has a 2xx status code
func (o *PcloudPvminstancesVolumesPutConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes put conflict response has a 3xx status code
func (o *PcloudPvminstancesVolumesPutConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes put conflict response has a 4xx status code
func (o *PcloudPvminstancesVolumesPutConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud pvminstances volumes put conflict response has a 5xx status code
func (o *PcloudPvminstancesVolumesPutConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud pvminstances volumes put conflict response a status code equal to that given
func (o *PcloudPvminstancesVolumesPutConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the pcloud pvminstances volumes put conflict response
func (o *PcloudPvminstancesVolumesPutConflict) Code() int {
	return 409
}

func (o *PcloudPvminstancesVolumesPutConflict) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutConflict %s", 409, payload)
}

func (o *PcloudPvminstancesVolumesPutConflict) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutConflict %s", 409, payload)
}

func (o *PcloudPvminstancesVolumesPutConflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesPutConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudPvminstancesVolumesPutInternalServerError creates a PcloudPvminstancesVolumesPutInternalServerError with default headers values
func NewPcloudPvminstancesVolumesPutInternalServerError() *PcloudPvminstancesVolumesPutInternalServerError {
	return &PcloudPvminstancesVolumesPutInternalServerError{}
}

/*
PcloudPvminstancesVolumesPutInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudPvminstancesVolumesPutInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud pvminstances volumes put internal server error response has a 2xx status code
func (o *PcloudPvminstancesVolumesPutInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud pvminstances volumes put internal server error response has a 3xx status code
func (o *PcloudPvminstancesVolumesPutInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud pvminstances volumes put internal server error response has a 4xx status code
func (o *PcloudPvminstancesVolumesPutInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud pvminstances volumes put internal server error response has a 5xx status code
func (o *PcloudPvminstancesVolumesPutInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud pvminstances volumes put internal server error response a status code equal to that given
func (o *PcloudPvminstancesVolumesPutInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the pcloud pvminstances volumes put internal server error response
func (o *PcloudPvminstancesVolumesPutInternalServerError) Code() int {
	return 500
}

func (o *PcloudPvminstancesVolumesPutInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVolumesPutInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /pcloud/v1/cloud-instances/{cloud_instance_id}/pvm-instances/{pvm_instance_id}/volumes/{volume_id}][%d] pcloudPvminstancesVolumesPutInternalServerError %s", 500, payload)
}

func (o *PcloudPvminstancesVolumesPutInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudPvminstancesVolumesPutInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
