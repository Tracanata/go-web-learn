package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Title": "Template if",
	})
}

func TestTemplateIf(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htpp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/perbandingan.gohtml"))

	t.ExecuteTemplate(writer, "perbandingan.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 70,
	})
}

func TestTemplateComparator(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htpp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Hobbies": []string{
			"Gaming", "Reading", "Coding",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htpp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Name": "Fatra",
		"Address": map[string]interface{}{
			"Street": "Jalan jalan",
			"City":   "Mars",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "htpp://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
