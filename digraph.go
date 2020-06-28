package digraph

import (
    "github.com/pkg/errors" // "errors"
)

type Digraph struct {
    V        int     // number of vertices in this digraph
    E        int     // number of edges in this digraph
    adj      [][]int // adj[v] = adjacency list for vertex v
    indegree []int   // indegree[v] = indegree of vertex v
}

func NewDigpraph() Digrapher {
    dg := Digraph{}

    return &dg
}

func (dg *Digraph) AddEdge(v, w int) (err error) {
    if err = dg.validateVertex(v); err != nil {
        return err
    }
    if err = dg.validateVertex(w); err != nil {
        return err
    }
    dg.adj[v] = append(dg.adj[v], w)
    dg.indegree[w]++
    dg.E++

    return err
}

// validateVertex returns error unless v > 0 and v < V
func (dg *Digraph) validateVertex(v int) (err error) {
    if v < 0 || v >= dg.V {
        err = errors.Wrap(err, "illegal argument")
        return err
    }
}
