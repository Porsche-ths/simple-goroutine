# Simple Goroutines

Welcome to **Simple Goroutines**, a personal project for learning and experimenting with Goroutines in Go. Feel free to explore and enjoy reading through it!

## Description

### Overview

Goroutines are lightweight threads managed by the Go runtime, designed to optimize business logic through concurrent programming. This repository demonstrates how Goroutines can be used to optimize the processing of a .csv file containing 1,000 random float values by computing their average.

### Calculation

- **Brute Force**
    Iterates through all rows in the file sequentially to calculate the average value.

- **Concurrent Algorithm (4 Modules)**
    Splits the 50M rows into `n` equal parts, calculates the sum of each part concurrently, and then combines the partial sums to compute the final average.

### Modules

This project features four modules, each using a different approach to calculate the average, enabling a comparison of execution times:

1. **Simple Module**:  
    Computes the average sequentially without using Goroutines.
2. **Wait Group Module**:  
    Implements Goroutines using the WaitGroup mechanism to synchronize concurrent tasks.
3. **Channel Module**:  
    Utilizes Go channels to implement Goroutines and manage concurrent communication.
4. **Worker Pool Module**:  
    Uses a worker pool pattern to process rows concurrently with Goroutines.

## Getting Started

Follow the steps below to run the project:

1. **Generate the `data.csv` file**:  
    Run the Python script to create a `.csv` file with 50M random float values.

    ```bash
    python "./fileGenerator.py"
    ```

2. **Run the Go program**:  
    Execute the main program to calculate averages.
    
    ```bash
    go run ./cmd/main.go
    ```

3. **Input the number of jobs**:  
    After running the program, enter the number of jobs to divide the 50M rows equally. Each Goroutine will process a portion of the rows in parallel. The prompt will look like this:

    ```bash
    Enter Jobs Number: 
    ```


## References

- [Goroutines (Concurrency)ใน Golang มันใช้ยังไง?](https://medium.com/@rayato159/goroutines-concurrency-%E0%B9%83%E0%B8%99-golang-%E0%B8%A1%E0%B8%B1%E0%B8%99%E0%B9%83%E0%B8%8A%E0%B9%89%E0%B8%A2%E0%B8%B1%E0%B8%87%E0%B9%84%E0%B8%87-7e45a4a85187)
- [Reading in Console Input in Golang](https://tutorialedge.net/golang/reading-console-input-golang/)