# cli

A personal portfolio served over SSH. Instead of a website, visitors connect via `ssh` and interact with an animated terminal UI featuring auto-completion, command-routed views, streaming text, and ASCII art.

## Tech Stack

- **Go** with [Wish](https://github.com/charmbracelet/wish) (SSH server), [Bubble Tea](https://github.com/charmbracelet/bubbletea) (TUI framework), and [Lip Gloss](https://github.com/charmbracelet/lipgloss) (styling)

## Project Structure

```
cmd/ssh-server/main.go    # SSH server entrypoint
internal/ssh/server.go     # Wish middleware & config
internal/ui/model.go       # Main TUI state & update loop
internal/ui/commands.go    # Tick generators (animation & typing)
internal/ui/view_home.go   # Home view with ASCII animation
internal/ui/view_about.go  # About view with streaming text
internal/ui/components/    # Reusable UI components
assets/frames/             # ASCII animation frames (1.txt–10.txt)
```

## Running

```bash
go run cmd/ssh-server/main.go
```

Then connect:

```bash
ssh localhost -p 2222
```

### Environment Variables

| Variable   | Default   | Description       |
|------------|-----------|-------------------|
| `SSH_HOST` | `0.0.0.0` | Bind address      |
| `SSH_PORT` | `2222`    | Listening port     |

## Commands

| Command       | Description              |
|---------------|--------------------------|
| `/home`       | Landing page + animation |
| `/about`      | Streaming bio text       |
| `/quit`       | Exit the session         |