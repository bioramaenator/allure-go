package allure

import (
	"fmt"
)

type ILink interface {
}

// Link is an implementation of the Link entity used by Allure to specify the links needed for test reports.
// Such as:
// - A link to a task in Issue tracker.
// - A link to a test case in the TMS
// - Any other link (e.g. a link to an environment pod)
type Link struct {
	Name string `json:"name"` // Link name
	Type string `json:"type"` // Link's Type (issue, test case or any other)
	URL  string `json:"url"`  // Link URL
}

// LinkTypes ...
type LinkTypes string

// LinkTypes constants
const (
	LINK     LinkTypes = "link"
	ISSUE    LinkTypes = "issue"
	TESTCASE LinkTypes = "test_case"
)

// NewLink Constructor. Builds and returns a new `allure.Link` object.
func NewLink(name string, _type LinkTypes, url string) Link {
	return Link{name, string(_type), url}
}

// TestCaseLink returns TESTCASE type link
func TestCaseLink(testCase string) Link {
	linkName := fmt.Sprintf("TestCase[%s]", testCase)
	return NewLink(linkName, TESTCASE, fmt.Sprintf(getTestCasePattern(), testCase))
}

// IssueLink returns ISSUE type link
func IssueLink(issue string) Link {
	linkName := fmt.Sprintf("Issue[%s]", issue)
	return NewLink(linkName, ISSUE, fmt.Sprintf(getIssuePattern(), issue))
}

// LinkLink returns LINK type link
func LinkLink(linkname, link string) Link {
	return NewLink(linkname, LINK, link)
}
