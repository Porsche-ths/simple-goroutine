# Simple Goroutines

Welcome to **Simple Goroutines**, a personal project for learning and experimenting with Goroutines in Go. Feel free to explore and enjoy reading through it!

## Description

Goroutines are lightweight threads managed by the Go runtime, designed to optimize business logic through concurrent programming. The main goal of this repository is to read random float values from a generated `.csv` file containing 1,000 rows and compute their average.

This project includes four modules, each implementing the same calculation method to compare execution times across different Goroutines approaches:

1. **Simple Module**: Average calculation without using Goroutines.
2. **Wait Group Module**: Implements Goroutines using a wait group style.
3. **Channel Module**: Uses channels to implement Goroutines.
4. **Worker Pool Module**: Implements Goroutines with a worker pool style.

## Getting Started

Follow the steps below to run the project:

1. **Generate the `data.csv` file**:  
    Run the Python script to create a `.csv` file with 1,000 random float values.

    ```bash
    python "./fileGenerator.py"
    ```

2. **Run the Go program**:  
    Execute the main program to calculate averages.
    
    ```bash
    go run ./cmd/main.go
    ```

3. **Input the number of jobs**:  
    After running the program, enter the number of jobs to divide the 1,000 rows equally. Each Goroutine will process a portion of the rows in parallel. The prompt will look like this:

    ```bash
    Enter Jobs Number: 
    ```


## References

- [Goroutines (Concurrency)ใน Golang มันใช้ยังไง?](https://medium.com/@rayato159/goroutines-concurrency-%E0%B9%83%E0%B8%99-golang-%E0%B8%A1%E0%B8%B1%E0%B8%99%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%A2%E0%B8%B1%E0%B8%87%E0%B9%84%E0%B8%87-7e45a4a85187)
- [Reading in Console Input in Golang](https://tutorialedge.net/golang/reading-console-input-golang/)