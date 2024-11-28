package main

import (
	"testing"

	mocks "cotisationCalculator/mocks"

	utils "cotisationCalculator/utils"

	"github.com/golang/mock/gomock"
)

func TestGetCotisationsComposition(t *testing.T) {
	// Create a new controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockFetcher := mocks.NewMockPayDataProvider(ctrl)
	for _, cotisation := range AllCotisations {
		mockFetcher.EXPECT().GetCotisation(cotisation.ToUrssaf(), utils.InfoEntreprise{Name: "my_company", SalarieCadre: true, ContratInformation: "CDI"}, float32(3000)).Return(float32(100.23), nil).AnyTimes()
	}
	var timeStub = utils.NewTestTime()
	adapter := PayCotisations{
		urssafAdapter:  mockFetcher,
		localProvider:  mockFetcher,
		timeProvider:   &timeStub,
		infoEntreprise: utils.InfoEntreprise{Name: "my_company", SalarieCadre: true, ContratInformation: "CDI"},
	}
	cotisations := adapter.CotisationPatronaleForAPI(3000)

	if cotisations["allocations_familiales_employeur"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
	if cotisations["assurance_chomage_employeur"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
	if cotisations["atmp"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
	if cotisations["maladie_employeur"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
	if cotisations["prevoyances_sante_employeur"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
	if cotisations["retraite_complementaire_employeur"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
	if cotisations["vieillesse_plafonnee_employeur"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
	if cotisations["vieillesse_deplafonnee_employeur"] != float32(100.23) {
		t.Fatalf(`want %v, error`, "e")
	}
}
