package paginator

import (
	"net/url"
	"strconv"
	"testing"
)

func TestNew(t *testing.T) {
	cases := []struct {
		page        int
		perPage     int
		wantPage    int
		wantPerPage int
	}{
		{1, 10, 1, 10},
		{0, 10, 1, 10},
		{1, 0, 1, PerPage},
		{1, PerPageMax + 1, 1, PerPageMax},
	}

	for _, c := range cases {
		q := generateURLQuery(c.page, c.perPage)

		paginator := new(q)

		if c.wantPage != paginator.CurrentPage {
			t.Errorf("%s: want %v, got %v", t.Name(), c.wantPage, paginator.CurrentPage)
		}

		if c.wantPerPage != paginator.PerPage {
			t.Errorf("%s: want %v, got %v", t.Name(), c.wantPerPage, paginator.PerPage)
		}
	}
}

func TestLimit(t *testing.T) {
	cases := []struct {
		page    int
		perPage int
		want    int
	}{
		{1, 10, 10},
		{1, 0, PerPage},
		{1, -1, PerPage},
		{1, PerPageMax + 1, PerPageMax},
	}

	for _, c := range cases {
		q := generateURLQuery(c.page, c.perPage)

		paginator := new(q)

		if limit := paginator.Limit(); limit != c.want {
			t.Errorf("%s: want %v, got %v", t.Name(), c.want, limit)
		}
	}

}

func TestOffset(t *testing.T) {
	cases := []struct {
		page    int
		perPage int
		want    int
	}{
		{3, 50, 100},
		{3, 100, 200},
		{3, 0, PerPage * 2},
		{0, 0, 0},
		{0, 50, 0},
		{2, PerPageMax + 1, PerPageMax},
	}

	for _, c := range cases {
		q := generateURLQuery(c.page, c.perPage)
		paginator := new(q)
		if offset := paginator.Offset(); offset != c.want {
			t.Errorf("%s: want %v, got %v", t.Name(), c.want, offset)
		}
	}
}

func TestTotalPage(t *testing.T) {
	cases := []struct {
		total int
		limit int
		want  int
	}{
		{100, 10, 10},
		{101, 10, 11},
		{109, 10, 11},
		{110, 10, 11},
		{0, 10, 0},
		{100, 0, 1},
		{-100, 10, 0},
		{100, -10, 1},
	}

	for _, c := range cases {
		total := TotalPage(c.total, c.limit)
		if total != c.want {
			t.Errorf("%s: want %v, got %v", t.Name(), c.want, total)
		}
	}
}

func generateURLQuery(page, perPage int) url.Values {
	q := url.Values{}
	q.Set("page", strconv.Itoa(page))
	q.Set("per_page", strconv.Itoa(perPage))
	return q
}
