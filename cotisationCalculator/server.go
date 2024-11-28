package main

import (
	"cotisationCalculator/adapter"
	"cotisationCalculator/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
)

type PayApi struct {
	PayService PayCotisations
}

type GenerateFormData struct {
	EmployeeSalary string `json:"employee_salary"`
}

func (p PayApi) generateCotisations(w http.ResponseWriter, r *http.Request) {
	slog.Info("Request GET /generate")
	var form GenerateFormData

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		w.Header()["Content-Type"] = []string{"application/json"}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid form"}`))
		return
	}

	// Validate mandatory fields
	if form.EmployeeSalary == "" {
		w.Header()["Content-Type"] = []string{"application/json"}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Mandatory fields are missing"}`))
		return
	}

	// Process the form data
	employee_salary, err := strconv.ParseFloat(form.EmployeeSalary, 32)
	if err != nil {
		w.Header()["Content-Type"] = []string{"application/json"}
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid employee salary"}`))
		return
	}

	// Calculate cotisations
	var result = make(map[string]interface{})
	result["cotisations_patronales"] = p.PayService.CotisationPatronaleForAPI(float32(employee_salary))
	w.Header()["Content-Type"] = []string{"application/json"}
	w.WriteHeader(http.StatusOK)
	resultData, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal server error"}`))
		return
	}
	w.Write(resultData)

	slog.Info("Response GET /generate")
}

func (p PayApi) generateHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		p.generateCotisations(w, r)
	default:
		http.Error(w, "405 method not allowed.", http.StatusMethodNotAllowed)
	}
}

func main() {
	infoEntreprise := utils.InfoEntreprise{Name: "my_company", ContratInformation: "CDI", SalarieCadre: true}

	urssafAdapter := adapter.UrssafAdapter{
		Client:  adapter.CreateHTTPClient("adapter/marketware-root-cert.pem"),
		BaseURL: "https://mon-entreprise.urssaf.fr/api/v1",
	}
	l := adapter.LocalPayCalculator{Name: "hello"}
	payCalculator := PayCotisations{urssafAdapter, l, &utils.Time{}, infoEntreprise}
	fmt.Println(payCalculator)

	t := PayApi{PayService: payCalculator}

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	http.HandleFunc("/generate", t.generateHandler)

	//http.HandleFunc("/", resinput.GetProductInputs)

	http.ListenAndServe(":8080", nil)

}
