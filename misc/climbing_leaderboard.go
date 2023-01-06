// We have a leaderboard that works like this:
// The player with the highest score is ranked number  1
// on the leaderboard.
// Players who have equal scores receive the same ranking number, and
// the next player(s) receive the immediately following ranking number.

//climbingLeaderboard has the following parameter(s):

//    int ranked[n]: the leaderboard scores
//    int player[m]: the player's scores

//Returns:

//    int[m]: the player's rank after each new score

package misc

import (
	"github.com/pratikluitel/o_riley_algorithm_design_manual_excercises/utils"
)

func climbingLeaderboard(ranked []int32, player []int32) []int32 {
	// Write your code here
	var new_rank []int32
	new_rank = make([]int32, len(player))

	combined := append(ranked, player...) // combining the two slices and sorting

	// removing duplicate repitition
	nodup_combined := []int32{}
	for _, val := range combined {
		in, _ := utils.IsIn(nodup_combined, val)
		if !in {
			nodup_combined = append(nodup_combined, val)
		}
	}

	nodup_combined, _ = utils.Sort(nodup_combined)

	for idx, val := range nodup_combined {
		in, i := utils.IsIn(player, val)
		if in {
			new_rank[i] = int32(idx) + 1
		}
	}
	return new_rank
}
