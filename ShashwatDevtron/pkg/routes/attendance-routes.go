// create the routes where the users will hit from postman
package routes

import (
	"github.com/ShashwatDevtron/attendance/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterAttendanceMangementRoutes = func(router *mux.Router) {
	router.HandleFunc("/student/add", controllers.CreateStudent).Methods("POST")                          //creates a student
	router.HandleFunc("/teacher/add", controllers.CreateTeacher).Methods("POST")                          //creates a teacher
	router.HandleFunc("/student/delete/{Id}", controllers.DeleteStudent).Methods("DELETE")                    // delete a student
	router.HandleFunc("/teacher/delete/{Id}", controllers.DeleteTeacher).Methods("DELETE")                    // delete a teacher
	router.HandleFunc("/student/getAttendance/{Id}/{month}/{year}", controllers.GetStudentAttendance).Methods("GET") //gets the attendance of a student of a particular month
	router.HandleFunc("/teacher/getAttendance/{Id}/{month}/{year}", controllers.GetTeacherAttendance).Methods("GET") // gets the attendance of the teacher for a particular month
	router.HandleFunc("/student/class/{class}/{date}/{month}/{year}", controllers.GetAttendanceOfClass).Methods("GET")       //gets the attendance of students of a particular class of a particular month
	router.HandleFunc("/student/punchin/{Id}", controllers.StudentPunchIn).Methods("POST")     //students will be able to punchin and punchout fot a particular day
    router.HandleFunc("/teacher/punchin/{Id}", controllers.TeacherPunchIn).Methods("POST")
	router.HandleFunc("/student/punchout/{Id}", controllers.StudentPunchOut).Methods("PUT")
	router.HandleFunc("/teacher/punchout/{Id}", controllers.TeacherPunchOut).Methods("PUT")
}
