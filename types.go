package main

// Response is the type of the response from the express api
type Response []struct {
	ID                  string `json:"id"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	DishID              string `json:"dishID"`
	MealProvidingUnitID string `json:"mealProvidingUnitID"`
	MealProvidingUnit   struct {
		ID                    string  `json:"id"`
		MealProvidingUnitName string  `json:"mealProvidingUnitName"`
		OrganizationID        string  `json:"organizationID"`
		ShowFoods             bool    `json:"showFoods"`
		ShowArticles          bool    `json:"showArticles"`
		Longitude             float64 `json:"longitude"`
		Latitude              float64 `json:"latitude"`
		SevenDayWeek          bool    `json:"sevenDayWeek"`
		DisplayNameCategories []struct {
			DisplayNameCategoryID   string `json:"displayNameCategoryID"`
			DisplayNameCategoryName string `json:"displayNameCategoryName"`
			SortOrder               int    `json:"sortOrder"`
			MealProvidingUnitID     string `json:"mealProvidingUnitID"`
		} `json:"displayNameCategories"`
		DishTypes []struct {
			ID                  string  `json:"id"`
			DishTypeName        string  `json:"dishTypeName"`
			DishTypeNameEnglish string  `json:"dishTypeNameEnglish"`
			IsEnabled           bool    `json:"isEnabled"`
			Price               float64 `json:"price"`
			HasPrice            bool    `json:"hasPrice"`
		} `json:"dishTypes"`
		Settings struct {
			ID                     string      `json:"id"`
			HasLibrary             bool        `json:"hasLibrary"`
			HasNyckelhalsLabel     bool        `json:"hasNyckelhalsLabel"`
			HasWwfApprovedLabel    bool        `json:"hasWwfApprovedLabel"`
			HasPricesLabel         bool        `json:"hasPricesLabel"`
			DayMenuSettings        interface{} `json:"dayMenuSettings"`
			WeekMenuSettings       interface{} `json:"weekMenuSettings"`
			IndividualMenuSettings interface{} `json:"individualMenuSettings"`
		} `json:"settings"`
	} `json:"mealProvidingUnit"`
	DisplayNames []struct {
		DishID                interface{} `json:"dishID"`
		ID                    string      `json:"id"`
		DisplayNameCategoryID string      `json:"displayNameCategoryID"`
		DishDisplayName       string      `json:"dishDisplayName"`
		DishOccurrenceID      string      `json:"dishOccurrenceID"`
		DisplayNameCategory   struct {
			DisplayNameCategoryID   string `json:"displayNameCategoryID"`
			DisplayNameCategoryName string `json:"displayNameCategoryName"`
			SortOrder               int    `json:"sortOrder"`
			MealProvidingUnitID     string `json:"mealProvidingUnitID"`
		} `json:"displayNameCategory"`
	} `json:"displayNames"`
	DishTypeID string `json:"dishTypeID"`
	DishType   struct {
		ID                  string  `json:"id"`
		DishTypeName        string  `json:"dishTypeName"`
		DishTypeNameEnglish string  `json:"dishTypeNameEnglish"`
		IsEnabled           bool    `json:"isEnabled"`
		Price               float64 `json:"price"`
		HasPrice            bool    `json:"hasPrice"`
	} `json:"dishType"`
	Dish struct {
		ID      string `json:"id"`
		Recipes []struct {
			ID         string `json:"id"`
			RecipeName string `json:"recipeName"`
			Allergens  []struct {
				ID           string `json:"id"`
				AllergenCode string `json:"allergenCode"`
				RecipeID     string `json:"recipeID"`
				AllergenURL  string `json:"allergenURL"`
			} `json:"allergens"`
			Ingredients []struct {
				ID                   string  `json:"id"`
				Amount               float64 `json:"amount"`
				FoodID               string  `json:"foodID"`
				RegionID             string  `json:"regionID"`
				ProductionActivityID string  `json:"productionActivityID"`
				IsEcological         bool    `json:"isEcological"`
				MeasurementUnitID    string  `json:"measurementUnitID"`
				RecipeID             string  `json:"recipeID"`
				TotalEmission        float64 `json:"totalEmission"`
			} `json:"ingredients"`
			BasedOn       interface{}   `json:"basedOn"`
			Instructions  interface{}   `json:"instructions"`
			Portions      int           `json:"portions"`
			CookbookIDs   []interface{} `json:"cookbookIDs"`
			TotalEmission float64       `json:"totalEmission"`
		} `json:"recipes"`
		DishName          string        `json:"dishName"`
		DisplayNames      interface{}   `json:"displayNames"`
		DishType          interface{}   `json:"dishType"`
		CookbookIDs       []interface{} `json:"cookbookIDs"`
		EditableByDefault bool          `json:"editableByDefault"`
		TotalEmission     float64       `json:"totalEmission"`
		MainRecipe        interface{}   `json:"mainRecipe"`
		Price             float64       `json:"price"`
		Prices            string        `json:"prices"`
		TotalEmissionURL  string        `json:"totalEmissionURL"`
	} `json:"dish"`
	AvailableDishTypes []struct {
		ID                  string  `json:"id"`
		DishTypeName        string  `json:"dishTypeName"`
		DishTypeNameEnglish string  `json:"dishTypeNameEnglish"`
		IsEnabled           bool    `json:"isEnabled"`
		Price               float64 `json:"price"`
		HasPrice            bool    `json:"hasPrice"`
	} `json:"availableDishTypes"`
	EditableByDefault        bool          `json:"editableByDefault"`
	SortOrder                int           `json:"sortOrder"`
	NumberOfPreparedPortions int           `json:"numberOfPreparedPortions"`
	NumberOfServedPortions   int           `json:"numberOfServedPortions"`
	Allergens                []interface{} `json:"allergens"`
	OverrideAllergens        bool          `json:"overrideAllergens"`
	DishMenuSettings         interface{}   `json:"dishMenuSettings"`
	DayMenuSettings          interface{}   `json:"dayMenuSettings"`
	WeekMenuSettings         interface{}   `json:"weekMenuSettings"`
	Nyckelhalmarkt           bool          `json:"nyckelhalmarkt"`
	WwfApproved              bool          `json:"wwfApproved"`
}
