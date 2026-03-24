package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/mattn/go-isatty"
	"github.com/mem0ai/mem0/mem0-cli/internal/api"
)

type Printer struct {
	Format string
	IsTTY  bool
	Writer io.Writer
}

func NewPrinter(format string) *Printer {
	isTTY := isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd())
	if format == "" {
		if isTTY {
			format = "table"
		} else {
			format = "json"
		}
	}
	return &Printer{
		Format: format,
		IsTTY:  isTTY,
		Writer: os.Stdout,
	}
}

func (p *Printer) PrintMemories(memories []api.Memory) {
	if p.Format == "json" {
		p.PrintJSON(memories)
		return
	}

	w := tabwriter.NewWriter(p.Writer, 0, 0, 2, ' ', 0)
	if p.IsTTY {
		bold := color.New(color.Bold)
		bold.Fprintf(w, "ID\tMEMORY\tUSER ID\tCREATED AT\n")
	} else {
		fmt.Fprintf(w, "ID\tMEMORY\tUSER ID\tCREATED AT\n")
	}
	for _, m := range memories {
		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", m.ID, truncate(m.Memory, 60), m.UserID, m.CreatedAt)
	}
	w.Flush()
}

func (p *Printer) PrintMemory(memory *api.Memory) {
	if p.Format == "json" {
		p.PrintJSON(memory)
		return
	}

	fmt.Fprintf(p.Writer, "ID:         %s\n", memory.ID)
	fmt.Fprintf(p.Writer, "Memory:     %s\n", memory.Memory)
	fmt.Fprintf(p.Writer, "User ID:    %s\n", memory.UserID)
	fmt.Fprintf(p.Writer, "Created At: %s\n", memory.CreatedAt)
	fmt.Fprintf(p.Writer, "Updated At: %s\n", memory.UpdatedAt)
	if len(memory.Metadata) > 0 {
		data, _ := json.Marshal(memory.Metadata)
		fmt.Fprintf(p.Writer, "Metadata:   %s\n", string(data))
	}
}

func (p *Printer) PrintEntities(entities []api.Entity) {
	if p.Format == "json" {
		p.PrintJSON(entities)
		return
	}

	w := tabwriter.NewWriter(p.Writer, 0, 0, 2, ' ', 0)
	if p.IsTTY {
		bold := color.New(color.Bold)
		bold.Fprintf(w, "ID\tTYPE\tNAME\n")
	} else {
		fmt.Fprintf(w, "ID\tTYPE\tNAME\n")
	}
	for _, e := range entities {
		fmt.Fprintf(w, "%s\t%s\t%s\n", e.ID, e.Type, e.Name)
	}
	w.Flush()
}

func (p *Printer) PrintMessage(msg string) {
	fmt.Fprintln(p.Writer, msg)
}

func (p *Printer) PrintJSON(v any) {
	data, _ := json.MarshalIndent(v, "", "  ")
	fmt.Fprintln(p.Writer, string(data))
}

func truncate(s string, max int) string {
	runes := []rune(s)
	if len(runes) <= max {
		return s
	}
	return string(runes[:max-3]) + "..."
}
