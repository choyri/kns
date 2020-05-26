package controller

import (
	"encoding/json"
	"fmt"
	"github.com/choyri/kns/model"
	"github.com/choyri/kns/service"
	"github.com/choyri/kns/store"
	"github.com/choyri/kns/support"
	"golang.org/x/sync/errgroup"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

const defaultPerPage = 20

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var (
		err     error
		queries = r.URL.Query()
		q       = queries.Get("q")
		page    = queries.Get("page")
		perPage = queries.Get("per_page")
		records []model.Order
	)

	records, err = search(q)
	if err != nil {
		http.Error(w, fmt.Sprintf("检索时出现了错误：%s", err.Error()), http.StatusInternalServerError)
		return
	}

	ret, headers := disposeResult(records, page, perPage)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Expose-Headers", "*")

	for k, v := range headers {
		w.Header().Set(k, v)
	}

	_ = json.NewEncoder(w).Encode(ret)
}

func search(q string) ([]model.Order, error) {
	var (
		err           error
		g             errgroup.Group
		db            = store.GetMySQL()
		keywords      = strings.Fields(q)
		retKeywordMap = make(map[string]map[uint]model.Order)
		retIDMap      = make(map[uint][]model.Order)
		ret           = model.Orders{}
		rawRet        = make(chan []model.Order)
	)

	for _, keyword := range keywords {
		for _, field := range service.SearchFields {
			field, keyword := field, keyword

			g.Go(func() error {
				var (
					err     error
					records []model.Order
					query   = fmt.Sprintf("%s LIKE ?", field)
				)

				err = db.Where(query, "%"+keyword+"%").Order("id DESC").Find(&records).Error
				if err != nil {
					return err
				}

				for k := range records {
					records[k].Keyword = keyword
				}

				rawRet <- records
				return nil
			})
		}
	}

	go func() {
		err = g.Wait()
		close(rawRet)
	}()

	for v := range rawRet {
		for k, vv := range v {
			if _, exists := retKeywordMap[vv.Keyword]; !exists {
				retKeywordMap[vv.Keyword] = make(map[uint]model.Order)
			}
			// 一个关键词可能搜到多条相同记录 需要去重
			if _, exists := retKeywordMap[vv.Keyword][vv.ID]; !exists {
				retKeywordMap[vv.Keyword][vv.ID] = v[k]
			}
		}
	}

	if err = g.Wait(); err != nil {
		return nil, fmt.Errorf("查询出错：%w", err)
	}

	// keywordMap 转 IDMap
	for _, v := range retKeywordMap {
		for k, vv := range v {
			retIDMap[vv.ID] = append(retIDMap[vv.ID], v[k])
		}
	}

	// 搜索是取并集
	for _, v := range retIDMap {
		if len(v) < len(keywords) {
			continue
		}
		ret = append(ret, v[0])
	}

	sort.Sort(ret)

	return ret, nil
}

func disposeResult(records []model.Order, pageStr, perPageStr string) ([]model.Order, map[string]string) {
	page := int(support.Str2Uint(pageStr))
	if page == 0 {
		page = 1
	}

	perPage := int(support.Str2Uint(perPageStr))
	if perPage == 0 {
		perPage = defaultPerPage
	}

	totalPage := int(math.Ceil(float64(len(records)) / float64(perPage)))
	if page > totalPage {
		page = totalPage
	}

	startIndex, endIndex := perPage*(page-1), perPage*page
	if endIndex > len(records) {
		endIndex = len(records)
	}

	ret := make([]model.Order, len(records[startIndex:endIndex]))
	copy(ret, records[startIndex:endIndex])

	nextPage := totalPage
	if nextPage > page {
		nextPage = page + 1
	}

	prevPage := page - 1
	if prevPage == 0 {
		prevPage = page
	}

	headers := map[string]string{
		"X-Total":       strconv.Itoa(len(records)),
		"X-Total-Pages": strconv.Itoa(totalPage),
		"X-Per-Page":    strconv.Itoa(perPage),
		"X-Page":        strconv.Itoa(page),
		"X-Next-Page":   strconv.Itoa(nextPage),
		"X-Prev-Page":   strconv.Itoa(prevPage),
	}

	return ret, headers
}
