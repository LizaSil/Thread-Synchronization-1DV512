package main

import (
  "context"
  "fmt"
  "sync"
  "time"

  "golang.org/x/sync/semaphore"
)

var Wg sync.WaitGroup

// How many can access concurrently
var Sem = semaphore.NewWeighted(1)

func main() {
  fmt.Println("Select a task to run:")
  fmt.Println("1. Task 1")
  fmt.Println("2. Task 2")
  fmt.Println("3. Task 3")
  fmt.Println("------------------")

  var input int
  fmt.Scanln(&input)

  switch input {
  case 1:
    Task_1()
  case 2:
    Task_2()
  case 3:
    Task_3()
  default:
    fmt.Println("Invalid input")
  }

}

func Task_1() {
  // Create two semaphores with a capacity of 1
  var semA, semB = semaphore.NewWeighted(1), semaphore.NewWeighted(1)
  semB.Acquire(context.Background(), 1) // Lock semB

  // Launch two goroutines
  Wg.Add(2)
  go func() {
    for i := 0; i < 10; i++ {
      semA.Acquire(context.Background(), 1) // Lock semA
      fmt.Print("A")
      semB.Release(1) // release lock on semB
    }
    Wg.Done()
  }()
  go func() {
    for i := 0; i < 10; i++ {
      semB.Acquire(context.Background(), 1) // Lock semB
      fmt.Print("B")
      semA.Release(1) // release lock on semA
    }
    Wg.Done()
  }()

  // Wait for the goroutines to finish
  Wg.Wait()
  fmt.Println()
}

var amount = 10

func Task_2() {
  pattern := 0
  var semA, semB, semC, semD = semaphore.NewWeighted(1), semaphore.NewWeighted(1), semaphore.NewWeighted(1), semaphore.NewWeighted(1)
  semA.Acquire(context.Background(), 1) // Lock semA
  semB.Acquire(context.Background(), 1) // Lock semB
  semC.Acquire(context.Background(), 1) // Lock semC
  semD.Acquire(context.Background(), 1) // Lock semD

  // A
  semA.Release(1) // Start with A
  go func() {
    for i := 0; i < amount; i++ {
      switch pattern {
      case 1:
        {
          semA.Acquire(context.Background(), 1) // Lock semA
          fmt.Print("A")
          pattern = 0
          semC.Release(1)
        }
      case 0:
        {
          semA.Acquire(context.Background(), 1) // Lock semA
          fmt.Print("A")
          pattern = 1
          semB.Release(1) // release lock on semB
        }
      }
    }

  }()
  // B
  go func() {
    for i := 0; i < amount; i++ {
      semB.Acquire(context.Background(), 1) // Lock semB
      fmt.Print("B")
      semA.Release(1) // release lock on semC
    }
  }()
  // C
  go func() {
    for i := 0; i < amount; i++ {
      switch pattern {
      case 1:
        {
          semC.Acquire(context.Background(), 1) // Lock semC
          pattern = 0
          fmt.Print("C")
          semA.Release(1) // release lock on semA
        }
      case 0:
        {
          semC.Acquire(context.Background(), 1) // Lock semC
          pattern = 1
          fmt.Print("C")
          semD.Release(1) // release lock on semD
        }
      }
    }
  }()
  // D
  go func() {
    for i := 0; i < amount; i++ {
      semD.Acquire(context.Background(), 1) // Lock semD
      fmt.Print("D")
      semC.Release(1) // release lock on semA
    }
  }()

  // Wait for the goroutines to finish
  time.Sleep(1 * time.Second)
  fmt.Println()

}
