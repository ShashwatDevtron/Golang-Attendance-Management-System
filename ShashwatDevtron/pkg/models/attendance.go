package models

import (
	"github.com/ShashwatDevtron/attendance/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

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

type  Teacher struct{
	ID 		  uint		`gorm:"primary_key;auto_increment" json:"id"`
	Name      string   `gorm:"not null" json:"name"`
	TAttendance [] TeacherAttendance
	
}

type TeacherAttendance struct{
	TeacherID    uint       `json:"id"`
	Year         int        `json:"year"`
	Month        time.Month        `json:"month"`
	Date         int        `json:"date"`
	Punchin      time.Time       `json:"punchin"`
	Punchout     time.Time       `json:"punchout"`
	
}



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


func (s Student) CreateStudent() Student{
	db.NewRecord(s)
	db.Create(&s)
	return s
}
func (t Teacher) CreateTeacher() Teacher{
	db.NewRecord(t)
	db.Create(&t)
	return t
}


func (s *Student) DeleteStudent(ID int64) Student{
	var student Student
	db.Where("ID=?", uint(ID)).Delete(student)
	return student
}
 
func (t *Teacher) DeleteTeacher(ID int64) Teacher{
	var teacher Teacher
	db.Where("ID=?", uint(ID)).Delete(teacher)
	return teacher
}


func (a *StudentAttendance) GetStudentAttendance(Id int64, month time.Month, year int)(*[]StudentAttendance , *gorm.DB){
	var getStudentAt[] StudentAttendance
	db := db.Where("student_id=? AND month=? AND year=?" ,uint(Id),month ,year).Find(&getStudentAt)
	return &getStudentAt, db
}


func (a *TeacherAttendance) GetTeacherAttendance(Id int64, month time.Month, year int)(*[]TeacherAttendance , *gorm.DB){
	var getTeacherAt[] TeacherAttendance
	db := db.Where("teacher_id=? AND month=? AND year=?" ,uint(Id),month ,year).Find(&getTeacherAt)
	return &getTeacherAt, db
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
func (a TeacherAttendance) TeacherPunchIn() TeacherAttendance{
	db.NewRecord(a)
	db.Create(&a)
	return a
}



func (a *StudentAttendance) GetStudentsTodayAttendance(Id int64, date int,month time.Month, year int)(*StudentAttendance , *gorm.DB){
	var getStudentAt StudentAttendance
	db := db.Where("student_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Last(&getStudentAt)
	return &getStudentAt, db
}



func (a *TeacherAttendance) GetTeachersTodayAttendance(Id int64, date int,month time.Month, year int)(*TeacherAttendance , *gorm.DB){
	var getTeacherAt TeacherAttendance
	db := db.Where("teacher_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Last(&getTeacherAt)
	return &getTeacherAt, db
}

func (a *StudentAttendance) DeleteTodaysStudentAttendance(Id int64, date int,month time.Month, year int){
	var getStudentAt StudentAttendance
	 db.Where("student_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Delete(&getStudentAt)	
}

func (a *TeacherAttendance) DeleteTodaysTeacherAttendance(Id int64, date int,month time.Month, year int){
	var getTeacherAt TeacherAttendance
	 db.Where("student_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Delete(&getTeacherAt)	
}



  func (s *Student) GetStudent(Id int64)(*Student , *gorm.DB){
	var getStudent Student
	db := db.Where("ID=?",uint(Id)).Find(&getStudent)
	return &getStudent, db
}





