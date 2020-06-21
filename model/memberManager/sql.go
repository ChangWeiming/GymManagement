package membermanager

import (
	mysql "GymManagement/MySQL"
	"GymManagement/model/member"
	"errors"
)

func saveMember(mem *member.Member) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("INSERT INTO member (name, password, gender, age, phone_number, address, term) VALUES (?, ?, ?, ?, ?, ?, ?)"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(mem.Name, mem.Password, mem.Gender, mem.Age, mem.PhoneNumber, mem.Address, mem.Term)
		return err
	}
}

func deleteMember(mem *member.Member) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("DELETE FROM member WHERE id = ?"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(mem.ID)
		return err
	}
}

func getMember(mem *member.Member) (map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT * FROM member WHERE id = ?", mem.ID)
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

func putMember(mem *member.Member) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("UPDATE member SET name = ?, password = ?, gender = ?, age = ?, phone_number = ?, address = ?, term = ?)"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(mem.Name, mem.Password, mem.Gender, mem.Age, mem.PhoneNumber, mem.Address, mem.Term)
		if err != nil {
			return err
		}
	}
	return nil
}
