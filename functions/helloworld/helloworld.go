package helloworld

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
)

type (
	SystemSettings struct {
		MaintenanceMode MyMaintenanceMode `firestore:"maintenance_mode"`
	}
	MyMaintenanceMode struct {
		IsMaintenance bool `firestore:"is_maintenance"`
	}
)

func HelloHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	project := "mindpowerbank-dev"
	client, err := firestore.NewClient(ctx, project)
	defer client.Close()
	if err != nil {
		log.Fatal("Error ", err)
	}
	dsnap, err := client.Collection("system_settings").Doc("maintenance_mode").Get(ctx)
	if err != nil {
		log.Fatal("err", err)
	}
	m := dsnap.Data()
	log.Printf("Document data: %#v\n", m)
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, World!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
