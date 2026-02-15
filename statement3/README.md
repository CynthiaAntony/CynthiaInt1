# Statement 3-String List manager with trie

Program takes string input and performs operations based on the input: List, Add, Remove, Check, Quit. The data structure used is a Trie (`map[rune]*Node`, `bool`) which stores the strings in a tree-like structure. Each node contains a character, a boolean flag indicating the end of a word, and pointers to children nodes.

---

## Prerequisites

- Go installed  
Check using:

```bash
go version
```

---

## How to Run

Navigate to folder:

```bash
cd statement3
```

Run program:

```bash
go run main.go
```

---

## Run Tests

```bash
go test -v
```

---

## Sample Input

```
-----String Manager-----

Enter:
Add
Remove
Check
List
Exit: add
Enter string to add: Hi
String added
Added: Hi


Enter:
Add
Remove
Check
List
Exit: Add
Enter string to add: High
String added
Added: High


Enter:
Add
Remove
Check
List
Exit: Check
Enter string to check: Hi
String exists

Enter:
Add
Remove
Check
List
Exit: List
Stored strings:
[Hi High]


Enter:
Add
Remove
Check
List
Exit: Remove
Enter string to remove: Hi
Removed: Hi


Enter:
Add
Remove
Check
List
Exit: List
Stored strings:
[High]


Enter:
Add
Remove
Check
List
Exit: Exit
Exiting...
```
