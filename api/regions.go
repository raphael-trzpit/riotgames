package api

// Region is a string representing a server region.
type Region string

const (
	DefaultHost = "api.riotgames.com"

	Brazil            Region = "br1"
	EuropeNorthEast          = "eun1"
	EuropeWest               = "euw1"
	Japan                    = "jp1"
	Korea                    = "kr"
	LatinAmericaNorth        = "la1"
	LatinAmericaSouth        = "la2"
	NorthAmerica             = "na1"
	Oceania                  = "oc1"
	Turkey                   = "tr1"
	Russia                   = "ru"
	PBE                      = "pbe1"
	Americas                 = "americas"
	Asia                     = "asia"
	Europe                   = "europe"
)

var (
	Regions = []Region{
		Brazil,
		EuropeNorthEast,
		EuropeWest,
		Japan,
		Korea,
		LatinAmericaNorth,
		LatinAmericaSouth,
		NorthAmerica,
		Oceania,
		Turkey,
		Russia,
		PBE,
		Americas,
		Asia,
		Europe,
	}
)

func (r Region) GetHost() string {
	return string(r) + "." + DefaultHost
}
