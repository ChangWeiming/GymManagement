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

func getMember(coachID int) ([]map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT member.*, course.name as course_name, course.time as course_time FROM course,member,study,(SELECT teach.* FROM teach WHERE coach_id = ?) as teach_relation WHERE teach_relation.course_id = study.course_id AND teach_relation.course_id = course.id AND study.member_id = member.id", coachID)
	if err != nil {
		return nil, err
	}
	json, err := mysql.GetResult(res)
	if err != nil {
		return nil, err
	}

	return json, nil
}

func putMember(mem *member.Member) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("UPDATE member SET phone_number = ? WHERE id = ?"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(mem.PhoneNumber, mem.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func putTerm(mem *member.Member) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("UPDATE member SET term = ? WHERE phone_number = ?"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(mem.Term, mem.PhoneNumber)
		return err
	}
}

func getMemberList() ([]map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT * FROM member")
	if err != nil {
		return nil, err
	}
	if retType, err := mysql.GetResult(res); err != nil {
		return nil, err
	} else {
		return retType, nil
	}
}

func postStartTime(memberID int, courseID string, startTime string) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("INSERT INTO study_record (course_id, member_id, start_time) VALUES (?, ?, ?)"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(courseID, memberID, startTime)
		if err != nil {
			return err
		}
	}
	return nil
}

func postLeaveTime(memberID int, courseID string, startTime string, leaveTime string) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("UPDATE study_record SET leave_time = ? WHERE course_id = ? AND member_id = ? AND start_time = ?"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(leaveTime, courseID, memberID, startTime)
		if err != nil {
			return err
		}
	}
	return nil
}

func getTerm(mem *member.Member) (map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT term FROM member WHERE phone_number = ?", mem.PhoneNumber)
	if err != nil {
		return nil, err
	}
	json, err := mysql.GetResult(res)

	if err != nil {
		return nil, err
	}

	if len(json) == 0 {
		return map[string]string{"term": ""}, nil
	}
	return json[0], nil
}
