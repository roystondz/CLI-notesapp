package main

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	vaultDir    string
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	docStyle    = lipgloss.NewStyle().Margin(1, 2)
)

type model struct {
	newFileInput           textinput.Model
	createFileInputVisible bool
	curretFile             *os.File
	noteTextArea           textarea.Model
	list                   list.Model
	showinglist            bool
}

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error getting home Directory : ", err)
	}
	vaultDir = fmt.Sprintf("%s/.terminotes", homeDir)
}

func intializeModel() model {
	err := os.MkdirAll(vaultDir, 0750)
	if err != nil {

	}

	//TextInput
	ti := textinput.New()
	ti.Placeholder = "What would you like to call it"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50
	ti.Cursor.Style = cursorStyle
	ti.PromptStyle = cursorStyle
	ti.TextStyle = cursorStyle

	//TextArea
	ta := textarea.New()
	ta.Placeholder = "Express your secrets here"
	ta.Focus()

	//List
	noteList := listFiles()
	finalList := list.New(noteList, list.NewDefaultDelegate(), 0, 0)
	finalList.Title = "All notes"
	finalList.Styles.Title = lipgloss.NewStyle().
		Foreground(lipgloss.Color("16")).Background(lipgloss.Color("254")).Padding(0, 1)
	finalList.SetShowStatusBar(false)

	return model{
		newFileInput:           ti,
		createFileInputVisible: false,
		noteTextArea:           ta,
		list:                   finalList,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-v, msg.Height-h)
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		case "ctrl+n":
			m.createFileInputVisible = true
			return m, nil
		case "enter":
			if m.curretFile != nil {
				break
			}
			if m.showinglist {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					filepath := fmt.Sprintf("%s/%s", vaultDir, item.title)
					content, err := os.ReadFile(filepath)
					if err != nil {
						log.Printf("Error reading file : ", err)
						return m, nil
					}
					m.noteTextArea.SetValue(string(content))
					f, err := os.OpenFile(filepath, os.O_RDWR, 0644)
					if err != nil {
						log.Printf("Error reading file: %v", err)
						return m, nil
					}
					m.curretFile = f
					m.showinglist = false

				}
				return m, nil

			}
			filename := m.newFileInput.Value()
			if filename != "" {
				filepath := fmt.Sprintf("%s/%s.md", vaultDir, filename)
				if _, err := os.Stat(filepath); err == nil {
					return m, nil
				}
				f, err := os.Create(filepath)
				if err != nil {
					log.Fatalf("%v", err)
				}
				m.curretFile = f
				m.createFileInputVisible = false
				m.newFileInput.SetValue("")

			}
			return m, nil
		case "ctrl+s":
			if m.curretFile == nil {
				break
			}
			if err := m.curretFile.Truncate(0); err != nil {
				fmt.Println("Cannot save file now !!")
				return m, nil
			}

			if _, err := m.curretFile.Seek(0, 0); err != nil {
				fmt.Println("Cannot save file now !!")
				return m, nil
			}

			if _, err := m.curretFile.WriteString(m.noteTextArea.Value()); err != nil {
				fmt.Println("Cannot save file now !!")
				return m, nil
			}

			if err := m.curretFile.Close(); err != nil {
				fmt.Println("Cannot close file now !!")
			}
			m.curretFile = nil
			m.noteTextArea.SetValue("")
			return m, nil
		case "ctrl+l":
			notelist := listFiles()
			m.list.SetItems(notelist)
			m.showinglist = true
			return m, nil
		case "esc":
			if m.createFileInputVisible {
				m.createFileInputVisible = false
			}
			if m.curretFile != nil {
				m.noteTextArea.SetValue("")
				m.curretFile = nil
			}
			if m.showinglist {
				if m.list.FilterState() == list.Filtering {
					break
				}
				m.showinglist = false
			}
			return m, nil
		case "ctrl+d":
			//TODO: Create a Delete hotkey
			if m.curretFile != nil {
				filepath := m.curretFile.Name()
				err := m.curretFile.Close()
				if err != nil {
					log.Fatal("Unable to close file")
				}
				if err := os.Remove(filepath); err != nil {
					log.Fatal("Error deleting file : ", err)
					return m, nil
				}
				m.curretFile = nil
				m.noteTextArea.SetValue("")
				m.list.SetItems(listFiles())
			}
			if m.showinglist {
				item, ok := m.list.SelectedItem().(item)
				if ok {
					filepath := fmt.Sprintf("%s/%s", vaultDir, item.title)
					err := os.Remove(filepath)
					if err != nil {
						log.Printf("Error deleting file : ", err)
						return m, nil
					}
					m.curretFile = nil
					m.list.SetItems(listFiles())

				}
				return m, nil

			}
			return m, nil
		}

	}

	if m.createFileInputVisible {
		m.newFileInput, cmd = m.newFileInput.Update(msg)
	}
	if m.curretFile != nil {
		m.noteTextArea, cmd = m.noteTextArea.Update(msg)
	}
	if m.showinglist {
		m.list, cmd = m.list.Update(msg)
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, cmd
}

func (m model) View() string {

	var style = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("16")).
		Background(lipgloss.Color("205")).
		PaddingLeft(4).
		PaddingRight(4)
	welcome := style.Render("Welcome to Termi-Notes")
	help_keys := "Ctrl+N: New file, Ctrl+L: List files, Esc: Back,\nCtrl+S: Save, Ctrl+D: Delete, Ctrl+C/Q: Quit"

	view := ""
	if m.createFileInputVisible {
		view = m.newFileInput.View()
	}
	if m.curretFile != nil {
		view = m.noteTextArea.View()
	}
	if m.showinglist {
		view = m.list.View()
	}
	return fmt.Sprintf("\n%s\n\n%s\n\n%s", welcome, view, help_keys)
}

func main() {
	p := tea.NewProgram(intializeModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("There has been a issue")
		os.Exit(1)
	}
}

func listFiles() []list.Item {
	items := make([]list.Item, 0)
	entries, err := os.ReadDir(vaultDir)
	if err != nil {
		log.Fatal("Erro")
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				continue
			}
			modTime := info.ModTime().Format("2006-06-06 15:05")
			items = append(items, item{
				title: entry.Name(),
				desc:  fmt.Sprintf("Modified: %s", modTime),
			})
		}

	}
	return items
}
