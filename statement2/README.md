#Statement 2-string list manager

Program takes string input and performs operations based on the input. List, add remove, check, quit.
The data structure used is a map(string, int) which stores the string and the count

##Prerequisites
 - Go installed
 Check using: go version

##How to run
1. Navigate to folder
 cd statement2

2. Run program 
 go run main.go

3 Run tests
 go test -v
##Sample input
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
Quit: add

Enter new string: hi
String added



Choose: 
List
Add
Remove
Check
Quit: list

String: Count
hi: 2
hello: 1



Choose: 
List
Add
Remove
Check
Quit: check

Enter string to check: hello
String 'hello' is in the list



Choose: 
List
Add
Remove
Check
Quit: remove

Enter string to remove: hi
String removed



Choose: 
List
Add
Remove
Check
Quit: list

String: Count
hello: 1
hi: 1



Choose: 
List
Add
Remove
Check
Quit: quit

Exiting...