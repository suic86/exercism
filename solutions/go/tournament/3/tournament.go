package tournament

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type teamStatistics struct {
	team    string
	matches int
	wins    int
	draws   int
	losses  int
}

type tournament map[string]*teamStatistics

func (ts teamStatistics) String() string {
	return fmt.Sprintf("%-30s |  %d |  %d |  %d |  %d |  %d\n", ts.team, ts.matches, ts.wins, ts.draws, ts.losses, ts.points())
}

func (ts teamStatistics) points() int {
	return 3*ts.wins + ts.draws
}

func (t tournament) update(firstTeam string, secondTeam string, result string) {

	if _, exists := t[firstTeam]; !exists {
		t[firstTeam] = &teamStatistics{team: firstTeam}
	}

	if _, exists := t[secondTeam]; !exists {
		t[secondTeam] = &teamStatistics{team: secondTeam}
	}

	ft, st := t[firstTeam], t[secondTeam]
	ft.matches++
	st.matches++

	switch result {
	case "win":
		ft.wins++
		st.losses++
	case "draw":
		ft.draws++
		st.draws++
	case "loss":
		ft.losses++
		st.wins++
	}
}

func (t tournament) sortTeams() []*teamStatistics {
	st := make([]*teamStatistics, 0)
	for _, v := range t {
		st = append(st, v)
	}
	sort.Slice(st, func(i, j int) bool {
		return st[i].points() > st[j].points() || (st[i].points() == st[j].points() && st[i].team < st[j].team)
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
