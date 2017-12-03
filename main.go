package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Recipes []Recipe

var recipes Recipes

func loadRecipes() Recipes {
	raw, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err.Error)
		os.Exit(1)
	}
	var new_recipes Recipes
	json.Unmarshal(raw, &new_recipes)
	fmt.Printf("%v\n", new_recipes)
	return new_recipes
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func listRecipes(w http.ResponseWriter, r *http.Request) {
	recipe_listing := RecipeListing{}
	recipe_listing.Init(recipes)
	json.NewEncoder(w).Encode(recipe_listing)
}

func showRecipe(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idx, err := strconv.Atoi(vars["id"])
	if err != nil || idx > len(recipes) {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Invalid identifier")
		return
	}
	json.NewEncoder(w).Encode(recipes[idx])
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/recipes", listRecipes)
	router.HandleFunc("/recipes/{id:[0-9]+}", showRecipe)

	recipes = loadRecipes()

	if err := http.ListenAndServe(":2665", router); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
