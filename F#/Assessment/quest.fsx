#r "nuget: Microsoft.Data.Sqlite, 8.0.11"
open System
open Microsoft.Data.Sqlite

type StudentStatus =
    | Active
    | Inactive

// Student Record
type Student = {
    Id: int
    Name: string
    Age: int
    Course: string
    Email: string
    Marks: float
    Status: StudentStatus
}



let getGrade marks =
    match marks with
    | m when m >= 90.0 -> "A+"
    | m when m >= 80.0 -> "A"
    | m when m >= 70.0 -> "B"
    | m when m >= 60.0 -> "C"
    | m when m >= 50.0 -> "D"
    | _ -> "F"

let getStatus status =
    match status with
    | Active -> "Active"
    | Inactive -> "Inactive"


let getStudentNames students =
    students |> List.map (fun s -> s.Name)

let getPassedStudents students =
    students |> List.filter (fun s -> s.Marks >= 50.0)

let getStudentsByCourse (course: string) (students: Student list) =
    students
    |> List.filter (fun (s: Student) ->
        s.Course.ToLower() = course.ToLower())
let getActiveStudents students =
    students
    |> List.filter (fun s ->
        match s.Status with
        | Active -> true
        | Inactive -> false
    )



let connectionString = "Data Source=students.db"

// Create database and table
let initializeDatabase () =
    async {
        use connection = new SqliteConnection(connectionString)

        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()

        command.CommandText <-
            """
            CREATE TABLE IF NOT EXISTS Students (
                Id INTEGER PRIMARY KEY AUTOINCREMENT,
                Name TEXT NOT NULL,
                Age INTEGER NOT NULL,
                Course TEXT NOT NULL,
                Email TEXT NOT NULL UNIQUE,
                Marks REAL NOT NULL,
                Status TEXT NOT NULL
            );
            """

        let! _ =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        printfn "Database initialized successfully!"
    }



let addStudent student =
    async {
        use connection = new SqliteConnection(connectionString)

        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()

        command.CommandText <-
            """
            INSERT INTO Students
            (Name, Age, Course, Email, Marks, Status)
            VALUES
            (@Name, @Age, @Course, @Email, @Marks, @Status)
            """

        command.Parameters.AddWithValue("@Name", student.Name) |> ignore
        command.Parameters.AddWithValue("@Age", student.Age) |> ignore
        command.Parameters.AddWithValue("@Course", student.Course) |> ignore
        command.Parameters.AddWithValue("@Email", student.Email) |> ignore
        command.Parameters.AddWithValue("@Marks", student.Marks) |> ignore
        command.Parameters.AddWithValue("@Status", getStatus student.Status) |> ignore

        let! result =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        return result > 0
    }



let getAllStudents () =
    async {
        use connection = new SqliteConnection(connectionString)

        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()

        command.CommandText <- "SELECT * FROM Students"

        use! reader =
            command.ExecuteReaderAsync()
            |> Async.AwaitTask

        let students = ResizeArray<Student>()

        let mutable reading = true

        while reading do
            let! hasRow =
                reader.ReadAsync()
                |> Async.AwaitTask

            if hasRow then

                let status =
                    match reader.GetString(6) with
                    | "Active" -> Active
                    | _ -> Inactive

                let student = {
                    Id = reader.GetInt32(0)
                    Name = reader.GetString(1)
                    Age = reader.GetInt32(2)
                    Course = reader.GetString(3)
                    Email = reader.GetString(4)
                    Marks = reader.GetDouble(5)
                    Status = status
                }

                students.Add(student)

            else
                reading <- false

        return students |> Seq.toList
    }



let updateStudent id name age course email marks status =
    async {
        use connection = new SqliteConnection(connectionString)

        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()

        command.CommandText <-
            """
            UPDATE Students
            SET Name=@Name,
                Age=@Age,
                Course=@Course,
                Email=@Email,
                Marks=@Marks,
                Status=@Status
            WHERE Id=@Id
            """

        command.Parameters.AddWithValue("@Id", id) |> ignore
        command.Parameters.AddWithValue("@Name", name) |> ignore
        command.Parameters.AddWithValue("@Age", age) |> ignore
        command.Parameters.AddWithValue("@Course", course) |> ignore
        command.Parameters.AddWithValue("@Email", email) |> ignore
        command.Parameters.AddWithValue("@Marks", marks) |> ignore
        command.Parameters.AddWithValue("@Status", getStatus status) |> ignore

        let! result =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        return result > 0
    }



let deleteStudent id =
    async {
        use connection = new SqliteConnection(connectionString)

        do! connection.OpenAsync() |> Async.AwaitTask

        use command = connection.CreateCommand()

        command.CommandText <-
            "DELETE FROM Students WHERE Id=@Id"

        command.Parameters.AddWithValue("@Id", id) |> ignore

        let! result =
            command.ExecuteNonQueryAsync()
            |> Async.AwaitTask

        return result > 0
    }



let readText message =
    printf "%s" message
    Console.ReadLine()

let readInt message =
    let mutable value = 0
    let mutable valid = false

    while not valid do
        let input = readText message
        valid <- Int32.TryParse(input, &value)

        if not valid then
            printfn "Enter a valid number."

    value

let readFloat message =
    let mutable value = 0.0
    let mutable valid = false

    while not valid do
        let input = readText message
        valid <- Double.TryParse(input, &value)

        if not valid then
            printfn "Enter valid marks."

    value

let readStatus () =
    match (readText "Enter Status (Active/Inactive): ").ToLower() with
    | "active" -> Active
    | _ -> Inactive


let displayStudent student =
    printfn "\n------------------------------"
    printfn "ID      : %d" student.Id
    printfn "Name    : %s" student.Name
    printfn "Age     : %d" student.Age
    printfn "Course  : %s" student.Course
    printfn "Email   : %s" student.Email
    printfn "Marks   : %.2f" student.Marks
    printfn "Grade   : %s" (getGrade student.Marks)
    printfn "Status  : %s" (getStatus student.Status)
    printfn "------------------------------"


// Add Student
let createStudent () =
    async {
        printfn "\n===== ADD STUDENT ====="

        let student = {
            Id = 0
            Name = readText "Enter Name: "
            Age = readInt "Enter Age: "
            Course = readText "Enter Course: "
            Email = readText "Enter Email: "
            Marks = readFloat "Enter Marks: "
            Status = readStatus()
        }

        let! result = addStudent student

        if result then
            printfn "Student added successfully!"
        else
            printfn "Unable to add student."
    }

// View Students
let viewStudents () =
    async {
        let! students = getAllStudents()

        if List.isEmpty students then
            printfn "No students found."
        else
            students |> List.iter displayStudent
    }

// Update Student
let modifyStudent () =
    async {
        printfn "\n===== UPDATE STUDENT ====="

        let id = readInt "Enter Student ID: "
        let name = readText "Enter New Name: "
        let age = readInt "Enter New Age: "
        let course = readText "Enter New Course: "
        let email = readText "Enter New Email: "
        let marks = readFloat "Enter New Marks: "
        let status = readStatus()

        let! result =
            updateStudent id name age course email marks status

        if result then
            printfn "Student updated successfully!"
        else
            printfn "Student not found or update failed."
    }

// Delete Student
let removeStudent () =
    async {
        let id = readInt "Enter Student ID: "

        let! result = deleteStudent id

        if result then
            printfn "Student deleted successfully!"
        else
            printfn "Student not found."
    }

let demonstrateFunctions () =
    async {
        let! students = getAllStudents()

        printfn "\n===== MAP: STUDENT NAMES ====="

        students
        |> getStudentNames
        |> List.iter (printfn "%s")

        printfn "\n===== FILTER: PASSED STUDENTS ====="

        students
        |> getPassedStudents
        |> List.iter displayStudent

        printfn "\n===== FILTER: COMPUTER SCIENCE ====="

        students
        |> getStudentsByCourse "Computer Science"
        |> List.iter displayStudent

        printfn "\n===== FILTER: ACTIVE STUDENTS ====="

        students
        |> getActiveStudents
        |> List.iter displayStudent

        printfn "\n===== PATTERN MATCHING: GRADES ====="

        students
        |> List.iter (fun s ->
            printfn "%s : %s" s.Name (getGrade s.Marks)
        )
    }


let main () =

    initializeDatabase() |> Async.RunSynchronously

    let mutable running = true

    while running do

        printfn "\n================================"
        printfn "     STUDENT MANAGEMENT SYSTEM"
        printfn "================================"
        printfn "1. Add Student"
        printfn "2. View All Students"
        printfn "3. Update Student"
        printfn "4. Delete Student"
        printfn "5. Map, Filter and Pattern Matching"
        printfn "6. Exit"
        printfn "================================"

        let choice = readInt "Enter your choice: "

        match choice with

        | 1 ->
            createStudent() |> Async.RunSynchronously

        | 2 ->
            viewStudents() |> Async.RunSynchronously

        | 3 ->
            modifyStudent() |> Async.RunSynchronously

        | 4 ->
            removeStudent() |> Async.RunSynchronously

        | 5 ->
            demonstrateFunctions() |> Async.RunSynchronously

        | 6 ->
            running <- false
            printfn "Thank you!"

        | _ ->
            printfn "Invalid choice!"

main()