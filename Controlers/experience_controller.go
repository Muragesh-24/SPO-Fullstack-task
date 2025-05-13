package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	config "spoBackend/Config"
	model "spoBackend/Models"
	"time"

)

func SaveExperience(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var experience model.Experience
    if err := json.NewDecoder(r.Body).Decode(&experience); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := config.GetCollection("experiences")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _, err := collection.InsertOne(ctx, experience)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]string{"message": "Form submitted successfully"})
}

func GetAllExperiences(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    collection := config.GetCollection("experiences")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := collection.Find(ctx, map[string]interface{}{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(ctx)

    var experiences []model.Experience
    for cursor.Next(ctx) {
        var experience model.Experience
        if err := cursor.Decode(&experience); err != nil {
            log.Println("Decode error:", err)
            continue
        }
        experiences = append(experiences, experience) 
    }

    json.NewEncoder(w).Encode(experiences)
}
func GetStats(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    collection := config.GetCollection("experiences")
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    count, err := collection.CountDocuments(ctx, map[string]interface{}{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(map[string]int64{"total": count})
}
