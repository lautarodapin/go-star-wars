package starships

const LIST_URL = "https://swapi.dev/api/starships"
const DETAIL_URL = "https://swapi.dev/api/starships/%s"

type Starship struct {
	Name                 *string   `json:"name"`
	Model                *string   `json:"model"`
	Manufacturer         *string   `json:"manufacturer"`
	CostInCredits        *string   `json:"cost_in_credits"`
	Length               *string   `json:"length"`
	MaxAtmospheringSpeed *string   `json:"max_atmosphering_speed"`
	Crew                 *string   `json:"crew"`
	Passengers           *string   `json:"passengers"`
	CargoCapacity        *string   `json:"cargo_capacity"`
	Consumables          *string   `json:"consumables"`
	HyperdriveRating     *string   `json:"hyperdrive_rating"`
	MGLT                 *string   `json:"MGLT"`
	StarshipClass        *string   `json:"starship_class"`
	Pilots               *[]string `json:"pilots"`
	Films                *[]string `json:"films"`
	Detail               *string   `json:"url"`
}
