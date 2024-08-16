# Setup

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

Copy `.env.dist` file to `~/.config/paysera-log-time/env`

```shell
mkdir -p ~/.config/paysera-log-time
cp .env.dist ~/.config/paysera-log-time/env
```

- Fill `PATH_TO_INPUT_FILE` - it should be a relative path from your home directory
- Fill `PATH_TO_INPUT_FILE` - it should be a relative path from your home directory

Example:

```markdown
PATH_TO_INPUT_FILE="Icloud/Documents/IA-writer/2. My day.md"
OUTPUT_SHELL_FILE=".config/paysera-log-time/output.sh"
```

Run the application:

```shell
tt
```

# Usage

Filled .md file with the data:

```markdown
#  Work logs

##  Time | test-505 10m
- Communication
- Daily

##  1 - Custom system rules - initial research | test-444 2h10m
- Research&Investigation
- Did something useful
- And did another useful stuff

##  2 - Some another ticket | test-234 3h10m
- Engineering activities
- Did something useful
- And did another useful stuff
```

Command output

```shell
+---------------------------------------------------------------+
| Log works for today                                           |
+---------------------------------------------------------------+
| 1. 10m   | 10m   | test-505 | Communication - Daily           |
| 2. 2h10m | 2h10m | test-444 | Research&Investigation - Did... |
| 3. 2h10m | 2h10m | test-444 | Engineering activities - Did... |
+---------------------------------------------------------------+
| Total time: 4h 30m | Left: 3h 30m                             |
+---------------------------------------------------------------+
| Action keys: R-Reload | L-Send work logs (double press) | ... |
+---------------------------------------------------------------+
```

Then you can press `l` two times to send work logs to Jira

- App will generate a shell script

```shell
#!/bin/bash

echo "Command number: 1" && please log-time test-505 "10m" "Communication
- Daily"

echo "Command number: 2" && please log-time test-444 "2h10m" "Research&Investigation
- Did something useful
- And did another useful stuff"

echo "Command number: 3" && please log-time test-444 "2h10m" "Engineering activities
- Did something useful
- And did another useful stuff"
```

And app will run it. You will see the output of the command

```shell
Command number: 1
Unable to add worklog.
Command number: 2
Unable to add worklog.
Command number: 3
Unable to add worklog.
```

- In my case I don't want to send a real work logs to Jira, so I set wrong ticket numbers.
- With real data you can see other messages from `please log-time`

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
