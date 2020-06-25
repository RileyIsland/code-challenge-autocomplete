package main

import (
    "bufio"
    "fmt"
    "math"
    "net/http"
    "os"
    "sort"
    "strings"
)

type getTermError struct {
    S string
}
func (gte *getTermError) Error() string {
    return fmt.Sprintf("%v", gte.S)
}

type match struct {
    count int
    word string
}
func (m match) String() string {
    return fmt.Sprintf("\"%v\"", m.word)
}

type matches []match
func (ms matches) Len() int {
    return len(ms)
}
func (ms matches) Less (i, j int) bool {
    return ms[i].count < ms[j].count
}
func (ms matches) String() string {
    str := ""
    length := len(ms)
    if (length == 0) {
        return "[]"
    }
    for _, m := range ms[0:int(math.Min(25.0, float64(len(ms))))] {
        str = str + fmt.Sprintf("%v", m) + ","
    }
    return fmt.Sprintf("[%v]", string([]rune(str)[0:len(str)-1]))
}
func (ms matches) Swap (i, j int) {
    ms[i], ms[j] = ms[j], ms[i]
}

func main() {
    http.HandleFunc("/autocomplete", autocomplete)
    if err := http.ListenAndServe(":9000", nil); err != nil {
        panic(err)
    }
}

func autocomplete(responseWriter http.ResponseWriter, httpRequest *http.Request) {
    terms, ok := httpRequest.URL.Query()["term"]
    if !ok || len(terms[0]) < 1 {
        responseWriter.Write([]byte("Error retrieving term: term is missing"))
        return
    }
    term := terms[0]

    matches, err := getMatches(term)
    if  err != nil {
        responseWriter.Write([]byte(fmt.Sprintf("Error retrieving matches map: %v", err)))
        return
    }

    responseWriter.Write([]byte(fmt.Sprintf("%v", matches)))
}

func getMatches(term string) (matches, error) {
    matchesMap, err := getMatchesMap(term)
    if err != nil {
        return nil, err
    }

    var ms matches
    for k, v := range matchesMap {
        ms = append(ms, match{count: v, word: k})
    }
    sort.Sort(sort.Reverse(ms))

    return ms, nil
}

func getMatchesMap(term string) (map[string]int, error) {
    file, err := os.Open("./shakespeare-complete.txt")
    if err != nil {
        return nil, err
    }
    defer file.Close()
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)
    matchesMap := make(map[string]int)
    for scanner.Scan() {
        str := strings.ToLower(scanner.Text())
        if !isNonWord(str) && stringContains(str, term) {
            matchesMap[str]++
        }
    }
    return matchesMap, nil
}

func isNonWord(str string) bool {
    for i := 0; i < len(str); i++ {
        c := str[i]
        if ('a' <= c && c <= 'z') ||
            ('A' <= c && c <= 'Z') ||
             c == '\'' || c == '-' {
            continue
        }
        return true
    }
    return false
}

func stringContains(str, term string) bool {
    return strings.Contains(strings.ToLower(str), strings.ToLower(term))
}
