package intervals_union

import "sort"

type EndpointWithPosition struct {
	Endpoint
	IsLeft bool
}

func UnionOfIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return []Interval{}
	}

	endpoints := make([]EndpointWithPosition, 0, 2*len(intervals))
	for _, i := range intervals {
		endpoints = append(endpoints,
			EndpointWithPosition{Endpoint: i.Left, IsLeft: true},
			EndpointWithPosition{Endpoint: i.Right, IsLeft: false},
		)
	}

	sort.Slice(endpoints, func(i, j int) bool {
		return lessEndpoint(endpoints[i], endpoints[j])
	})

	result := make([]Interval, 0)

	var currentLeft Endpoint
	currentlyOpen := 0
	for _, e := range endpoints {
		if e.IsLeft {
			if currentlyOpen == 0 {
				currentLeft = e.Endpoint
			}
			currentlyOpen++
		} else {
			currentlyOpen--
			if currentlyOpen == 0 {
				result = append(result, Interval{Left: currentLeft, Right: e.Endpoint})
			}
		}
	}

	return result
}

func lessEndpoint(e1, e2 EndpointWithPosition) bool {
	if e1.Val != e2.Val {
		return e1.Val < e2.Val
	}

	if e1.IsLeft == e2.IsLeft {
		if e1.IsLeft {
			// for left endpoints, closed come before open
			return e1.IsClosed
		} else {
			// for right endpoints, open come before closed
			return !e1.IsClosed
		}
	}

	// there is one left and one right endpoint
	if !e1.IsClosed && !e2.IsClosed {
		// for open endpoints we want to have the closing before the
		// opening as they don't intersect
		return !e1.IsLeft
	}

	// for one closed endpoint we always want the left endpoint before
	// the right one so that they intersect
	return e1.IsLeft
}
