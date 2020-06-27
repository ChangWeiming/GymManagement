package coursemanager

import (
	mysql "GymManagement/MySQL"
	"GymManagement/model/course"
	"errors"
	"strconv"
)

func saveCourse(courseTmp *course.Course) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("INSERT INTO course (name, time, people_number) VALUES (?, ?, ?)"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(courseTmp.Name, courseTmp.Time, courseTmp.PeopleNumber)
		return err
	}
}

func deleteCourse(courseTmp *course.Course) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("DELETE FROM course WHERE id = ?"); err != nil {
		return err
	} else {
		_, err = stmt.Exec(courseTmp.ID)
		return err
	}
}

func getCourse(courseTmp *course.Course) (map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT * FROM course WHERE id = ?", courseTmp.ID)
	if err != nil {
		return nil, err
	}
	json, err := mysql.GetResult(res)
	if err != nil {
		return nil, err
	}
	if len(json) == 0 {
		return nil, errors.New("course not found")
	}
	return json[0], nil
}

func putCourse(courseTmp *course.Course) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("UPDATE course SET name = ?, time = ?, people_number = ? WHERE id = ?"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(courseTmp.Name, courseTmp.Time, courseTmp.PeopleNumber, courseTmp.ID)
		if err != nil {
			return err
		}
	}
	return nil
}

func getCourseList() ([]map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT * FROM course")
	if err != nil {
		return nil, err
	}
	if retType, err := mysql.GetResult(res); err != nil {
		return nil, err
	} else {
		return retType, nil
	}
}

func selectCourse(courseID string, memberID int) error {
	db := mysql.GetDB()
	if stmt, err := db.Prepare("INSERT INTO study (course_id, member_id) VALUES (?, ?)"); err != nil {
		return err
	} else {
		_, err := stmt.Exec(courseID, memberID)
		if err != nil {
			return err
		}
	}
	return nil
}

func getCoachCourse(coachID int) ([]map[string]string, error) {
	db := mysql.GetDB()
	res, err := db.Query("SELECT course.* FROM course, teach WHERE course.id = teach.course_id AND teach.coach_id = ?", coachID)
	if err != nil {
		return nil, err
	}
	if retType, err := mysql.GetResult(res); err != nil {
		return nil, err
	} else {
		return retType, nil
	}
}

func getSelectedCourse(memberID int) ([]map[string]string, error) {
	db := mysql.GetDB()
	//fmt.Println(memberID)
	res, err := db.Query("SELECT course.*,coach.name as coach_name FROM course,teach,coach,study WHERE course.id = teach.course_id AND teach.coach_id = coach.id AND study.course_id = course.id AND study.member_id = ?", memberID)
	if err != nil {
		return nil, err
	}
	if retType, err := mysql.GetResult(res); err != nil {
		return nil, err
	} else {
		return retType, nil
	}
}

func getUnelectCourse(memberID int) ([]map[string]string, error) {
	db := mysql.GetDB()
	unselectRes, err := db.Query("SELECT course.* FROM course WHERE course.id NOT IN (SELECT course_id FROM study WHERE member_id = ?)", memberID)
	if err != nil {
		return nil, err
	}

	unselect, err := mysql.GetResult(unselectRes)
	if err != nil {
		return nil, err
	}
	teachRes, err := db.Query("SELECT COUNT(member_id) AS people_number, course_id FROM study GROUP BY course_id")
	if err != nil {
		return nil, err
	}

	teach, err := mysql.GetResult(teachRes)
	if err != nil {
		return nil, err
	}

	for _, v := range unselect {
		for _, t := range teach {
			if t["course_id"] == v["id"] {
				vv, _ := strconv.Atoi(v["people_number"])
				tt, _ := strconv.Atoi(t["people_number"])
				v["people_number"] = strconv.FormatInt(int64(vv-tt), 10)
			}
		}
	}
	return unselect, nil
}
