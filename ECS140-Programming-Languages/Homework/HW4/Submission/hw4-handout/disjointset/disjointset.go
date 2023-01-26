package disjointset

// DisjointSet is the interface for the disjoint-set (or union-find) data
// structure.
// Do not change the definition of this interface.
type DisjointSet interface {
	// UnionSet(s, t) merges (unions) the sets containing s and t,
	// and returns the representative of the resulting merged set.
	UnionSet(int, int) int
	// FindSet(s) returns representative of the class that s belongs to.
	FindSet(int) int
}

// TODO: implement a type that satisfies the DisjointSet interface.

// NewDisjointSet creates a struct of a type that satisfies the DisjointSet interface.

type Graph struct {
	mapper map[int]Subset
}

type Subset struct {
	num, parent, rank int
}

func (e *Subset) updateParent(x int) {
	e.parent = x
}

func (e *Subset) addRank(x int) {
	e.rank += x
}

func MakeSubset(pNum int) Subset {
	var s Subset
	s.num = pNum
	s.parent = pNum
	s.rank = 1
	return s
}

func MakeGraph() Graph {
	var newGraph Graph
	newGraph.mapper = make(map[int]Subset)

	return newGraph
}

func (g Graph) FindSet(x int) int {
	val, found := g.mapper[x]
	if found == true {
		if val.parent != x {
			val.updateParent(g.Find(x))
			g.mapper[x] = val
			return val.parent
		} else {
			return val.parent
		}
	} else {
		var newSub Subset = MakeSubset(x)
		g.mapper[x] = newSub
		return newSub.parent
	}
}

func (g *Graph) Find(x int) int {
	var store = g.mapper[x]
	if store.parent != x {
		store.parent = g.Find(store.parent)
	}
	return store.parent
}

func (g Graph) UnionSet(x int, y int) int {
	xRoot := g.mapper[g.FindSet(x)]
	yRoot := g.mapper[g.FindSet(y)]

	if xRoot.parent == yRoot.parent {
		return xRoot.parent
	}

	if xRoot.rank >= yRoot.rank {
		yRoot.updateParent(xRoot.parent)
		xRoot.addRank(yRoot.rank)
		g.mapper[x] = xRoot
		g.mapper[y] = yRoot
		return xRoot.parent
	} else {
		xRoot.updateParent(yRoot.parent)
		yRoot.addRank(xRoot.rank)

		g.mapper[x] = xRoot
		g.mapper[y] = yRoot
		return yRoot.parent
	}
}

func NewDisjointSet() DisjointSet {
	var graphSet Graph
	graphSet = MakeGraph()
	return graphSet
}
