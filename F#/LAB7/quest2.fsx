open System

let workertask (taskid: int) (runtime: int) =
    async {
        printfn "Starting task %d" taskid
        do! Async.Sleep(runtime)
        printfn "Completed task %d" taskid
        return taskid
    }

printfn "Starting concurrent tasks"

let jobs = [
    workertask 1 2000
    workertask 2 3000
    workertask 3 4000
]

let results =
    jobs
    |> Async.Parallel
    |> Async.RunSynchronously

results |> Array.iter (fun result ->
    printfn "Result: %d" result
)