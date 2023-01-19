package models

import (
	"github.com/ShashwatDevtron/attendance/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Student{})
	db.AutoMigrate(&StudentAttendance{})
	db.AutoMigrate(&Teacher{})
	db.AutoMigrate(&TeacherAttendance{})
	db.Model(&StudentAttendance{}).AddForeignKey("student_id", "students(id)", "cascade", "cascade")
	db.Model(&TeacherAttendance{}).AddForeignKey("teacher_id", "teachers(id)", "cascade", "cascade")
}


type Student struct{
	ID 		  uint		`gorm:"primary_key;auto_increment" json:"id"`
	Name      string   `gorm:"not null" json:"name"`
	Class     int      `json:"class"`
	StudentAttendances [] StudentAttendance
}
type StudentAttendance struct{
	StudentID    uint       `json:"id"`
	StudentClass int        `json:"class"`
	Year         int        `json:"year"`
	Month        time.Month       `json:"month"`
	Date         int        `json:"date"`
	Punchin      time.Time       `json:"punchin"`
	Punchout     time.Time       `json:"punchout"`
	
}


func (s Student) CreateStudent() Student{
	db.NewRecord(s)
	db.Create(&s)
	return s
}

func (s *Student) DeleteStudent(ID int64) Student{
	var student Student
	db.Where("ID=?", uint(ID)).Delete(student)
	return student
}
 

func (a *StudentAttendance) GetStudentAttendance(Id int64, month time.Month, year int)(*[]StudentAttendance , *gorm.DB){
	var getStudentAt[] StudentAttendance
	db := db.Where("student_id=? AND month=? AND year=?" ,uint(Id),month ,year).Find(&getStudentAt)
	return &getStudentAt, db
}

func (a *StudentAttendance) GetAttendanceOfClass(class int, date int, month time.Month, year int )(*[]StudentAttendance , *gorm.DB){
	var getClassAt[] StudentAttendance
	db := db.Where("student_class=? AND date=? AND month=? AND year=?" ,class,date,month,year).Find(&getClassAt)
	return &getClassAt, db
}

func (a StudentAttendance) StudentPunchIn() StudentAttendance{
	db.NewRecord(a)
	db.Create(&a)
	return a
}

func (a *StudentAttendance) GetStudentsTodayAttendance(Id int64, date int,month time.Month, year int)(*StudentAttendance , *gorm.DB){
	var getStudentAt StudentAttendance
	db := db.Where("student_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Last(&getStudentAt)
	return &getStudentAt, db
}

func (a *StudentAttendance) DeleteTodaysStudentAttendance(Id int64, date int,month time.Month, year int){
	var getStudentAt StudentAttendance
	 db.Where("student_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Delete(&getStudentAt)	
}

  func (s *Student) GetStudent(Id int64)(*Student , *gorm.DB){
	var getStudent Student
	db := db.Where("ID=?",uint(Id)).Find(&getStudent)
	return &getStudent, db
}

