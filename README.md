# Jira Work Log Sender ✨

A simple, interactive CLI tool to help you log your daily work in Jira quickly and efficiently. This application reads a Markdown file containing your work logs and sends them to Jira, allowing you to easily track and manage your time.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Requirements](#requirements)
4. [Setup](#setup)
5. [Usage](#usage)
6. [Command Reference](#command-reference)
7. [Troubleshooting](#troubleshooting)
8. [Contributing](#contributing)
9. [License](#license)

---

## Introduction

Tired of logging your hours manually in Jira? **Jira Work Log Sender** automates the process by parsing a Markdown file and creating work log entries in Jira. This reduces the friction of time-tracking, so you can focus on more important tasks.

**Why you’ll love it**:

- Minimal configuration required
- Simple and customizable workflow
- Quick feedback on your logged hours

---

## Features

- **Interactive UI:** Navigate through your parsed logs and confirm changes before sending them.
- **Time Modification:** Proportionally adjust all your work logs to reach your desired total daily hours (e.g., 8 hours).
- **Status Feedback:** See the total time logged per Jira ticket, along with sending statuses.
- **Keyboard Shortcuts:** Quickly move through logs or send them with a couple of keystrokes.

---

## Requirements

- Go 1.23+ (for building the project)
- A valid Jira account (for logging work)
- Jira API token
- macOS, Linux, or Windows (with some tweaks)

---

## Setup

The application now uses a YAML-based configuration file for improved clarity and flexibility. Replace the .env file with config.yml as described below.

1. Create the YAML configuration file:

```shell
mkdir -p ~/.config/jira-work-log-sender
cp config.dist.yml ~/.config/jira-work-log-sender/config.yml
```

2. Edit config.yml (~/.config/jira-work-log-sender/config.yml) to match your environment. Here’s a breakdown of the fields:
```yaml
jira:
  baseUrl: "https://jira.example.com"   # Base URL for your Jira instance.
  username: ""                          # Your Jira username or email.
  apiToken: ""                          # API token for authenticating requests.

issueHighlight:
  highlightAfterHours: 16               # Number of hours after which tickets are highlighted.
  excludedNumbers:                      # List of Jira issue IDs excluded from highlighting.
    - "ISSUE-123"

timeModification:
  enabled: true                         # Enable or disable time modification.
  excludedNumbers:                      # List of Jira issue IDs excluded from time modification.
    - "ISSUE-456"

pathToInputFile: "Icloud/Documents/IA-writer/2. My day.md"
                                        # Relative path to your daily Markdown work log file.
                                        # A relative path from your home directory.

cacheDir: ".config/jira-work-log-sender/cache"
                                        # Relative path to the cache directory.
                                        # A relative path from your home directory pointing to the cache directory (this directory must exist).

allowedTags:                            # You may leave it empty and app won't validate tags.
  - "[Engineering activities]"          # List of allowed tags for work log descriptions.
  - "[Documentation]"
  - "[Deployment&Monitoring]"
  - "[Research&Investigation]"
  - "[Code review]"
  - "[Communication]"
  - "[Environment Issue]"
  - "[Operational work]"
  - "[Other]"
```

3. **Install Go dependencies**:

```shell
go mod download
```

4. **Build the project**:

```shell
make build
```

5. **Install (move) the binary** to your `$PATH` (example for macOS):

    sudo mv ./bin/app /usr/local/bin/tt

   *You can rename `app` to any alias you prefer, such as `tt`.*

6. **Run the application**:

```shell
tt
```

Congratulations! You’re ready to start logging your work.

---

## Usage

1. **Create or open your `.md` file** (pointed to by `PATH_TO_INPUT_FILE`) and fill it with your daily tasks.  
   For example:

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

   > **Tip:** The `# Worklogs` heading is required. Each work log entry starts with `##`, followed by `Name | ISSUE-ID Time`.

2. **Run the application**:

```shell
tt
```

3. **View the parsed logs** and confirm before sending:

```
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
```

   - **Name**: A descriptive name (for your reference).
   - **T**: The original time you wanted to log.
   - **MT**: The modified time (if you have leftover hours to reach 8 hours, it’s automatically distributed).
   - **Issue**: The Jira ticket ID.
   - **Description**: The tags or activities you performed.

4. **Send work logs** to Jira by pressing `l` twice. You’ll see a result screen like this:

```
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
```

---

## Command Reference

- **R** — Reload the file from disk.
- **LL** — Send work logs to Jira
- **J** or **Arrow Down** — Move down in the list.
- **K** or **Arrow Up** — Move up in the list.
- **GG** — Jump to the top of the list.
- **G** — Jump to the bottom of the list.
- **Y** — Copy the selected line/title to clipboard.
- **M** — Toggle "modify time" mode for the selected ticket.
- **Shift+M** — Toggle "modify time" for all tickets at once.
- **Q/Space/Return/Esc** — Exit the application.

---

## Troubleshooting

- **File Not Found**: Check that your `PATH_TO_INPUT_FILE` in `~/.config/jira-work-log-sender/env` is correct and that the file exists.
- **Incorrect Time Format**: Ensure you specify time in the format `XhYm` (e.g., `7h40m` or just `30m`).
- **Cache Directory Issues**: Confirm `CACHE_DIR` points to a valid directory.
- **Invalid Jira Credentials**: Double-check your `JIRA_BASE_URL`, `JIRA_USERNAME`, and `JIRA_API_TOKEN`.

---

## Contributing

Contributions are welcome! Feel free to open issues and pull requests to help improve this project. Whether it's bug fixes, new features, or documentation improvements, every bit counts.

1. **Fork** the project
2. **Create** your feature branch (`git checkout -b feature/amazing-improvement`)
3. **Commit** your changes (`git commit -m 'Add some amazing improvements'`)
4. **Push** to the branch (`git push origin feature/amazing-improvement`)
5. **Open** a Pull Request

---

## License

This project is licensed under the MIT License — see the [LICENSE](./LICENSE) file for details.

---

**Happy Logging!** 🚀

If you run into any issues or have questions, please open an [issue on GitHub](#) (if applicable) or send a message. We’d love to hear your feedback and suggestions!
