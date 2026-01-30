# Input Handling & System Robustness

## 1. Multiple Return Values
- **The Concept**: Unlike C/Java, Go functions can return multiple values at once.
- **Example**: `fmt.Scan` returns two values:
  1. `n` (int): The number of items successfully scanned.
  2. `err` (error): A report if something went wrong (or `nil` if successful).

## 2. The Blank Identifier (`_`)
- **The Problem**: Go considers unused variables a compilation error.
- **The Solution**: The underscore `_` acts as a "trash can." It tells the compiler to discard a return value we don't need.
- **Usage**: `_, err := fmt.Scan(&var)` means "Discard the count (`n`), but keep the error (`err`)."

## 3. Error Handling
- **Errors are Values**: Go does not use `try/catch`. Errors are just values checked with `if` statements.
- **The Check**: `if err != nil` is the standard way to check for failure.
  - `nil` means "No error" (Success).
  - Not `nil` means "Something broke."

## 4. Pointers in Scanning
- **Why `&`?**: `fmt.Scan(&variable)` requires a pointer (memory address).
- **Logic**: The function needs to know *where* in memory your variable lives so it can go there and update the value. Passing just `variable` would send a copy, and the original would remain empty.

## 5. System Interaction
- **Time**: Never hardcode dates. Use `time.Now().Year()` to ensure code longevity.
- **Exit Codes**:
  - `os.Exit(0)`: The program finished successfully.
  - `os.Exit(1)`: The program crashed or failed. Orchestration tools (like Docker/Kubernetes) use this code to decide if they should restart the container.