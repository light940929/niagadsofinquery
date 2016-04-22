package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"
)

//type A string //check for id type

type snplist struct {
	Title      string `json:"title"`
	Chr        string `json:"chr"`
	VariantID  string `json:"variant_id"`
	Coordinate string `json:"coordinate"`
}

type individualist struct {
	Title        string `json:"title"`
	FamilyID     string `json:"family_id"`
	IndividualID string `json:"individual_id"`
}

func main() {
	//snplistfile, err := os.Open("snpID.txt") only id
	snplistfile, err := os.Open("snplist.txt")
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
	// for snprd.Scan() {
	// 	line := snprd.Text()
	// 	//fmt.Printf("%s\n", line)
	// 	snplists := &snplist{line}
	// 	out, err := json.Marshal(snplists)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println(string(out))
	// }

	//{"title":"cbd","chr":"17","variant_id":"rs12185268","coordinate":"12185268"}
	for snprd.Scan() {
		line := snprd.Text()
		regchr := regexp.MustCompile(`\d{1,2} ,`)
		chr1 := regchr.FindString(line)
		chr := strings.Trim(chr1, " ,")
		//fmt.Printf("%q\n", chr)

		regVariantID := regexp.MustCompile(`[[:lower:]]+\d+`)
		variantID := regVariantID.FindString(line)
		//fmt.Printf("%q\n", variantID)

		regCoordinate := regexp.MustCompile(`\d{8}`) //^/d{8}
		coordinate := regCoordinate.FindString(line)
		//fmt.Printf("%q\n", coordinate)

		snplists := &snplist{
			Title:      "cbd",
			Chr:        chr,
			VariantID:  variantID,
			Coordinate: coordinate}
		out, err := json.Marshal(snplists)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		url := "http://localhost:9000/api/genotypes"
		var jsonStr = []byte(out)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6Ik5JQUdBRFMiLCJleHAiOjE0NjEzNDMxNDl9.M1K4fiH-jKoT-flBbla79A4q4aSM9qVOp3Q7xMtVe_8")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
	}
	//{"title":"cbd","family_id":"0","individual_id":"CBD0001"}
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
			Title:        "cbd",
			FamilyID:     familyID,
			IndividualID: individualID}
		out, err := json.Marshal(individualists)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
		url := "http://localhost:9000/api/phenotypes"
		var jsonStr = []byte(out)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6Ik5JQUdBRFMiLCJleHAiOjE0NjEzNDMxNDl9.M1K4fiH-jKoT-flBbla79A4q4aSM9qVOp3Q7xMtVe_8")
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		fmt.Println("response Status:", resp.Status)
	}

}
