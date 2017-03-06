package controllers

import (
	"fmt"

	"github.com/lara-go/app/app/http/requests"
	"github.com/lara-go/larago/http"
	"github.com/lara-go/larago/http/errors"
	"github.com/lara-go/larago/http/responses"
)

// Form request.
type Form struct {
	Text string `schema:"text" json:"text"`
}

// APIController struct.
type APIController struct{}

// Index sends plain text as a response.
func (c *APIController) Index(request *http.Request) string {
	return fmt.Sprintf("%s: Index", request.Method())
}

// Redirect does redirect via route name and param.
func (c *APIController) Redirect() responses.Response {
	return responses.NewRedirect(302).Route("foobar", map[string]string{"bar": "baz"})
}

// FooBar accepts bar param and returns plain text.
func (c *APIController) FooBar(bar string) responses.Response {
	return responses.NewText(200, "Foo Bar: %s", bar)
}

// BarBaz accepts param, requesting it from unmarshalling to struct.
func (c *APIController) BarBaz(request *http.Request) responses.Response {
	var form Form
	request.ReadParams(&form)

	return responses.NewText(200, "Bar Baz: %s", form.Text)
}

// Form accepts POST form request.
func (c *APIController) Form(request *http.Request) responses.Response {
	var form Form
	request.ReadForm(&form)

	return responses.NewHTML(200, "Form: <br>%+v", form)
}

// JSON returns JSON formatted application/json response.
func (c *APIController) JSON(request *requests.PostNews) responses.Response {
	return responses.NewJSON(200, map[string]string{"foo": request.Text})
}

// HTTPError sends prepared 404 HTTPError with additional meta info.
func (c *APIController) HTTPError() *errors.HTTPError {
	return errors.NotFoundHTTPError().
		WithMeta(map[string]string{"foo": "Bar"}).
		WithContext("Custom context. Can be anything.")
}

// CustomError returns an error instance that will be converted to InternalServerError.
func (c *APIController) CustomError() error {
	return fmt.Errorf("Custom error")
}

// Panic sends some generic panic that will be catched and converted to InternalServerError.
func (c *APIController) Panic() {
	panic("Panic!")
}
