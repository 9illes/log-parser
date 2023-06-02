package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	cf "github.com/9illes/log-parser/cloudfront"
	"github.com/9illes/log-parser/domain"
	"github.com/spf13/cobra"
)

var site string
var inputFilename string
var basedir string

func init() {
	preprocessCmd.Flags().StringVarP(&site, "site", "s", "default", "site directory")
	preprocessCmd.Flags().StringVarP(&inputFilename, "in", "i", "", "Input filename")
	preprocessCmd.Flags().StringVarP(&basedir, "out", "o", "", "base dirname")

	rootCmd.AddCommand(preprocessCmd)
}

var preprocessCmd = &cobra.Command{
	Use:   "preprocess",
	Short: "Preprocess requests",
	Long:  `Group request by type`,
	Run:   preprocessJournal,
}

func preprocessJournal(cmd *cobra.Command, args []string) {

	in, err := os.Open(inputFilename)
	if err != nil {
		log.Print(err)
		return
	}
	defer in.Close()

	parser := cf.NewParser()

	lines, err := parser.ParseFile(in)
	if err != nil {
		panic(err)
	}

	requests, err := domain.ParseLines(
		lines,
		// qualifiers
		domain.IsMediaRequest,
		domain.IsBotRequest,
		domain.IsSlowRequest(0.3),
		domain.IsCacheHitRequest)
	if err != nil {
		panic(err)
	}

	inputFilename := filepath.Base(inputFilename)
	filteredRequest := []*domain.FilteredRequests{
		domain.NewFilteredRequests(getSuffixedFullpath(basedir, site, "all", inputFilename), domain.FilterAlwaysTrue),
		domain.NewFilteredRequests(getSuffixedFullpath(basedir, site, "assets", inputFilename), domain.FilterIsAssets),
		domain.NewFilteredRequests(getSuffixedFullpath(basedir, site, "pages", inputFilename), domain.FilterIsPage),
		domain.NewFilteredRequests(getSuffixedFullpath(basedir, site, "bot", inputFilename), domain.FilterIsBot),
		domain.NewFilteredRequests(getSuffixedFullpath(basedir, site, "google", inputFilename), domain.FilterIsGoogleBot),
	}

	for _, r := range requests {
		for _, f := range filteredRequest {
			f.Append(r)
		}
	}

	for _, f := range filteredRequest {
		log.Printf("writing %d request(s) in file %s", len(f.Requests), f.Filename)
		writeJsonFile(f.Filename, f.Requests)
	}
}

// TODO move to Journal
func getSuffixedFullpath(basedir, site, suffix, inputFilename string) string {
	return fmt.Sprintf("%s/%s/%s-%s.json",
		basedir,
		site,
		strings.TrimSuffix(inputFilename, filepath.Ext(inputFilename)),
		suffix)
}

// writeJsonFile writes requests as JSON in a file
func writeJsonFile(fullpath string, requests []*domain.Request) error {

	b, err := json.MarshalIndent(requests, "", "\t")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile(fullpath, []byte(string(b)), 0666); err != nil {
		log.Fatal(err)
	}

	return err
}
