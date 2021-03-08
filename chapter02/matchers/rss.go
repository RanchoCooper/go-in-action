package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/RanchoCooper/go-in-action/chapter02/search"
)

// item defines the fields associated with the item tag
type item struct {
	XMLName xml.Name `xml:"item"`
	PubDate     string   `xml:"pubDate"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	GUID        string   `xml:"guid"`
	GeoRssPoint string   `xml:"georss:point"`
}

// image defines the fields associated with the image tag
// in the rss document.
type image struct {
	XMLName xml.Name `xml:"image"`
	URL     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
}

// channel defines the fields associated with the channel tag
// in the rss document.
type channel struct {
	XMLName        xml.Name `xml:"channel"`
	Title          string   `xml:"title"`
	Description    string   `xml:"description"`
	Link           string   `xml:"link"`
	PubDate        string   `xml:"pubDate"`
	LastBuildDate  string   `xml:"lastBuildDate"`
	TTL            string   `xml:"ttl"`
	Language       string   `xml:"language"`
	ManagingEditor string   `xml:"managingEditor"`
	WebMaster      string   `xml:"webMaster"`
	Image          image    `xml:"image"`
	Item           []item   `xml:"item"`
}

// rssDocument defines the fields associated with the rss document.
type rssDocument struct {
	XMLName xml.Name `xml:"rss"`
	Channel channel  `xml:"channel"`
}

// rssMatcher implements the Matcher interface.
type rssMatcher struct {

}

// init registers the matcher with the program.
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Search looks at the document for the specified search term.
func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	// retrieve the data to search
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// check the title for the search term.
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		// save result if we found a match.
		if matched {
			results = append(results, &search.Result{
				Field:   "Title",
				Content: channelItem.Title,
			})
		}

		// check the description for the search term.
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		// save result if we found a match.
		if matched {
			results = append(results, &search.Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}

	return results, nil
}

// retrieve performs a HTTP GET request for the rss feed and decodes the results.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("no rss feed URI provided")
	}

	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error %d\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}