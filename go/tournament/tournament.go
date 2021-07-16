package tournament

import (
	"bufio"
	"errors"
	"io"
	"sort"
	"strings"
)

type TeamRecord struct {
	Name   string
	Wins   int
	Losses int
	Draws  int
}

func (record TeamRecord) GetPoints() int {
	return record.Wins*3 + record.Draws
}

func (record TeamRecord) GetMatchesPlayed() int {
	return record.Wins + record.Draws + record.Losses
}

type ByPoints []TeamRecord

func (a ByPoints) Len() int { return len(a) }
func (a ByPoints) Less(i, j int) bool {
	// TODO: lexical sort in case of tie
	return a[i].GetPoints() > a[j].GetPoints()
}
func (a ByPoints) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func Tally(r io.Reader, w io.Writer) error {
	records := make(map[string]*TeamRecord) // team name -> record
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		result := strings.Split(scanner.Text(), ";")
		// create new team records if none exist yet
		if len(result) != 3 {
			return errors.New("invalid line of input")
		}
		if _, ok := records[result[0]]; !ok {
			records[result[0]] = &TeamRecord{Name: result[0]}
		}
		if _, ok := records[result[1]]; !ok {
			records[result[1]] = &TeamRecord{Name: result[1]}
		}
		// update tally
		if result[2] == "win" {
			records[result[0]].Wins++
			records[result[1]].Losses++
		} else if result[2] == "loss" {
			records[result[0]].Losses++
			records[result[1]].Wins++
		} else if result[2] == "draw" {
			records[result[1]].Draws++
			records[result[0]].Draws++
		} else {
			return errors.New("invalid game outcome")
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	sortedRecords := []TeamRecord{}
	for _, v := range records {
		sortedRecords = append(sortedRecords, *v)
	}

	sort.Sort(ByPoints(sortedRecords))

	return nil
}
