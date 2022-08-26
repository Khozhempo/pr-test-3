package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
	"pr-test-3-client/proto"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())

	client := proto.NewUsersServiceClient(conn)

	args := os.Args
	help := "Need argument: " + os.Args[0] + " a/d/l (add, delete, list) and username (for 'a' and 'd' cmd)\n Example: ... a testUsername"

	if len(args) < 2 || len(args) > 3 {
		log.Fatal(help)
	}

	if len(args) == 2 && os.Args[1] != "l" {
		log.Fatal(help)
	}

	if len(args) == 3 && (os.Args[1] != "a" && os.Args[1] != "d") {
		log.Fatal(help)
	}

	switch args[1] {
	case "a":
		_, err := client.AddUser(context.Background(), &proto.UserInfoRequest{Username: args[2]})
		if err != nil {
			log.Printf("Got error message: %s", err.Error())
		}
	case "d":
		_, err := client.RemoveUser(context.Background(), &proto.UserInfoRequest{Username: args[2]})
		if err != nil {
			log.Printf("Got error message: %s", err.Error())
		}
	case "l":
		AllUsers, err := client.ListAllUsers(context.Background(), &proto.NoArgs{})
		if err != nil {
			log.Printf("Got error message: %s", err.Error())
		}
		log.Println(AllUsers.Usernames)
	}
}
