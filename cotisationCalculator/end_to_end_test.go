//go:build integration
// +build integration

package main

import (
	adapter "cotisationCalculator/adapter"
	utils "cotisationCalculator/utils"
	"encoding/json"
	"log/slog"
	"testing"
)

func TestCalculateComposition(t *testing.T) {
	urssafAdapter := adapter.UrssafAdapter{
		Client:  adapter.CreateHTTPClient(),
		BaseURL: "https://mon-entreprise.urssaf.fr/api/v1",
	}
	var l = adapter.LocalPayCalculator{}
	payCalculator := PayCotisations{urssafAdapter, l, &utils.Time{}}
	cotisations := payCalculator.CotisationPatronaleForAPI(2000)
	jsonData, _ := json.Marshal(cotisations)
	slog.Info(string(jsonData))
	slog.Info("finish")
}
