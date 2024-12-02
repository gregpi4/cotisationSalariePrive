//go:build adapter
// +build adapter

package adapter

import (
	"cotisationCalculator/data"
	"testing"
)

func TestGetCotisationsComposition(t *testing.T) {
	adapter := UrssafAdapter{
		Client:  CreateHTTPClient("marketware-root-cert.pem"),
		BaseURL: "https://mon-entreprise.urssaf.fr/api/v1",
	}
	body, _ := adapter.getCotisationsComposition()
	t.Log(body)

}

func TestGetCotisationInformation(t *testing.T) {
	adapter := UrssafAdapter{
		Client:  CreateHTTPClient("marketware-root-cert.pem"),
		BaseURL: "https://mon-entreprise.urssaf.fr/api/v1",
	}
	adapter.getRule("salarié . contrat . statut cadre")
}

func TestGetCotisation(t *testing.T) {
	adapter := UrssafAdapter{
		Client:  CreateHTTPClient("marketware-root-cert.pem"),
		BaseURL: "https://mon-entreprise.urssaf.fr/api/v1",
	}
	value, err := adapter.GetCotisation("maladie . employeur", data.InfoEntreprise{Name: "my-company"}, float32(1000))
	if value != 70 || err != nil {
		t.Errorf("got %f, wanted %f", value, float32(70))
	}
}
