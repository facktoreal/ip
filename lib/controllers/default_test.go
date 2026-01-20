package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	"github.com/facktoreal/ip/lib/controllers"
	"github.com/facktoreal/ip/lib/services"
	"github.com/facktoreal/ip/lib/views"
	"html/template"
	"io"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func TestDefaultController_Public(t *testing.T) {
	// Setup
	os.Setenv("HOSTNAME", "http://localhost:8080")
	defer os.Unsetenv("HOSTNAME")

	e := echo.New()
	e.Renderer = &Template{
		templates: template.Must(template.New("index.html").Parse(views.DefaultLayout)),
	}
	statsSrv := services.NewStatsService()
	ctl := controllers.NewDefaultController(statsSrv)

	t.Run("JSON response via Accept header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctl.Public(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), `{"ip":`)
		}
	})

	t.Run("XML response via Accept header", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationXML)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctl.Public(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), `<PublicIpResponse>`)
			assert.Contains(t, rec.Body.String(), `<ip>`)
		}
	})

	t.Run("XML response via format query param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?format=xml", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctl.Public(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), `<PublicIpResponse>`)
		}
	})

	t.Run("XML response via xml query param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?xml=1", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctl.Public(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), `<PublicIpResponse>`)
		}
	})

	t.Run("Plain text response via format query param", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/?format=text", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctl.Public(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			// RemoteAddr is empty in httptest.NewRequest by default, but RealIP() handles it
			assert.NotEmpty(t, rec.Body.String())
		}
	})

	t.Run("HTML response (default)", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if assert.NoError(t, ctl.Public(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Contains(t, rec.Body.String(), "<!DOCTYPE html>")
			assert.Contains(t, rec.Body.String(), "Your Public IP Address")
		}
	})
}
