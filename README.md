# Setup

Copy `.env.dist` file to `~/.config/paysera-log-time/env`

```shell
mkdir -p ~/.config/paysera-log-time
cp .env.dist ~/.config/paysera-log-time/env
```

- Fill `PATH_TO_INPUT_FILE` - it should be a relative path from your home directory
- Fill `CACHE_DIR` - it should be a relative path from your home directory. And this directory should exist

Example:

```markdown
PATH_TO_INPUT_FILE="Icloud/Documents/IA-writer/2. My day.md"
CACHE_DIR=".config/paysera-log-time/cache"
JIRA_BASE_URL="https://jira.paysera.net"
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

##  Time | TIME-505 10m
[Communication]
- RocketChat communication / emails reading
- Daily

##  1 - Separate bank charge creation | COMP-904 3h
[Engineering activities]
- Worked on something important
- Added something new
- Fixed something broken
```

Command output

```shell
+---------------------------------------------------------------------------------------+
| Work logs for today                                                                   |
+---------------------------------------------------------------------------------------+
| 1. 10m | TIME-505 | [Communication] - RocketChat communication / emails r...          |
| 2. 3h  | COMP-904 | [Engineering activities] - Worked on something import...          |
+---------------------------------------------------------------------------------------+
| Total time: 3h 10m | Left: 4h 50m                                                     |
+---------------------------------------------------------------------------------------+
| Action keys: R-Reload | L-Send work logs (double press) | [Q/Space/Return/Esc]-Exit   |
+---------------------------------------------------------------------------------------+
```

Then you can press `l` two times to send work logs to Jira

- App will send worklogs to jira and print you total time logged to the ticket (only your worklogs)

```shell
+---------------------------------------------------------------------------------------+
| Send work logs                                                                        |
+---------------------------------------------------------------------------------------+
| 1. TIME-505 | 10m | Done! | Total: 657h 21m                                           |
| 2. COMP-904 | 3h  | Done! | Total: 13h 29m                                            |
+---------------------------------------------------------------------------------------+
| Action keys: R-Reload | L-Send work logs (double press) | [Q/Space/Return/Esc]-Exit   |
+---------------------------------------------------------------------------------------+
```

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
