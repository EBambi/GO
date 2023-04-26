package main

import (
	"html/template"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type SortData struct {
	ArrayToSort []int
	SortedArray []int
}

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/sort", handleSort)
	log.Fatal(http.ListenAndServe(":3333", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleSort(w http.ResponseWriter, r *http.Request) {
	arrayToSortStr := r.FormValue("arrayToSort")
	arrayToSort, err := parseArray(arrayToSortStr)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	sort.Ints(arrayToSort)
	data := SortData{ArrayToSort: arrayToSort, SortedArray: arrayToSort}
	tmpl := template.Must(template.ParseFiles("sorted.html"))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func parseArray(arrayStr string) ([]int, error) {
	arrayStr = strings.ReplaceAll(arrayStr, " ", "")
	if arrayStr == "" {
		return nil, nil
	}
	arrayStrSlice := strings.Split(arrayStr, ",")
	arrayIntSlice := make([]int, len(arrayStrSlice))
	for i, s := range arrayStrSlice {
		val, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		arrayIntSlice[i] = val
	}
	return arrayIntSlice, nil
}
