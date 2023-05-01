/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/antoniomedeiros1/anki-notion-syncer/pkg/notionclient"
	"github.com/spf13/cobra"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Syncronizes Notion with Anki.",
	Long:  `Receives a path to a page in Notion wich has the table Flashcards and updates Anki deck.`,
	Run: func(cmd *cobra.Command, args []string) {

		pageID := args[0]
		// deckName := args[1]

		notion := notionclient.NotionAPI{
			BaseURL: "https://api.notion.com",
			APIKey:  os.Getenv("NOTION_KEY"),
		}

		res, err := notion.GetBlockChildren(pageID)
		if err != nil {
			fmt.Println(err)
		} else {
			if res["results"] != nil {

				results, ok := res["results"].([]interface{})
				if !ok {
					panic("Unexpected type for results field")
				}

				for _, block := range results {
					blockMap, ok := block.(map[string]interface{})
					if !ok {
						panic("Unexpected type for block")
					}
					blockType, ok := blockMap["type"].(string)
					if !ok {
						panic("Unexpected type for object field")
					}
					fmt.Println(blockType)
				}

			} else {
				fmt.Println(nil)
			}
		}

		// iterate over the content blocks and print them out
		// for _, block := range content {
		// 	blockMap, ok := block.(map[string]interface{})
		// 	if !ok {
		// 		panic("Unexpected type for block")
		// 	}
		// 	objectType, ok := blockMap["object"].(string)
		// 	if !ok {
		// 		panic("Unexpected type for object field")
		// 	}
		// 	switch objectType {
		// 	case "paragraph":
		// 		text := blockMap["paragraph"].(map[string]interface{})["text"].([]interface{})[0].(map[string]interface{})
		// 		content := text["content"].(string)
		// 		fmt.Printf("Paragraph: %s\n", content)
		// 	case "heading_1":
		// 		text := blockMap["heading_1"].(map[string]interface{})["text"].([]interface{})[0].(map[string]interface{})
		// 		content := text["content"].(string)
		// 		fmt.Printf("Heading 1: %s\n", content)
		// 	case "table":
		// 		// handle table blocks
		// 		fmt.Println("Table:")
		// 	default:
		// 		fmt.Printf("Unknown block type: %s\n", objectType)
		// 	}
		// }
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// syncCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
