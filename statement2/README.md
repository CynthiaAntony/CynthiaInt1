#Statement 2-string list manager

Program takes string input and performs operations based on the input. List, add remove, check, quit.
The data structure used is a map(string, int)

##Prerequisites
 - Go installed
 Check using: 
 ```bash
 go version
 ```

##How to run
1. Navigate to folder
```bash
 cd statement2
 ```

2. Run program 
```bash
 go run main.go
 ```

3 Run tests
```bash
 go test -v
 ```
 
##Sample input
```bash
 -------String list manager-------



Choose: 
List
Add
Remove
Check
Quit: add

Enter new string: hi
String added



Choose: 
List
Add
Remove
Check
Quit: add

Enter new string: hello
String added



Choose: 
List
Add
Remove
Check
Quit: list

Strings:
hi
hello



Choose: 
List
Add
Remove
Check
Quit: check

Enter string to check: hi
String is in the list



Choose: 
List
Add
Remove
Check
Quit: check

Enter string to check: Hallo
String not found



Choose: 
List
Add
Remove
Check
Quit: remove

Enter string to remove: hi
String Removed



Choose: 
List
Add
Remove
Check
Quit: list

Strings:
hello



Choose: 
List
Add
Remove
Check
Quit: quit

Exiting...
```