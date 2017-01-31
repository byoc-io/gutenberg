package main

import (
	"fmt"
	"github.com/byoc-io/gutenberg/pkg/gutenberg"
	"github.com/byoc-io/gutenberg/pkg/storage/database"
	"github.com/spf13/cobra"
	"gopkg.in/mgo.v2"
	"os"
)

func commandServe() *cobra.Command {
	cmd := cobra.Command{
		Use:   "serve",
		Short: "Start HTTP Server",
		Run: func(cmd *cobra.Command, args []string) {
			c, err := readConfig()
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			if err := serve(c); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
		},
	}

	return &cmd
}

func serve(c *Config) error {
	session, err := mgo.Dial(c.Database.URL)
	if err != nil {
		return err
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	mgoDB := session.DB(c.Database.Name)

	if err := database.EnsureIndex(mgoDB); err != nil {
		return fmt.Errorf("unable to create database indexes")
	}

	logger, err := gutenberg.NewLogger(c.Logger.Level, c.Logger.Format)
	if err != nil {
		return fmt.Errorf("invalid config: %v", err)
	}
	if c.Logger.Level != "" {
		logger.Infof("config using log level: %s", c.Logger.Level)
	}

	srvCfg := gutenberg.Config{
		Repository: database.New(mgoDB),
		Logger:     logger,
	}

	srv, err := gutenberg.NewServer(srvCfg)
	if err != nil {
		return err
	}

	return srv.RunHTTPServer(c.Server.HTTP)
}
