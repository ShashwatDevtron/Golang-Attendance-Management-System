package models;

import (
	
	"github.com/jinzhu/gorm"
	"time"
)

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

func (t Teacher) CreateTeacher() Teacher{
	db.NewRecord(t)
	db.Create(&t)
	return t
}
func (t *Teacher) DeleteTeacher(ID int64) Teacher{
	var teacher Teacher
	db.Where("ID=?", uint(ID)).Delete(teacher)
	return teacher
}

func (a *TeacherAttendance) GetTeacherAttendance(Id int64, month time.Month, year int)(*[]TeacherAttendance , *gorm.DB){
	var getTeacherAt[] TeacherAttendance
	db := db.Where("teacher_id=? AND month=? AND year=?" ,uint(Id),month ,year).Find(&getTeacherAt)
	return &getTeacherAt, db
}

func (a TeacherAttendance) TeacherPunchIn() TeacherAttendance{
	db.NewRecord(a)
	db.Create(&a)
	return a
}


func (a *TeacherAttendance) GetTeachersTodayAttendance(Id int64, date int,month time.Month, year int)(*TeacherAttendance , *gorm.DB){
	var getTeacherAt TeacherAttendance
	db := db.Where("teacher_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Last(&getTeacherAt)
	return &getTeacherAt, db
}

func (a *TeacherAttendance) DeleteTodaysTeacherAttendance(Id int64, date int,month time.Month, year int){
	var getTeacherAt TeacherAttendance
	 db.Where("teacher_id=? AND date=? AND month=? AND year=?" ,uint(Id),date,month ,year).Delete(&getTeacherAt)	
}

func (t *Teacher) GetTeacher(Id int64)(*Teacher , *gorm.DB){
	var getTeacher Teacher
	db := db.Where("ID=?",uint(Id)).Find(&getTeacher)
	return &getTeacher, db
}
