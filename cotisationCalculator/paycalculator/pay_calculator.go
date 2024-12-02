package paycalculator

import (
	"log/slog"
	"strconv"
	"time"

	"cotisationCalculator/data"
	utils "cotisationCalculator/utils"
)

type PayDataProvider interface {
	GetCotisation(cotisation string, infoEntreprise data.InfoEntreprise, salaire float32) (float32, error)
}

type Cotisation interface {
	ToUrssaf() string
	ToString() string
}

type CotisationProvider interface {
	GetCotisation(cotisation string) (float64, error)
}

// enum of employer cotisation
type CotisationPatronale string

const (
	CotisationPatronaleMaladie                CotisationPatronale = "maladie_employeur"
	CotisationPatronalePrevoyanceSante        CotisationPatronale = "prevoyances_sante_employeur"
	CotisationPatronaleATMP                   CotisationPatronale = "atmp"
	CotisationPatronaleVieillessePlafonnee    CotisationPatronale = "vieillesse_plafonnee_employeur"
	CotisationPatronaleVieillesseDeplafonee   CotisationPatronale = "vieillesse_deplafonnee_employeur"
	CotisationPatronaleRetraiteComplementaire CotisationPatronale = "retraite_complementaire_employeur"
	CotisationPatronaleAllocationFamiliale    CotisationPatronale = "allocations_familiales_employeur"
	CotisationPatronaleAssuranceChomage       CotisationPatronale = "assurance_chomage_employeur"
	CotisationPatronaleAPEC                   CotisationPatronale = "apec"
	CotisationPatronaleIncapaciteInvalidite   CotisationPatronale = "incapacite_invalidite_employeur"
	CotisationPatronaleAutres                 CotisationPatronale = "autres_employeur"
)

// enum of employee cotisation
type CotisationSalariale string

const (
	CotisationSalarialeMaladie                CotisationSalariale = "maladie_salarie"
	CotisationSalarialePrevoyanceSante        CotisationSalariale = "prevoyances_sante_salarie"
	CotisationSalarialeVieillessePlafonnee    CotisationSalariale = "vieillesse_plafonnee_salarie"
	CotisationSalarialeVieillesseDeplafonee   CotisationSalariale = "vieillesse_deplafonnee_salarie"
	CotisationSalarialeRetraiteComplementaire CotisationSalariale = "retraite_complementaire_salarie"
	CotisationSalarialeAPEC                   CotisationSalariale = "apec"
	CotisationSalarialeCSGDeductible          CotisationSalariale = "csg_deductible_salarie"
	CotisationSalarialeCSGCDRSImposable       CotisationSalariale = "csg_imposable_salarie"
	CotisationSalarialeCSGCDRSNonImposable    CotisationSalariale = "csg_non_imposable_salarie"
	CotisationSalarialeIncapaciteInvalidite   CotisationSalariale = "incapacite_invalidite_salarie"
	CotisationSalarialeAutres                 CotisationSalariale = "autres_salarie"
)

var cotisation_patronale_model_to_urssaf_map = map[CotisationPatronale]string{
	CotisationPatronaleMaladie:                "maladie . employeur",
	CotisationPatronalePrevoyanceSante:        "prévoyances . santé . employeur",
	CotisationPatronaleATMP:                   "ATMP",
	CotisationPatronaleVieillessePlafonnee:    "vieillesse . plafonnée . employeur",
	CotisationPatronaleVieillesseDeplafonee:   "vieillesse . déplafonnée . employeur",
	CotisationPatronaleRetraiteComplementaire: "retraite complémentaire-CEG-CET . employeur",
	CotisationPatronaleAllocationFamiliale:    "allocations familiales",
	CotisationPatronaleAssuranceChomage:       "assurance chômage",
	CotisationPatronaleAPEC:                   "APEC . employeur",
	CotisationPatronaleIncapaciteInvalidite:   "prévoyances . incapacité invalidité décès . employeur",
	CotisationPatronaleAutres:                 "autres employeur",
}

func (c CotisationPatronale) ToUrssaf() string {
	return cotisation_patronale_model_to_urssaf_map[c]
}

func (c CotisationSalariale) ToUrssaf() string {
	// TODO: implement cotisation_salariale_model_to_urssaf_map
	return string(c)
}

func (c CotisationPatronale) ToString() string {
	return string(c)
}

func (c CotisationSalariale) ToString() string {
	return string(c)
}

type PayCotisations struct {
	UrssafAdapter  PayDataProvider
	LocalProvider  PayDataProvider
	TimeProvider   utils.TimeOperations
	InfoEntreprise data.InfoEntreprise
}

type job struct {
	cotisation string
	value      float32
}

type jobError struct {
	cotisation string
	err        error
}

var AllCotisations = []CotisationPatronale{
	CotisationPatronaleMaladie,
	CotisationPatronalePrevoyanceSante,
	CotisationPatronaleATMP,
	CotisationPatronaleVieillessePlafonnee,
	CotisationPatronaleVieillesseDeplafonee,
	CotisationPatronaleRetraiteComplementaire,
	CotisationPatronaleAllocationFamiliale,
	CotisationPatronaleAssuranceChomage,
}

type PayCotisationsInterface interface {
	CotisationPatronaleForAPI(salaire int) map[string]float64
}

func (p *PayCotisations) cotisationValue(delay time.Duration, cotisation Cotisation, salaire float32, resultChannel chan job, errChan chan<- jobError) {
	slog.Info("calculate " + cotisation.ToString())

	p.TimeProvider.Sleep(delay)
	value, err := p.UrssafAdapter.GetCotisation(cotisation.ToUrssaf(), p.InfoEntreprise, salaire)
	if err != nil {
		value, err = p.LocalProvider.GetCotisation(cotisation.ToUrssaf(), p.InfoEntreprise, salaire)
		if err != nil {
			slog.Info("envoie dans canal pour" + cotisation.ToString())
			errChan <- jobError{cotisation: cotisation.ToString(), err: err}
			return
		}
	}
	slog.Info("envoie dans canal pour " + cotisation.ToString())
	resultChannel <- job{cotisation: cotisation.ToString(), value: value}
}

func (p *PayCotisations) CotisationPatronaleForAPI(salaire float32) map[string]float32 {
	var delay = time.Second
	var nbLine = 0

	var resultChannel = make(chan job)
	var errChan = make(chan jobError)
	cotisations := make(map[string]float32)
	for _, cotisation := range AllCotisations {
		go p.cotisationValue(time.Duration(nbLine)*delay, cotisation, salaire, resultChannel, errChan)
		nbLine += 1
	}

	for i := 0; i < nbLine; i++ {
		select {
		case result := <-resultChannel:
			cotisations[result.cotisation] = result.value
		case jobErr := <-errChan:
			if jobErr.err != nil {
				slog.Error(jobErr.cotisation + jobErr.err.Error())
			}
		case <-time.After(1 * time.Minute):
			slog.Error("timeout for index " + strconv.Itoa(i))
		}
	}
	return cotisations
}
