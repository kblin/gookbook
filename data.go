package main

import "encoding/json"

type Ingredient struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

type _Ingredient Ingredient

func (i *Ingredient) UnmarshalJSON(b []byte) (err error) {
	_i, n, a, u := _Ingredient{}, "", 0, ""
	if err = json.Unmarshal(b, &_i); err == nil {
		*i = Ingredient(_i)
		return
	}
	if err = json.Unmarshal(b, &n); err == nil {
		i.Name = n
		return
	}
	if err = json.Unmarshal(b, &a); err == nil {
		i.Amount = a
		return
	}
	if err = json.Unmarshal(b, &u); err == nil {
		i.Unit = u
		return
	}
	return
}

type Recipe struct {
	Name        string       `json:"name"`
	Ingredients []Ingredient `json:"ingredients"`
}

type _Recipe Recipe

func (r *Recipe) UnmarshalJSON(b []byte) (err error) {
	_r, s := _Recipe{}, ""
	if err = json.Unmarshal(b, &_r); err == nil {
		*r = Recipe(_r)
		return
	}
	if err = json.Unmarshal(b, &s); err == nil {
		r.Name = s
		return
	}
	return
}

type RecipeListEntry struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type RecipeListing struct {
	Recipes []RecipeListEntry `json:"recipes"`
}

func (r *RecipeListing) Init(recipes []Recipe) (err error) {
	for idx, recipe := range recipes {
		r.Recipes = append(r.Recipes, RecipeListEntry{idx, recipe.Name})
	}
	return
}
