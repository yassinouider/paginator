// Package paginator provides utilities to setup a paginator within the context of a http request.
package paginator

import (
	"math"
	"net/http"
	"net/url"
	"strconv"
)

//Default values
var (
	PerPage     = 50
	PerPageMax  = 1000
	PerPageName = "per_page"
	PageName    = "page"
)

type Paginator struct {
	Total       int `json:"total"`
	Count       int `json:"count"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
}

//New Instantiates a Paginator struct for the current http request.
func New(r *http.Request) *Paginator {
	q := r.URL.Query()
	paginator := new(q)
	return paginator
}

func new(q url.Values) *Paginator {
	paginator := Paginator{}
	var err error

	paginator.PerPage, err = strconv.Atoi(q.Get(PerPageName))
	if err != nil || paginator.PerPage <= 0 {
		paginator.PerPage = PerPage
	}

	if PerPageMax > 0 {
		if paginator.PerPage > PerPageMax {
			paginator.PerPage = PerPageMax
		}
	}

	paginator.CurrentPage, err = strconv.Atoi(q.Get(PageName))
	if err != nil || paginator.CurrentPage <= 0 {
		paginator.CurrentPage = 1
	}

	return &paginator
}

func (p *Paginator) Limit() int {
	if p.PerPage <= 0 {
		return PerPage
	}

	if PerPageMax > 0 {
		if p.PerPage > PerPageMax {
			return PerPageMax
		}
	}

	return p.PerPage
}

//Offset Returns the current offset.
func (p *Paginator) Offset() int {
	return Offset(p.CurrentPage, p.Limit())
}

func (p *Paginator) SetCount(count int) *Paginator {
	p.Count = count
	return p
}

func (p *Paginator) SetTotal(total int) *Paginator {
	p.Total = total
	p.TotalPage = TotalPage(p.Total, p.Limit())
	return p
}

func Offset(page, limit int) int {
	if page > 0 && limit > 0 {
		return (page - 1) * limit
	}

	return 0
}

func TotalPage(total, limit int) int {
	if total <= 0 {
		return 0
	}

	if limit <= 0 {
		return 1
	}

	return int(math.Ceil(float64(total) / float64(limit)))
}
