package adapter

import (
	utils "cotisationCalculator/utils"
	"errors"
)

type LocalPayCalculator struct {
	// ...
	Name string
}

func (l LocalPayCalculator) GetCotisation(cotisation string, infoEntreprise utils.InfoEntreprise, salaire float32) (float32, error) {
	if cotisation == "APEC . employeur" {
		//Cela advient quand l'employe n'est pas un cadre
		return 0, nil
	} else {
		return 0, errors.New(cotisation + " not found")
	}
}
