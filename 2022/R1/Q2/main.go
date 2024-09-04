package main

import (
	"fmt"
	"sort"
    	)


var eds [25][6]int


var Rp, Bp, Re, Be int


func Control(hex int) int {
	r, b := 0, 0
	for i := 0; i < 6; i++ {
		
		if eds[hex][i] > 0 {
			r++
		} else if eds[hex][i] < 0 {
			b++
		}
	}

	
	if r == b {
		return 0
	}
	if r > b {
		return 1
	}
	return -1
}


func Opposite(ed int) int {
	return (ed + 3) % 6
}


func Own(v, hex, ed int) bool {
    //
	switched := false

	
	if eds[hex][ed] != v && eds[hex][ed] != 0 {
		switched = true
	}

	
	eds[hex][ed] = v

	
	h := hex
	hex += 1
	e := Opposite(ed)

	
	row := h / 5
	x, y := 0, 0

	if row % 2 == 1 {
		y = 1
		x = 0
	} else {
		y = 0
		x = 1
	}

	
	nhex := -1

	switch ed {
	case 0:
		if hex > 5 && (hex % 5 > 0 || row % 2 == 0) {
			nhex = h - 4 - x
		}
	case 1:
		if hex % 5 > 0 { 
            		nhex = h + 1
		}
	case 2:
		if hex < 21 && (hex % 5 > 0 || row % 2 == 0) {
			nhex = h + 5 + y
		}
	case 3:
		if hex < 21 && (hex % 5 != 1 || row % 2 == 1) {
			nhex = h + 4 + y
		}
	case 4:
		if hex % 5 != 1 {
			nhex = h - 1
		}
	case 5:
		if hex > 5 && (hex % 5 != 1 || row % 2 == 1) {
			nhex = h - 5 - x
		}
	}

	
	if nhex != -1 {
		if eds[nhex][e] != v && eds[nhex][e] != 0 {
			switched = true
		}
		eds[nhex][e] = v
	}

	return switched
}


func Init() {
	Rp = 0   
	Bp = 24  
	Re = 0  
	Be = 5  

	
	for i := 0; i < 25; i++ {
		for j := 0; j < 6; j++ {
			eds[i][j] = 0
		}
	}
}


func Skirmish(r, b int) bool {
	
	switchRed := Own(1, Rp, Re)
	Re = (Re + 1) % 6 
	Rp = (Rp + r) % 25  

	
	switchBlue := Own(-1, Bp, Be)
	Be = (Be - 1 + 6) % 6 
	Bp = (Bp + b) % 25      

	return switchRed || switchBlue
}


func Score() (int, int) {
	ro, bo := 0, 0

	
	for hex := 0; hex < 25; hex++ {
		c := Control(hex)
		if c == 1 {
			ro++
		} else if c == -1 {
			bo++
		}
	}

	return ro, bo
}


func Feud() {
	
	e := []struct{ cntrb, ncntrb, hex, ed int }{}
    
	rs, bs := Score()
    
	for hex := 0; hex < 25; hex++ {
		for ed := 0; ed < 6; ed++ {
			if eds[hex][ed] != 0 {
				continue
			}
			Own(1, hex, ed) 
			nr, nb := Score()
			cntrb := nr - rs  
			ncntrb := bs - nb 
			Own(0, hex, ed)   
			e = append(e, struct{ cntrb, ncntrb, hex, ed int }{cntrb, ncntrb, -hex, -ed})
		}
	}

	
	sort.Slice(e, func(i, j int) bool {
		if e[i].cntrb != e[j].cntrb {
			return e[i].cntrb < e[j].cntrb
		}
		return e[i].ncntrb < e[j].ncntrb
	})

	if len(e) > 0 {
		hex, ed := -e[len(e)-1].hex, -e[len(e)-1].ed
		Own(1, hex, ed)
	}
	
	e = nil
	rs, bs = Score()

	
	for hex := 0; hex < 25; hex++ {
		for ed := 0; ed < 6; ed++ {
			if eds[hex][ed] != 0 {
				continue
			}
			Own(-1, hex, ed) 
			nr, nb := Score()
			cntrb := nb - bs  
			ncntrb := rs - nr 
			Own(0, hex, ed)   
			e = append(e, struct{ cntrb, ncntrb, hex, ed int }{cntrb, ncntrb, hex, ed})
		}
	}
	
	sort.Slice(e, func(i, j int) bool {
		if e[i].cntrb != e[j].cntrb {
			return e[i].cntrb < e[j].cntrb
		}
		return e[i].ncntrb < e[j].ncntrb
	})

	if len(e) > 0 {
		hex, ed := e[len(e)-1].hex, e[len(e)-1].ed
		Own(-1, hex, ed)
	}
}

func main() {
	
	var r, b, s, f int
	fmt.Scan(&r, &b) 
	fmt.Scan(&s, &f)
	
	Init()
    
	for i := 0; i < s; i++ {
		Skirmish(r, b)
	}
    
	for i := 0; i < f; i++ {
		Feud()
	}
	
    ro, bo := Score()
	fmt.Println(ro)
	fmt.Println(bo)
}
