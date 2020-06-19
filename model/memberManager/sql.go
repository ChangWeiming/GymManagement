package membermanager

import (
	mysql "GymManagement/MySQL"
	"GymManagement/model/member"
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
