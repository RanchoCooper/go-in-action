package search

import "log"

// Result contains the result of a search.
type Result struct {
	Field string
	Content string
}

type Matcher interface {
	Search(feed *Feed, searchTerm string) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, results chan<- *Result) {
	// perform the search against the specified matcher.
	searchResults, err := matcher.Search(feed, searchTerm)
	if err != nil {
		log.Println(err)
		return
	}

	// write the results to the channel.
	for _, result := range searchResults {
		results <- result
	}

}

// Display writes results to the console window as they are received by the individual goroutines.
func Display(results chan *Result) {
	// the channel blocks until a result is written to the channel.
	// once the channel is closed the for loop terminates.
	for result := range results {
		log.Println("%s:%n%s\n\n", result.Field, result.Content)
	}
}