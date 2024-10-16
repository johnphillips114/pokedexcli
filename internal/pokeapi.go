package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Location struct {
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	GameIndicies []struct {
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
		GameIndex uint `json:"game_index"`
	} `json:"game_indices"`
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

type LocationArea struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	GameIndex         int    `json:"game_index"`
	PokemonEncounters []struct {
		VersionDetails []struct {
			EncounterDetails []struct {
				Method struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				MinLevel        int   `json:"min_level"`
				MaxLevel        int   `json:"max_level"`
				Chance          int   `json:"chance"`
			} `json:"encounter_details"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance int `json:"max_chance"`
		} `json:"version_details"`
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
	EncounterMethodRates []struct {
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			Rate int `json:"rate"`
		} `json:"version_details"`
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
	} `json:"encounter_method_rates"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
}

type Berry struct {
	Flavors []struct {
		Potency uint `json:"potency"`
		Flavor  struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"flavor"`
	} `json:"flavors"`
	Firmness struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"firmness"`
	Item struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"item"`
	NaturalGiftType struct {
		Name string `json:"name"`
		url  string `json:"URL"`
	} `json:"natural_gift_type"`
	Name             string `json:"name:`
	ID               uint   `json:"id"`
	GrowthTime       uint   `json:"growth_time"`
	MaxHarvest       uint   `json:"max_harvest"`
	NaturalGiftPower uint   `json:"natural_gift_power"`
	Size             uint   `json:"size"`
	Smoothness       uint   `json:"smoothness"`
	SoilDryness      uint   `json:"soil_dryness"`
}

type BerryFirmness struct {
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	Berries []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"berries"`
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type BerryFlavor struct {
	Berries []struct {
		Berry struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"berry"`
		Potency uint `json:"potency"`
	} `json:"berries"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	ContestType struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"contest_type"`
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

type ContestType struct {
	Names []struct {
		Name     string `json:"name"`
		Color    string `json:"color"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	BerryFlavor struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"berry_flavor"`
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

type ContestEffect struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	EffectEntries []struct {
		Effect    string `json:"effect"`
		Lanaguage struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"effect_entries"`
	ID     uint `json:"id"`
	Appeal uint `json:"appeal"`
	Jam    uint `json:"jam"`
}

type SuperContestEffect struct {
	FlavorTextEntries []struct {
		FlavorText string `json:"flavor_text"`
		Language   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"flavor_text_entries"`
	Moves []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"moves"`
	ID     uint `json:"id"`
	Appeal uint `json:"appeal"`
}

type EncounterMethod struct {
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	Name  string `json:"name"`
	ID    uint   `json:"id"`
	Order uint   `json:"order"`
}

type EncounterCondition struct {
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	Values []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"values"`
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

func GetLocations(ID int) ([]LocationArea, error) {
	locationAreas := make([]LocationArea, 20)
	endpoint := "https://pokeapi.co/api/v2/location-area/" + string(ID)
	res, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(body, &locationAreas)
	if err != nil {
		return nil, err
	}
	//for i := 0; i < 20; i++ {
	//	var location LocationArea
	//	locationAreas[i] = location
	//}

	return locationAreas, nil
}
