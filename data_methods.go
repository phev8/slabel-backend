package main

import (
	"errors"

	"github.com/jinzhu/gorm"
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
	if DB.Preload("Labels.Children", func(db *gorm.DB) *gorm.DB {
		return DB.Select("*")
	}).First(&labelset, id).RecordNotFound() {
		return labelset, errors.New("label set doesn't exist")
	}

	return labelset, nil
}

func UpdateLabelTemplate(lt LabelTemplate) (LabelTemplate, error) {
	result := DB.Model(&lt).Updates(lt)

	if result.Error != nil {
		return lt, errors.New("label template couldn't be updated")
	}

	DB.First(&lt, lt.ID)
	return lt, nil
}

func DeleteLabelTemplate(id uint) error {
	var lt LabelTemplate
	if DB.Preload("Children").First(&lt, id).RecordNotFound() {
		return errors.New("label template doesn't exist")
	}
	for _, item := range lt.Children {
		if err := DeleteLabelTemplate(item.ID); err != nil {
			return err
		}
	}

	DB.Delete(&lt)
	return nil
}

func CreateNewLabelTemplate(lt LabelTemplate) (LabelTemplate, error) {
	if lt.LabelSetID <= 0 {
		return lt, errors.New("label set not defined")
	}
	result := DB.Create(&lt)

	if result.Error != nil {
		return lt, errors.New("label template couldn't be created")
	}
	return lt, nil
}

func CreateNewSession(session Session) (Session, error) {
	result := DB.Create(&session)

	if result.Error != nil {
		return session, errors.New("session couldn't be created")
	}

	return session, nil
}

func UpdateSession(session Session) (Session, error) {
	result := DB.Model(&session).Updates(session)

	if result.Error != nil {
		return session, errors.New("session couldn't be updated")
	}

	DB.First(&session, session.ID)
	return session, nil
}

func GetSessions() (sessions []Session, err error) {
	DB.Find(&sessions)
	return sessions, nil
}

func DeleteSession(id uint) error {
	var ls Session
	if DB.First(&ls, id).RecordNotFound() {
		return errors.New("session doesn't exist")
	}

	DB.Delete(&ls)
	return nil
}

func GetSingleSession(id uint) (Session, error) {
	var session Session
	if DB.Preload("Labels").First(&session, id).RecordNotFound() {
		return session, errors.New("session doesn't exist")
	}
	return session, nil
}
