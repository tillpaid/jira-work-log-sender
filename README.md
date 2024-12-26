# Setup

Copy `.env.dist` file to `~/.config/jira-work-log-sender/env`

```shell
mkdir -p ~/.config/jira-work-log-sender
cp .env.dist ~/.config/jira-work-log-sender/env
```

- Fill `PATH_TO_INPUT_FILE` - it should be a relative path from your home directory
- Fill `CACHE_DIR` - it should be a relative path from your home directory. And this directory should exist

Example:

```markdown
PATH_TO_INPUT_FILE="Icloud/Documents/IA-writer/2. My day.md"
CACHE_DIR=".config/jira-work-log-sender/cache"
JIRA_BASE_URL="https://jira.example.com"
JIRA_USERNAME="your_email"
JIRA_API_TOKEN="your_token"
```

Install the dependencies:

```shell
go mod download
```

Build the project:

```shell
make build
```

Move the binary to the bin folder:

> Example for macOS. For example, I would like to call it `tt`

```shell
sudo mv ./bin/app /usr/local/bin/tt
```

Run the application:

```shell
tt
```

# Usage

Filled .md file with the data:

```markdown
#  Worklogs

##  T1 - Time | ISSUE-987 10m
[Communication]
- Communication with team. Emails reading
- Daily meeting

##  1 - You task title (just for you) | ISSUE-12345 7h40m
[Engineering activities]
- Worked on something important
- Added something new
- Fixed something broken
```

Command output

```shell
┌─  Work Logs for Today  ──────────────────────────────────────────────────────┐
│ ┌────────────────────────────┬────────┬────────┬─────────────┬─────────────┐ │
│ │ Name                       │ T      │ MT     │ Issue       │ Description │ │
│ ├────────────────────────────┼────────┼────────┼─────────────┼─────────────┤ │
│ │ T1 - Time | ISSUE-987 10m  │ 10m    │ 10m    │ ISSUE-987   │ [Communicat │ │
│ │ 1 - You task title (just f │ 7h 40m │ 7h 50m │ ISSUE-12345 │ [Engineerin │ │
│ └────────────────────────────┴────────┴────────┴─────────────┴─────────────┘ │
│   Total time: 7h 50m │ Left: 0h 10m │ Total modified time: 8h 0m             │
│                                                                              │
│                                                                              │
└──────────────────────────────────────────────────────────────────────────────┘

   Action keys: R-Reload | L-Send work logs | [Q/Space/Return/Esc]-Exit
```

- Name - name of the work log (you can use it for your convenience)
- T - time you want to log
- MT - modified time. If you have time left, the system will increase time of all tickets to 8 hours in total proportionally to the time you have left
- Issue - ticket number in Jira
- Description - description of what you did

Then you can press `l` two times to send work logs to Jira

- App will send workLogs to jira and print you total time logged to the ticket (only your workLogs)

```shell
┌─  Send Work Logs  ──────────────────────────────────────────────────────────────┐
│ ┌────┬─────────────┬────────┬───────────────┬───────────────┐                   │
│ │ #  │ Issue       │ MT     │ Send status   │ Total time    │                   │
│ ├────┼─────────────┼────────┼───────────────┼───────────────┤                   │
│ │ 1  │ ISSUE-987   │ 10m    │ Done!         │ n/a           │                   │
│ │ 2  │ ISSUE-12345 │ 7h 50m │ Done!         │ n/a           │                   │
│ └────┴─────────────┴────────┴───────────────┴───────────────┘                   │
│                                                                                 │
│                                                                                 │
└─────────────────────────────────────────────────────────────────────────────────┘

   Action keys: R-Reload | [Q/Space/Return/Esc]-Exit
```

- `#` - number of the work log
- Issue - ticket number in Jira
- MT - modified time. If you have time left, the system will increase time of all tickets to 8 hours in total proportionally to the time you have left
- Send status - status of sending work log to Jira (Done! or Error)
- Total time - total time logged to the ticket (only your workLogs)

### Another action keys

- `r` - reload data from the file.
- `ll` - send work logs to Jira.
- `j` or `arrow down` - move down.
- `k` or `arrow up` - move up.
- `gg` - move to the top of the list.
- `G` - move to the end of the list.
- `yy` - copy the selected line to the clipboard (will copy title like `T1 - Time | ISSUE-987`).
- `m` - disable or enable modify time. When you disable it, you can't change the time of the work log. When you enable it, the system will increase time of all tickets to 8 hours in total proportionally to the time you have left.
- `M` - toggle modify time for all tickets.
- `q` or `space` or `return` or `esc` - exit.

## Input file explanation

- The heading `#` and the empty line after it are required
- The file separated by sections. Section is a separate work log
- Each section should have a title with the `##` symbol
- Format of title:
  - Between `##` and `|` it's a name for you, app doesn't use it
  - After `|` it's a ticket number in jira to which you want to log time
  - After ticket number it's a time you want to log. Allowed `h` and `m`. You can use only `h` or only `m` or both
- After title goes a list of descriptions of what you did
  - Each description should start with `-` symbol
  - The first row used as a `TAG` for the work log and will be validated against predefined tags we use in Jira

In case is file not valid, you forgot to fill some required fields, app will show you an error message. You for sure will not send wrong data to Jira.
