package pagi

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

func GetPagination(r *http.Request) (Request, []SortField) {
	var page uint64
	pageStr := chi.URLParam(r, "page")
	if pageStr != "" {
		n, err := strconv.Atoi(pageStr)
		if err != nil {
			page = 1
		}
		page = uint64(n)
	}
	var size uint64
	sizeStr := chi.URLParam(r, "size")
	if sizeStr != "" {
		n, err := strconv.Atoi(sizeStr)
		if err != nil {
			size = 20
		}
		size = uint64(n)
	}

	var sortFields []SortField
	if sortStr := r.URL.Query().Get("sort"); sortStr != "" {
		parts := strings.Split(sortStr, ",")
		for _, p := range parts {
			ascend := true
			field := p
			if strings.HasPrefix(p, "-") {
				ascend = false
				field = strings.TrimPrefix(p, "-")
			}
			sortFields = append(sortFields, SortField{
				Field:  field,
				Ascend: ascend,
			})
		}
	}

	return Request{
		Page: page,
		Size: size,
	}, sortFields

}
