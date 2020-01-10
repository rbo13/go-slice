package main

import "fmt"

// Counter sets the count of every pagination
type Counter interface {
	Count() uint
}

type paginator struct {
	nextPage       int
	prevPage       int
	pageStartIndex int
	page           int
	totalList      []int
}

func main() {
	fmt.Println("vim-go")
}

func computePagination(counter Counter, currentPage, displayLimit int) *paginator {

	page := 0
	nextPage := 1

	var pageStartIndex = 0

	if currentPage > 0 {
		page = currentPage - 1
	}

	prevPage := page

	if currentPage > 1 {
		pageStartIndex = (displayLimit * (pageStartIndex - 1)) + 1
	}

	var tList []int
	listIndex := 1
	totalList := counter.Count() / displayLimit

	if counter.Count()%displayLimit > 0 {
		totalList++
	}

	displayLimitStart := int(5)

	displayPageStart := int(0)
	displayPageEnd := ((page / displayLimit) + 1) * displayLimit

	if displayPageEnd/displayLimit > 1 {
		displayPageStart = displayPageEnd - displayLimit
	}

	if (page > 0 && page%displayLimit == 0) && displayPageEnd == page {
		displayPageEnd = page
		displayPageStart = page - displayLimitStart

		if page < displayLimitStart {
			displayPageStart = 0
		}
	}

	if displayPageEnd >= totalList {
		displayPageStart = totalList - displayLimitStart
		if totalList < displayLimitStart {
			displayPageStart = 0
		}

		displayPageEnd = totalList
	}

	if currentPage == 0 {
		nextPage = 2
	} else if currentPage < totalList {
		nextPage = currentPage + 1
	}

	for listIndex <= int(totalList) {
		tList = append(tList, listIndex)
		listIndex++
	}
	return &paginator{
		nextPage:       nextPage,
		prevPage:       prevPage,
		pageStartIndex: int(pageStartIndex),
		page:           currentPage,
		totalList:      tList[displayPageStart:displayPageEnd],
	}
}

func hasNextPage(currentPage, nextPage int) bool {
	if currentPage < nextPage {
		return true
	}
	return false
}

func hasPreviousPage(prevPage int, currentPage string) bool {
	if (prevPage >= 1) && (currentPage != "") {
		return true
	}
	return false
}
