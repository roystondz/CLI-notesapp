# 📝 Termi-Notes

A fast, lightweight, and beautiful terminal-based note-taking application (TUI) built with Go and [Charm](https://charm.sh) libraries ([Bubble Tea](https://github.com/charmbracelet/bubbletea), [Bubbles](https://github.com/charmbracelet/bubbles), and [Lip Gloss](https://github.com/charmbracelet/lipgloss)).

---

## ✨ Features

- **⚡ Lightweight & Fast**: Quick note-taking directly from your command line without leaving the terminal.
- **📁 Local Markdown Vault**: Automatically saves notes as standard `.md` files in `~/.terminotes`.
- **🔍 Built-in Search & Filter**: Easily filter and fuzzy search through your notes by title.
- **🗑️ Note Deletion**: Quickly delete notes using `Ctrl + D` directly from the list view or while editing.
- **🕒 Timestamps**: View modification dates for all saved notes.
- **🎨 Stylish TUI**: Clean interface with modern styling powered by Lip Gloss.

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.25 or higher recommended)
- `make` (optional, for convenience)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/roystondz/CLI-notesapp.git
   cd CLI-notesapp
   ```

2. **Build the binary:**
   ```bash
   # Using Makefile
   make build

   # Or using Go directly
   go build -o terminotes .
   ```

3. **Run the app:**
   ```bash
   # Using Makefile
   make run

   # Or run the built binary
   ./terminotes
   ```

*(Optional)* Install globally to your PATH:
```bash
go install
```

---

## ⌨️ Keybindings & Controls

| Key Shortcut | Action |
| :--- | :--- |
| `Ctrl + N` | Create a new note (enter filename and press `Enter`) |
| `Ctrl + L` | List all existing notes |
| `Enter` | Open selected note / Confirm note creation |
| `Ctrl + S` | Save current note to disk |
| `Ctrl + D` | Delete current note (while editing) or selected note (in list view) |
| `Esc` | Return to main view / Cancel current prompt |
| `/` *(in list view)* | Filter / search through notes |
| `Ctrl + C` / `q` | Quit application |

---

## 📂 Vault Storage

All notes are stored in your user home directory under:
```
~/.terminotes/
```
Since they are stored as standard `.md` files, you can freely view, edit, back up, or sync them with external tools (Git, cloud storage, Obsidian, etc.).

---

## 🛠️ Built With

- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)** — The Elm-architecture TUI framework for Go
- **[Bubbles](https://github.com/charmbracelet/bubbles)** — TUI components (list, textinput, textarea)
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)** — Style definitions & layout styling

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
