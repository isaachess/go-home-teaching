package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Positions struct {
	Elders      []Position `json:"ELDERS_QUORUM"`
	HighPriests []Position `json:"HIGH_PRIESTS_GROUP"`
}

type Position struct {
	Id               int    `json:"id"`
	IndividualId     int    `json:"individualId"`
	OrgId            int    `json:"orgId"`
	OrgTypeId        int    `json:"orgTypeId"`
	OrgTypeName      string `json:"orgTypeName"`
	PositionTypeId   string `json:"positionTypeId"`
	PositionTypeName string `json:"PositionTypeName"`
}

type District struct {
	AuxiliaryId                int             `json:"auxiliaryId"`
	Companionships             []Companionship `json:"companionships"`
	DistrictLeaderId           int             `json:"districtLeaderId"`
	DistrictLeaderIndividualId int             `json:"districtLeaderIndividualId"`
	Id                         int             `json:"id"`
	name                       string          `json:"name"`
}

type Companionship struct {
	Assignments []Assignment `json:"assignments"`
	DistrictId  int          `json:"districtId"`
	Id          int          `json:"id"`
	StartDate   int          `json:"startDate"`
	Teachers    []Teacher    `json:"teachers"`
}

type Assignment struct {
	AssignmentType  string `json:"assignmentType"`
	CompanionshipId int    `json:"companionshipId"`
	Id              int    `json:"id"`
	IndividualId    int    `json:"individualId"`
}

type Teacher struct {
	CompanionshipId int `json:"companionshipId"`
	Id              int `json:"id"`
	IndividualId    int `json:"individualId"`
}

type Families struct {
	Families []Family `json:"families"`
}

type Family struct {
	EmailAddress        string `json:"emailAddress"`
	FormattedCoupleName string `json:"formattedCoupleName"`
	HeadOfHouse         Person `json:"headOfHouse"`
	IsAssignedHT        bool   `json:"isAssignedHT"`
	Phone               bool   `json:"phone"`
	Spouse              Person `json:"spouse"`
}

type Person struct {
	Email                   string `json:"email"`
	FormattedName           string `json:"formattedName"`
	Gender                  string `json:"gender"`
	GivenName1              string `json:"givenName1"`
	HeadOfHouseIndividualId int    `json:"headOfHouseIndividualId"`
	ImageId                 string `json:"imageId"`
	IndividualId            int    `json:"individualId"`
	IsAdult                 bool   `json:"isAdult"`
	IsAssignedVT            bool   `json:"isAssignedVT"`
	Phone                   string `json:"phone"`
	Surname                 string `json:"surname"`
}

func main() {
	userIds := []int{20597359310}
	api := &DistrictApi{
		DistrictsPath: "./json_files/757359_2.json",
		PositionsPath: "./json_files/positions.json",
	}
	reporter := Report{DistrictApi: api}
	for _, id := range userIds {
		if err := reporter.Message(id); err != nil {
			panic(err)
		}
	}

	//jsonData := readJson("./json_files/757359_2.json")
	//postitionData := readJson("./json_files/positions.json")
	//var districts []*District
	//var positions Positions
	//json.Unmarshal(jsonData, &districts)
	//json.Unmarshal(postitionData, &positions)
	//positionIds := getPositionIds(&positions, userId)
	//fmt.Println("positionIds", positionIds)
	//userDistricts := getDistricts(districts, positionIds)
	//messageDistricts(userDistricts)
	//fmt.Printf("%#v", userDistricts)
}

func readJson(filePath string) []byte {
	file, e := ioutil.ReadFile(filePath)
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}
	return file
}

func getPositionIds(positions *Positions, userId int) []int {
	var positionIds []int
	for _, position := range positions.Elders {
		if position.IndividualId == userId {
			positionIds = append(positionIds, position.Id)
		}
	}
	for _, position := range positions.HighPriests {
		if position.IndividualId == userId {
			positionIds = append(positionIds, position.Id)
		}
	}
	return positionIds
}

func getDistricts(districts []*District, positionIds []int) []*District {
	var userDistricts []*District
	for _, district := range districts {
		if intInSlice(district.DistrictLeaderId, &positionIds) == true {
			fmt.Println("DistrictLeaderId", district.DistrictLeaderId)
			fmt.Println("DistrictId", district.Id)
			userDistricts = append(userDistricts, district)
		}
	}
	return userDistricts
}

func intInSlice(a int, list *[]int) bool {
	for _, b := range *list {
		if a == b {
			return true
		}
	}
	return false
}

func messageDistricts(districts []*District) {
	for _, district := range districts {
		thing := *district
		fmt.Println("district", thing.Id)
		messageDistrict(*district)
	}
}

func messageDistrict(district District) {
	fmt.Println(district.Id)
}
