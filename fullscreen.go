package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type model struct {
	table table.Model
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter":
			return m, tea.Batch(
				tea.Printf("Let's go to %s!", m.table.SelectedRow()[1]),
			)
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return baseStyle.Render(m.table.View()) + "\n"
}

func main() {
	columns := []table.Column{
		{Title: "Status", Width: 6},
		{Title: "Url", Width: 10},
		{Title: "Elapsed Time", Width: 10},
	}

	rows := []table.Row{
		{"Tokyo", "Japan", "37,274,000"},
		{"Delhi", "India", "32,065,760"},
		{"Shanghai", "China", "28,516,904"},
		{"Dhaka", "Bangladesh", "22,478,116"},
		{"São Paulo", "Brazil", "22,429,800"},
		{"Mexico City", "Mexico", "22,085,140"},
		{"Cairo", "Egypt", "21,750,020"},
		{"Beijing", "China", "21,333,332"},
		{"Mumbai", "India", "20,961,472"},
		{"Osaka", "Japan", "19,059,856"},
		{"Chongqing", "China", "16,874,740"},
		{"Karachi", "Pakistan", "16,839,950"},
		{"Istanbul", "Turkey", "15,636,243"},
		{"Kinshasa", "DR Congo", "15,628,085"},
		{"Lagos", "Nigeria", "15,387,639"},
		{"Buenos Aires", "Argentina", "15,369,919"},
		{"Kolkata", "India", "15,133,888"},
		{"Manila", "Philippines", "14,406,059"},
		{"Tianjin", "China", "14,011,828"},
		{"Guangzhou", "China", "13,964,637"},
		{"Rio De Janeiro", "Brazil", "13,634,274"},
		{"Lahore", "Pakistan", "13,541,764"},
		{"Bangalore", "India", "13,193,035"},
		{"Shenzhen", "China", "12,831,330"},
		{"Moscow", "Russia", "12,640,818"},
		{"Chennai", "India", "11,503,293"},
		{"Bogota", "Colombia", "11,344,312"},
		{"Paris", "France", "11,142,303"},
		{"Jakarta", "Indonesia", "11,074,811"},
		{"Lima", "Peru", "11,044,607"},
		{"Bangkok", "Thailand", "10,899,698"},
		{"Hyderabad", "India", "10,534,418"},
		{"Seoul", "South Korea", "9,975,709"},
		{"Nagoya", "Japan", "9,571,596"},
		{"London", "United Kingdom", "9,540,576"},
		{"Chengdu", "China", "9,478,521"},
		{"Nanjing", "China", "9,429,381"},
		{"Tehran", "Iran", "9,381,546"},
		{"Ho Chi Minh City", "Vietnam", "9,077,158"},
		{"Luanda", "Angola", "8,952,496"},
		{"Wuhan", "China", "8,591,611"},
		{"Xi An Shaanxi", "China", "8,537,646"},
		{"Ahmedabad", "India", "8,450,228"},
		{"Kuala Lumpur", "Malaysia", "8,419,566"},
		{"New York City", "United States", "8,177,020"},
		{"Hangzhou", "China", "8,044,878"},
		{"Surat", "India", "7,784,276"},
		{"Suzhou", "China", "7,764,499"},
		{"Hong Kong", "Hong Kong", "7,643,256"},
		{"Riyadh", "Saudi Arabia", "7,538,200"},
		{"Shenyang", "China", "7,527,975"},
		{"Baghdad", "Iraq", "7,511,920"},
		{"Dongguan", "China", "7,511,851"},
		{"Foshan", "China", "7,497,263"},
		{"Dar Es Salaam", "Tanzania", "7,404,689"},
		{"Pune", "India", "6,987,077"},
		{"Santiago", "Chile", "6,856,939"},
		{"Madrid", "Spain", "6,713,557"},
		{"Haerbin", "China", "6,665,951"},
		{"Toronto", "Canada", "6,312,974"},
		{"Belo Horizonte", "Brazil", "6,194,292"},
		{"Khartoum", "Sudan", "6,160,327"},
		{"Johannesburg", "South Africa", "6,065,354"},
		{"Singapore", "Singapore", "6,039,577"},
		{"Dalian", "China", "5,930,140"},
		{"Qingdao", "China", "5,865,232"},
		{"Zhengzhou", "China", "5,690,312"},
		{"Ji Nan Shandong", "China", "5,663,015"},
		{"Barcelona", "Spain", "5,658,472"},
		{"Saint Petersburg", "Russia", "5,535,556"},
		{"Abidjan", "Ivory Coast", "5,515,790"},
		{"Yangon", "Myanmar", "5,514,454"},
		{"Fukuoka", "Japan", "5,502,591"},
		{"Alexandria", "Egypt", "5,483,605"},
		{"Guadalajara", "Mexico", "5,339,583"},
		{"Ankara", "Turkey", "5,309,690"},
		{"Chittagong", "Bangladesh", "5,252,842"},
		{"Addis Ababa", "Ethiopia", "5,227,794"},
		{"Melbourne", "Australia", "5,150,766"},
		{"Nairobi", "Kenya", "5,118,844"},
		{"Hanoi", "Vietnam", "5,067,352"},
		{"Sydney", "Australia", "5,056,571"},
		{"Monterrey", "Mexico", "5,036,535"},
		{"Changsha", "China", "4,809,887"},
		{"Brasilia", "Brazil", "4,803,877"},
		{"Cape Town", "South Africa", "4,800,954"},
		{"Jiddah", "Saudi Arabia", "4,780,740"},
		{"Urumqi", "China", "4,710,203"},
		{"Kunming", "China", "4,657,381"},
		{"Changchun", "China", "4,616,002"},
		{"Hefei", "China", "4,496,456"},
		{"Shantou", "China", "4,490,411"},
		{"Xinbei", "Taiwan", "4,470,672"},
		{"Kabul", "Afghanistan", "4,457,882"},
		{"Ningbo", "China", "4,405,292"},
		{"Tel Aviv", "Israel", "4,343,584"},
		{"Yaounde", "Cameroon", "4,336,670"},
		{"Rome", "Italy", "4,297,877"},
		{"Shijiazhuang", "China", "4,285,135"},
		{"Montreal", "Canada", "4,276,526"},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := model{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
