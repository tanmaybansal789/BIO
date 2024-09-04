package main
import (

"fmt"
"math"
)

type set[T comparable] struct {
    Add          func(T)
    Contains     func(T) bool
    Length       func() int
    Remove       func(T)
    Union        func(set[T]) set[T]
    Intersection func(set[T]) set[T]
    Equals       func(set[T]) bool // New method for comparison
    GetData      func() map[T]struct{}
}

func Set[T comparable]() set[T] {
    data := make(map[T]struct{})
    return set[T]{
        Add: func(i T) {
            data[i] = struct{}{}
        },
        Contains: func(i T) bool {
            _, ok := data[i]
            return ok
        },
        Length: func() int {
            return len(data)
        },
        Remove: func(i T) {
            delete(data, i)
        },
        Union: func(s set[T]) set[T] {
            result := Set[T]()
            for k := range data {
                result.Add(k)
            }
            for k := range s.GetData() {
                result.Add(k)
            }
            return result
        },
        Intersection: func(s set[T]) set[T] {
            result := Set[T]()
            for k := range data {
                if s.Contains(k) {
                    result.Add(k)
                }
            }
            return result
        },
        GetData: func() map[T]struct{} {
            return data
        },
    }
}

func ToSet[T comparable](data []T) set[T] {
    s := Set[T]()
    for _, d := range data {
        s.Add(d)
    }
    return s
}

// Point class
type Point struct {
    x, y int
}

type PointSet = set[Point]

var pentominoes = map[rune]PointSet{
    'F': ToSet([]Point{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 2}}),
    'G': ToSet([]Point{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}),
    'I': ToSet([]Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}}),
    'L': ToSet([]Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 0}}),
    'J': ToSet([]Point{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {1, 3}}),
    'N': ToSet([]Point{{0, 0}, {0, 1}, {0, 2}, {1, 2}, {1, 3}}),
    'M': ToSet([]Point{{0, 2}, {0, 3}, {1, 0}, {1, 1}, {1, 2}}),
    'P': ToSet([]Point{{0, 0}, {0, 1}, {0, 2}, {1, 1}, {1, 2}}),
    'Q': ToSet([]Point{{0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}}),
    'T': ToSet([]Point{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 2}}),
    'U': ToSet([]Point{{0, 0}, {0, 1}, {1, 0}, {2, 0}, {2, 1}}),
    'V': ToSet([]Point{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}),
    'W': ToSet([]Point{{0, 1}, {0, 2}, {1, 0}, {1, 1}, {2, 0}}),
    'X': ToSet([]Point{{0, 1}, {1, 0}, {1, 1}, {1, 2}, {2, 1}}),
    'Z': ToSet([]Point{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}}),
    'S': ToSet([]Point{{0, 0}, {1, 0}, {1, 1}, {1, 2}, {2, 2}}),
    'Y': ToSet([]Point{{0, 2}, {1, 0}, {1, 1}, {1, 2}, {1, 3}}),
    'A': ToSet([]Point{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {1, 2}}),
}

const singlePentominoSize = 5
const combinedPentominoSize = 10

func Transform(p Point, t PointSet) PointSet {
    result := Set[Point]()
    for k := range t.GetData() {
        result.Add(Point{k.x + p.x, k.y + p.y})
    }
    return result
}

func MinXYZero(t PointSet) PointSet{
    // Get the minimum x and y values
    minX, minY := math.MaxInt, math.MaxInt
    for k := range t.GetData() {
        if k.x < minX {
            minX = k.x
        }
        if k.y < minY {
            minY = k.y
        }
    }

    result := Set[Point]()
    for k := range t.GetData() {
        result.Add(Point{k.x - minX, k.y - minY})
    }

    return result
}

func IsValidCombinedShape(p PointSet) bool {
    if p.Length() != combinedPentominoSize {
        return false
    }

    visited := Set[Point]()

    var dfs func(Point)
    dfs = func(k Point) {
        if !p.Contains(k) || visited.Contains(k) {
            return
        }
        visited.Add(k)
        dfs(Point{k.x + 1, k.y})
        dfs(Point{k.x - 1, k.y})
        dfs(Point{k.x, k.y + 1})
        dfs(Point{k.x, k.y - 1})
    }

    // Get some point to start the dfs
    var start Point
    for k := range p.GetData() {
        start = k
        break
    }
    dfs(start)

    return visited.Length() == combinedPentominoSize
}

func PentominoCombinations(a, b rune) []PointSet {
    pa := pentominoes[a]
    pb := pentominoes[b]

    distinct := []PointSet{}


    for ox := -singlePentominoSize; ox < singlePentominoSize + 1; ox++ {
        for oy := -singlePentominoSize; oy < singlePentominoSize + 1; oy++ {
            tb := Transform(Point{ox, oy}, pb)

            cmb := MinXYZero(pa.Union(tb))

            if IsValidCombinedShape(cmb) {
                // Check if the shape is already in the slice
                found := false
                for _, v := range distinct {
                    if v.Intersection(cmb).Length() == combinedPentominoSize {
                        found = true
                        break
                    }
                }
                if !found {
                    distinct = append(distinct, cmb)
                }
            }
        }
    }

    return distinct
}

func main() {
    var s string
    fmt.Scan(&s)
    a, b := rune(s[0]), rune(s[1])

    combinations := PentominoCombinations(a, b)

    fmt.Println(len(combinations))
}