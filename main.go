package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    //"reflect"
    "strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

// get grades based on totalScore
func getGrade(totalScore int) Grade {
    var grade Grade 
    switch true {
        case totalScore < 140 :
            grade = F 

        case totalScore >= 140 && totalScore < 200 :
            grade = C

        case totalScore >= 200 && totalScore < 280 :
            grade = B

        default : 
            grade = A 
    }
    return grade
}

//parse the grades.csv file and prepare a slice of structs
func parseCSV(filePath string) []student {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file" + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    students, err := csvReader.ReadAll()

    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }
    //fmt.Println(students)

    // type of records is [][]string
    // so we loop through all records and create a student variable for each student
    // and append it / assign it to ith slot in studentRecords variable
    var studentRecords = make([]student, 0, len(students))

    // skipping the first record, as they are column headers
    for _, std := range(students[1:]) {
        //fmt.Println(reflect.TypeOf(std[4]))

        var scores [4]int 
        for i:=3 ; i<=6 ; i++ {
            t1, err := strconv.Atoi(std[i])
            if err != nil {
                log.Fatal("unable to convert " + std[i] + " to integer ", err)
            }
            scores[i-3] = t1
        }
        s := student{
                firstName:      std[0],
                lastName:       std[1],
                university:     std[2],
                test1Score:     scores[0], 
                test2Score:     scores[1],
                test3Score:     scores[2],
                test4Score:     scores[3],
           }
        studentRecords = append(studentRecords , s)
    }

    return studentRecords 
}

// 
func calculateGrade(students []student) []studentStat {
    fmt.Println("Inside calculateGrade")
    // given details along with scores of all students
    // find totalScore and derive the grade based on below grade scale
    /*
    Grade is based on the final score.
    If final score is < 35, then student is graded as F (failed)
    If final score is >= 35 and < 50, then student is graded as C
    If final score is >= 50 and < 70, then student is graded as B
    If final score is >= 70, then student is graded as A
        type studentStat struct {
                student
                finalScore float32
                grade      Grade
        }
     */	

    var stdStats = make([]studentStat, 0, len(students))

    for _, std := range(students) {
        totalScore := std.test1Score + std.test2Score + std.test3Score + std.test4Score 
        //fmt.Printf("%v",float32(float32(totalScore)/float32(4)))
        stdStats = append(stdStats, studentStat{
                    student: std,
                    finalScore : float32(totalScore)/float32(4),
                    grade : getGrade(totalScore)})
    }

    return stdStats
}

// find the student with highest marks
func findOverallTopper(gs []studentStat) studentStat {

    fmt.Println("Inside findOverallTopper")
    var topper studentStat
    var mm = float32(0)
    for _, std := range(gs) {
        if std.finalScore > mm {
            mm = std.finalScore
            topper = studentStat{
                    student:        std.student,
                    finalScore :    std.finalScore, 
                    grade :         std.grade}
        } 
    }
    return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
    fmt.Println("Inside findTopperPerUniversity")

    var utm = make(map[string]studentStat)
    for _, std := range(gs) {
        currTop, exist := utm[std.student.university]

        if exist {
            if std.finalScore > currTop.finalScore {
               utm[std.student.university] = std 
            }
        } else {
           utm[std.student.university] = std 
        }
         
    }
    return utm
}
/*
func main() {
    students := parseCSV("grades.csv")

    //fmt.Println("Type of records is ", calculateGrade(students))
}
*/
