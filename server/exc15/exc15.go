// There are n kids with candies. You are given an integer array candies,
//where each candies[i] represents the number of candies the ith kid has,
//and an integer extraCandies, denoting the number of extra candies that you have.

// Return a boolean array result of length n, where result[i] is true if,
//after giving the ith kid all the extraCandies,
//they will have the greatest number of candies among all the kids, or false otherwise.

// Note that multiple kids can have the greatest number of candies.

package exc15

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxcandies := 0
	for _, candy := range candies {
		if candy > maxcandies {
			maxcandies = candy
		}
	}

	result := make([]bool, len(candies))
	for i, candy := range candies {
		result[i] = (candy + extraCandies) >= maxcandies
	}
	return result
}
