package main

import (
	"errors"
	"fmt"
	"strings"
)

type districtGetter interface {
	GetDistrict(userId int) (*District, error)
}

type Report struct {
	DistrictApi districtGetter
}

func (r *Report) Message(userId int) error {
	district, err := r.DistrictApi.GetDistrict(userId)
	if err != nil {
		return err
	}

	if err := r.messageDistrict(district); err != nil {
		return err
	}

	return nil
}

func (r *Report) messageDistrict(d *District) error {
	for _, comp := range d.Companionships {
		if err := r.messageCompanionship(&comp); err != nil {
			return err
		}
	}
	return nil
}

func (r *Report) messageCompanionship(c *Companionship) error {
	for _, teacher := range c.Teachers {
		if err := r.messageTeacher(&teacher, c.Assignments); err != nil {
			return err
		}
	}
	return nil
}

func (r *Report) messageTeacher(teacher *Teacher, assignments []Assignment) error {
	//surnames, err := r.formatSurnames(people)
	//if err != nil {
	//return err
	//}

	//message := fmt.Sprintf("Please give me your home teaching report. We have you down for the %s.", surnames)
	//fmt.Println(message)
	return nil
}

func (r *Report) formatSurnames(people []Person) (string, error) {
	familyWord, err := r.getFamilyWord(people)
	if err != nil {
		return "", err
	}

	names, err := r.getNames(people)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", names, familyWord), nil
}

func (r *Report) getFamilyWord(people []Person) (string, error) {
	if len(people) < 1 {
		return "", errors.New("Cannot format surnames, list too short")
	} else if len(people) == 1 {
		return "family", nil
	} else {
		return "families", nil
	}
}

func (r *Report) getNames(people []Person) (string, error) {
	var (
		names []string
		sep   string
	)
	for _, person := range people {
		names = append(names, person.Surname)
	}
	if len(names) > 1 {
		names[len(names)-1] = "and " + names[len(names)-1]
	}
	if len(names) > 2 {
		sep = ", "
	} else {
		sep = " "
	}
	return strings.Join(names, sep), nil
}
