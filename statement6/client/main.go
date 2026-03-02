package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	pb "github.com/CynthiaAntony/CynthiaInt1/statement6/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	addr := "trie-server-service:50051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	fmt.Println("Connected to server")
	c := pb.NewTrieServiceClient(conn)
	sc := bufio.NewScanner(os.Stdin)
	fmt.Println("---String List Manager---")
	for {
		fmt.Print("Enter choice: 1. Add\n2. Remove\n3. Check\n4. List\n5. Exit: ")
		if !sc.Scan() {
			break
		}
		choice := strings.ToLower(strings.TrimSpace(sc.Text()))

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if choice == "list" {
			res, err := c.List(ctx, &pb.Empty{})
			if err != nil {
				fmt.Printf("Error Listing words: %v\n", err)
			} else {
				fmt.Printf("Words : %v\n", res.GetWords())
			}
			continue
		}
		if choice == "exit" {
			fmt.Println("Exiting...")
			break
		}
		switch choice {
		case "add":
			fmt.Print("Enter word: ")
			if !sc.Scan() {
				break
			}
			word := strings.TrimSpace(sc.Text())
			_, err := c.Add(ctx, &pb.WordRequest{Word: word})
			if err != nil {
				fmt.Printf("Error adding word: %v\n", err)
			} else {
				fmt.Println("Word added")
			}
		case "remove":
			fmt.Print("Enter word: ")
			if !sc.Scan() {
				break
			}
			word := strings.TrimSpace(sc.Text())
			res, err := c.Remove(ctx, &pb.WordRequest{Word: word})
			if err != nil {
				fmt.Printf("Error removing word: %v\n", err)
			} else if res.GetExists() {
				fmt.Println("Word removed")
			} else {
				fmt.Println("Word not found")
			}
		case "check":
			fmt.Print("Enter word: ")
			if !sc.Scan() {
				break
			}
			word := strings.TrimSpace(sc.Text())
			res, err := c.Check(ctx, &pb.WordRequest{Word: word})
			if err != nil {
				fmt.Printf("Error checking word: %v\n", err)
			} else if res.GetExists() {
				fmt.Println("Word Exists")
			} else {
				fmt.Println("Word does not exist")
			}
		default:
			fmt.Println("Invalid Choice")
		}
	}
}
