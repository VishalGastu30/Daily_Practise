# Daily Golang Practice Plan: Day 1
**Target Goal:** Campus Interview Preparation & Practical Software Engineering Skills

---

## Program 1: Logical Thinking & DSA (Optimization Focus)

### Problem Statement: Maximum Sum Subarray of Size K under Threshold
Given an array of positive integers `nums`, an integer `K` (window size), and a maximum threshold limit `M`, write a Go function to find the maximum contiguous subarray sum of size `K` that is **less than or equal to** `M`. If no valid subarray exists, return `0`.

### Technical Constraints & Objectives
- **Time Complexity Goal:** $O(N)$ amortized time.
- **Space Complexity Goal:** $O(1)$ auxiliary space.
- **Key Concepts:** Sliding Window algorithm, Boundary Checks, Efficient Slice Iteration.

### Input/Output Example
```text
Input:  nums = [2, 1, 5, 1, 3, 2], K = 3, M = 8
Output: 7  (Subarrays of size 3: [2,1,5]=8, [1,5,1]=7, [5,1,3]=9, [1,3,2]=6. 
            The maximum sum <= 8 is 7 from [1,5,1].)
```

### Guidance for Optimization
- Do not recompute the sum of the window from scratch for every index. Compute the initial window sum of $K$ elements, then slide the window by adding the incoming element and subtracting the outgoing element.

---

## Program 2: Go Language Mastery (Concurrency & Idioms)

### Problem Statement: Concurrent URL Status Checker
Design and implement a concurrent URL status checker in Go using a **Worker Pool** pattern. Given a list of target URLs, distribute the task of fetching their HTTP status codes across a fixed number of goroutines (e.g., $W = 4$ workers).

### Technical Constraints & Objectives
- **Key Concepts:** Goroutines, Channels, `sync.WaitGroup`, `context.WithTimeout`, Data Race Prevention.
- Use a **Worker Pool** design to prevent spawning unbounded goroutines.
- Implement a timeout per HTTP request using `context.Context` (e.g., 2 seconds max per request).
- Ensure safe concurrent data collection into a output slice or map.
- Test the application using `go run -race main.go` to ensure no race conditions exist.

### Input/Output Example
```text
Input URLs: ["https://golang.org", "https://github.com", "https://invalid.url"]
Worker Count: 2

Output:
[SUCCESS] https://golang.org - Status: 200 OK (Execution Time: 120ms)
[SUCCESS] https://github.com - Status: 200 OK (Execution Time: 150ms)
[ERROR]   https://invalid.url - Error: Get "https://invalid.url": dial tcp: lookup invalid.url: no such host
```

---

## Program 3: Real-World Backend System Problem

### Problem Statement: In-Memory Token Bucket Rate Limiter
Implement a thread-safe **Token Bucket Rate Limiter** struct in Go. This pattern is widely used in API gateways, microservices, and network routers to prevent system overload.

### Requirements & Design Specifications
1. **Struct Parameters:**
   - `capacity` (int): Maximum number of tokens the bucket can hold.
   - `refillRate` (float64): Number of tokens added per second.
   - `tokens` (float64): Current available tokens.
   - `lastRefill` (`time.Time`): Timestamp of the last refill operation.
   - `mu` (`sync.Mutex`): Mutex to guard shared state.

2. **Core Method:**
   - `Allow() bool`: Checks if a request can be processed. If tokens are available ($\ge 1.0$), decrement the token count by $1$ and return `true`. Otherwise, return `false`.

3. **Concurrency & Logic Rules:**
   - Lazy refill: Instead of running a continuous background ticker to add tokens, dynamically calculate the elapsed time since `lastRefill` inside `Allow()` and add the appropriate fraction of tokens before evaluating.
   - Must handle multiple concurrent calls to `Allow()` safely.

### Input/Output Simulation Example
```text
Capacity = 3, RefillRate = 1.0 token/sec

Time T=0.0s -> Allow() -> true  (Tokens left: 2.0)
Time T=0.0s -> Allow() -> true  (Tokens left: 1.0)
Time T=0.0s -> Allow() -> true  (Tokens left: 0.0)
Time T=0.0s -> Allow() -> false (Bucket empty)
Time T=1.5s -> Allow() -> true  (Refilled 1.5 tokens, consumed 1.0, left: 0.5)
```

---

## Daily Practice Guidelines
1. **Time Allocation:** Spend ~30 minutes per problem.
2. **Coding Standards:** Follow `gofmt` code formatting and idiomatic naming conventions.
3. **Verification:** Write basic standard library unit tests (`*_test.go`) for edge cases (e.g., empty slices, negative limits, concurrency edge cases).