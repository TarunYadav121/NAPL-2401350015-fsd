open System

printfn "Enter marks for 4 subjects:"

let a = Convert.ToInt32(Console.ReadLine())
let b = Convert.ToInt32(Console.ReadLine())
let c = Convert.ToInt32(Console.ReadLine())
let d = Convert.ToInt32(Console.ReadLine())

let sum = a + b + c + d
let avg = sum / 4

let result =
    if avg >= 60 then "First Division"
    elif avg >= 50 then "Second Division"
    elif avg >= 40 then "Third Division"
    else "Fail"

printfn "Total = %d" sum
printfn "Percentage = %d%%" avg
printfn "Result = %s" result