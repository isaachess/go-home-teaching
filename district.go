package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Reader interface {
	Read(filePath string) ([]byte, error)
}

type DistrictApi struct {
	DistrictsPath string
	PositionsPath string
	Reader        Reader
}

func (d *DistrictApi) GetDistrict(userId int) (*District, error) {
	var allDistricts []*District
	if err := d.readAndUnmarshal(d.DistrictsPath, allDistricts); err != nil {
		return nil, err
	}

	positionId, err := d.getPositionId(userId)
	if err != nil {
		return nil, err
	}

	for _, district := range allDistricts {
		if positionId == district.DistrictLeaderId {
			return district, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Cannot find district for ID %d", userId))
}

func (d *DistrictApi) getPositionId(userId int) (int, error) {
	var positions Positions
	if err := d.readAndUnmarshal(d.PositionsPath, &positions); err != nil {
		return 0, err
	}

	for _, position := range positions.Elders {
		if userId == position.IndividualId {
			return position.Id, nil
		}
	}
	for _, position := range positions.HighPriests {
		if userId == position.IndividualId {
			return position.Id, nil
		}
	}
	return 0, errors.New(fmt.Sprintf("Cannot find position for ID %d", userId))
}

func (d *DistrictApi) readAndUnmarshal(path string, v interface{}) error {
	jsonData, err := d.Reader.Read(path)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, v); err != nil {
		return err
	}
	return nil
}

func containsInt(ints []int, test int) bool {
	for _, num := range ints {
		if num == test {
			return true
		}
	}
	return false
}
