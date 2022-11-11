/*
	Code Author	: github.com/pratikluitel
	Book		: The Algorithm Design Manual, written by Steven S Skiena
	Chapter 	: 1

	Q			: Implement the two TSP heuristics of Section 1.1 (page 5). Which of them
				  gives better solutions in practice? Can you devise a heuristic that works better
				  than both of them?

	Problem		: Robot Tour Optimization
	Input		: A set S of n points in the plane.
	Output		: What is the shortest cycle tour that visits each point in the set S?
*/

package main

import (
	"fmt"
	"math"
)

type Node struct { // A node is a point in a plane where the robot needs to move
	x float64
	y float64
}

type ChainableNode struct {
	x            float64
	y            float64
	connected_to *ChainableNode // this is used  in chaining nodes
}

type Visit struct { // represents a point and it's visit status
	point         Node
	visited       bool    // the bool denotes if the node/point has been visited
	zero_distance float64 // distance from the 0 point, so the round trip can be made
}

func eucledian_distance(p1, p2 Node) float64 {
	distance := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))

	return distance
}

func eucledian_distance_chainable(p1, p2 ChainableNode) float64 {
	distance := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))

	return distance
}

// Heuristic 1	: Nearest Neighbor Heuristic

func tour_nearest_neighbor(set []Node) float64 { // The assumption here is that the 1st element of the set is the start point
	fmt.Printf("Starting the tour using nearest neighbor heuristic.\n")

	visits := []Visit{}
	visit_idx := 0 // visit the start node/point first
	current_visit := Visit{set[visit_idx], true, 0.0}

	// state initialization
	for idx, point := range set {
		current_distance := eucledian_distance(current_visit.point, point)
		visited := idx == 0 //the first point is set to as visited
		visits = append(visits, Visit{point, visited, current_distance})
	}

	fmt.Printf("Visiting point %v\n\n", current_visit.point)
	fmt.Printf("Current state:\n%+v\n\n", visits)

	all_nodes_visited := false
	total_distance := 0.0 //tracking the total distance covered

	for !all_nodes_visited { // run till all points are visited

		min_distance := math.Inf(1) // initializing the minimum distance from point to inf
		previous_nodes_visited := true

		var next_visit Visit

		for idx, visit := range visits {
			if !visit.visited { // we only run for not visited points
				distance := eucledian_distance(visit.point, current_visit.point)
				if distance < min_distance {
					min_distance = distance
					next_visit = Visit{visit.point, true, visit.zero_distance}
					visit_idx = idx
				}
			}
		}
		current_visit = next_visit // the final next visit obtained after the loop exits is the min distance visit
		total_distance += min_distance
		visits[visit_idx] = current_visit //updating the visit in the original visits slice

		for _, visit := range visits {
			all_nodes_visited = visit.visited && previous_nodes_visited
			previous_nodes_visited = all_nodes_visited
		}

		fmt.Printf("Visiting point %v, Total distance covered: %f\n\n", current_visit.point, total_distance)
		fmt.Printf("Current state:\n%+v\n\n", visits)
	}
	// returning to the start point
	total_distance += current_visit.zero_distance
	current_visit = visits[0]
	fmt.Printf("Completing round trip to point %v, Total distance covered: %f\n\n", current_visit.point, total_distance)
	fmt.Printf("Current state:\n%+v\n\n", visits)

	return total_distance
}

// Heuristic 2: Closest Pair Heuristic
// TODO

// Devising a better heuristic
// TODO

func main() {
	set := []Node{ //set of points in 2D
		{0, 0}, //starting point on the 0th index
		{-1, 0},
		{1, 0},
		{-5, 0},
		{3, 0},
		{-21, 0},
		{11, 0},
	}

	nearest_neighbor_dist := tour_nearest_neighbor(set)
	fmt.Printf("Using nearest neighbor - %f", nearest_neighbor_dist)
}
