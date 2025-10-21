package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/juishuyeh/todo_cli/internal/app"
	"github.com/juishuyeh/todo_cli/pkg/models"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[37m"
	colorBold   = "\033[1m"
)

func main() {
	application, err := app.New()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing app: %v\n", err)
		os.Exit(1)
	}
	defer application.Close()

	if len(os.Args) < 2 {
		runInteractive(application)
		return
	}

	command := os.Args[1]
	switch command {
	case "add", "a":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <task title>")
			os.Exit(1)
		}
		title := strings.Join(os.Args[2:], " ")
		if err := application.AddTask(title); err != nil {
			fmt.Fprintf(os.Stderr, "Error adding task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s✓%s Task added successfully!\n", colorGreen, colorReset)

	case "list", "ls", "l":
		listTasks(application, false)

	case "done", "do", "d":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo done <task id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid task ID: %v\n", err)
			os.Exit(1)
		}
		if err := application.ToggleTask(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error toggling task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s✓%s Task status updated!\n", colorGreen, colorReset)

	case "delete", "del", "rm":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo delete <task id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid task ID: %v\n", err)
			os.Exit(1)
		}
		if err := application.DeleteTask(id); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s✓%s Task deleted!\n", colorGreen, colorReset)

	case "help", "h", "--help", "-h":
		printHelp()

	case "interactive", "i":
		runInteractive(application)

	default:
		fmt.Printf("Unknown command: %s\n", command)
		printHelp()
		os.Exit(1)
	}
}

func listTasks(app *app.App, interactive bool) {
	tasks, err := app.ListTasks()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error listing tasks: %v\n", err)
		os.Exit(1)
	}

	if len(tasks) == 0 {
		if !interactive {
			fmt.Println("No tasks yet. Add one with: todo add <task>")
		}
		return
	}

	if !interactive {
		fmt.Printf("\n%s%s📝 Todo List%s\n", colorBold, colorCyan, colorReset)
		fmt.Println(strings.Repeat("─", 50))
	}

	for _, task := range tasks {
		displayTask(task)
	}

	if !interactive {
		fmt.Println(strings.Repeat("─", 50))
		fmt.Printf("\nTotal: %d tasks\n\n", len(tasks))
	}
}

func displayTask(task *models.Task) {
	checkbox := "☐"
	color := colorYellow
	if task.Done {
		checkbox = "✓"
		color = colorGray
	}

	fmt.Printf("%s[%d] %s %s%s\n",
		color,
		task.ID,
		checkbox,
		task.Title,
		colorReset,
	)
}

func printHelp() {
	help := fmt.Sprintf(`
%s%s📝 Todo CLI - 終端機待辦事項管理工具%s

%s使用方式:%s
  todo                    啟動互動模式
  todo add <任務>         新增任務
  todo list              列出所有任務
  todo done <id>         切換任務完成狀態
  todo delete <id>       刪除任務
  todo help              顯示此幫助訊息

%s別名:%s
  add  -> a
  list -> ls, l
  done -> do, d
  delete -> del, rm

%s範例:%s
  todo add 買牛奶
  todo list
  todo done 1
  todo delete 2

%s互動模式快捷鍵:%s
  a - 新增任務
  d - 刪除選中的任務
  Space - 切換任務完成狀態
  j/↓ - 下移
  k/↑ - 上移
  q - 退出

`, colorBold+colorCyan, "🎯 ", colorReset,
		colorBold, colorReset,
		colorBold, colorReset,
		colorBold, colorReset,
		colorBold, colorReset)

	fmt.Print(help)
}

func runInteractive(app *app.App) {
	reader := bufio.NewReader(os.Stdin)
	selectedIndex := 0

	for {
		clearScreen()
		printHeader()

		tasks, err := app.ListTasks()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			return
		}

		if len(tasks) == 0 {
			fmt.Printf("%s還沒有任何任務%s\n\n", colorGray, colorReset)
		} else {
			for i, task := range tasks {
				prefix := "  "
				if i == selectedIndex {
					prefix = fmt.Sprintf("%s>%s ", colorCyan, colorReset)
				}

				checkbox := "☐"
				taskColor := colorYellow
				if task.Done {
					checkbox = "✓"
					taskColor = colorGray
				}

				fmt.Printf("%s%s[%d] %s %s%s\n",
					prefix,
					taskColor,
					task.ID,
					checkbox,
					task.Title,
					colorReset,
				)
			}

			// Ensure selectedIndex is within bounds
			if selectedIndex >= len(tasks) {
				selectedIndex = len(tasks) - 1
			}
			if selectedIndex < 0 && len(tasks) > 0 {
				selectedIndex = 0
			}
		}

		printFooter()

		char, err := readChar(reader)
		if err != nil {
			return
		}

		switch char {
		case "q", "Q":
			fmt.Println("Goodbye!")
			return

		case "a", "A":
			fmt.Printf("\n%s新增任務:%s ", colorCyan, colorReset)
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title != "" {
				if err := app.AddTask(title); err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					waitForEnter(reader)
				}
				selectedIndex = 0 // Reset to top after adding
			}

		case "d", "D":
			if len(tasks) > 0 && selectedIndex < len(tasks) {
				task := tasks[selectedIndex]
				fmt.Printf("\n%s確定要刪除「%s」嗎? (y/n):%s ", colorRed, task.Title, colorReset)
				confirm, _ := reader.ReadString('\n')
				if strings.ToLower(strings.TrimSpace(confirm)) == "y" {
					if err := app.DeleteTask(task.ID); err != nil {
						fmt.Fprintf(os.Stderr, "Error: %v\n", err)
						waitForEnter(reader)
					}
					if selectedIndex >= len(tasks)-1 {
						selectedIndex = len(tasks) - 2
					}
					if selectedIndex < 0 {
						selectedIndex = 0
					}
				}
			}

		case " ":
			if len(tasks) > 0 && selectedIndex < len(tasks) {
				task := tasks[selectedIndex]
				if err := app.ToggleTask(task.ID); err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
					waitForEnter(reader)
				}
			}

		case "j", "J":
			if selectedIndex < len(tasks)-1 {
				selectedIndex++
			}

		case "k", "K":
			if selectedIndex > 0 {
				selectedIndex--
			}
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func printHeader() {
	fmt.Printf("\n%s%s╔═══════════════════════════════════════════════════════╗%s\n", colorBold, colorCyan, colorReset)
	fmt.Printf("%s%s║         📝 Todo CLI - 待辦事項管理工具               ║%s\n", colorBold, colorCyan, colorReset)
	fmt.Printf("%s%s╚═══════════════════════════════════════════════════════╝%s\n\n", colorBold, colorCyan, colorReset)
}

func printFooter() {
	fmt.Printf("\n%s─────────────────────────────────────────────────────────%s\n", colorGray, colorReset)
	fmt.Printf("%s[a]%s 新增  %s[d]%s 刪除  %s[Space]%s 完成/未完成  %s[j/k]%s 上下移動  %s[q]%s 退出\n",
		colorCyan, colorReset,
		colorCyan, colorReset,
		colorCyan, colorReset,
		colorCyan, colorReset,
		colorCyan, colorReset,
	)
}

func readChar(reader *bufio.Reader) (string, error) {
	char, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(char), nil
}

func waitForEnter(reader *bufio.Reader) {
	fmt.Print("Press Enter to continue...")
	reader.ReadString('\n')
}
