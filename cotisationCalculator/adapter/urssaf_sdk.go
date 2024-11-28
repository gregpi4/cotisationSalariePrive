package adapter

import (
	"bytes"
	utils "cotisationCalculator/utils"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"runtime"
)

type UrssafAdapter struct {
	Client  *http.Client
	BaseURL string
}

type WTF struct {
	Number int
}

func CreateHTTPClient(certPath string) *http.Client {
	rootCAs, _ := x509.SystemCertPool()
	certs, _ := os.ReadFile(certPath)
	rootCAs.AppendCertsFromPEM(certs)
	tlsConfig := &tls.Config{
		RootCAs: rootCAs,
	}

	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{
		Transport: transport,
	}
	return client
}

func (payDataProvider UrssafAdapter) getCotisationsComposition() ([]string, error) {
	// Create a new HTTP client with the custom transport
	resp, err := payDataProvider.Client.Get("https://mon-entreprise.urssaf.fr/api/v1/rules/salarié%20.%20cotisations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// Process the response here...
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err.Error())
		return make([]string, 0), errors.New("read error")
	}
	var jsonBody map[string]interface{}
	json.Unmarshal(body, &jsonBody)

	generic_cot_employeur := jsonBody["rawNode"].(map[string]interface{})["avec"].(map[string]interface{})["employeur"].(map[string]interface{})["somme"].([]interface{})
	cot_employer := make([]string, len(generic_cot_employeur))
	for index, cot := range generic_cot_employeur {
		cot_employer[index] = cot.(string)
	}
	return cot_employer, nil
}

func (payDataProvider UrssafAdapter) getCotisationInformation(cotisation string) {
	// Create a new HTTP client with the custom transport
	u, err := url.Parse(payDataProvider.BaseURL + "/rules/salarié . cotisations . " + url.PathEscape(string(cotisation)))
	if err != nil {
		panic(err)
	}

	slog.Debug(u.String())

	resp, err := payDataProvider.Client.Get(u.String())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// Process the response
	body, err := io.ReadAll(resp.Body)

	slog.Info(string(body))
}

type Salaire struct {
	Valeur float32 `json:"valeur"`
	Unite  string  `json:"unité"`
}

type Situation struct {
	EntrpriseCategorieJuridique string  `json:"entreprise . catégorie juridique"`
	SalarieContratSalaireBrut   Salaire `json:"salarié . contrat . salaire brut"`
	SalarieContrat              string  `json:"salarié . contrat"`
	SalarieStatutCadre          string  `json:"salarié . contrat . statut cadre"`
}

type Expressions struct {
	SalarieRemunerationNetAPayerAvantImpot string `json:"salarié . rémunération . net . à payer avant impôt"`
}

type Data struct {
	Situation   Situation `json:"situation"`
	Expressions []string  `json:"expressions"`
}

func parseStatutCadre(statutCadre bool) string {
	if statutCadre {
		return "oui"
	}
	return "non"
}

func (payDataProvider UrssafAdapter) GetCotisation(cotisation string, infoEntreprise utils.InfoEntreprise, salaire float32) (ret float32, ret_err error) {
	// example post body
	// {
	// 	"situation": {
	// 	  "salarié . contrat . salaire brut": {
	// 		"valeur": 4200,
	// 		"unité": "€/mois"
	// 	  },
	// 	  "salarié . contrat": "'CDI'"
	// 	},
	// 	"expressions": [
	// 	  "salarié . rémunération . net . à payer avant impôt"
	// 	]
	//   }

	// Prepare the request body
	statutCadre := parseStatutCadre(infoEntreprise.SalarieCadre)
	data := Data{Situation: Situation{SalarieStatutCadre: statutCadre, EntrpriseCategorieJuridique: "'SARL'", SalarieContratSalaireBrut: Salaire{Valeur: salaire, Unite: "€/mois"}, SalarieContrat: "'CDI'"}, Expressions: []string{"salarié . cotisations . " + string(cotisation)}}

	jsonData, err := json.Marshal(data)
	if err != nil {
		slog.Error("Error marshalling JSON:" + err.Error())
		return 0, err
	}

	reader := bytes.NewReader(jsonData)

	// Send the POST request
	resp, err := payDataProvider.Client.Post("https://mon-entreprise.urssaf.fr/api/v1/evaluate", "application/json", reader)
	if err != nil {
		slog.Error(err.Error())
		return 0, err
	}
	defer resp.Body.Close()

	// Process the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err.Error())
		return 0, errors.New("read error")
	}
	var jsonBody map[string]interface{}
	json.Unmarshal(body, &jsonBody)

	var error_unparsing = errors.New("unparsing failed")
	error_unparsing = nil

	defer func() {
		if r := recover(); r != nil {
			error_unparsing = errors.New("bad json treatment: " + string(r.(runtime.Error).Error()))
			ret = 0
			ret_err = error_unparsing
		}
	}()
	var cotisationCalculated = jsonBody["evaluate"].([]interface{})[0].(map[string]interface{})["nodeValue"]

	if cotisationCalculated == nil {
		return 0, errors.New("cotisation not found")
	}

	return float32(cotisationCalculated.(float64)), error_unparsing
}

func (payDataProvider UrssafAdapter) getRule(rule string) {
	resp, err := payDataProvider.Client.Get("https://mon-entreprise.urssaf.fr/api/v1/rules/" + rule)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	// Process the response here...
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read error:", err)
		return
	}

	fmt.Println(string(body))
}
