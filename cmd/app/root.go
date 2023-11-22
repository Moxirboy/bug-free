package app

import (
	configs "bug-free/internal/config"
	http2 "bug-free/internal/http"
	"bug-free/internal/server"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use: "http-server",
	Run: func(cmd *cobra.Command, args []string) {
		srv := new(server.Server)
		port := configs.Configuration().RunPort

		if err := srv.Run(port, http2.Router()); err != nil {
			log.Fatalf("error occured")
		}
	},
}

func Execute() {
	//if err := rootCmd.Execute(); err != nil {
	//	fmt.Fprintf(os.Stderr, "error while executing your CLI '%s'", err)
	//	os.Exit(1)
	//}
	srv := new(server.Server)
	port := "3000"
	log.Print("started")
	if err := srv.Run(port, http2.Router()); err != nil {
		log.Fatalf("error occured")
	}
}
