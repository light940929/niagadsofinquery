package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

//type A string //check for id type

type snplist struct {
	VariantID string
}

type individualist struct {
	FamilyID     string
	IndividualID string
}

func main() {
	snplistfile, err := os.Open("snpID.txt")
	individualistfile, err := os.Open("individualist.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer snplistfile.Close()
	defer individualistfile.Close()
	snprd := bufio.NewScanner(snplistfile)
	individualrd := bufio.NewScanner(individualistfile)
	snprd.Split(bufio.ScanLines)
	individualrd.Split(bufio.ScanLines)
	//{"VariantID":"rs12185268"}
	for snprd.Scan() {
		line := snprd.Text()
		//fmt.Printf("%s\n", line)
		snplists := &snplist{line}
		out, err := json.Marshal(snplists)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}
	//{"FamilyID":"0","IndividualID":"CBD0001"}
	for individualrd.Scan() {
		line := individualrd.Text()
		regfamilyID := regexp.MustCompile(`\d ,`)
		familyID1 := regfamilyID.FindString(line)
		familyID := strings.Trim(familyID1, " ,")
		//fmt.Printf("%q\n", familyID)
		// check id type
		// m := map[string]interface{}{
		// 	"a": A(familyID),
		// }
		//
		// fmt.Println(reflect.ValueOf(m["a"]).Type().Kind())

		regindividualID := regexp.MustCompile(`[[:upper:]]+\d+`)
		individualID := regindividualID.FindString(line)
		//fmt.Printf("%q\n", individualID)

		individualists := &individualist{
			FamilyID:     familyID,
			IndividualID: individualID}
		out, err := json.Marshal(individualists)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	}

}
