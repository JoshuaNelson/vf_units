package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Unit struct {
	name string
	strength int
	agility int
	constitution int
	intelligence int
	wisdom int
	charisma int
	occupation string
	id int
}

func cmp_melee_combat(unit Unit) int {
	if !can_Knight(unit) {
		return 0
	}
	return (unit.strength + unit.strength + unit.agility) / 3
}

func can_Knight(unit Unit) bool {
	return (((unit.strength + unit.wisdom + unit.charisma) / 3) > 10)
}

func import_unit_csv(units map[int] Unit) {
	file, err := os.Open("D:/Games/LoL/ockford_units.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(bufio.NewReader(file))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		name := record[0]
		st, _ := strconv.Atoi(record[1])
		ag, _ := strconv.Atoi(record[2])
		co, _ := strconv.Atoi(record[3])
		in, _ := strconv.Atoi(record[4])
		wi, _ := strconv.Atoi(record[5])
		ch, _ := strconv.Atoi(record[6])
		id, _ := strconv.Atoi(record[7])
		oc := record[8]

		units[id] = Unit{name, st, ag, co, in, wi, ch, oc, id}
	}
}

func print_unit(unit Unit) {
	fmt.Printf("%-15v %3v %3v %3v %3v %3v %3v %7v %-10v\n",
		unit.name,
		unit.strength,
		unit.agility,
		unit.constitution,
		unit.intelligence,
		unit.wisdom,
		unit.charisma,
		unit.id,
		unit.occupation,
	)
}

func main() {
	fmt.Println("Running vf_units_main...")

	//var unit_id_list []int
	units := make(map[int]Unit)

	//Import units
	import_unit_csv(units)

	type smap struct {
		id int
		val int
	}

	var knight_smap []smap
	for id := range units {
		knight_smap = append(knight_smap, smap{id, cmp_melee_combat(units[id])})
	}

	sort.Slice(knight_smap, func(i, j int) bool {
		return knight_smap[i].val > knight_smap[j].val
	})

	var i int = 0
	for _, s := range knight_smap {
		if i >= 10 {
			break
		}

		if units[s.id].occupation != "Settlers" {
			continue
		}

		print_unit(units[s.id])
		i++
	}

	/* Print all units
	for id := range units {
		unit_id_list = append(unit_id_list, id)
	}
	for _, id := range unit_id_list {
		print_unit(units[id])
	}
	*/

}
