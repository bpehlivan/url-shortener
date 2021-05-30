package commands

import (
	"fmt"
	"log"
	"net/http"
	"src/api"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start application server",
	Long:  "Starts the application server at configured ip address and port",
	Run:   serve,
}

func serve(cmd *cobra.Command, args []string) {
	hostAddress := viper.Get("APPLICATION_HOST").(string)
	port := viper.Get("APPLICATION_PORT").(string)
	servingUrl := hostAddress + ":" + port

	router := mux.NewRouter()
	api.SetUrlHandlers(router)

	err := router.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		pathTemplate, err := route.GetPathTemplate()
		if err == nil {
			fmt.Println("ROUTE:", pathTemplate)
		}
		pathRegexp, err := route.GetPathRegexp()
		if err == nil {
			fmt.Println("Path regexp:", pathRegexp)
		}
		queriesTemplates, err := route.GetQueriesTemplates()
		if err == nil {
			fmt.Println("Queries templates:", strings.Join(queriesTemplates, ","))
		}
		queriesRegexps, err := route.GetQueriesRegexp()
		if err == nil {
			fmt.Println("Queries regexps:", strings.Join(queriesRegexps, ","))
		}
		methods, err := route.GetMethods()
		if err == nil {
			fmt.Println("Methods:", strings.Join(methods, ","))
		}
		fmt.Println()
		return nil
	})
	if err != nil {
		fmt.Println(err)
	}


	fmt.Printf("Serving at: %s \n", servingUrl)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(servingUrl, nil))
}
