package is_array_dominated

import "sort"

func ValidPlacementExists(team0 Team, team1 Team) bool {
	sort.Slice(team0.Players, func(i int, j int) bool {
		return team0.Players[i].Height < team0.Players[j].Height
	})

	sort.Slice(team1.Players, func(i int, j int) bool {
		return team1.Players[i].Height < team1.Players[j].Height
	})

	for i := 0; i < len(team0.Players) && i < len(team1.Players); i++ {
		if team0.Players[i].Height >= team1.Players[i].Height {
			return false
		}
	}

	return true
}
