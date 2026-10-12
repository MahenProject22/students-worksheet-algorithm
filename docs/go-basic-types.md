# Go Basics — Data Types, Constants and Operators

**Task 2 · Part B · International Class · 2026**

> **Choosing a type is choosing which operations make sense.**

Every value in a program has a **type**: a rule about what the value is and what you may do with it. This guide covers the types you will use all semester: `int`, `float64` (real numbers), `bool`, `string` and `rune`/`byte` (characters), plus **constants** and **operators**.

Go is *statically typed*: the compiler checks types before the program runs. This feels strict at first, but it catches many mistakes early.

If your pseudocode course uses the words `integer`, `real`, `boolean`, `string`, `character`, here is how they map to Go:

| Pseudocode | Go type | Example value |
|---|---|---|
| integer | `int` | `19`, `-3`, `0` |
| real | `float64` | `62.5`, `3.14`, `-0.001` |
| boolean | `bool` | `true`, `false` |
| string | `string` | `"Hello"`, `"00123"` |
| character | `rune` (any Unicode character) or `byte` (one ASCII byte) | `'A'`, `'7'`, `'é'` |

> Go has no type literally called `real` or `char`. Use `float64` and `rune`.

---

## 1. Declaring variables

```go
var age int = 19          // full form
var mass float64 = 62.5   // type given explicitly
checkedIn := true         // short form: Go infers bool
studentID := "00123"      // inferred string
```

- `var name type = value` can be used anywhere. `:=` works **only inside functions**.
- `:=` declares a *new* variable; `=` assigns to an existing one.
- Go refuses to compile if a declared variable is never used. That is a feature, not a bug.

**Zero values.** A variable declared without a value is not garbage; it gets a default:

| Type | Zero value |
|---|---|
| `int` | `0` |
| `float64` | `0` |
| `bool` | `false` |
| `string` | `""` (empty) |

```go
var count int
var ok bool
fmt.Println(count, ok) // 0 false
```

**Predict:** what does `var s string; fmt.Println(len(s))` print?

---

## 2. `int` — whole numbers

Use for counts, positions, ages, quantities.

```go
a, b := 7, 2
fmt.Println(a + b) // 9
fmt.Println(a / b) // 3   <- integer division drops the fraction
fmt.Println(a % b) // 1   <- remainder (modulo)
```

Key points:

- `/` on two `int`s **truncates** toward zero: `7 / 2` is `3`, and `-7 / 2` is `-3`.
- `%` gives the remainder and works only on integers. `y % x == 0` means *x divides y*.
- Dividing an integer by zero crashes the program at run time. Check first.
- An `int` has limits (typically 64-bit on modern computers). Very large results can overflow silently.

---

## 3. `float64` — real numbers

Use for measurements: mass, price, temperature, averages.

```go
mass := 62.5
price := 18000.0
fmt.Println(mass * 2)        // 125
fmt.Println(7.0 / 2.0)       // 3.5
fmt.Printf("%.2f\n", mass)   // 62.50  (two decimals)
```

Key points:

- Floating-point numbers are stored in binary, so some decimals are approximate. with variables, `x, y := 0.1, 0.2` makes `x + y` print `0.30000000000000004` and `x + y == 0.3` is `false`. (Typing `0.1 + 0.2` directly prints `0.3` because Go computes constants exactly.) Never test two `float64` values for exact equality; compare their difference to a small tolerance instead.
- `%` does not work on floats.
- Use `float64` unless you have a specific reason to use `float32`.

---

## 4. `bool` — true or false

Use for yes/no decisions: *is checked in?*, *is valid?*, *is divisible?*

```go
passed := true
open := false
fmt.Println(passed && open) // false
```

Booleans usually come from **comparisons** (section 8) and are combined with **logical operators**.

---

## 5. `string` — text

```go
greeting := "Hello" + ", " + "Go"
fmt.Println(greeting)       // Hello, Go
fmt.Println(len(greeting))  // 9
fmt.Println(greeting[0])    // 72  <- a byte, NOT the letter!
```

Key points:

- Strings use double quotes. Single quotes are for single characters (`rune`).
- `+` joins two strings. `"5" + "3"` is `"53"`, not `8`.
- Strings are **immutable**: you cannot change one character in place; build a new string.
- `len` counts **bytes**, not visible characters. `len("é")` is `2`.
- A student ID such as `"00123"` is a `string`: you never add two IDs, and leading zeros matter.

---

## 6. `rune` and `byte` — characters

```go
letter := 'A'                  // type rune (an int32 number: the Unicode code point)
fmt.Println(letter)            // 65
fmt.Printf("%c\n", letter)     // A
next := letter + 1
fmt.Printf("%c\n", next)       // B
```

- A `rune` is just a number that stands for a character: `'A'` is `65`, `'a'` is `97`, `'0'` is `48`.
- `byte` is a smaller number (0–255) used for ASCII data.
- Use `%c` to print as a character, `%d` to print as a number.
- To loop over characters safely, use `for _, r := range text`.

**Predict:** what does `'a' - 'A'` evaluate to? Why is that useful?

---

## 7. Constants — values that never change

```go
const pi = 3.14
const maxSeats = 30
const school = "Campus A"

const (
    small  = 1
    medium = 2
    large  = 3
)
```

- A `const` is fixed at compile time. Trying `pi = 3.15` produces a compile error, which protects you from accidental changes.
- Constants replace *magic numbers*: write `price * taxRate` instead of `price * 0.11`, so the meaning is readable and can be changed in one place.
- Constants can be **untyped**: `pi` above can be used wherever a `float64` is expected.
- Naming: choose clear names, e.g. `taxRate`, not `x`.

---

## 8. Operators

### Arithmetic

| Operator | Meaning | Works on | Example → result |
|---|---|---|---|
| `+` | add / join strings | numbers, strings | `3 + 4` → `7` |
| `-` | subtract | numbers | `3 - 4` → `-1` |
| `*` | multiply | numbers | `3 * 4` → `12` |
| `/` | divide | numbers | `7 / 2` → `3`; `7.0 / 2.0` → `3.5` |
| `%` | remainder | integers only | `7 % 2` → `1` |
| `++` `--` | add/subtract 1 (a statement) | numbers | `count++` |

### Comparison (result is `bool`)

`==` equal · `!=` not equal · `<` less · `<=` less or equal · `>` greater · `>=` greater or equal

```go
fmt.Println(5 == 5)       // true
fmt.Println("a" < "b")    // true  (alphabetical order)
```

Remember: `=` assigns, `==` compares.

### Logical (on `bool`)

| Operator | Meaning | Truth rule |
|---|---|---|
| `&&` | and | true only if **both** sides are true |
| `\|\|` | or | true if **at least one** side is true |
| `!` | not | flips true/false |

```go
score, attendance := 72, 70
passed := score >= 60 && attendance >= 75
fmt.Println(passed) // false
```

`&&` and `||` are *short-circuit*: if the left side already decides the answer, the right side is not evaluated. That is why `x > 0 && y % x == 0` is safe from division by zero.

### Assignment

`=` `+=` `-=` `*=` `/=` `%=`

```go
total := 10
total += 5   // same as total = total + 5 → 15
```

### Precedence (who goes first?)

`*` `/` `%` run before `+` `-`; comparisons run before `&&`; `&&` runs before `||`. Use parentheses when unsure: `(a + b) * c`.

---

## 9. Converting types

Go **never converts types for you** when numbers meet.

```go
a, b := 7, 2
// ratio := a / b       // 3 (integer division)
// bad := a * 1.5       // compile error: 1.5 (untyped float constant) truncated to int
// x := 1.5; bad := a * x // compile error: mismatched types int and float64
ratio := float64(a) / float64(b) // 3.5
f := 3.9
whole := int(f)                  // 3  (truncated, NOT rounded)
```

Converting between numbers and strings uses the `strconv` package:

```go
import "strconv"

n, err := strconv.Atoi("42")        // string -> int
if err != nil {
    fmt.Println("not a number")
}
text := strconv.Itoa(42)            // int -> string "42"
```

A common mistake: `string(65)` gives `"A"` (a character), **not** `"65"`. Use `strconv.Itoa(65)` for the digits.

---

## 10. Printing with `fmt`

| Verb | Prints | Example |
|---|---|---|
| `%d` | integer | `fmt.Printf("%d\n", 19)` → `19` |
| `%.2f` | float, 2 decimals | `62.50` |
| `%t` | bool | `true` |
| `%s` | string | `Hello` |
| `%q` | string with quotes | `"Hello"` |
| `%c` | character from rune | `A` |
| `%v` | any value, default format | |
| `%T` | the **type** | `float64` |

`\n` ends the line. A quick debugging trick: `fmt.Printf("%v %T\n", x, x)` shows the value and its type together.

---

## 11. Your Task 2 exercise

Open `student/10_task2_types/main.go`. Replace each placeholder marked `TODO`, run, and compare with the expected output.

**Predict first, in your notebook:** write the output you expect for each line *before* running your code.

```bash
cd student/10_task2_types
gofmt -w main.go
go vet main.go
go run main.go
```

**Expected output when finished:**

```text
int: 7 / 2 = 3, 7 % 2 = 1
float64: 3.5
bool: passed = false
string: Hello, Go (len 9)
rune: A -> 65 -> B
const: area = 78.5
```

| # | Concept | What you do |
|---|---|---|
| 1 | `int` | Compute quotient and remainder of `7` and `2`. |
| 2 | `float64` | Convert before dividing to get `3.5`. |
| 3 | `bool` | Combine two conditions with `&&`. |
| 4 | `string` | Join pieces with `+` and read `len`. |
| 5 | `const` | Try to reassign `pi`; read the compiler error; remove the line. |
| 6 | `rune` | Move `'A'` one letter forward. |
| 7 | `const` + `float64` | Compute the area of a circle of radius `5`. |

### Extension (required for the explanation)

Change the inputs and test:

- a **normal** case (e.g. `score = 90`, `attendance = 80`) → expect `passed = true`;
- an **edge** case (e.g. `score = 60`, `attendance = 75`, exactly on the boundary) → what do you expect, and why?
- an **invalid** case: what happens with `a / 0` when `b = 0`? Do not leave it in your final code; describe it in your explanation.

### Explain (write in `EXPLANATION.md`)

1. Why is `7 / 2` equal to `3` but `float64(7) / float64(2)` equal to `3.5`?
2. Why is a student ID better stored as a `string` than an `int`?
3. What does a `const` protect you from? Give one example from this exercise.
4. In `x > 0 && y % x == 0`, why is the order of the two conditions important?
5. What would be the output of `fmt.Println('A')`, and why is it not `A`?

---

## 12. Submit

Commit your work, push to your fork, and submit through the **Task 2** issue form. The full workflow is in [github-cli-intro.md](github-cli-intro.md).

**Rubric (10 pts):** program output 3 · Git workflow evidence 2 · tests (normal, edge, invalid) 2 · explanation of types and operators 3.

Back: **[← GitHub CLI intro](github-cli-intro.md)**
