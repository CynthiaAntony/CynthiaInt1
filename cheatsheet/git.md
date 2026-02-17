# Git Commands Cheatsheet 

### 1. Intiialize a repository - Create a new repository in the current folder
```bash
git init
```
Creates a new Git repository inside the current directory and adds a hidden `.git` folder that tracks all version history. Used when starting a brand-new project.

---

### 2. Clone an existing repository
```bash
git clone <repo-url>
```
Downloads an existing remote repository to your local machine and copies all files, branches, and commit history.

---

### 3. Clone an existing repository in a folder
```bash
git clone <repo-url> <folder-name>
```
Clones the repository but places it inside a folder with a custom name instead of the default repository name.

---

### 4. Show commit history
```bash
git log
```
Displays the full commit history including commit hash, author, date, and message.

---

### 5. Show status of all the files (staged, unstaged, tracked)
```bash
git status
```
Shows the current state of the working directory.  
It tells:
- Which files are staged
- Which files are modified
- Which files are untracked

---

### 6. Show all the staged changes
```bash
git diff --staged
```
Displays differences between the staged files and the last commit. Useful before committing to verify what will be included.

---

### 7. Show all the unstaged changes
```bash
git diff
```
Shows changes made in files that are not yet staged.

---

### 8. Stage a specific file
```bash
git add <file-name>
```
Moves a specific file’s changes into the staging area in preparation for committing.

---

### 9. Stage the changes in the working directory
```bash
git add .
```
Stages all changes in the current directory (new, modified files).

---

### 10. Stage everything including the deletions
```bash
git add -A
```
Stages all changes across the entire repository, including deleted files.

---

### 11. Commit your changes
```bash
git commit -m "message"
```
Creates a new commit with the staged changes and attaches a message describing the changes.

---

### 12. Change a commit (message or content)
```bash
git commit --amend
```
Modifies the most recent commit. Can update the commit message or include newly staged changes.

---

### 13. List branches
```bash
git branch
```
Displays all local branches. The current branch is marked with `*`.

---

### 14. Create a new branch
```bash
git branch <branch-name>
```
Creates a new branch but does not switch to it.

---

### 15. Create a new branch and switch to it
```bash
git checkout -b <branch-name>
```
Creates a new branch and immediately switches to it.

---

### 16. Just switch to a branch
```bash
git checkout <branch-name>
```
Switches from the current branch to another existing branch.

---

### 17. Delete a branch - only when the branch is merged upstream
```bash
git branch -d <branch-name> (Substitute with -D for force delete, deletes irrespective of merge status)
```
Deletes a local branch.  
`-d` ensures it is safely merged before deletion.  
`-D` forces deletion even if not merged.

---

### 18. Get changes / Download changes without merge
```bash
git fetch
```
Downloads changes from the remote repository but does not merge them into your branch.

---

### 19. Fetch the changes + Merge
```bash
git pull (use --rebase for a cleaner history)
```
Fetches changes from the remote and automatically merges them into your current branch.  
`--rebase` replays your commits on top instead of creating a merge commit.

---

### 20. Pull a specific branch
```bash
git pull <branch-name>
```
Pulls updates from a specific branch from the remote repository.

---

### 21. Push current branch
```bash
git push
```
Uploads your local commits to the remote repository.

---

### 22. Set upstream and push, better tracking
```bash
git push -u origin <branch-name>
```
Pushes the branch and sets it to track the remote branch, allowing future simple `git push` or `git pull`.

---

### 23. List remotes
```bash
git remote -v
```
Shows remote repositories linked to your project along with their URLs.

---

### 24. Stash uncommitted changes (eg. moving to another branch without committing changes)
```bash
git stash (use -u if you want to add untracked files as well)
```
Temporarily saves uncommitted changes so you can switch branches without committing.

---

### 25. Apply and remove stash
```bash
git stash pop
```
Restores the most recent stash and removes it from the stash list.

---

### 26. Apply but not remove stash
```bash
git stash apply
```
Restores stashed changes but keeps the stash saved for reuse.

---

### 27. Reapply commits on top of another branch
```bash
git rebase <branch>
```
Moves your current branch’s commits on top of another branch to create a linear history.

---

### 28. Rebase interactive mode
```bash
git rebase -i HEAD~3 (It says you can perform rebase on last 3 commits)
```
Opens interactive mode to edit, squash, reorder, or remove the last 3 commits.

---

### 29. Merge branch with the current branch
```bash
git merge <branch>
```
Combines another branch into the current branch.

---

### 30. Undo commits (soft reset) but no unstaging
```bash
git reset --soft HEAD~1
```
Moves the branch pointer back one commit but keeps changes staged.

---

### 31. Undo commits and unstage changes (Basically remove the 'A' sign)
```bash
git reset HEAD~1 (1 is the latest ever commit) (No deletion of changes - use --hard for deletion)
```
Moves the branch pointer back one commit and unstages changes.  
Changes remain in your working directory.

---

### 33. Unstage a file (opposite of git add <file>)
```bash
git restore --staged <file-name>
```
Removes a file from the staging area without deleting changes.

---

### 34. Discard changes made to a file 
```bash
git restore <file-name>
```
Reverts the file back to its last committed state.

---

### 35. Inspect a commit 
```bash
git show <commit>
```
Displays detailed information about a specific commit, including changes made.

---

### 36. See who changed what in a specific file
```bash
git blame <file-name>
```
Shows line-by-line authorship for a file, helping identify who made specific changes.

