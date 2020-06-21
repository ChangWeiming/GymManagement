package coachmanager

import (
	mysql "GymManagement/MySQL"
	"GymManagement/model/coach"
	"errors"
)

func saveCoach(coachTmp *coach.Coach) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("INSERT INTO admin (name, gender, password, age, phone_number, address, personal_profile) VALUES (?, ?, ?, ?, ?, ?, ?)"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(coachTmp.Name, coachTmp.Gender, coachTmp.Password, coachTmp.Age, coachTmp.PhoneNumber, coachTmp.Address, coachTmp.PersonalProfile)
		return err
	}
}

func deleteCoach(coachTmp *coach.Coach) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("DELETE FROM coach WHERE id = ?"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(coachTmp.ID)
		return err
	}
}

func getCoach(coachTmp *coach.Coach) (map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT * FROM coach WHERE id = ?", coachTmp.ID)
	if err != nil {
		return nil, err
	}
	json, err := mysql.GetResult(res)
	if err != nil {
		return nil, err
	}
	if len(json) == 0 {
		return nil, errors.New("user not found")
	}
	return json[0], nil
}

func putCoach(coachTmp *coach.Coach) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("UPDATE coach SET name = ?, gender = ?, password = ?, age = ?, phone_number = ?, address = ?, personal_profile = ?"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(coachTmp.Name, coachTmp.Gender, coachTmp.Password, coachTmp.Age, coachTmp.PhoneNumber, coachTmp.Address, coachTmp.PersonalProfile)
		if err != nil {
			return err
		}
	}
	return nil
}
