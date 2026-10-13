# Student Worksheet — Go, Sequencing & Introduction to Genetic Algorithms

**International Class · 2026 · shiddiqazis**

> **A program that produces an answer is not always a program that understands the problem.**

**Mission:** For each exercise, **predict → implement → run → explain**. State inputs, processing steps, output, and any assumptions *before* writing code.

## Set up your workspace

1. Fork [this teaching repository](https://github.com/shiddiqeuy/students-worksheet-algorithm).
2. Work in your personal fork.
3. Enter an exercise directory, e.g. `cd student/01_data_types`.
4. Run `go run main.go`. Fill in each `TODO`, then run again.
5. Use `gofmt -w main.go` and make a small commit for each completed exercise.

### 01 — Campus Festival Registration: What is your data?

**Story.** Maya builds a digital check-in desk. Student age is a count, mass is a measurement, registration is a true/false status, and student ID may begin with zeros.

**Task.** Declare and print age `19` (`int`), mass `62.5` (`float64`), checked-in `true` (`bool`), and student ID `"00123"` (`string`).

**Expected:** `19 | 62.5 | true | 00123`. Keep the leading zeros.

**Explain:** Why is a student ID not a number even when it contains digits? What errors could happen if you treated it as an `int`?

**Reflection:** *Choosing a type is choosing which operations make sense.*

[Go starter](student/01_data_types/main.go)

### 02 — Campus Café: Does order matter?

**Story.** A student orders two coffees at Rp18,000 each and receives a Rp5,000 discount.

**Task.** Calculate the subtotal **before** applying the discount; display both values. Change the quantity to three and recompute.

**Expected with quantity 2:**

```text
Subtotal: 36000
Total: 31000
```

**Explain:** What should happen if the discount exceeds the subtotal? Which variable must be computed first?

**Reflection:** *The right calculation in the wrong order may still lead to the wrong decision.*

[Go starter](student/02_cafe_sequence/main.go)

### 03 — Equal Volunteer Teams: Can everyone fit?

**Story.** Maya wants to split `y` volunteers into groups of `x`, without leaving anyone out.

**Task.** Read two positive integers with `fmt.Scan`. Print whether `x` divides `y` using `y % x == 0`. Validate before dividing.

| Input | Expected |
|---|---|
| `5 20` | `true` |
| `20 5` | `false` |
| `123 123` | `true` |
| `0 20` | `invalid input` |

**Explain:** Why is checking `x > 0` *before* modulo essential?

[Go starter](student/03_factor_checker/main.go)

### 04 — Population Dashboard: A model is a simplification

**Story.** A city wants a population estimate after births, incoming residents, deaths and people leaving.

**Rule:** `final = initial + births + arrivals - deaths - departures`.

**Input data:** 1000, 20, 5, 10, 4. **Expected:** `Population: 1011`.

**Explain:** What does a negative result imply about the data or model? Which quantities cannot reasonably be negative?

**Reflection:** *A useful mathematical model is not the whole of reality; its assumptions matter.*

[Go starter](student/04_population/main.go)

### 05 — A Composer's Pattern: Number sequences

**Story.** A composer gives two notes represented by numbers. Each new number is the sum of the previous two.

**Task:** For two starting values, calculate the third, fourth, and fifth values, **first without a loop**.

| Starting values | Next three |
|---|---|
| `1 1` | `2 3 5` |
| `5 3` | `8 11 19` |
| `10 20` | `30 50 80` |

**Explain:** How is *generating a number sequence* different from *executing instructions sequentially*? Optional: implement the same rule with a `for` loop.

[Go starter](student/05_number_sequence/main.go)

### 06 — Gravity on Mars: Why units matter

**Story.** An astronaut's mass remains unchanged between planets, but gravitational force does not.

**Rules:** Earth weight `W = massKg × 9.8`; Mars weight `W = massKg × 9.8 × 0.38`.

**Input:** `70.0` kg. **Expected (one decimal place):**

```text
Earth: 686.0 N
Mars: 260.7 N
```

**Explain:** Why is the input measured in kg while the output is in newtons (N)? Why use `float64`?

[Go starter](student/06_gravity/main.go)

### 07 — Optional AI Preview: Tiny genetic algorithm operations

**Story.** A four-switch controller can be any four-bit string (such as `1001`). The objective is to reach `1111`.

**Vocabulary:** candidate (possible solution), fitness (number of `1` bits), selection (choosing candidates), crossover (mixing parts), mutation (changing a bit).

**Manual walkthrough:**

```text
Parent A: 1001  fitness = 2
Parent B: 0110  fitness = 2
Cut after 2 bits → child: 1010   fitness = 2
Change second bit 0 → 1: 1110    fitness = 3
```

**Task:** Implement `fitness(candidate string) int`, demonstrate the crossover and one mutation, then print the child's string and score **before and after** mutation.

**Explain:** Why is this a demonstration of GA *operations*, not a full genetic algorithm? Does mutation always improve fitness? (No.)

**Reflection:** *Evolution changes candidates; the objective function defines what counts as better.*

[Go starter](student/07_genetic_intro/main.go)

### 08 — Cube-root Check: A precise rule

**Story.** A game unlocks a door when a submitted integer `x` really is the cube root of `y`.

**Task:** Read `x` and `y`; verify `x*x*x == y` using multiplication, not floating-point cube-root approximation.

| Input | Expected |
|---|---|
| `5 125` | `true` |
| `-3 -27` | `true` |
| `12 12` | `false` |

**Explain:** Why can a negative number have a negative cube? What happens for very large integers?

[Go starter](student/08_cube_root/main.go)

### 09 — Advanced Bonus: Evolve a population

**Story.** Instead of changing only one bit string once, the controller starts with a *population* of candidates.

**Task:** Write a small population-based genetic search over four-bit strings. Calculate fitness, select parents, cross over, sometimes mutate, keep or regenerate candidates, and repeat. Print generation number and best fitness.

**Acceptance:** Include an explicit generation limit or other stopping condition. Explain why the algorithm **does not guarantee** finding `1111` within a chosen budget.

**Explain:** Why might random exploration help? How could it make a candidate worse?

[Go starter](student/09_bonus_ga_loop/main.go)

## Task 2 — Fill in the blank

Practice for [GitHub CLI intro](docs/github-cli-intro.md) and [Go basic types](docs/go-basic-types.md). Replace every `____` with the missing word, symbol or value. **Predict first, then check** each Go answer by running it. Write your answers in your fork's `EXPLANATION.md` under *Task 2 — Fill in the blank*. Answers are not published here.

### A — Git and GitHub (terminal)

**Word bank:** `add` · `clone` · `diff` · `fork` · `main` · `origin` · `push` · `status` · `switch`

1. A ____ is your own online copy of someone else's repository.
2. `git ____ <url>` downloads a repository to your computer.
3. `git ____ -c task-2` creates a new branch and moves you onto it.
4. `git ____` lists the files you changed, and `git ____` shows the exact lines you added or removed.
5. `git ____ student/10_task2_types/main.go` puts a file in the box before you seal it with `git commit`.
6. `git ____ -u origin task-2` sends your commits to your fork on GitHub.
7. `git log --oneline ____..task-2` lists only the commits you made on your branch.
8. The remote that points to **your** fork is normally called ____.

### B — Go types, constants and operators

9. `7 / 2` is ____ because both numbers are ____ values, so the fraction is dropped.
10. `7 % 2` is ____. The `%` operator gives the ____ of a division.
11. `float64(7) / float64(2)` is ____. Go does **not** convert types ____ (automatically).
12. Go has no `char` type. A single character is stored as a ____ (or as a ____ for one ASCII byte).
13. A student ID such as `"00123"` should be a ____ because we never do arithmetic on it and the leading ____ must stay.
14. The keyword for a value that can never change is ____. Reassigning it gives a ____ error.
15. `true && false` is ____, `true || false` is ____, and `!true` is ____.
16. In `x > 0 && y % x == 0`, when `x` is `0` the right side is ____ evaluated, so the program does not divide by zero.
17. In Go, `=` ____ a value and `==` ____ two values.
18. `2 + 3 * 4` is ____ but `(2 + 3) * 4` is ____.
19. With `x, y := 0.1, 0.2`, the expression `x + y == 0.3` is ____ because `float64` stores some decimals only ____.
20. `string(65)` gives ____, so use `strconv.____(65)` when you want the text `"65"`.

### C — Complete the program

```go
package main

import "fmt"

____ taxRate = 0.1

func main() {
	price := 20000.0
	qty := 3
	subtotal := price * ____(qty)
	total := subtotal + subtotal*taxRate
	isBig := total ____ 50000
	fmt.Println(total, isBig)
}
```

21. Fill the three blanks so the program compiles: ____ , ____ , ____ .
22. Before running, predict the output: ____ ____
23. Which line would stop the program from compiling if you removed the conversion blank? Explain in one sentence: ____

**Check:** run it with `go run main.go`, then `gofmt -l .` and `go vet main.go` should print nothing.

## Submission & mini-defense

Submit the link to **your fork**, not just a screenshot, through [the submission issue form](https://github.com/shiddiqeuy/students-worksheet-algorithm/issues/new/choose). Add `EXPLANATION.md` and terminal evidence for at least three labs in your fork. Do not post personal student IDs in this public repository.

Your two-minute defense should answer: **Which types did you choose? Why does the execution order matter? What normal and edge cases did you test? What did you learn from a wrong answer?**

| Criterion | Points |
|---|---:|
| Correct Go behavior | 4 |
| Data-type and logical justification | 2 |
| Normal and boundary-case testing | 2 |
| Code readability and mini-defense | 2 |

> **Do not just write code. Explain what the code means.**

*© 2026 shiddiqazis · International Class. The original five computational exercises remain the foundation; the Go scenarios and genetic-algorithm introduction are instructor-created enrichment.*
