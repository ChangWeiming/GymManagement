package adminmanager

import (
	mysql "GymManagement/MySQL"
	"GymManagement/model/admin"
	"errors"
)

func saveAdmin(administrator *admin.Admin) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("INSERT INTO admin (name, password, phone_number) VALUES (?, ?, ?)"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(administrator.Name, administrator.Password, administrator.PhoneNumber)
		return err
	}
}

func deleteAdmin(administrator *admin.Admin) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("DELETE FROM admin WHERE id = ?"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(administrator.ID)
		return err
	}
}

func getAdmin(administrator *admin.Admin) (map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT * FROM admin WHERE id = ?", administrator.ID)
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

func putAdmin(administrator *admin.Admin) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("UPDATE admin SET name = ?, password = ?, phone_number = ?)"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(administrator.Name, administrator.Password, administrator.PhoneNumber)
		if err != nil {
			return err
		}
	}
	return nil
}
