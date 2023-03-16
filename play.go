package goplay

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func bruteForceBinsRules(rules, bins []string) []string {
	subset := make([]string, 0, len(bins))
	for _, bin := range bins {
		for _, rule := range rules {
			if strings.HasPrefix(bin, rule) {
				subset = append(subset, bin)
			}
		}
	}
	return subset
}

var (
	mapSet  = false
	ruleSet = map[string]bool{}
)

// technically an O(n) + O(m) operation
func mapBinsRules(rules, bins []string) []string {
	subset := make([]string, 0, len(bins))
	initRules(rules)
	for _, bin := range bins {
		if fulfilsRules(bin) {
			subset = append(subset, bin)
		}
	}
	return subset
}

// O(m) operation
func initRules(rules []string) {
	if mapSet {
		return
	}
	for _, rule := range rules {
		ruleSet[rule] = true
	}
	mapSet = true
}

// Technically an O(1) operation
func fulfilsRules(bin string) bool {
	max := 6
	if len(bin) < 6 {
		max = len(bin) // setting length to be min(len(bin), 6)
	}

	var ok bool
	for i := 2; i < max+1; i++ { // O(5)
		if ok = ruleSet[bin[:i]]; !ok { // O(1)
			continue
		}
		return true
	}
	return false
}

type foo struct {
	Bar string `json:"bar"`
	Baz int    `json:"baz"`
}

func sleeper(ctx context.Context) string {
	time.Sleep(1 * time.Second)
	return "hi"
}

var (
	token = ""
	url   = ""
)

func CallFwF() {
	cw := initCsv()
	today := time.Now()

	for _, project := range []string{"pandora", "fintech"} {
		for _, flag := range flags {
			client := &http.Client{}
			req, err := http.NewRequest("GET", fmt.Sprintf(url, project, flag), nil)
			if err != nil {
				log.Fatal(err)
			}
			req.Header.Set("authority", "")
			req.Header.Set("accept", "*/*")
			req.Header.Set("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
			req.Header.Set("app-version", "13.9.1")
			req.Header.Set("authorization", token)
			req.Header.Set("cache-control", "no-cache")
			req.Header.Set("content-type", "application/json")
			req.Header.Set("origin", "")
			req.Header.Set("pragma", "no-cache")
			req.Header.Set("referer", "/")
			req.Header.Set("sec-ch-ua", `"Google Chrome";v="111", "Not(A:Brand";v="8", "Chromium";v="111"`)
			req.Header.Set("sec-ch-ua-mobile", "?0")
			req.Header.Set("sec-ch-ua-platform", `"macOS"`)
			req.Header.Set("sec-fetch-dest", "empty")
			req.Header.Set("sec-fetch-mode", "cors")
			req.Header.Set("sec-fetch-site", "same-site")
			req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			bodyText, err := io.ReadAll(resp.Body)

			var response FwFResponse
			json.Unmarshal(bodyText, &response)
			if err != nil {
				log.Fatal(err)
			}

			changes := len(response.Result)
			log.Printf("%s flag has %d results", flag, changes)
			if changes == 0 {
				continue
			}
			lastChange, _ := time.Parse(time.RFC3339, response.Result[0].Date)
			firstChange, _ := time.Parse(time.RFC3339, response.Result[changes-1].Date)

			difference := lastChange.Sub(firstChange)
			differenceDays := int(difference.Hours() / 24)

			daysSinceLast := int(today.Sub(lastChange).Hours() / 24)
			daysSinceFirst := int(today.Sub(firstChange).Hours() / 24)

			writeToCsv(cw, []string{
				flag,
				strconv.Itoa(changes),
				response.Result[0].Date,
				strconv.Itoa(differenceDays / changes),
				strconv.Itoa(daysSinceLast),
				strconv.Itoa(daysSinceFirst / changes)},
			)
			time.Sleep(10 * time.Millisecond)
		}
	}

}

func initCsv() *csv.Writer {
	file, err := os.Create("flag_usage.csv")
	if err != nil {
		panic(err)
	}
	csvwriter := csv.NewWriter(file)
	csvwriter.Write([]string{"flag", "changes", "last date amended", "days per change", "days since last change", "days per change till current"})
	csvwriter.Flush()

	return csvwriter
}

func writeToCsv(w *csv.Writer, s []string) {
	w.Write(s)
	w.Flush()
}

type FwFResponse struct {
	Result []details `json:"result"`
}
type details struct {
	Date string `json:"date"`
}
