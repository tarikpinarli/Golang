# Go Basics & Syntax Learnings

## 1. Clean Line Endings
- **No Semicolons (;)**: Unlike C++, statements in Go do not end with a semicolon. The compiler inserts them automatically at the end of each line.

## 2. Variable Declaration (`:=`)
- **Short Declaration Operator**: The `:=` syntax is used to declare and initialize a variable in one step.
- **Type Inference**: You do not need to specify the type (e.g., `float64`). The compiler analyzes the value on the right side and assigns the correct type automatically.
  - *Example:* `name := "Tarik"` (String)
  - *Example:* `temp := 25.0` (Float64)
- **Constraint**: This operator can **only** be used inside functions.

## 3. Formatting Output (`%.2f`)
- **Printf Function**: `fmt.Printf` allows you to format strings with placeholders (verbs), unlike `fmt.Println` which just prints raw values.
- **Float Formatting**:
  - `%f`: The standard verb for floating-point numbers.
  - `%.2f`: The `.2` modifier restricts the output to exactly **two decimal places** (rounding if necessary).