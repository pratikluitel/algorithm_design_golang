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

package chapter1

import (
	"fmt"
	"math"
	"time"
)

type Node struct { // A node is a point in a plane where the robot needs to move
	x float64
	y float64
}

type Edge struct {
	start Node
	end   Node
}

type Chain struct { // a chain is a collection of joined nodes, with a start node and an end node
	start Node
	end   Node
	edges []Edge
}

type Visit struct { // represents a point and it's visit status
	point               Node
	visited             bool    // the bool denotes if the node/point has been visited
	distance_from_start float64 // distance from the 0 point, so the round trip can be made
}

func eucledian_distance(p1, p2 Node) float64 {
	distance := math.Sqrt(math.Pow(p1.x-p2.x, 2) + math.Pow(p1.y-p2.y, 2))

	return distance
}

// Heuristic 1	: Nearest Neighbor Heuristic

func tour_nearest_neighbor(set []Node) ([]Node, float64) { // The assumption here is that the 1st element of the set is the start point

	fmt.Printf("Starting the tour using nearest neighbor heuristic.\n")

	visits := []Visit{}
	visit_idx := 0 // visit the start node/point first
	path := []Node{set[visit_idx]}
	current_visit := Visit{set[visit_idx], true, 0.0}

	// state initialization
	for idx, point := range set {
		zero_distance := eucledian_distance(current_visit.point, point)
		visited := idx == 0 //the first point is set to as visited
		visits = append(visits, Visit{point, visited, zero_distance})
	}

	all_nodes_visited := false
	total_distance := 0.0 //tracking the total distance covered

	for !all_nodes_visited { // run till all points are visited

		min_distance := math.Inf(1) // initializing the minimum distance from point to inf
		previous_nodes_visited := true

		for idx, visit := range visits {
			if !visit.visited { // we only run for not visited points
				distance := eucledian_distance(visit.point, current_visit.point)
				if distance < min_distance {
					min_distance = distance
					visit_idx = idx
				}
			}
		}
		visits[visit_idx].visited = true  //updating the visit in the original visits slice
		current_visit = visits[visit_idx] // the final next visit obtained after the loop exits is the min distance visit
		total_distance += min_distance
		path = append(path, current_visit.point)

		for _, visit := range visits {
			all_nodes_visited = visit.visited && previous_nodes_visited
			previous_nodes_visited = all_nodes_visited
		}

	}
	// returning to the start point
	path = append(path, visits[0].point)
	total_distance += current_visit.distance_from_start

	return path, total_distance
}

// Heuristic 2: Closest Pair Heuristic
// The book is way too vague on the heuristic, I used a local variation of this
// similar to nearest neighbor

func tour_closest_pair(set []Node) (Chain, float64) {

	fmt.Printf("Starting the tour using closest pair heuristic.\n")

	chain := Chain{set[0], set[0], []Edge{}}
	current_visit := Visit{set[0], true, 0.0}

	visits := []Visit{}
	visit_idx := 0 // visit the start node/point first

	// state initialization
	for idx, point := range set {
		zero_distance := eucledian_distance(current_visit.point, point)
		visited := idx == 0 //the first point is set to as visited
		visits = append(visits, Visit{point, visited, zero_distance})
	}

	total_distance := 0.0 //tracking the total distance covered
	all_nodes_visited := false

	for !all_nodes_visited { // run till all points are visited

		previous_nodes_visited := true

		min_distance := math.Inf(1) // initializing the minimum distance from point to inf
		var start_bool bool         // denotes if we need to add an edge from the start or the end of the chain

		for idx := 0; idx < len(visits); idx++ {
			if !visits[idx].visited {
				start_distance := eucledian_distance(visits[idx].point, chain.start)
				end_distance := eucledian_distance(visits[idx].point, chain.end)
				if (start_distance < min_distance || end_distance < min_distance) && (start_distance != 0 && end_distance != 0) {
					start_bool = start_distance < end_distance
					if start_bool {
						min_distance = start_distance
						visit_idx = idx
					} else {
						min_distance = end_distance
						visit_idx = idx
					}
				}
			}

		}
		visits[visit_idx].visited = true //updating the visit in the original visits slice
		if start_bool {
			chain.edges = append(chain.edges, Edge{visits[visit_idx].point, chain.start})
			chain.start = visits[visit_idx].point
		} else {
			chain.edges = append(chain.edges, Edge{chain.end, visits[visit_idx].point})
			chain.end = visits[visit_idx].point
		}

		current_visit = visits[visit_idx] // the final next visit obtained after the loop exits is the min distance visit
		total_distance += min_distance

		for _, visit := range visits {
			all_nodes_visited = visit.visited && previous_nodes_visited
			previous_nodes_visited = all_nodes_visited
		}
	}
	// returning to the start point
	chain.edges = append(chain.edges, Edge{chain.end, chain.start})
	total_distance += eucledian_distance(chain.end, chain.start)

	return chain, total_distance
}

// Devising a better heuristic
// TODO

func Problem_30() {
	set := []Node{ //set of points in 2D
		{0, 0}, //starting point on the 0th index
		{-1, 0},
		{1, 0},
		{-5, 0},
		{3, 0},
		{-21, 0},
		{11, 0},
	}
	// set := []Node{
	// 	{0, 0.9},
	// 	{0, 0},
	// 	{1.1, 0},
	// 	{1.1, 0.9},
	// 	{2.2, 0.9},
	// 	{2.2, 0},
	// 	{3.2, 4},
	// 	{3.3, 0},
	// 	{10, 0},
	// 	{4, 12},
	// }

	start1 := time.Now()
	nearest_neighbor_path, nearest_neighbor_dist := tour_nearest_neighbor(set)
	fmt.Printf("Using nearest neighbor -\n Path:%v\nTotal Distance: %f, executed in: %s\n\n", nearest_neighbor_path, nearest_neighbor_dist, time.Since(start1))

	start2 := time.Now()
	closest_pair_path, closest_pair_dist := tour_closest_pair(set)
	fmt.Printf("Using closest pair -\n Path:%v\nTotal Distance: %f, executed in: %s\n\n", closest_pair_path.edges, closest_pair_dist, time.Since(start2))
}
