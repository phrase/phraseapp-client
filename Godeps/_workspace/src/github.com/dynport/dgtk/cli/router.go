package cli

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// The basic datastructure used in cli. Actions are added for different paths to a router via the "Register" and
// "RegisterFunc" methods. These actions can be executed using the "Run" and "RunWithArgs" methods.
type Router struct {
	root *routingTreeNode

	initFailed bool
}

func New(routes ...func(*Router) error) (*Router, error) {
	router := NewRouter()
	for _, route := range routes {
		if err := route(router); err != nil {
			return nil, err
		}
	}
	return router, nil
}

// Create a new router that will be used to register and run the actions of the application.
func NewRouter() *Router {
	r := &Router{}
	r.root = &routingTreeNode{children: map[string]*routingTreeNode{}}
	return r
}

// Run the given arguments against the registered actions, i.e. try to find a matching route and run the according
// action.
func (r *Router) Run(args ...string) (e error) {
	if r.initFailed {
		fmt.Fprintln(Stderr, "errors found during initialization")
		os.Exit(1)
	}
	// Find action and parse args.
	node, args := r.findNode(args, true)
	if node != nil && node.action != nil {
		if e := node.action.parseArgs(args); e != nil {
			node.showHelp()
			return e
		}
	} else { // Failed to find node.
		node.showHelp()
		return ErrorNoRoute
	}

	return node.action.runner.Run()
}

// Run the arguments from the commandline (aka os.Args) against the registered actions, i.e. try to find a matching
// route and run the according action.
func (r *Router) RunWithArgs() (e error) {
	return r.Run(os.Args[1:]...)
}

type annonymousAction struct {
	runner      func() error
	description string
}

func (aA *annonymousAction) Run() error {
	return aA.runner()
}

// Register the given function as handler for the given route. This is a shortcut for actions that don't need options or
// arguments. A description can be provided as an optional argument.
func (r *Router) RegisterFunc(path string, f func() error, desc string) {
	aA := &annonymousAction{runner: f}
	r.Register(path, aA, desc)
}

// Register the given action (some struct implementing the Runner interface) for the given route.
func (r *Router) Register(path string, runner Runner, desc string) {
	a, e := newAction(path, runner, desc)
	if e != nil {
		fmt.Fprintln(Stderr, e)
		r.initFailed = true
		return
	}

	pathSegments := strings.Split(a.path, "/")
	node, pathSegments := r.findNode(pathSegments, false)
	if node != nil {
		if node.action != nil {
			fmt.Fprintf(Stderr, "failed to register action for path %q: action for path %q already registered\n", a.path, node.action.path)
			r.initFailed = true
			return
		} else if len(pathSegments) == 0 && len(node.children) > 0 {
			fmt.Fprintf(Stderr, "failed to register action for path %q: longer paths with this prefix exist\n", a.path)
			r.initFailed = true
			return
		}
	} else {
		node = r.root
	}

	for _, p := range pathSegments {
		newNode := &routingTreeNode{children: map[string]*routingTreeNode{}}
		node.children[p] = newNode
		node = newNode
	}

	node.action = a
}

func (r *Router) showHelp() {
}

// A tree used for easy access to the matching action. An action can only be set if there are no children, i.e. only
// leaf nodes can have actions.
type routingTreeNode struct {
	children map[string]*routingTreeNode
	action   *action
}

func (rt *routingTreeNode) showHelp() {
	if rt.action != nil {
		rt.action.showHelp()
	} else {
		t := &table{}
		rt.showTabularHelp(t)
		fmt.Fprintln(Stderr, t)
	}
}

func (rt *routingTreeNode) showTabularHelp(t *table) {
	if rt.action != nil {
		rt.action.showTabularHelp(t)
	} else {
		pathSegments := make([]string, 0, len(rt.children))
		for k, _ := range rt.children {
			pathSegments = append(pathSegments, k)
		}
		sort.Strings(pathSegments)

		for _, ps := range pathSegments {
			rt.children[ps].showTabularHelp(t)
		}
	}
}

// Find the node matching most segments of the given path. Will return the according tree node and the remaining (non
// matched) path segments.
func (r *Router) findNode(pathSegments []string, fuzzy bool) (*routingTreeNode, []string) {
	node := r.root
	for i, p := range pathSegments {
		if c, found := node.children[p]; found {
			node = c
		} else {
			if fuzzy { // try fuzzy search
				candidates := []string{}
				for key, _ := range node.children {
					if strings.HasPrefix(key, p) {
						candidates = append(candidates, key)
					}
				}
				if len(candidates) == 1 {
					node = node.children[candidates[0]]
					continue
				}
			}
			return node, pathSegments[i:]
		}
	}
	return node, nil
}
