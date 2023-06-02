package cmd

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/9illes/log-parser/api"
	"github.com/spf13/cobra"
)

var port string
var profiler bool

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "listen port")
	serveCmd.Flags().BoolVarP(&profiler, "profiler", "i", false, "start profiler")

	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start http server",
	Long:  `Start web UI`,
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {

	if profiler {
		go func() {
			log.Println(http.ListenAndServe("localhost:6060", nil))
		}()
	}

	HTTP := api.NewHTTP()

	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/", fileServer)
	http.HandleFunc("/v1/load", HTTP.LoadHandler)

	fmt.Println("Starting server at port " + port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
