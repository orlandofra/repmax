package main

import (
	"log"
	"fmt"
	"os"
	"math"

	"strconv"
	"strings"
	"errors"
	)

func round(f float64) float64 {
	return math.Round(f*2) / 2
}


func checkWeight (part string) bool {
	return strings.HasSuffix(part, "kg")
}

func stripReps (part string) (int, error) {
	return strconv.Atoi(part)
}

func stripRPE (part string) (int, error) {
	return strconv.Atoi(part)
}

func stripWeight (part string) (float64, error) {
	if checkWeight(part) {
		s := strings.TrimSuffix(part, "kg")
		return strconv.ParseFloat(s, 64)
	} else {
		return strconv.ParseFloat(part, 64)
	}
}

func locateWeight(prcalc string) (int, error) {
	parts := strings.Split(prcalc, "x")

	if len(parts) > 2 {
		return 0, errors.New("input not valid")
	}

	if len(parts) == 1 || checkWeight(parts[0]) {
		return 0, nil
	} else if checkWeight(parts[1]) {
		return 1, nil
	} else {
		return -1, errors.New("input not valid")
	}

}

func parsePrcalc(prcalc string) (float64, int, error){

	weightIndex, err := locateWeight(prcalc)
	if err != nil {
		return 0, 0, err
	}

	parts := strings.Split(prcalc, "x")

	if len(parts) == 1 {
		weight, err := stripWeight(parts[weightIndex])
		return weight, 1, err
	}

	weight, err := stripWeight(parts[weightIndex])
	if err != nil {
		return 0, 0, err
	}

	reps, err := stripReps(parts[1 - weightIndex])
	if err != nil {
		return 0, 0, err
	}
	return weight, reps, nil
}

func parseRpecalc(weight float64, rpecalc string) (float64, int, int, error) {
	/* @ delimiter */
	parts := strings.Split(rpecalc, "@")
	reps, err:= stripReps(parts[0])
	if err != nil {
		return 0, 0, 0, err
	}

	if reps <= 0{
		return 0, 0, 0, errors.New("reps not valid")
	}

	rpe, err := stripRPE(parts[1])
	if err != nil {
		return 0, 0, 0, err
	}

	if rpe > 10 || rpe <= 0{
		return 0, 0, 0, errors.New("rpe specified not valid")
	}

	buffer := 10 - rpe
	return ReverseOneRM(weight, reps + buffer), reps, rpe, nil
}


func main(){

	if len(os.Args) < 2 || len(os.Args) > 3 {
		log.Fatalf("Usage: %s <prcalc> [<rpecalc>]\n", os.Args[0])
	}

	prcalc := os.Args[1]
	weight, reps, err := parsePrcalc(prcalc)
	if err != nil {
		log.Fatalf("repmax: %v\n", err)
	}
	pr := OneRM(weight, reps)
	fmt.Printf("%.1fkg", round(pr))

	if len(os.Args) == 3 {
		rpecalc := os.Args[2]
		weight, reps, rpe, err := parseRpecalc(pr, rpecalc)
		if err != nil {
			log.Fatalf("repmax: %v\n", err)
		}
		fmt.Printf(",%dx%.1fkg(@%d)", reps, round(weight), rpe)
	}

	fmt.Printf("\n")

}
