package emailcheck

import (
	"errors"
	"net"
	"sort"
	"strings"
)

type sorter struct {
	records []*net.MX
}

func (s *sorter) Len() int {
	return len(s.records)
}

func (s *sorter) Less(i, j int) bool {
	return s.records[i].Pref < s.records[j].Pref
}

func (s *sorter) Swap(i, j int) {
	t := s.records[i]
	s.records[i] = s.records[j]
	s.records[j] = t
}

func sortRecords(records []*net.MX) []*net.MX {
	s := &sorter{
		records: records,
	}

	sort.Sort(s)

	return s.records
}

func getRecords(email string) ([]*net.MX, error) {
	if !CheckRegex(email) {
		return nil, errors.New("Invalid email address")
	}

	domain := strings.SplitN(email, "@", 2)[1]

	records, e := net.LookupMX(domain)

	if e != nil {
		dnserr := e.(*net.DNSError)

		if dnserr.IsTemporary { //temporary would signal connectivity problems
			return nil, e
		}
	}

	return sortRecords(records), nil
}
