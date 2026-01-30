# Arrays vs. Slices (Memory & Architecture)

## 1. Arrays (`[N]T`)
- **Fixed Size**: The size is part of the type. `[2]int` and `[3]int` are completely different types and cannot be mixed.
- **Value Semantics**: Unlike C/C++, arrays in Go are **values**, not pointers. If you pass an array to a function, Go **copies the entire array** (memory intensive).
- **Stack Allocation**: Usually allocated on the stack. fast, but inflexible.
- **Syntax**: `[2]string{"A", "B"}` (Size is explicit).

## 2. Slices (`[]T`)
- **Dynamic View**: A slice is a lightweight "window" or view into an underlying array. This is what is used 99% of the time in Go.
- **Internal Structure**: A slice is actually a small struct with 3 fields:
  1. **Pointer**: Address of the first element in the underlying array.
  2. **Length (`len`)**: The number of elements currently in the slice.
  3. **Capacity (`cap`)**: The total size of the underlying array before it needs to resize.
- **Syntax**: `[]string{"A", "B"}` (Brackets are empty).

## 3. The `append` Function
- **Growth Logic**: You cannot simply write `list[3] = "new"` if the slice only has 3 items. You must use `append`.
- **Memory Reallocation**:
  1. If `len < cap`, `append` just places the value in the next open slot.
  2. If `len == cap`, Go **allocates a new, larger array** (usually double size), copies all existing data to it, adds the new item, and updates the slice pointer.
- **Usage**: `list = append(list, "Item")` (You must reassign the result back to the variable because the memory address might have changed).