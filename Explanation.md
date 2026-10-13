# EXPLANATION — My Algorithm Worksheet

> Copy this file to `EXPLANATION.md` in your own fork. Write in English. Do not put private student details in a public repository.

**GitHub username:** MahenProject22  
**Exercises completed:** 01–09 and Task 2 Go basic types (`10_task2_types`)  
**Most surprising bug:** The first terminal command used an incomplete folder path, so PowerShell could not find the exercise directory. Including the repeated `students-worksheet-algorithm-main` folder fixed the path; the Go programs then ran.

## Exercise 01 — Campus Festival Registration
- **Story in my words:** Print a student's age, mass, check-in status, and ID.
- **Input(s) and their types:** No runtime input; `age` is `int`, `mass` is `float64`, `checkedIn` is `bool`, and `studentID` is `string`.
- **Processing / logical rule:** Store each value using the type appropriate for its meaning, then print them separated by `|`.
- **Expected output (before coding):** `19 | 62.5 | true | 00123`
- **Actual output:** `19 | 62.5 | true | 00123`
- **Normal test:** Ran the program with the worksheet values; output matched.
- **Boundary or invalid test:** The ID `"00123"` retains its leading zeros because it is text.
- **One assumption or limitation:** Values are fixed in the source code; the program does not ask for user input.
- **What I changed in `main.go`:** Filled in the four values from the worksheet.
- **Why I believe the result is correct:** The output matches the expected values and preserves the ID's leading zeros.

## Exercise 02 — Campus Café: Does order matter?
- **Story in my words:** Calculate the café order subtotal and subtract its discount.
- **Input(s) and their types:** Fixed integer values: unit price `18000`, quantity `3`, and discount `5000`.
- **Processing / logical rule:** Calculate `subtotal = unitPrice * quantity` first, then `total = subtotal - discount`.
- **Expected output (before coding):** With quantity 3: subtotal `54000`, total `49000`. The worksheet's original quantity-2 example is subtotal `36000`, total `31000`.
- **Actual output:** `Subtotal: 54000`; `Total: 49000`.
- **Normal test:** Ran the current quantity-3 version; output matched the calculation.
- **Boundary or invalid test:** A discount greater than the subtotal would produce a negative total; the program does not reject or cap that result.
- **One assumption or limitation:** The program assumes the price, quantity, and discount are valid non-negative values.
- **What I changed in `main.go`:** Set the quantity to 3, calculated the subtotal before the discount, and printed both amounts.
- **Why I believe the result is correct:** `18000 * 3 = 54000`, and `54000 - 5000 = 49000`.

## Exercise 03 — Equal Volunteer Teams: Can everyone fit?
- **Story in my words:** Check whether `y` volunteers can be divided into equal groups of `x`.
- **Input(s) and their types:** Two integers, `x` and `y`, read from standard input.
- **Processing / logical rule:** Reject non-positive values first; otherwise, `x` divides `y` when `y % x == 0`.
- **Expected output (before coding):** `5 20` → `true`; `20 5` → `false`; `0 20` → `invalid input`.
- **Actual output:** `5 20` → `true`; `20 5` → `false`; `0 20` → `invalid input`.
- **Normal test:** Tested `5 20` and `20 5`; both results matched.
- **Boundary or invalid test:** Tested `0 20`; it printed `invalid input` without evaluating the modulo.
- **One assumption or limitation:** The program assumes both integers can be successfully read; it does not check `fmt.Scan` errors.
- **What I changed in `main.go`:** Read two integers, added the positive-value guard, and printed the divisibility result.
- **Why I believe the result is correct:** The remainder is zero for `20 % 5`, but not for `5 % 20`; zero is rejected before modulo.

## Exercise 04 — Population Dashboard
- **Story in my words:** Estimate a city's final population after population increases and decreases.
- **Input(s) and their types:** Fixed integer counts: start `1000`, births `20`, arrivals `5`, deaths `10`, departures `4`.
- **Processing / logical rule:** Add births and arrivals to the starting population, then subtract deaths and departures.
- **Expected output (before coding):** `Population: 1011`
- **Actual output:** `Population: 1011`
- **Normal test:** Ran the worksheet's data; output matched.
- **Boundary or invalid test:** The program has no validation for negative counts or a result below zero.
- **One assumption or limitation:** All counts are assumed to be valid, non-negative whole numbers.
- **What I changed in `main.go`:** Applied the population formula and printed the result.
- **Why I believe the result is correct:** `1000 + 20 + 5 - 10 - 4 = 1011`.

## Exercise 05 — A Composer's Pattern: Number sequences
- **Story in my words:** Generate the next three values by adding the previous two values each time.
- **Input(s) and their types:** Fixed integer starting values `10` and `20`.
- **Processing / logical rule:** `third = first + second`, `fourth = second + third`, and `fifth = third + fourth`.
- **Expected output (before coding):** `30 50 80`
- **Actual output:** `30 50 80`
- **Normal test:** Ran the program with the worksheet's `10, 20` starting values; output matched.
- **Boundary or invalid test:** No input validation is needed for the fixed values; sufficiently large integers could overflow.
- **One assumption or limitation:** The starting values are hard-coded and only three new values are generated.
- **What I changed in `main.go`:** Added the three sequential additions and printed their results.
- **Why I believe the result is correct:** Each new number is the sum of the two immediately preceding numbers.

## Exercise 06 — Gravity on Mars
- **Story in my words:** Calculate the weight of a 70 kg mass on Earth and Mars.
- **Input(s) and their types:** Fixed `float64` mass `70.0`, Earth gravity `9.8`, and Mars factor `0.38`.
- **Processing / logical rule:** Earth weight is mass times gravity; Mars weight is Earth weight times the Mars factor.
- **Expected output (before coding):** `Earth: 686.0 N` and `Mars: 260.7 N`
- **Actual output:** `Earth: 686.0 N`; `Mars: 260.7 N`.
- **Normal test:** Ran the worksheet's mass value; both formatted results matched.
- **Boundary or invalid test:** The program does not validate negative mass or unusual gravity values.
- **One assumption or limitation:** The calculation assumes the given gravity and Mars factor are correct and uses one-decimal-place display.
- **What I changed in `main.go`:** Calculated Earth weight first, then scaled it for Mars.
- **Why I believe the result is correct:** `70 * 9.8 = 686`; `686 * 0.38 = 260.68`, which rounds to `260.7`.

## Exercise 07 — Optional AI Preview: Tiny genetic algorithm operations
- **Story in my words:** Demonstrate how two four-bit candidates can be crossed over and then mutated.
- **Input(s) and their types:** Fixed strings `a = "1001"` and `b = "0110"`.
- **Processing / logical rule:** Fitness counts the `'1'` bits; crossover joins the first two bits of `a` with the last two of `b`; mutation changes the second child bit.
- **Expected output (before coding):** Child `1010` has fitness 2; mutated child `1110` has fitness 3.
- **Actual output:** Parent A `1001`; Parent B `0110`; Child `1010`, fitness `2`; Mutated `1110`, fitness `3`.
- **Normal test:** Ran the fixed example; the child and fitness scores matched the manual walkthrough.
- **Boundary or invalid test:** A string with no ones has fitness 0; a string with four ones has fitness 4.
- **One assumption or limitation:** The code assumes the candidates have at least two bits and demonstrates only one fixed crossover and mutation.
- **What I changed in `main.go`:** Implemented the fitness count, crossover, mutation, and output.
- **Why I believe the result is correct:** Counting the bits gives two ones in `1010` and three in `1110`.

## Exercise 08 — Cube-root Check
- **Story in my words:** Check whether the first integer is exactly the cube root of the second.
- **Input(s) and their types:** Two integers, `x` and `y`, read from standard input.
- **Processing / logical rule:** Compare `x*x*x` with `y` using integer arithmetic.
- **Expected output (before coding):** `5 125` → `true`; `-3 -27` → `true`; `12 12` → `false`.
- **Actual output:** `5 125` → `true`; `-3 -27` → `true`; `12 12` → `false`.
- **Normal test:** Tested `5 125`; it printed `true`.
- **Boundary or invalid test:** Tested the negative cube `-3 -27` and the non-cube `12 12`; outputs were `true` and `false`.
- **One assumption or limitation:** Very large `x` values can overflow when cubed, and scan errors are not checked.
- **What I changed in `main.go`:** Read both integers and compared `x*x*x` with `y`.
- **Why I believe the result is correct:** A cube preserves the sign of a negative integer, and the equality check uses exact integer arithmetic for these sample values.

## Exercise 09 — Advanced Bonus: Evolve a population
- **Story in my words:** Repeatedly evolve a group of four-bit candidates toward the all-ones target `1111`.
- **Input(s) and their types:** No runtime input; a population of six strings, four bits per string, at most ten generations, and mutation rate `0.1`.
- **Processing / logical rule:** Score each candidate by its number of ones, select fitter parents, combine their bits, sometimes mutate a bit, and stop on fitness 4 or the generation limit.
- **Expected output (before coding):** Up to ten generation summaries; the search may find `1111`, but it is not guaranteed.
- **Actual output:** One run found `1111` at generation 7, printed `Optimal solution found!`, and stopped. The candidate sequence can differ on another run.
- **Normal test:** Ran the program; it printed generation summaries and stopped when it found the target.
- **Boundary or invalid test:** If the target is not found, the generation limit stops the search after generation 10.
- **One assumption or limitation:** The search is stochastic; reaching the optimal candidate within ten generations is not guaranteed.
- **What I changed in `main.go`:** Added candidate generation, fitness scoring, parent selection, crossover, mutation, and a bounded generation loop.
- **Why I believe the result is correct:** The reported best fitness is the count of ones, and the loop stops at fitness 4 or the stated limit.

## Task 2 — GitHub CLI and Go basic types
- **Fork URL and branch:** `https://github.com/MahenProject22/students-worksheet-algorithm`, checked-out branch `main`.
- **Git evidence (`git remote -v`, `git log --oneline main..task-2`, at least 3 commits):** `origin` points to the fork URL above. At the time I checked, only local branch `main` existed; `task-2` was missing, so `git log --oneline main..task-2` reported an unknown revision. I still need to create/push `task-2` and provide the required three-commit evidence.
- **Program output (`go run main.go`):**
  ```text
  int: 7 / 2 = 3, 7 % 2 = 1
  float64: 3.5
  bool: passed = false
  string: Hello, Go (len 9)
  rune: A -> 65 -> B
  const: area = 78.5
  ```
- **Normal test (inputs and result):** This program has fixed values rather than runtime input. Running it with `a = 7`, `b = 2`, `score = 72`, and `attendance = 70` produced the output above; `passed` is false because attendance is below 75.
- **Boundary test (inputs and result):** With the current fixed values, integer division of `7 / 2` truncates to `3` and leaves remainder `1`. The program does not accept values interactively.
- **Invalid case I considered (e.g. dividing by zero) and what happens:** The current values do not divide by zero. If `b` were changed to zero, integer division by that variable would panic at runtime; this program has no guard for that case.
- **Q1. Why is `7 / 2` equal to `3` but `float64(7) / float64(2)` equal to `3.5`?** The first division uses integers, so Go truncates the fractional part. Converting both operands to `float64` makes the division preserve the fractional result.
- **Q2. Why is a student ID better stored as a `string` than an `int`?** An ID identifies a student rather than representing a quantity. A string preserves leading zeros, and arithmetic is not needed.
- **Q3. What does a `const` protect you from? Give one example from this exercise.** A constant cannot be reassigned, which protects fixed values from accidental changes. Here, `const pi = 3.14` keeps the circle calculation's value fixed.
- **Q4. In `x > 0 && y % x == 0`, why is the order of the two conditions important?** Go evaluates `&&` from left to right and short-circuits. If `x` is zero, `x > 0` is false, so Go does not evaluate `y % x`, avoiding division by zero.
- **Q5. What would `fmt.Println('A')` print, and why is it not `A`?** It prints `65`, because `'A'` is an untyped rune/integer character value and `fmt.Println` formats it numerically. Use `fmt.Printf("%c", 'A')` to print the character.

## Task 2 — Fill in the blank
- **Part A (questions 1-8):** 1. `fork`; 2. `clone`; 3. `switch`; 4. `status`, `diff`; 5. `add`; 6. `push`; 7. `main`; 8. `origin`.
- **Part B (questions 9-20):** 9. `3`, `int`; 10. `1`, remainder; 11. `3.5`, automatically; 12. `rune`, `byte`; 13. `string`, zeros; 14. `const`, compile-time; 15. `false`, `true`, `false`; 16. not; 17. assigns, compares; 18. `14`, `20`; 19. `false`, approximately; 20. `"A"`, `Itoa`.
- **Part C (questions 21-23):** 21. `const`, `float64`, `>`; 22. `66000 true`; 23. `subtotal := price * qty` would not compile because Go does not allow multiplying a `float64` by an `int` without converting one operand.

**Terminal evidence (text):** The outputs for Exercises 01, 03, and Task 2 are recorded above. Exercise 09's run reached fitness 4 at generation 7; its output may vary because the search is random. Add screenshots from your own terminal if your instructor requires screenshots.
