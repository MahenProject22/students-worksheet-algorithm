# GitHub Intro for Students — Using the Terminal (CLI)

**Task 2 · Part A · International Class · 2026**

> **Predict first. Run second. Explain third.**  
> Before you press Enter on a command, say out loud what you expect it to change.

By the end of this guide you can take the teaching repository, work on your own copy, save your progress in small steps, and publish it so your instructor can read it. You will do all of it in the **terminal** (a text window where you type commands; also called the *CLI*, command-line interface).

---

## 1. Four ideas you need first

| Word | Plain meaning | Analogy |
|---|---|---|
| **Git** | A program on your computer that remembers every saved version of your files. | A save-game system with unlimited slots. |
| **GitHub** | A website that stores Git projects online so people can share them. | A cloud drive built for Git. |
| **Repository (repo)** | A project folder that Git is tracking. | A notebook with a full history. |
| **Fork** | Your own online copy of someone else's repository. | Photocopying a worksheet so you can write on it. |

Three places hold the same project:

```text
 instructor's repo  --fork-->  your fork on GitHub  --clone-->  your computer
   (upstream)                       (origin)                     (local folder)
        ^                              ^   |
        |                              |   |  git push
        +---- you never push here -----+   v
                                    your commits travel up
```

**Rule of this course:** never edit the instructor's `main` branch. Your work lives in your fork.

---

## 2. Check your tools

Open a terminal (Windows: *PowerShell* or *Git Bash*; macOS/Linux: *Terminal*; or the terminal inside VS Code with `` Ctrl+` ``).

```bash
git --version
go version
```

**Expected:** two version lines, for example `git version 2.x.x` and `go version go1.2x ...`. If you see *command not found* or *not recognized*, install Git from <https://git-scm.com> and Go from <https://go.dev/dl>, then open a **new** terminal.

### Tell Git who you are (once per computer)

```bash
git config --global user.name "Your Name"
git config --global user.email "your-github-email@example.com"
git config --global init.defaultBranch main
```

> **Privacy:** this email is written into every commit you publish, and this repository is public. On GitHub open *Settings → Emails* and tick *Keep my email addresses private*, then use the `...@users.noreply.github.com` address shown there.  
> Never put your student ID, grades, passwords or tokens in any file, commit message or issue.

---

## 3. Terminal survival kit

| Command | What it does |
|---|---|
| `pwd` | Show the folder you are in. (PowerShell: `pwd` or `Get-Location`.) |
| `ls` | List files. (Command Prompt: `dir`.) |
| `cd folder_name` | Enter a folder. |
| `cd ..` | Go up one folder. |
| `cd student/10_task2_types` | Enter a nested folder in one step. |
| `Tab` key | Auto-complete a file or folder name. |
| `↑` key | Bring back the previous command. |

**Predict:** if you are in `students-worksheet-algorithm` and type `cd student`, then `ls`, what will you see? Write down your guess, then check.

---

## 4. Step-by-step workflow

### Step 1 — Fork (in the browser, one time)

1. Sign in to GitHub.
2. Open <https://github.com/shiddiqeuy/students-worksheet-algorithm>.
3. Click **Fork → Create fork**. You now own `https://github.com/<your-handle>/students-worksheet-algorithm`.

### Step 2 — Clone your fork

```bash
git clone https://github.com/<your-handle>/students-worksheet-algorithm.git
cd students-worksheet-algorithm
```

`git clone` downloads the project and its full history. The remote called **origin** now points to your fork.

If Git asks you to sign in, choose the browser sign-in option. **Do not paste a password or token into a file or chat.**

### Step 3 — Link the instructor's repo (optional but useful)

```bash
git remote add upstream https://github.com/shiddiqeuy/students-worksheet-algorithm.git
git remote -v
```

`origin` is your fork; `upstream` is the original. Later, if the instructor adds fixes, switch to `main` first (`git switch main`) and get them with:

```bash
git fetch upstream
git merge upstream/main
```

### Step 4 — Create a branch for your work

```bash
git switch -c task-2
git branch
```

A **branch** is a separate line of work. The `*` marks the one you are on. Working on `task-2` keeps `main` clean.

### Step 5 — Do the work, then look before you save

Open `student/10_task2_types/main.go`, fill in the `TODO` lines (see [go-basic-types.md](go-basic-types.md)), then:

```bash
cd student/10_task2_types
gofmt -w main.go
go vet main.go
go run main.go
cd ../..
```

Now ask Git what changed:

```bash
git status
git diff
```

- `git status` lists files that are new, changed, or staged.
- `git diff` shows the exact lines you added (`+`) and removed (`-`).

### Step 6 — Stage and commit

```bash
git add student/10_task2_types/main.go
git commit -m "Complete task 2 types exercise"
```

Think of **staging** (`git add`) as putting items in a box and **committing** as sealing and labelling the box. A good message completes the sentence *"This commit will…"* and stays under about 60 characters.

Check your history:

```bash
git log --oneline main..task-2
```

`main..task-2` means "commits on `task-2` that are not on `main`", so you see only your own work.

Make **small commits**: one per finished exercise.

### Step 7 — Push to your fork

```bash
git push -u origin task-2
```

`-u` links your local `task-2` to the same-named branch on `origin`, so next time `git push` is enough. Open your fork in the browser and refresh: your branch and commits are now online.

### Step 8 — Submit with the issue form

1. Go to the **Issues** tab of the *instructor's* repository and click **New issue**.
2. Choose **Task 2 — GitHub CLI and Go Basic Types**.
3. Paste the link to your fork (and the branch name `task-2`), fill the fields, tick the checklist, and submit.

You do **not** open a Pull Request in this course. A Pull Request is a request to merge your branch into someone else's repository. You will meet it in later courses.

---

## 5. Evidence you must collect

Copy the terminal output (or take a screenshot) for each item, and put it in your `EXPLANATION.md` or in an `evidence/` folder in your fork:

1. `git --version` and `go version`
2. `git remote -v` (shows `origin` pointing to **your** fork)
3. `git status` showing a changed file, before committing
4. `git log --oneline main..task-2` showing at least **three** commits (this lists only *your* commits, not the instructor's history)
5. `go run main.go` output of your finished exercise

Hide any private information before sharing a screenshot.

---

## 6. When something goes wrong

| Message | Likely cause | What to try |
|---|---|---|
| `fatal: not a git repository` | You are outside the project folder. | `pwd`, then `cd` into the cloned folder. |
| `nothing to commit, working tree clean` | You have not changed anything, or forgot to save the file in your editor. | Save the file, run `git status` again. |
| `Author identity unknown` | Step 2 of section 2 was skipped. | Run the two `git config` commands. |
| `error: failed to push some refs` | Your fork has commits you lack locally. | `git pull origin task-2`, then push again. If unsure, ask. |
| `Permission denied` / `403` on push | You are pushing to the instructor's repo instead of your fork. | `git remote -v`; `origin` must contain **your** handle. |
| You committed the wrong file | The file was staged by mistake. | `git restore --staged file_name` (before commit). Ask for help after commit. |

**Safe habits:** run `git status` before every `add`; never copy-paste commands containing `--force`, `reset --hard`, or `rm -rf` unless your instructor tells you to.

---

## 7. Practice checklist (self-check)

- [ ] I can explain the difference between Git and GitHub in one sentence.
- [ ] I can explain fork versus clone.
- [ ] I created a branch named `task-2` and confirmed it with `git branch`.
- [ ] I ran `git status` and `git diff` before each commit.
- [ ] I have at least three commits with meaningful messages.
- [ ] My `origin` is my own fork.
- [ ] I submitted via the Task 2 issue form without private information.

## 8. Mini-defense questions (2 minutes)

1. What is the difference between `git add` and `git commit`?
2. Why do we work on a branch instead of `main`?
3. What does `git push` change, and where?
4. Which command shows what you are about to commit?

---

Next: **[Go basic types, constants and operators →](go-basic-types.md)**
