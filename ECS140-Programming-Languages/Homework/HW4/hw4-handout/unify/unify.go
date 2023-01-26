package unify

import (
	"errors"
	"hw4/disjointset"
	"hw4/term"
)

// ErrUnifier is the error value returned by the Parser if the string is not a
// valid term.
// See also https://golang.org/pkg/errors/#New
// and // https://golang.org/pkg/builtin/#error
var ErrUnifier = errors.New("unifier error")

// UnifyResult is the result of unification. For example, for a variable term
// `s`, `UnifyResult[s]` is the term which `s` is unified with.
type UnifyResult map[*term.Term]*term.Term

// Unifier is the interface for the term unifier.
// Do not change the definition of this interface
type Unifier interface {
	Unify(*term.Term, *term.Term) (UnifyResult, error)
}

type UnifierImpl struct {
	// disjoint set dataset
	dj disjointset.DisjointSet
	// schema function map
	sigma map[*term.Term]*term.Term
	// visited and acyclic flags
	visited, acyclic map[*term.Term]bool
	// list of vars
	vars map[*term.Term][]*term.Term

	// result
	unif UnifyResult

	// encoder map
	encode map[*term.Term]int
	// decoder map
	decode map[int]*term.Term
	// current max index
	imax int
}


func (u *UnifierImpl) Unify(s *term.Term, t *term.Term) (UnifyResult, error) {

	s_comb, err := u.UnifClosure(s, t)
	if err == nil {
		err, _ := u.FindSolution(s_comb, false)
		if err == nil {
			// additional cycle checking
			for k, v := range u.unif {
				if contains(getArgs(v), k) {
					// current solution is cyclical
					return nil, ErrUnifier
				}
			}
			return u.unif, nil
		} else {
			return nil, err
		}
	} else {
		return nil, err
	}
}

func (u *UnifierImpl) Setup(s *term.Term, functor bool, s_o *term.Term) {
	if functor && s.Typ == term.TermVariable {
		u.sigma[s] = s_o
	} else {
		u.sigma[s] = s
	}

	// initialize the flags
	if u.visited[s] != true {
		u.visited[s] = false
	}
	if u.acyclic[s] != true {
		u.acyclic[s] = false
	}

	// setup vars list
	if u.vars[s] == nil {
		if s.Typ == term.TermCompound {
			u.vars[s] = make([]*term.Term, 0)
		} else {
			u.vars[s] = make([]*term.Term, 1)
			u.vars[s][0] = s
		}
		
		u.PopulateED(s)
	}

}

func (u *UnifierImpl) PopulateED(s *term.Term) error {
	u.imax = u.imax + 1
	u.encode[s] = u.imax
	u.decode[u.imax] = s
	return nil
}


func (u *UnifierImpl) SetupFunction(s *term.Term, functor bool, s_o *term.Term) {
	s_f := s.Functor
	s_a := s.Args
	if s_f != nil {
		u.Setup(s_f, functor, s_o)
		if s_a != nil {
			for _, a := range s.Args {
				u.Setup(a, functor, s_o)
				if a.Functor != nil {
					u.SetupFunction(a, functor, s_o)
				}
			}

		}
	}
}

func (u *UnifierImpl) UnifClosure(s *term.Term, t *term.Term) (*term.Term, error) {

	// setup for current terms
	u.Setup(s, false, s)
	u.SetupFunction(s, s.Typ != term.TermVariable, s)
	u.Setup(t, false, t)
	u.SetupFunction(t, t.Typ != term.TermVariable, t)


	// proceed with algorithm
	s_rep := u.decode[u.dj.FindSet(u.encode[s])]

	t_rep := u.decode[u.dj.FindSet(u.encode[t])]

	if s_rep == t_rep {
		// do nothing?
		return nil, nil
	} else {
		if (u.sigma[s_rep].Typ == u.sigma[t_rep].Typ) && ((u.sigma[s_rep].Typ == term.TermAtom) || (u.sigma[s_rep].Typ == term.TermNumber)) {
			return nil, ErrUnifier
		} else if ((u.sigma[s_rep].Typ != u.sigma[t_rep].Typ) && (u.sigma[s_rep].Typ != term.TermVariable && u.sigma[t_rep].Typ != term.TermVariable)) {
			return nil, ErrUnifier
		} else if (u.sigma[s_rep].Typ == term.TermCompound && len(u.sigma[s_rep].Args) >= 0) && (u.sigma[t_rep].Typ == term.TermCompound && len(u.sigma[t_rep].Args) >= 0) {
			s_f := u.sigma[s_rep].Functor
			t_f := u.sigma[t_rep].Functor

			s_args := u.sigma[s_rep].Args
			t_args := u.sigma[t_rep].Args
			
			// handle the cycle case where you have s = args[0]

			if (len(t_args) == 1 && s_rep == t_args[0]) || (len(s_args) == 1 && t_rep == s_args[0]) {
				return nil, ErrUnifier
			}
			
			if len(s_args) != len(t_args) {
				return nil, ErrUnifier
			}

			if s_f == t_f {
				s_rep = u.Union(s_rep, t_rep)

				for i := 0; i < len(s_args); i++ {
					_, err := u.UnifClosure(s_args[i], t_args[i])
					if err != nil {
						return nil, ErrUnifier
					}
				}
				return s_rep, nil
			} else {
				// symbol clash
				return nil, ErrUnifier
			}
		} else {
			s_rep = u.Union(s_rep, t_rep)
			return s_rep, nil
		}
	}
}

func (u *UnifierImpl) FindSolution(s *term.Term, args bool) (error, bool) {
	s_rep := u.decode[u.dj.FindSet(u.encode[s])]
	s_s := u.sigma[s_rep]

	if u.visited[s_s] == true {
		if s != s_s && (!contains(s_s.Args, s) && !contains(getArgs(s_s), s)) {
			return ErrUnifier, true
		} else {
			return nil, false
		}
	}
	if u.acyclic[s_s] {
		return nil, false
	}
	if args {
		if u.visited[s] {
			return ErrUnifier, false
		}
	}
	if s_s.Typ == term.TermCompound && len(s_s.Args) > 0 {
		s_s_args := s_s.Args
		u.visited[s] = true
		for i := 0; i < len(s_s_args); i++ {
			err, b := u.FindSolution(s_s_args[i], true)
			u.visited[s_s_args[i]] = true
			if err != nil && b == true {
				return err, true
			}
		}
		u.visited[s] = false
	}


	// mark the function and all associated arguments as acyclic
	u.acyclic[s_s] = true

	for _, x := range u.vars[u.decode[u.dj.FindSet(u.encode[s_s])]] {
		if x != s_s {
			if x.Typ != term.TermVariable && s_s.Typ == term.TermVariable {
				u.unif[s_s] = x
			} else if (x.Typ != term.TermVariable && s_s.Typ != term.TermVariable) {
				return ErrUnifier, true
			} else {
				u.unif[x] = s_s
			}
		}
	}
	return nil, false
}

func (u *UnifierImpl) Union(s *term.Term, t *term.Term) (rep *term.Term) {
	s_par := u.dj.FindSet(u.encode[s])

	merged_rep := u.decode[u.dj.UnionSet(u.encode[s], u.encode[t])]

	if u.encode[merged_rep] == s_par {
		u.vars[s] = append(u.vars[s], u.vars[t]...)
		// if schema(s) is a variable => s is not a functor type => schema(s) is NOT the functor (there shouldn't even be a functor?)
		if u.sigma[s].Typ != term.TermCompound {
			u.sigma[s] = u.sigma[t]
		}
	} else {
		u.vars[t] = append(u.vars[s], u.vars[t]...)
		if u.sigma[t].Typ != term.TermCompound {
			u.sigma[t] = u.sigma[s]
		}
	}

	return merged_rep
}

// NewUnifier creates a struct of a type that satisfies the Unifier interface.
func NewUnifier() Unifier {
	return &UnifierImpl{
		dj:      disjointset.NewDisjointSet(),
		sigma:   make(map[*term.Term]*term.Term),
		visited: make(map[*term.Term]bool),
		acyclic: make(map[*term.Term]bool),
		vars:    make(map[*term.Term][]*term.Term),
		encode:  make(map[*term.Term]int),
		decode:  make(map[int]*term.Term),
		unif:    make(UnifyResult),
		imax:    0,
	}
}

func getArgs(f *term.Term) []*term.Term {
	f_a := f.Args
	if f_a == nil {
		return nil
	} else {
		ret := make([]*term.Term, 0)
		for _, a := range f_a {
			if a.Typ == term.TermCompound {
				ret = append(ret, getArgs(a)...)
			} else {
				ret = append(ret, a)
			}
		}
		return ret
	}
}

func contains(s []*term.Term, elem *term.Term) bool {
	for _, v := range s {
		if v == elem {
			return true
		}
	}
	return false
}