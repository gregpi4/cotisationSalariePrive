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
	cotisation_value := float32(100.23)
	for _, cotisation := range AllCotisations {
		mockFetcher.EXPECT().GetCotisation(cotisation.ToUrssaf(), utils.InfoEntreprise{Name: "my_company", SalarieCadre: true, ContratInformation: "CDI"}, float32(3000)).Return(cotisation_value, nil).AnyTimes()
		cotisation_value++
	}
	var timeStub = utils.NewTestTime()
	adapter := PayCotisations{
		urssafAdapter:  mockFetcher,
		localProvider:  mockFetcher,
		timeProvider:   &timeStub,
		infoEntreprise: utils.InfoEntreprise{Name: "my_company", SalarieCadre: true, ContratInformation: "CDI"},
	}
	cotisations := adapter.CotisationPatronaleForAPI(3000)

	expected_cotisation_value := float32(100.23)
	if cotisations["maladie_employeur"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["vieillesse_plafonnee_employeur"])
	}
	expected_cotisation_value++
	if cotisations["prevoyances_sante_employeur"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["prevoyances_sante_employeur"])
	}
	expected_cotisation_value++
	if cotisations["atmp"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["atmp"])
	}
	expected_cotisation_value++
	if cotisations["vieillesse_plafonnee_employeur"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["vieillesse_plafonnee_employeur"])
	}
	expected_cotisation_value++
	if cotisations["vieillesse_deplafonnee_employeur"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["vieillesse_deplafonnee_employeur"])
	}
	expected_cotisation_value++
	if cotisations["retraite_complementaire_employeur"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["retraite_complementaire_employeur"])
	}
	expected_cotisation_value++
	if cotisations["allocations_familiales_employeur"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["allocations_familiales_employeur"])
	}
	expected_cotisation_value++
	if cotisations["assurance_chomage_employeur"] != expected_cotisation_value {
		t.Fatalf(`want %f, got %f`, expected_cotisation_value, cotisations["assurance_chomage_employeur"])
	}
}
