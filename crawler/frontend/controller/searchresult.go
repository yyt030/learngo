package controller

import (
	"context"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"learngo/crawler/engine"
	"learngo/crawler/frontend/model"
	"learngo/crawler/frontend/view"

	"github.com/olivere/elastic"
)

// TODO:
// support search button
// support paging
// add start page

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view:   view.CreateSearchResultView(template),
		client: client,
	}
}

func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}
	var page model.SearchResult
	page, err = h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	err = h.view.Render(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult
	result.Query = q
	searchService := h.client.Search("dating_profile").From(from)
	if q != "" {
		searchService = searchService.Query(elastic.NewQueryStringQuery(rewriteQueryString(q)))
	}
	resp, err := searchService.Do(context.Background())
	if err != nil {
		return result, err
	}
	result.Hits = resp.TotalHits()
	result.Start = from
	result.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Za-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
