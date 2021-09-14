/*
Copyright © 2021 Dimitri Acosta <info@dimitriacosta.xyz>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "This command will get the desired Gopher",
	Long:  `This command will call GitHub repository in order to return the desired Gopher.`,
	Run: func(cmd *cobra.Command, args []string) {
		var gopherName = "dr-who"

		if len(args) >= 1 && args[0] != "" {
			gopherName = args[0]
		}

		URL := "https://github.com/scraly/gophers/raw/main/" + gopherName + ".png"

		fmt.Printf("Try to get '%s Gopher...\n", gopherName)

		// Get the data
		response, err := http.Get(URL)
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Create the file
			out, err := os.Create(gopherName + ".png")
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			// Write the body to file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Perfect! Just saved in", out.Name(), "!")
		} else {
			log.Fatalf("Error: %s doesn't exists! :-(\n", gopherName)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
