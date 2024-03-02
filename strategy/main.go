// Strategy Pattern
// It's a behavioral design pattern that let us create a family of related algorithms, that execute a common task in different ways
// and make them interchangeble in runtime at a context object

// Why?
// Sometimes our application performs a given task in many different ways, depending only how it was requested. This may lead
// to a higher code complexity

// How?
// We can extract each of these behaviors into a family of related classes, so that the client code can swap these behavors at
// runtime, just informing a context object which one should be used. The context only relates to the strategies via a common interface
// and the client must known which behavior should be used during runtime

// Conceptual example
// Suppose we have a Navigator App that allow users to build routes based on their transportation preferences: Bicicle, Bus or Cars.

package main

import "fmt"

// Route is a helper struct to allow us to illustrate the conceptual example
type Route struct {
	milestones []string
}

// Router is the common interface to be implemented by all different strategies
type IRouter interface {
	buildRoute(origin string, destination string) *Route
}

// defyning our different strategies
type BicicleRouter struct{}

func (b *BicicleRouter) buildRoute(origin string, destination string) *Route {
	fmt.Printf("[BicicleRouter] building route from %s to %s...\n", origin, destination)
	route := &Route{}
	route.milestones = append(route.milestones, origin)
	route.milestones = append(route.milestones, "milestone1")
	route.milestones = append(route.milestones, "milestone2")
	route.milestones = append(route.milestones, "milestone3")
	route.milestones = append(route.milestones, "milestone4")
	route.milestones = append(route.milestones, "milestone5")
	route.milestones = append(route.milestones, destination)

	return route
}

type BusRouter struct{}

func (b *BusRouter) buildRoute(origin string, destination string) *Route {
	fmt.Printf("[BusRouter] building route from %s to %s...\n", origin, destination)
	route := &Route{}
	route.milestones = append(route.milestones, origin)
	route.milestones = append(route.milestones, "milestone1")
	route.milestones = append(route.milestones, "milestone2")
	route.milestones = append(route.milestones, "milestone3")
	route.milestones = append(route.milestones, destination)

	return route
}

type CarRouter struct{}

func (c *CarRouter) buildRoute(origin string, destination string) *Route {
	fmt.Printf("[CarRouter] building route from %s to %s...\n", origin, destination)
	route := &Route{}
	route.milestones = append(route.milestones, origin)
	route.milestones = append(route.milestones, "milestone1")
	route.milestones = append(route.milestones, "milestone2")
	route.milestones = append(route.milestones, destination)

	return route
}

// Our context object
type Navigator struct {
	router IRouter
}

func (n *Navigator) setRouter(router IRouter) {
	n.router = router
}

func (n *Navigator) buildRoute(origin string, destination string) *Route {
	fmt.Println("[Navigator] starting route building process...")
	return n.router.buildRoute(origin, destination)
}

// helper function to just display the route.milestones field
func displayRoute(route *Route) {
	for _, v := range route.milestones {
		fmt.Printf("- %s\n", v)
	}
}

func main() {
	navigator := &Navigator{}

	navigator.setRouter(&BicicleRouter{})
	route := navigator.buildRoute("PointA", "PointB")
	displayRoute(route)

	navigator.setRouter(&BusRouter{})
	route = navigator.buildRoute("PointA", "PointC")
	displayRoute(route)

	navigator.setRouter(&CarRouter{})
	route = navigator.buildRoute("PointD", "PointA")
	displayRoute(route)
}
