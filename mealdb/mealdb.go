/*
 * Copyright 2022 XXIV
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package mealdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	httpClient "net/http"
	"net/url"
)

type meals struct {
	Meals []Meal `json:"meals"`
}

type Meal struct {
	DateModified                string		 `json:"dateModified"`
	IDMeal                      string      `json:"idMeal"`
	StrArea                     string      `json:"strArea"`
	StrCategory                 string      `json:"strCategory"`
	StrCreativeCommonsConfirmed string		 `json:"strCreativeCommonsConfirmed"`
	StrDrinkAlternate           string		 `json:"strDrinkAlternate"`
	StrImageSource              string		 `json:"strImageSource"`
	StrIngredient1              string      `json:"strIngredient1"`
	StrIngredient10             string      `json:"strIngredient10"`
	StrIngredient11             string      `json:"strIngredient11"`
	StrIngredient12             string      `json:"strIngredient12"`
	StrIngredient13             string      `json:"strIngredient13"`
	StrIngredient14             string      `json:"strIngredient14"`
	StrIngredient15             string      `json:"strIngredient15"`
	StrIngredient16             string		 `json:"strIngredient16"`
	StrIngredient17             string		 `json:"strIngredient17"`
	StrIngredient18             string		 `json:"strIngredient18"`
	StrIngredient19             string		 `json:"strIngredient19"`
	StrIngredient2              string      `json:"strIngredient2"`
	StrIngredient20             string		 `json:"strIngredient20"`
	StrIngredient3              string      `json:"strIngredient3"`
	StrIngredient4              string      `json:"strIngredient4"`
	StrIngredient5              string      `json:"strIngredient5"`
	StrIngredient6              string      `json:"strIngredient6"`
	StrIngredient7              string      `json:"strIngredient7"`
	StrIngredient8              string      `json:"strIngredient8"`
	StrIngredient9              string      `json:"strIngredient9"`
	StrInstructions             string      `json:"strInstructions"`
	StrMeal                     string      `json:"strMeal"`
	StrMealThumb                string      `json:"strMealThumb"`
	StrMeasure1                 string      `json:"strMeasure1"`
	StrMeasure10                string      `json:"strMeasure10"`
	StrMeasure11                string      `json:"strMeasure11"`
	StrMeasure12                string      `json:"strMeasure12"`
	StrMeasure13                string      `json:"strMeasure13"`
	StrMeasure14                string      `json:"strMeasure14"`
	StrMeasure15                string      `json:"strMeasure15"`
	StrMeasure16                string		 `json:"strMeasure16"`
	StrMeasure17                string		 `json:"strMeasure17"`
	StrMeasure18                string		 `json:"strMeasure18"`
	StrMeasure19                string 		`json:"strMeasure19"`
	StrMeasure2                 string      `json:"strMeasure2"`
	StrMeasure20                string		 `json:"strMeasure20"`
	StrMeasure3                 string      `json:"strMeasure3"`
	StrMeasure4                 string      `json:"strMeasure4"`
	StrMeasure5                 string      `json:"strMeasure5"`
	StrMeasure6                 string      `json:"strMeasure6"`
	StrMeasure7                 string      `json:"strMeasure7"`
	StrMeasure8                 string      `json:"strMeasure8"`
	StrMeasure9                 string      `json:"strMeasure9"`
	StrSource                   string		 `json:"strSource"`
	StrTags                     string      `json:"strTags"`
	StrYoutube                  string      `json:"strYoutube"`
}

type ingredients struct {
	Meals []Ingredient `json:"meals"`
}

type Ingredient struct {
	IDIngredient   string `json:"idIngredient"`
	StrDescription string `json:"strDescription"`
	StrIngredient  string `json:"strIngredient"`
	StrType        string `json:"strType"`
}

type categories struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	IDCategory             string `json:"idCategory"`
	StrCategory            string `json:"strCategory"`
	StrCategoryDescription string `json:"strCategoryDescription"`
	StrCategoryThumb       string `json:"strCategoryThumb"`
}

type filters struct {
	Meals []Filter `json:"meals"`
}

type Filter struct {
	IDMeal       string `json:"idMeal"`
	StrMeal      string `json:"strMeal"`
	StrMealThumb string `json:"strMealThumb"`
}

type categoriesFilter struct {
	Meals []struct {
		StrCategory string `json:"strCategory"`
	} `json:"meals"`
}

type areaFilter struct {
	Meals []struct {
		StrArea string `json:"strArea"`
	} `json:"meals"`
}

func http(endpoint string) (string, error) {
	response, err := httpClient.Get(fmt.Sprintf("https://themealdb.com/api/json/v1/1/%s", endpoint))
	if err != nil {
		return "", errors.New("ERROR")
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("ERROR")
	}
	return string(body), nil
}

// Search meal by name
func Search(s string) ([]Meal, error) {
	response, err := http(fmt.Sprintf("search.php?s=%s",url.QueryEscape(s)))
	if err != nil {
		return []Meal{}, newError("error")
	}
	if len(response) == 0 {
		return []Meal{}, newError("error")
	}
	var meal meals
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []Meal{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []Meal{}, newError("error")
	}
	var responseSlice []Meal
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// SearchByLetter search meals by first letter
func SearchByLetter(b byte) ([]Meal, error) {
	response, err := http(fmt.Sprintf("search.php?f=%c", b))
	if err != nil {
		return []Meal{}, newError("error")
	}
	if len(response) == 0 {
		return []Meal{}, newError("error")
	}
	var meal meals
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []Meal{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []Meal{}, newError("error")
	}
	var responseSlice []Meal
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// SearchByID search meal details by id
func SearchByID(i int64) (Meal, error) {
	response, err := http(fmt.Sprintf("lookup.php?i=%d", i))
	if err != nil {
		return Meal{}, newError("error")
	}
	if len(response) == 0 {
		return Meal{}, newError("error")
	}
	var meal meals
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return Meal{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return Meal{}, newError("error")
	}
	ingredient := meal.Meals[0]
	return ingredient, nil
}

// Random meal
func Random() (Meal, error) {
	response, err := http("random.php")
	if err != nil {
		return Meal{}, newError("error")
	}
	if len(response) == 0 {
		return Meal{}, newError("error")
	}
	var meal meals
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return Meal{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return Meal{}, newError("error")
	}
	ingredient := meal.Meals[0]
	return ingredient, nil
}

// MealCategories List the meals categories
func MealCategories() ([]Category, error) {
	response, err := http("categories.php")
	if err != nil {
		return []Category{}, newError("error")
	}
	if len(response) == 0 {
		return []Category{}, newError("error")
	}
	var meal categories
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []Category{}, newError("error")
	}
	if len(meal.Categories) == 0 {
		return []Category{}, newError("error")
	}
	var responseSlice []Category
	for _, v := range meal.Categories {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// FilterByIngredient filter by ingredient
func FilterByIngredient(s string) ([]Filter, error) {
	response, err := http(fmt.Sprintf("filter.php?i=%s",url.QueryEscape(s)))
	if err != nil {
		return []Filter{}, newError("error")
	}
	if len(response) == 0 {
		return []Filter{}, newError("error")
	}
	var meal filters
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []Filter{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []Filter{}, newError("error")
	}
	var responseSlice []Filter
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// FilterByArea filter by area
func FilterByArea(s string) ([]Filter, error) {
	response, err := http(fmt.Sprintf("filter.php?a=%s",url.QueryEscape(s)))
	if err != nil {
		return []Filter{}, newError("error")
	}
	if len(response) == 0 {
		return []Filter{}, newError("error")
	}
	var meal filters
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []Filter{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []Filter{}, newError("error")
	}
	var responseSlice []Filter
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// FilterByCategory filter by category
func FilterByCategory(s string) ([]Filter, error) {
	response, err := http(fmt.Sprintf("filter.php?c=%s",url.QueryEscape(s)))
	if err != nil {
		return []Filter{}, newError("error")
	}
	if len(response) == 0 {
		return []Filter{}, newError("error")
	}
	var meal filters
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []Filter{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []Filter{}, newError("error")
	}
	var responseSlice []Filter
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// CategoriesFilter List the categories filter
func CategoriesFilter() ([]string, error) {
	response, err := http("list.php?c=list")
	if err != nil {
		return []string{}, newError("error")
	}
	if len(response) == 0 {
		return []string{}, newError("error")
	}
	var meal categoriesFilter
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []string{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []string{}, newError("error")
	}
	var responseSlice []string
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v.StrCategory)
	}
	return responseSlice, nil
}

// IngredientsFilter List the ingredients filter
func IngredientsFilter() ([]Ingredient, error) {
	response, err := http("list.php?i=list")
	if err != nil {
		return []Ingredient{}, newError("error")
	}
	if len(response) == 0 {
		return []Ingredient{}, newError("error")
	}
	var meal ingredients
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []Ingredient{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []Ingredient{}, newError("error")
	}
	var responseSlice []Ingredient
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v)
	}
	return responseSlice, nil
}

// AreaFilter List the area filter
func AreaFilter() ([]string, error) {
	response, err := http("list.php?a=list")
	if err != nil {
		return []string{}, newError("error")
	}
	if len(response) == 0 {
		return []string{}, newError("error")
	}
	var meal areaFilter
	jsonError := json.Unmarshal([]byte(response), &meal)
	if jsonError != nil {
		return []string{}, newError("error")
	}
	if len(meal.Meals) == 0 {
		return []string{}, newError("error")
	}
	var responseSlice []string
	for _, v := range meal.Meals {
		responseSlice= append(responseSlice,v.StrArea)
	}
	return responseSlice, nil
}