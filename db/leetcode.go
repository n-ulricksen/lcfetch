package db

import (
	"fmt"
	"net/http"

	"github.com/ulricksennick/lcfetch/parser"
	"github.com/ulricksennick/lcfetch/problem"
)

const (
	leetcodeApiUrl = "https://leetcode.com/api/problems/all/"
)

func main() {
	// TODO: write proper tests, dweeb..

	// dropProblems()

	// fetchAndParseProblems()

	// printAllProblems()

	testFilters()

}

func FetchAndParseProblems() {
	httpResp, err := http.Get(leetcodeApiUrl)
	must(err)

	htmlReader := httpResp.Body
	defer htmlReader.Close()

	problems, err := parser.ParseProblems(htmlReader)
	must(err)

	database, err := CreateDB()
	err = database.InsertProblems(problems)
	must(err)

	fmt.Printf("Fetched %d problems.\n", len(problems))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Test stuff
func dropProblems() {
	database, err := CreateDB()
	must(err)
	database.DropAllProblems()
}

func testFilters() {
	database, err := CreateDB()
	must(err)
	probs, err := database.GetAllProblems()

	difficulty := problem.HARD

	filtered := problem.FilterByDifficulty(probs, difficulty)
	filtered = problem.FilterOutPaid(filtered)
	for _, prob := range filtered {
		fmt.Printf("%v\t%v\n", prob.DisplayId, prob.Name)
	}
	fmt.Println()
	fmt.Printf("%d filtered problems.\n", len(filtered))
}

func printAllProblems() {
	database, err := CreateDB()
	must(err)
	probs, err := database.GetAllProblems()
	must(err)
	for _, prob := range probs {
		fmt.Printf("%v\t%v\n", prob.DisplayId, prob.Name)
	}
	fmt.Printf("Fetched %d problems.\n", len(probs))
}