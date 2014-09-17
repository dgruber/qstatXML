package main

import (
   "fmt"
   "github.com/dgruber/drmaa/gestatus"
   "os"
   "strconv"
)

func main() {

   if len(os.Args) <= 1 {
      fmt.Println("Error: Job id as argument required.")
      os.Exit(1)
   }

   if jobStatus, err := gestatus.GetJob(os.Args[1]); err != nil {
      fmt.Println("Job not found.")
      os.Exit(2)
   } else {
      fmt.Printf("Job Name: %s\n", jobStatus.JobName())
      fmt.Printf("Job Number: %d\n", jobStatus.JobId())
      fmt.Printf("Job Script: %s\n", jobStatus.JobScript())
      fmt.Printf("Job Args: %s\n", jobStatus.JobArgs())
      fmt.Printf("Job Owner: %s\n", jobStatus.JobOwner())
      fmt.Printf("Job Group: %s\n", jobStatus.JobGroup())
      fmt.Printf("Job UID: %d\n", jobStatus.JobUID())
      fmt.Printf("Job GID: %d\n", jobStatus.JobGID())
      fmt.Printf("Job accounting string: %s\n", jobStatus.JobAccountName())
      fmt.Printf("Job is now: %t\n", jobStatus.IsImmediateJob())
      fmt.Printf("Job is binary: %t\n", jobStatus.IsBinaryJob())
      fmt.Printf("Job has reservation: %t\n", jobStatus.HasReservation())
      fmt.Printf("Job is array job: %t\n", jobStatus.IsArrayJob())
      fmt.Printf("Job merges stderr %t\n", jobStatus.JobMergesStderr())
      fmt.Printf("Job has 'no shell' requested: %t\n", jobStatus.HasNoShell())
      fmt.Printf("Job has memory binding: %t\n", jobStatus.HasMemoryBinding())
      fmt.Printf("Job memory binding: %s\n", jobStatus.MemoryBinding())
      fmt.Printf("Job submission time: %s\n", jobStatus.SubmissionTime())
      fmt.Printf("Job start time: %s\n", jobStatus.StartTime())
      fmt.Printf("Job run time: %s\n", jobStatus.RunTime())
      fmt.Printf("Job deadline: %s\n", jobStatus.JobDeadline())
      fmt.Printf("Job mail options: %s\n", jobStatus.MailOptions())
      fmt.Printf("Job AR: %d\n", jobStatus.AdvanceReservationID())
      fmt.Printf("Job POSIX priority: %d\n", jobStatus.PosixPriority())
      fmt.Printf("Job class name: %s\n", jobStatus.JobClassName())
      fmt.Printf("Job mailing adresses: %s\n", jobStatus.MailAdresses())
      fmt.Printf("Job destination queue instance list: %s\n",
         jobStatus.DestinationQueueInstanceList())
      fmt.Printf("Job slots list: %s\n", jobStatus.DestinationSlotsList())
      fmt.Printf("Job destination host List: %s\n", jobStatus.DestinationHostList())
      var sum int
      for _, slot := range jobStatus.DestinationSlotsList() {
         if val, err := strconv.Atoi(slot); err == nil {
            sum += val
         }
      }
      fmt.Printf("Job slots sum: %d\n", sum)
      fmt.Printf("Job tasks: %d\n", jobStatus.TasksCount())
      fmt.Printf("Job PE request: %s %d-%d:%d\n", jobStatus.ParallelEnvironment(),
         jobStatus.ParallelEnvironmentMin(), jobStatus.ParallelEnvironmentMax(),
         jobStatus.ParallelEnvironmentStep())
      resources, usage := jobStatus.ResourceUsage(0)
      for i, _ := range resources {
         fmt.Printf("Job resource usage %s: %s\n", resources[i], usage[i])
      }
   }
}
