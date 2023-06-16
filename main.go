package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func main() {
	client := newFirestoreClient()
	defer client.Close()

	http.HandleFunc("/cron", manualTrigger(client))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error starting a server: %v", err)
	}
}

func newFirestoreClient() *firestore.Client {
	ctx := context.Background()
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}

	return client
}

func manualTrigger(client *firestore.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := execJob(r.Context(), client); err != nil {
			w.WriteHeader(500)
			w.Write([]byte(err.Error()))
			fmt.Printf("Failed to execute job: %v\n", err)
			return
		}

		w.WriteHeader(200)
		w.Write([]byte("Success!!!"))
		fmt.Println("Success!!! Yuhu~~~")
	}
}

func execJob(ctx context.Context, client *firestore.Client) error {
	collection := client.Collection("user_scan_count")

	iter := collection.Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return fmt.Errorf("error iterating over documents: %v", err)
		}

		_, err = doc.Ref.Update(ctx, []firestore.Update{
			{Path: "count", Value: 0},
		})

		if err != nil {
			return fmt.Errorf("failed to update document: %v", err)
		}

		fmt.Printf("Document with ID %s updated successfully.\n", doc.Ref.ID)
	}

	return nil
}
