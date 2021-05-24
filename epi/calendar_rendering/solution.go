package calendar_rendering

import "sort"

type EventEnd struct {
	value int
	start bool
}

func FindMaxSimultaneousEvents(a []Event) int {
	ends := make([]EventEnd, 0, 2*len(a))
	for _, e := range a {
		ends = append(ends, EventEnd{
			value: e.Start,
			start: true,
		})
		ends = append(ends, EventEnd{
			value: e.Finish,
			start: false,
		})
	}

	sort.Slice(ends, func(i, j int) bool {
		if ends[i].value < ends[j].value {
			return true
		}
		if ends[i].value > ends[j].value {
			return false
		}
		if ends[i].start != ends[j].start {
			return ends[i].start // we want the start to appear before the end for equal value
		}
		return false // the structs are equal
	})

	maxOVerlapping := 0
	currentOverlapping := 0
	for _, e := range ends {
		if e.start {
			currentOverlapping++
			if currentOverlapping > maxOVerlapping {
				maxOVerlapping = currentOverlapping
			}
		}
		if !e.start {
			currentOverlapping--
		}
	}

	return maxOVerlapping
}
