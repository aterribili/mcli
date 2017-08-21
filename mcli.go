package main

import (
	"fmt"
	"os"
	"strings"
	"io/ioutil"

	"github.com/urfave/cli"
)

// Swift
func generateViewController(name string) string {
	contents := "import UIKit\n\nclass {ViewController} : UIViewController {\n\n\toverride func viewDidLoad() {\n\t\tsuper.viewDidLoad()\n\t}\n\n}"
	return strings.Replace(contents, "{ViewController}", name, -1)
}

func createViewController(name string) {
	validate(name)

	fmt.Println("Generating view controller:", name)

	contents := generateViewController(name)
	generateFile(contents, name, "swift")
}

// React-Native
func generateComponent(name string) string {
	contents := "import React, { Component } from 'react';\nimport {\n  \n} from 'react-native';\n\nexport default class {ComponentName} extends Component {\n  state = {\n\n  };\n\n  render() {\n    return (\n\n    );\n  }\n}\n\nconst styles = {\n\n};\n\n{ComponentName}.propTypes = {\n\n};\n\n"
	return strings.Replace(contents, "{ComponentName}", name, -1)
}

func createReactNativeComponent(name string) {
	validate(name)

	fmt.Println("Generating component: ", name)

	contents := generateComponent(name)
	generateFile(contents, name, "js")
}

func generateStatelessComponent(name string) string {
	contents := "import React from 'react';\n\nimport {\n\n} from 'react-native';\n\nconst {ComponentName} = props => (\n\n);\n\n{ComponentName}.propTypes = {\n\n};\n\nexport default {ComponentName};\n"
	return strings.Replace(contents, "{ComponentName}", name, -1)
}

func createReactNativeStatelessComponent(name string) {
	validate(name)

	fmt.Println("Generating stateless component: ", name)

	contents := generateStatelessComponent(name)
	generateFile(contents, name, "js")
}

// General
func generateFile(content string, name string, extension string) {
	fileContent := []byte(content)
	filePath := fmt.Sprintf("%v.%v", name, extension)
	err := ioutil.WriteFile(filePath, fileContent, 0644)

	if (err != nil) {
		panic(err)
	}
}

func validate(input string) {
	if (input == "") {
		panic("invalid input")
	}
}

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name: "stateless",
			Aliases: []string{"s"},
			Usage: "create a stateless component",
			Action: func(c *cli.Context) error {
				name := c.Args().First()
				createReactNativeStatelessComponent(name)
				return nil
			},
		},
		{
			Name: "component",
			Aliases: []string{"c"},
			Usage: "create a component based on react-native's component",
			Action: func(c *cli.Context) error {
				name := c.Args().First()
				createReactNativeComponent(name)
				return nil
			},
		},
		{
			Name: "viewcontroller",
			Aliases: []string{"v"},
			Usage: "create a view controler based on uikit UIViewController",
			Action: func(c *cli.Context) error {
				name := c.Args().First()
				createViewController(name)
				return nil
			},
		},
	}

	app.Run(os.Args)
}

