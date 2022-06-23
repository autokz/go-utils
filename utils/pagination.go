package utils

import "math"

const defaultPerPage = 10

type Paginator interface {
	GetPage() uint32
	GetLimit() int32
	GetOffset() uint32
}

type Pagination struct {
	totalCount uint32
	page       uint32
	perPage    uint32
	offset     uint32
}

func NewPagination(p uint32, limit int32, offset, totalCounts uint32) *Pagination {
	pagination := Pagination{
		totalCount: totalCounts,
		page:       1,
		perPage:    uint32(limit),
	}

	if p != 0 {
		pagination.page = p
	}

	if limit < 0 {
		pagination.perPage = 0
		pagination.offset = 0
		pagination.page = 0
	}

	pagination.offset = pagination.getOffset()
	if offset != 0 {
		pagination.offset = offset
	}

	return &pagination
}

func (p *Pagination) Offset() uint32 {
	return p.offset
}

func (p *Pagination) PerPage() uint32 {
	return p.perPage
}

func (p *Pagination) GetPagesCount() uint32 {
	if p.perPage == 0 {
		return 0
	}

	totalCount := math.Ceil(float64(p.totalCount) / float64(p.perPage))
	return uint32(totalCount)
}

func (p *Pagination) getOffset() uint32 {
	return (p.page - 1) * p.perPage
}

func CreatePagination(req Paginator, countAll uint32) *Pagination {
	p := req.GetPage()
	limit := req.GetLimit()
	offset := req.GetOffset()

	var _perPage int32
	if limit == 0 {
		_perPage = defaultPerPage
	}
	if limit > 0 {
		_perPage = limit
	}
	pager := NewPagination(p, _perPage, offset, countAll)

	return pager
}
