package controllers

import(
	"encoding/json"
	"fmt"
	"time"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/ShashwatDevtron/attendance/pkg/utils"
	"github.com/ShashwatDevtron/attendance/pkg/models"
	
)

var NewStudent = models.Student{}//cretaing new student of type Student struct that we defined in model
var NewTeacher = models.Teacher{}
var StudentTodayAttendance = models.StudentAttendance{}
var TeacherTodayAttendance = models.TeacherAttendance{}





func CreateStudent(w http.ResponseWriter, r *http.Request){
	CreateStudent := &NewStudent
	utils.ParseBody(r, CreateStudent)//utils mai ek function tha parse body 
	s:= CreateStudent.CreateStudent()// s was somethig that was created in database and returned from the function create student
	res, _:= json.Marshal(s)//converting it to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request){
	CreateTeacher := &NewTeacher
	utils.ParseBody(r, CreateTeacher)//utils mai ek function tha parse body 
	s:= CreateTeacher.CreateTeacher()// s was somethig that was created in database and returned from the function create student
	res, _:= json.Marshal(s)//converting it to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func DeleteStudent(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	studentId := vars["Id"]
	ID, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	DoesStudentExist,_ := NewStudent.GetStudent(ID)
	if DoesStudentExist.ID == 0{
		fmt.Fprintln(w,"Student with given Id does not exist")
	}else{
	student := NewStudent.DeleteStudent(ID)
	//now giving response to user by postman in json format
	res, _ := json.Marshal(student)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	}
}



func DeleteTeacher(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	teacherId := vars["Id"]
	ID, err := strconv.ParseInt(teacherId, 0, 0)
	if err != nil{
		fmt.Println("error while parsing")
	}
	DoesTeacherExist,_ := NewTeacher.GetTeacher(ID)
	if DoesTeacherExist.ID == 0{
		fmt.Fprintln(w,"Teacher with given Id does not exist")
	}else{
	teacher := NewTeacher.DeleteTeacher(ID)
	//now giving response to user by postman in json format
	res, _ := json.Marshal(teacher)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	}
}

func GetStudentAttendance(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	studentId := vars["Id"]
	thisMonth := vars["month"]
	thisYear := vars["year"]

	ID , err := strconv.ParseInt(studentId,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	MONTH , err := strconv.ParseInt(thisMonth,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	YEAR , err := strconv.ParseInt(thisYear,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	DoesStudentExist,_ := NewStudent.GetStudent(ID)
	if DoesStudentExist.ID == 0{
		fmt.Fprintln(w,"Student with given Id does not exist")
	}else{
	studentDetails, _:= StudentTodayAttendance.GetStudentAttendance(ID,time.Month(MONTH), int(YEAR))
	res, _ := json.Marshal(studentDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	}
}


func GetTeacherAttendance(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	teacherId := vars["Id"]
	thisMonth := vars["month"]
	thisYear := vars["year"]

	ID , err := strconv.ParseInt(teacherId,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	MONTH , err := strconv.ParseInt(thisMonth,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	YEAR , err := strconv.ParseInt(thisYear,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	DoesTeacherExist,_ := NewTeacher.GetTeacher(ID)
	if DoesTeacherExist.ID == 0{
		fmt.Fprintln(w,"Teacher with given Id does not exist")
	}else{
	teacherDetails, _:= TeacherTodayAttendance.GetTeacherAttendance(ID,time.Month(MONTH), int(YEAR))
	res, _ := json.Marshal(teacherDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	}
}



func GetAttendanceOfClass(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	classIs := vars["class"]
	dateIs := vars["date"]
	monthIs := vars["month"]
	yearIs := vars["year"]
	CLASS , err := strconv.ParseInt(classIs,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}	
	DATE , err := strconv.ParseInt(dateIs,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}	
	MONTH , err := strconv.ParseInt(monthIs,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	YEAR , err := strconv.ParseInt(yearIs,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	classDetails, _:= StudentTodayAttendance.GetAttendanceOfClass(int(CLASS),int(DATE),time.Month(MONTH),int(YEAR))
	res, _ := json.Marshal(classDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}




func StudentPunchIn(w http.ResponseWriter, r *http.Request){
	CreateSTudentAttendance := StudentTodayAttendance	
	
	vars := mux.Vars(r)
	studentId := vars["Id"]
	ID , err := strconv.ParseInt(studentId,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}	  
	DoesStudentExist,_ := NewStudent.GetStudent(ID)
	if DoesStudentExist.ID == 0{
		fmt.Fprintln(w,"Student with given Id does not exist")
	}else{

	 studentTodayAttendenceDetail,_ := StudentTodayAttendance.GetStudentsTodayAttendance(ID,time.Now().Day(),time.Now().Month(),time.Now().Year())
	 if studentTodayAttendenceDetail.Punchin.IsZero(){		 
	CreateSTudentAttendance.StudentID =uint(ID)
	studentClassDetatails,_ :=NewStudent.GetStudent(ID)
	CreateSTudentAttendance.StudentClass = studentClassDetatails.Class
	CreateSTudentAttendance.Year = time.Now().Year()
	CreateSTudentAttendance.Month = time.Now().Month()
	CreateSTudentAttendance.Date = time.Now().YearDay()
	CreateSTudentAttendance.Punchin = time.Now()
	s:= CreateSTudentAttendance.StudentPunchIn()
	res, _:= json.Marshal(s)//converting it to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	}else{
		fmt.Fprintf(w,"You have already punched in for today")
	}
}
}


func TeacherPunchIn(w http.ResponseWriter, r *http.Request){
	CreateTeacherAttendance := TeacherTodayAttendance
	vars := mux.Vars(r)
	studentId := vars["Id"]
	ID , err := strconv.ParseInt(studentId,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}	
	DoesTeacherExist,_ := NewTeacher.GetTeacher(ID)
	if DoesTeacherExist.ID == 0{
		fmt.Fprintln(w,"Teacher with given Id does not exist")
	}else{

	teacherTodayAttendenceDetail,_ := TeacherTodayAttendance.GetTeachersTodayAttendance(ID,time.Now().Day(),time.Now().Month(),time.Now().Year())
	
	if teacherTodayAttendenceDetail.Punchin.IsZero(){
	CreateTeacherAttendance.TeacherID =uint(ID)	
	CreateTeacherAttendance.Year = time.Now().Year()
	CreateTeacherAttendance.Month = time.Now().Month()
	CreateTeacherAttendance.Date = time.Now().YearDay()
	CreateTeacherAttendance.Punchin = time.Now()
	s:= CreateTeacherAttendance.TeacherPunchIn()
	res, _:= json.Marshal(s)//converting it to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)
	}else{
		fmt.Fprintf(w,"You have already punched in for today")
	}
}
}

func StudentPunchOut(w http.ResponseWriter, r *http.Request){
	studentsTodaysAttendanceRecord := StudentTodayAttendance
	CreateSTudentAttendance := StudentTodayAttendance
	vars := mux.Vars(r)
	studentId := vars["Id"]
	ID , err := strconv.ParseInt(studentId,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}
	DoesStudentExist,_ := NewStudent.GetStudent(ID)
	if DoesStudentExist.ID == 0{
		fmt.Fprintln(w,"Student with given Id does not exist")
	}else{
	studentTodayAttendenceDetail,_ := StudentTodayAttendance.GetStudentsTodayAttendance(ID,time.Now().Day(),time.Now().Month(),time.Now().Year())
	if studentTodayAttendenceDetail.Punchin.IsZero(){
		fmt.Fprintln(w, "You have not punched in for today")
	}else if !studentTodayAttendenceDetail.Punchout.IsZero() && !studentTodayAttendenceDetail.Punchin.IsZero(){
		fmt.Fprintln(w, "You have already punched out for today")
	}else if studentTodayAttendenceDetail.Punchout.IsZero() && !studentTodayAttendenceDetail.Punchin.IsZero(){
	CreateSTudentAttendance.StudentID =uint(ID)	
	CreateSTudentAttendance.StudentClass = studentTodayAttendenceDetail.StudentClass
	CreateSTudentAttendance.Year = studentTodayAttendenceDetail.Year
	CreateSTudentAttendance.Month = studentTodayAttendenceDetail.Month
	CreateSTudentAttendance.Date = studentTodayAttendenceDetail.Date
	CreateSTudentAttendance.Punchin = studentTodayAttendenceDetail.Punchin
		//delete student
	StudentTodayAttendance.DeleteTodaysStudentAttendance(ID,time.Now().Day(),time.Now().Month(),time.Now().Year())
	studentsTodaysAttendanceRecord.StudentID = CreateSTudentAttendance.StudentID
	studentsTodaysAttendanceRecord.StudentClass = CreateSTudentAttendance.StudentClass
	studentsTodaysAttendanceRecord.Year = CreateSTudentAttendance.Year
	studentsTodaysAttendanceRecord.Date = CreateSTudentAttendance.Date
	studentsTodaysAttendanceRecord.Month = CreateSTudentAttendance.Month
	studentsTodaysAttendanceRecord.Punchin = CreateSTudentAttendance.Punchin
	studentsTodaysAttendanceRecord.Punchout = time.Now()
	s:= studentsTodaysAttendanceRecord.StudentPunchIn()
	res, _:= json.Marshal(s)//converting it to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)		
	} 
	}
}


func TeacherPunchOut(w http.ResponseWriter, r *http.Request){
	teacherssTodaysAttendanceRecord := TeacherTodayAttendance
	CreateTeacherAttendance := TeacherTodayAttendance
	vars := mux.Vars(r)
	studentId := vars["Id"]
	ID , err := strconv.ParseInt(studentId,10,32)
	if err != nil{
		fmt.Println("error while parsing")
	}	
	DoesTeacherExist,_ := NewTeacher.GetTeacher(ID)
	if DoesTeacherExist.ID == 0{
		fmt.Fprintln(w,"Teacher with given Id does not exist")
	}else{
	teacherTodayAttendenceDetail,_ := TeacherTodayAttendance.GetTeachersTodayAttendance(ID,time.Now().Day(),time.Now().Month(),time.Now().Year())
	if teacherTodayAttendenceDetail.Punchin.IsZero(){
		fmt.Fprintln(w, "You have not punched in for today")
	}else if !teacherTodayAttendenceDetail.Punchout.IsZero() && !teacherTodayAttendenceDetail.Punchin.IsZero(){
		fmt.Fprintln(w, "You have already punched out for today")
	}else if teacherTodayAttendenceDetail.Punchout.IsZero() && !teacherTodayAttendenceDetail.Punchin.IsZero(){
	CreateTeacherAttendance.TeacherID =uint(ID)	
	CreateTeacherAttendance.Year = teacherTodayAttendenceDetail.Year
	CreateTeacherAttendance.Month = teacherTodayAttendenceDetail.Month
	CreateTeacherAttendance.Date = teacherTodayAttendenceDetail.Date
	CreateTeacherAttendance.Punchin = teacherTodayAttendenceDetail.Punchin
		//delete student
	TeacherTodayAttendance.DeleteTodaysTeacherAttendance(ID,time.Now().Day(),time.Now().Month(),time.Now().Year())
	teacherssTodaysAttendanceRecord.TeacherID = CreateTeacherAttendance.TeacherID
	teacherssTodaysAttendanceRecord.Year = CreateTeacherAttendance.Year
	teacherssTodaysAttendanceRecord.Date = CreateTeacherAttendance.Date
	teacherssTodaysAttendanceRecord.Month = CreateTeacherAttendance.Month
	teacherssTodaysAttendanceRecord.Punchin = CreateTeacherAttendance.Punchin
	teacherssTodaysAttendanceRecord.Punchout = time.Now()
	s:= teacherssTodaysAttendanceRecord.TeacherPunchIn()
	res, _:= json.Marshal(s)//converting it to json
	w.WriteHeader(http.StatusOK)
	w.Write(res)		
	}
	}
}
