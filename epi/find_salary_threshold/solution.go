package find_salary_threshold

import "sort"

func FindSalaryCap(targetPayroll int, currentSalaries []int) float64 {
	sort.Ints(currentSalaries)
	n := len(currentSalaries)

	unadjustedSum := 0
	for i, v := range currentSalaries {
		adjustedPeople := n - i
		adjustedSum := adjustedPeople * v
		if unadjustedSum+adjustedSum >= targetPayroll {
			return float64(targetPayroll-unadjustedSum) / float64(adjustedPeople)
		}

		unadjustedSum += v
	}

	return -1
}
