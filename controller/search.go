package controller

import (
	"encoding/json"
	"fmt"
	"github.com/choyri/kns/model"
	"github.com/choyri/kns/service"
	"github.com/choyri/kns/store"
	"golang.org/x/sync/errgroup"
	"net/http"
	"sort"
	"strings"
)

func Search(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var (
		err     error
		queries = r.URL.Query()
		q       = queries.Get("q")
		records []model.Order
	)

	records, err = search(q)
	if err != nil {
		http.Error(w, fmt.Sprintf("检索时出现了错误：%s", err.Error()), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(records)
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
