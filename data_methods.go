package main

import (
	"errors"
)

func CreateNewLabelSet(labelset LabelSet) (LabelSet, error) {
	result := DB.Create(&labelset)

	if result.Error != nil {
		return labelset, errors.New("labelset couldn't be created")
	}

	return labelset, nil
}

// UpdateUser updates a user entry in the DB if it's comming from the same user
func UpdateLabelSet(labelset LabelSet) (LabelSet, error) {

	result := DB.Model(&labelset).Updates(labelset)

	if result.Error != nil {
		return labelset, errors.New("label set couldn't be updated")
	}

	DB.First(&labelset, labelset.ID)
	return labelset, nil
}

// FindCustomersForTeam is used to get all customers for a team
func GetLabelSets() (labelsets []LabelSet, err error) {
	DB.Find(&labelsets)
	return labelsets, nil
}

func DeleteLabelSet(id uint) error {
	var ls LabelSet
	if DB.First(&ls, id).RecordNotFound() {
		return errors.New("label set doesn't exist")
	}

	DB.Delete(&ls)
	return nil
}

func GetSingleLabelSet(id uint) (LabelSet, error) {
	var labelset LabelSet
	if DB.Preload("Labels").First(&labelset, id).RecordNotFound() {
		return labelset, errors.New("label set doesn't exist")
	}

	return labelset, nil
}
