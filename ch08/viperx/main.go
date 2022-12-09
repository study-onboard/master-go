package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	processEnv()
	printLine()
	readParameters()
	printLine()
	readJsonConfigFile()
	printLine()
	readYamlConfigFile()
}

// print line
func printLine() {
	fmt.Println("---------------------------------------------------------------")
}

// process environments
func processEnv() {
	key := "WHO_AM_I"

	viper.BindEnv(key)
	whoAmI := viper.Get(key)
	fmt.Printf("type of shoAmI: %T\n", whoAmI)

	if whoAmI == nil {
		viper.Set(key, "Kut Zhang")
	}

	whoAmI = viper.Get(key)
	fmt.Printf("type of shoAmI: %T\n", whoAmI)

	fmt.Printf("Who am I? %s\n", whoAmI)
}

// read parameters
func readParameters() {
	// define flags
	flag.String("name", "", "name of task")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	flag.Int("times", 1, "repeat times")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	// bind flags
	viper.BindPFlags(pflag.CommandLine)

	// use flags
	name := viper.GetString("name")
	times := viper.GetInt("times")
	fmt.Printf("name: %s, times: %d\n", name, times)
	for i := 0; i < times; i++ {
		fmt.Printf("%d. %s\n", i, name)
	}
}

// read json config file
func readJsonConfigFile() {
	viper.SetConfigType("json")
	viper.SetConfigFile("./config.json")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Read config file error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf(
		"manager name: %s, age: %d\n",
		viper.GetString("manager.name"),
		viper.GetInt("manager.age"),
	)
}

// read yaml config file
func readYamlConfigFile() {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("./config.yaml")
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Read config file error: %s\n", err.Error())
		os.Exit(1)
	}

	fmt.Printf(
		"boss name: %s, age: %d\n",
		viper.GetString("boss.name"),
		viper.GetInt("boss.age"),
	)
}
