package main

const (
	crawURL        = "https://pkg.go.dev/search?q=%s"
	querySeparator = " OR "
	scheme         = "https://"
)

var repos = []string{
	"golang.org",
	"github.com",
	"gopkg.in",
	"go.uber.org",
}
