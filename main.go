package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg struct{}
type resultMsg struct{ i, code int; err error; elapsed time.Duration }

type model struct {
	tbl           table.Model
	urls          []string
	start, doneAt []time.Time
	elapsed       []time.Duration
	done          []bool
	status        []string
}

func (m model) Init() tea.Cmd { return tea.Batch(m.startAll(), tickEvery(16*time.Millisecond)) }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case tea.KeyMsg:
		switch v.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.tbl.Focused() { m.tbl.Blur() } else { m.tbl.Focus() }
		}
	case tickMsg:
		now := time.Now()
		changed := false
		for i := range m.urls {
			if !m.done[i] && !m.start[i].IsZero() {
				m.elapsed[i] = now.Sub(m.start[i])
				changed = true
			}
		}
		if changed { m.refresh() }
		return m, tickEvery(16*time.Millisecond)
	case resultMsg:
		m.done[v.i] = true
		m.elapsed[v.i] = v.elapsed
		if v.err != nil || v.code >= 400 || v.code == 0 {
			if v.code > 0 { m.status[v.i] = fmt.Sprintf("FAIL (%d)", v.code) } else { m.status[v.i] = "FAIL" }
		} else {
			m.status[v.i] = fmt.Sprintf("OK (%d)", v.code)
		}
		m.refresh()
	}
	var cmd tea.Cmd
	m.tbl, cmd = m.tbl.Update(msg)
	return m, cmd
}

func (m model) View() string {
	frame := lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240"))
	return frame.Render(m.tbl.View()) + "\n"
}

func (m model) startAll() tea.Cmd {
	client := &http.Client{Timeout: 12 * time.Second}
	cmds := make([]tea.Cmd, 0, len(m.urls))
	for i, u := range m.urls {
		m.start[i] = time.Now()
		m.status[i] = "…"
		ii, url := i, u
		cmds = append(cmds, func() tea.Msg {
			start := time.Now()
			req, _ := http.NewRequest(http.MethodGet, url, nil)
			// keep it simple; just GET and close
			resp, err := client.Do(req)
			code := 0
			if resp != nil { code = resp.StatusCode; resp.Body.Close() }
			return resultMsg{ii, code, err, time.Since(start)}
		})
	}
	return tea.Batch(cmds...)
}

func (m *model) refresh() {
	rows := m.tbl.Rows()
	for i := range rows {
		rows[i] = table.Row{nz(m.status[i], "…"), m.urls[i], fmtDur(m.elapsed[i])}
	}
	m.tbl.SetRows(rows)
}

func tickEvery(d time.Duration) tea.Cmd { return tea.Tick(d, func(time.Time) tea.Msg { return tickMsg{} }) }

func fmtDur(d time.Duration) string {
	if d < 0 { d = 0 }
	return fmt.Sprintf("%02d:%02d.%03d",
		int(d/time.Minute),
		int(d/time.Second)%60,
		int(d/time.Millisecond)%1000,
	)
}
func nz(s, fb string) string { if s == "" { return fb }; return s }
func clamp(v, lo, hi int) int { if v < lo { return lo }; if v > hi { return hi }; return v }

func load() (urls []string, rows []table.Row, urlMax int) {
	b, err := os.ReadFile("urls.txt")
	if err != nil { fmt.Println("read urls.txt:", err); os.Exit(1) }
	for _, ln := range strings.Split(strings.TrimSpace(string(b)), "\n") {
		u := strings.TrimSpace(ln)
		if u == "" { continue }
		urls = append(urls, u)
		if l := utf8.RuneCountInString(u); l > urlMax { urlMax = l }
	}
	rows = make([]table.Row, len(urls))
	for i, u := range urls { rows[i] = table.Row{"…", u, "00:00.000"} }
	return
}

func newModel(urls []string, rows []table.Row, urlMax int) model {
	cols := []table.Column{
		{Title: "Status", Width: 12},
		{Title: "Url", Width: clamp(urlMax, 20, 80)},
		{Title: "Elapsed (ms)", Width: 14},
	}
	t := table.New(table.WithColumns(cols), table.WithRows(rows), table.WithFocused(true), table.WithHeight(clamp(len(rows)+3, 8, 20)))
	st := table.DefaultStyles()
	st.Header = st.Header.BorderStyle(lipgloss.NormalBorder()).BorderForeground(lipgloss.Color("240")).BorderBottom(true)
	st.Selected = st.Selected.Foreground(lipgloss.Color("229")).Background(lipgloss.Color("57"))
	t.SetStyles(st)

	n := len(urls)
	return model{
		tbl:     t,
		urls:    urls,
		start:   make([]time.Time, n),
		doneAt:  make([]time.Time, n),
		elapsed: make([]time.Duration, n),
		done:    make([]bool, n),
		status:  make([]string, n),
	}
}

func main() {
	urls, rows, maxLen := load()
	m := newModel(urls, rows, maxLen)
	if _, err := tea.NewProgram(m).Run(); err != nil { fmt.Println("Error:", err); os.Exit(1) }
}
