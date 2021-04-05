package max_teams_in_photograph

import "github.com/stefantds/go-epi-judge/test_utils"

func FindLargestNumberTeams(graph []GraphVertex) int {
	var maxLevel int
	for _, g := range graph {
		if g.MaxDistance == 0 {
			maxLevel = test_utils.Max(maxLevel, dfs(&g))
		}
	}

	return maxLevel
}

func dfs(g *GraphVertex) int {
	g.MaxDistance = 1
	for _, v := range g.Edges {
		var vMaxDist int
		if v.MaxDistance == 0 {
			vMaxDist = dfs(v)
		} else {
			vMaxDist = v.MaxDistance
		}
		g.MaxDistance = test_utils.Max(g.MaxDistance, vMaxDist+1)
	}

	return g.MaxDistance
}
