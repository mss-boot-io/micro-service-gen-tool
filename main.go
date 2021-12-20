package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/lwnmengjing/micro-service-gen-tool/pkg"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	flag.Parse()
}

var (
	service     = flag.String("service", "", "generate service name")
	templateUrl = flag.String("templateUrl", "", "from template")
	config      = flag.String("config", "config.yml", "config file path")
	createRepo  = flag.Bool("createRepo", false, "auto create repo to github")
)

func main() {
	v := viper.New()
	v.SetConfigFile(*config) // optionally look for config in the working directory
	err := v.ReadInConfig()  // Find and read the config file
	if err != nil {          // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	var c pkg.TemplateConfig
	if err = v.Unmarshal(&c); err != nil {
		log.Fatalln(err)
	}
	if *service != "" {
		c.Service = *service
	}
	if *templateUrl != "" {
		c.TemplateUrl = *templateUrl
	}
	if *createRepo {
		c.CreateRepo = *createRepo
	}
	if c.CreateRepo {
		if c.Github.Name == "" {
			c.Github.Name = c.Service
		}
	}
	if c.Destination == "" {
		c.Destination = "."
	}
	c.Destination = filepath.Join(c.Destination, c.Service)
	err = pkg.Generate(&c)
	if err != nil {
		log.Fatalln(err)
	}
}
