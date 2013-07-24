package main

import "fmt"

type Graph struct {
    Vertices []*Vertex
}

type Vertex struct {
    adjacencies []*Vertex
    Data interface{}
}

func (graph *Graph) Insert(vertex *Vertex) {
    graph.Vertices = append(graph.Vertices, vertex)
}

func (vertex *Vertex) IsConnectedTo(u *Vertex) bool {
    for i := range vertex.adjacencies {
        if(vertex.adjacencies[i] == u) {
            return true
        }
    }

    return false
}

func (vertex *Vertex) Connect(u *Vertex) {
    vertex.adjacencies = append(vertex.adjacencies, u)
}

func main() {
    fmt.Println("Creating test graph")
    v1 := new(Vertex)
    v1.Data = "Hello world!"

    v2 := new(Vertex)
    v2.Data = "Woah there, world!"

    v3 := new(Vertex)
    v3.Data = "One more test"

    graph := new(Graph)
    graph.Insert(v1)
    graph.Insert(v2)
    graph.Insert(v3)

    for i := range graph.Vertices {
        fmt.Println("\t[", i, "]", graph.Vertices[i])
    }

    v1.Connect(v2)
    fmt.Println("Are 1 and 2 connected?", v1.IsConnectedTo(v2))
    fmt.Println("Are 2 and 3 connected?", v2.IsConnectedTo(v3))
}
