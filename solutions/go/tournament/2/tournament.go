package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type teamStatistics struct {
	Team    string
	Matches int
	Wins    int
	Draws   int
	Losses  int
	Points  int
}

type tournament map[string]*teamStatistics

func newTeamStatistics(team string) teamStatistics {
	return teamStatistics{
		Team:    team,
		Matches: 0,
		Wins:    0,
		Draws:   0,
		Losses:  0,
		Points:  0,
	}
}

func (ts teamStatistics) String() string {
	return fmt.Sprintf("%-30s |  %d |  %d |  %d |  %d |  %d\n", ts.Team, ts.Matches, ts.Wins, ts.Draws, ts.Losses, ts.Points)
}

func (t tournament) update(firstTeam string, secondTeam string, result string) {
	_, exists := t[firstTeam]
	if !exists {
		nts := newTeamStatistics(firstTeam)
		t[firstTeam] = &nts
	}

	_, exists = t[secondTeam]
	if !exists {
		nts := newTeamStatistics(secondTeam)
		t[secondTeam] = &nts
	}

	ft := t[firstTeam]
	st := t[secondTeam]
	ft.Matches++
	st.Matches++
	switch result {
	case "win":
		ft.Wins++
		ft.Points += 3
		st.Losses++
	case "draw":
		ft.Draws++
		ft.Points++
		st.Draws++
		st.Points++
	case "loss":
		ft.Losses++
		st.Wins++
		st.Points += 3
	}
}

func (t tournament) sortTeams() []*teamStatistics {
	st := make([]*teamStatistics, 0)
	for _, v := range t {
		st = append(st, v)
	}
	sort.Slice(st, func(i, j int) bool {
		return st[i].Points > st[j].Points || (st[i].Points == st[j].Points && st[i].Team < st[j].Team)
	})
	return st
}

func (t tournament) printTable() string {
	var builder strings.Builder
	builder.WriteString("Team                           | MP |  W |  D |  L |  P\n")
	for _, v := range t.sortTeams() {
		builder.WriteString(v.String())
	}
	return builder.String()
}

// Tally the results of a small football competition in table format.
func Tally(r io.Reader, w io.Writer) error {
	b, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	tour := tournament{}

	for _, line := range strings.Split(string(b), "\n") {
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		fields := strings.Split(line, ";")
		if l := len(fields); l != 3 {
			return fmt.Errorf("invalid number of fields (found: %d, expected: 3): '%s'", l, line)
		}

		firstTeam, secondTeam, result := fields[0], fields[1], fields[2]

		switch result {
		case "win", "loss", "draw":
			tour.update(firstTeam, secondTeam, result)
		default:
			return fmt.Errorf("invalid match result: %s (win, loss or draw expected): %s", result, line)
		}
	}

	if _, err := fmt.Fprint(w, tour.printTable()); err != nil {
		return err
	}

	return nil
}
