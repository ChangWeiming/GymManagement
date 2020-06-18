package coursemanager

import (
	"GymManagement/model/course"
	"container/list"
)

var courseList list.List

func PutCourseList(ob course.Course) {
	courseList.PushBack(ob)

}

func GetCourseList() *list.List {
	return &courseList
}
