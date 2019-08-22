// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	errors "github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	strfmt "github.com/go-openapi/strfmt"
	swag "github.com/go-openapi/swag"
	validate "github.com/go-openapi/validate"
)

// AddTodoHandlerFunc turns a function with the right signature into a add todo handler
type AddTodoHandlerFunc func(AddTodoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTodoHandlerFunc) Handle(params AddTodoParams) middleware.Responder {
	return fn(params)
}

// AddTodoHandler interface for that can handle valid add todo params
type AddTodoHandler interface {
	Handle(AddTodoParams) middleware.Responder
}

// NewAddTodo creates a new http.Handler for the add todo operation
func NewAddTodo(ctx *middleware.Context, handler AddTodoHandler) *AddTodo {
	return &AddTodo{Context: ctx, Handler: handler}
}

/*AddTodo swagger:route POST /todos addTodo

Add a new todo to the list

*/
type AddTodo struct {
	Context *middleware.Context
	Handler AddTodoHandler
}

func (o *AddTodo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddTodoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

// AddTodoBadRequestBody add todo bad request body
// swagger:model AddTodoBadRequestBody
type AddTodoBadRequestBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this add todo bad request body
func (o *AddTodoBadRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddTodoBadRequestBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("addTodoBadRequest"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddTodoBadRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddTodoBadRequestBody) UnmarshalBinary(b []byte) error {
	var res AddTodoBadRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddTodoInternalServerErrorBody add todo internal server error body
// swagger:model AddTodoInternalServerErrorBody
type AddTodoInternalServerErrorBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this add todo internal server error body
func (o *AddTodoInternalServerErrorBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddTodoInternalServerErrorBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("addTodoInternalServerError"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddTodoInternalServerErrorBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddTodoInternalServerErrorBody) UnmarshalBinary(b []byte) error {
	var res AddTodoInternalServerErrorBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// AddTodoOKBody add todo o k body
// swagger:model AddTodoOKBody
type AddTodoOKBody struct {

	// message
	// Required: true
	Message *string `json:"message"`
}

// Validate validates this add todo o k body
func (o *AddTodoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddTodoOKBody) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("addTodoOK"+"."+"message", "body", o.Message); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddTodoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddTodoOKBody) UnmarshalBinary(b []byte) error {
	var res AddTodoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
