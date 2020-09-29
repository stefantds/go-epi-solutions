package task_pairing

import "sort"

type Task = int

func OptimumTaskAssignment(taskDurations []Task) [][2]Task {
	sort.Ints(taskDurations)

	optimumAssignment := make([][2]Task, 0)

	for i, j := 0, len(taskDurations)-1; i < j; i, j = i+1, j-1 {
		optimumAssignment = append(optimumAssignment, [2]Task{taskDurations[i], taskDurations[j]})
	}

	return optimumAssignment
}
