package main

import (
    "fmt"
    "reflect"
    "strings"
)

type queue[T any] struct {
    Push func(T)
    Pop func() T
    Length func() int
}

func Queue[T any]() queue[T] {
    var data []T
    return queue[T]{
        Push: func(item T) {
            data = append(data, item)
        },
        Pop: func() T {
            res := data[0]
            data = data[1:]
            return res
        },
        Length: func() int {
            return len(data)
        },
    }
}

const TowerConfigLength = 4

type TowerConfig = [TowerConfigLength][]int

type Move struct {
    config TowerConfig
    moves int
}

func NextTowerConfigs(config TowerConfig) []TowerConfig {
    res := []TowerConfig{}
    for i := 0; i < TowerConfigLength; i++ {
        for j := 0; j < TowerConfigLength; j++ {
            if i != j && len(config[i]) > 0 && (len(config[j]) == 0 || config[i][0] < config[j][0]) {
                newConfig := TowerConfig{}
                for k := 0; k < TowerConfigLength; k++ {
                    newConfig[k] = make([]int, len(config[k]))
                    copy(newConfig[k], config[k])
                }
                newConfig[j] = append(newConfig[j], newConfig[i][len(newConfig[i]) - 1])
                newConfig[i] = newConfig[i][:len(newConfig[i]) - 1]
                res = append(res, newConfig)
            }
        }
    }
    return res
}

func ToString(config TowerConfig) string {
	var builder strings.Builder
	for _, tower := range config {
		builder.WriteString(fmt.Sprintf("%v|", tower))
	}
	return builder.String()
}

func BetweenTowerConfigs(a TowerConfig, b TowerConfig) int {
    mq := Queue[Move]()
    mq.Push(Move{config: a, moves: 0})

    visited := map[string]bool{}

    for mq.Length() > 0 {
        mv := mq.Pop()
        moves, config := mv.moves, mv.config

        if reflect.DeepEqual(config, b) {
            return moves
        }

        for _, newConfig := range NextTowerConfigs(config) {
            configStr := ToString(newConfig)
            if _, ok := visited[configStr]; !ok {
                mq.Push(Move{config: newConfig, moves: moves + 1})
                visited[configStr] = true
            }
        }
    }

    return -1
}

func main() {
    var sa1, sa2, sa3, sa4 string
    var sb1, sb2, sb3, sb4 string
    fmt.Scanln(&sa1, &sa2, &sa3, &sa4)
    fmt.Scanln(&sb1, &sb2, &sb3, &sb4)

    a := TowerConfig{{}, {}, {}, {}}
    b := TowerConfig{{}, {}, {}, {}}

    for i, s := range []string{sa1, sa2, sa3, sa4} {
        if s != "0" {
            for _, c := range s {
                a[i] = append(a[i], int(c - '0'))
            }
        }
    }

    for i, s := range []string{sb1, sb2, sb3, sb4} {
        if s != "0" {
            for _, c := range s {
                b[i] = append(b[i], int(c - '0'))
            }
        }
    }

    fmt.Println(BetweenTowerConfigs(a, b))
}
