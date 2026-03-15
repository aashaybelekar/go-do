<div align="center">

# 📝 go-do (Yet Another To-Do App)

*A fast, interactive CLI-based task manager built with Go, featuring a customized REPL environment.*

[![Go Version](https://img.shields.io/badge/Go-1.24.0-00ADD8?style=flat&logo=go)](https://golang.org/)
[![Cobra CLI](https://img.shields.io/badge/CLI-Cobra-blue?style=flat)](https://github.com/spf13/cobra)

</div>

---

**go-do** (also known as go-to-do) is a lightweight Command Line Interface (CLI) application designed to help you quickly manage your daily tasks. Running in an interactive REPL shell, it automatically stores tasks locally in a CSV database and keeps track of task completion states and relative timestamps.

## 📑 Table of Contents

- [Features](#-features)
- [Technology Stack](#-technology-stack)
- [Getting Started](#-getting-started)
  - [Prerequisites](#prerequisites)
  - [Setup](#setup)
- [Configuration](#️-configuration)
- [Usage](#️-usage--commands)

---

## ✨ Features

- **Interactive REPL Interface:** Launch into a continuous looping shell prompt (`> `) to manage tasks without repeatedly calling the binary.
- **Local CSV Storage:** Automatically initializes and manages your to-do lists in a `.csv` database locally.
- **Relative Time Tracking:** See how long ago tasks were created using human-readable relative time parsing.
- **Dynamic Task Management:** Easily add, list, mark complete, and completely delete tasks while maintaining accurate IDs.
- **Cross-Platform Compatibility:** Native support for clearing the terminal screen on Linux, Darwin (macOS), and Windows.

---

## 🚀 Technology Stack

| Domain | Technologies |
| :--- | :--- |
| **Language** | Go 1.24.0 |
| **CLI Framework**| Cobra (`github.com/spf13/cobra`) |
| **Input Parsing**| Go Shellwords (`github.com/mattn/go-shellwords`) |
| **Data Storage** | Native `encoding/csv` module |
| **Configuration**| YAML (`gopkg.in/yaml.v2`) |

---

## 🏁 Getting Started

Follow these instructions to get a copy of the project up and running on your local machine.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.24.0 or higher)

### Setup

1. **Clone the repository:**
   ```bash
   git clone <your-repository-url>
   cd go-do
   ```

2. **Install Dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the Application:**
   ```bash
   go run todo.go
   ```

---

## ⚙️ Configuration

The database storage location is highly customizable. Configuration is stored inside a YAML file located at `./config/db.yaml`.

By default, the application configures the data layer to store tasks in the `artifacts/` folder:
```yaml
location: ./artifacts/data.csv
```
*(Ensure the directory structure allows for `.csv` file generation as initialized by the app's internal DB logic)*

---

## ▶️ Usage & Commands

Once you launch the app (`go run todo.go`), you will be greeted by the `go-do` ASCII logo and entered into the REPL shell (`> `).

Available commands within the shell:

| Command | Usage | Description |
| :--- | :--- | :--- |
| `add` | `add "[description]"` | Adds a new task to your to-do list with a specific description. |
| `list` | `list` / `list -a` | Lists all pending tasks. Use `-a` or `--all` flag to output everything, including completed tasks. |
| `complete` | `complete [id]` | Marks a specific task as "true" (completed) via its integer ID. |
| `delete` | `delete [id]` | Permanently deletes a task by its ID and shifts subsequent IDs back down to maintain order. |
| `clear` | `clear` | Clears the terminal screen. |
| `exit` | `exit`, `quit`, or `q`| Closes the REPL shell and exits the application. |

**Example session:**
```bash
> add "Buy groceries"
> list
> complete 1
> list -a
> exit
```

---
