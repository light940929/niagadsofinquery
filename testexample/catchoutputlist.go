package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

//type A string //check for id type

type tped struct {
	Chr        string `json:"chr"`
	VariantID  string `json:"variant_id"`
	Location   string `json:"location"`
	Coordinate string `json:"coordinate"`
	Call       string `json:"call"`
}

type tfam struct {
	FamilyID        string `json:"family_id"`
	IndividualID    string `json:"individual_id"`
	PaternalID      string `json:"paternal_id"`
	MaternalID      string `json:"maternal_id"`
	Sex             string `json:"sex"`
	AffectionStatus string `json:"affection_status"`
}

func main() {
	//tpedfile, err := os.Open("snpID.txt") only id
	tpedfile, err := os.Open("example.merged.tped")
	tfamfile, err := os.Open("example.merged.tfam")
	if err != nil {
		fmt.Println(err)
	}
	defer tpedfile.Close()
	defer tfamfile.Close()
	snprd := bufio.NewScanner(tpedfile)
	individualrd := bufio.NewScanner(tfamfile)
	snprd.Split(bufio.ScanLines)
	individualrd.Split(bufio.ScanLines)

	//{"title":"cbd","chr":"1","variant_id":"snp2","location":"0","coordinate":"5000830","call":"G T G T G G T T G T T T"}
	for snprd.Scan() {
		line := snprd.Text()
		regchr := regexp.MustCompile(`\d{1,2}`)
		chr1 := regchr.FindString(line)
		chr := strings.Trim(chr1, " ,")
		fmt.Printf("%q\n", chr)

		regVariantID := regexp.MustCompile(`[[:lower:]]+\d+`)
		variantID := regVariantID.FindString(line)
		fmt.Printf("%q\n", variantID)

		regCoordinate := regexp.MustCompile(`\d{7}`)
		coordinate := regCoordinate.FindString(line)
		fmt.Printf("%q\n", coordinate)

		regCall := regexp.MustCompile(`[[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+ [[:upper:]]+`)
		call := regCall.FindString(line)
		fmt.Printf("%q\n", call)

		tpeds := &tped{
			Chr:        chr,
			VariantID:  variantID,
			Location:   "0",
			Coordinate: coordinate,
			Call:       call}
		out, err := json.Marshal(tpeds)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))

		// url := "http://localhost:9000/api/genotypes"
		// var jsonStr = []byte(out)
		// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		// req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6Ik5JQUdBRFMiLCJleHAiOjE0NjEzNDMxNDl9.M1K4fiH-jKoT-flBbla79A4q4aSM9qVOp3Q7xMtVe_8")
		// client := &http.Client{}
		// resp, err := client.Do(req)
		// if err != nil {
		// 	panic(err)
		// }
		// defer resp.Body.Close()
		// fmt.Println("response Status:", resp.Status)
	}
	//{"title":"cbd","family_id":"2","individual_id":"1","paternal_id":"0","maternal_id":"0","sex":"unknown","affection_status":"1"}
	for individualrd.Scan() {
		line := individualrd.Text()
		regfamilyID := regexp.MustCompile(`\d `)
		familyID1 := regfamilyID.FindString(line)
		familyID := strings.Trim(familyID1, " ")
		fmt.Printf("%q\n", familyID)
		command := "awk '{print $2}' ./example.merged.tped | awk '{for(i=1;i<=NF;i=i+1){a[NR,i]=$i}}END{for(j=1;j<=NF;j++){str=a[1,j];for(i=2;i<=NR;i++){str=str ", "  a[i,j]}print str}}'"
		individualID := exec.Command("sh", command).Output()
		fmt.Printf("%q\n", individualID)
		// regindividualID := regexp.MustCompile(` \d`)
		// individualID := regindividualID.FindString(line)
		// fmt.Printf("%q\n", individualID)

		regStatus := regexp.MustCompile(` \d`)
		status1 := regStatus.FindString(line)
		status := strings.Trim(status1, " ")
		fmt.Printf("%q\n", status)

		tfams := &tfam{
			FamilyID:        familyID,
			IndividualID:    individualID,
			PaternalID:      "0",
			MaternalID:      "0",
			Sex:             "0",
			AffectionStatus: status}
		out, err := json.Marshal(tfams)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
		// url := "http://localhost:9000/api/phenotypes"
		// var jsonStr = []byte(out)
		// req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		// req.Header.Set("Content-Type", "application/json")
		// req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJRCI6Ik5JQUdBRFMiLCJleHAiOjE0NjEzNDMxNDl9.M1K4fiH-jKoT-flBbla79A4q4aSM9qVOp3Q7xMtVe_8")
		// client := &http.Client{}
		// resp, err := client.Do(req)
		// if err != nil {
		// 	panic(err)
		// }
		// defer resp.Body.Close()
		// fmt.Println("response Status:", resp.Status)
	}

}
