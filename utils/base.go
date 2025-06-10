package utils

import (
	"fmt"
	"io"

	"github.com/cloudwego/eino/schema"
)

// print the stream reader
// func StreamPrint[T any](sr *schema.StreamReader[T]) error {
func StreamPrint(sr *schema.StreamReader[*schema.Message]) error {
	// Print the response
	fmt.Print("🤖 : ")
	for {
		msg, err := sr.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error receiving message: %v\n", err)
			break
		}
		fmt.Print(msg.Content)
	}
	fmt.Println()
	fmt.Println()
	return nil
}
