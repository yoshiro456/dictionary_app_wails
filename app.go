package main

import (
	"context"
	"dictionary_app/models"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	// Setup Database
	models.Connect()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) EMFindById(id uint) *models.EMDictionary {
	m := models.EMDictionary{}
	return m.FindById(id)
}

func (a *App) EMFindManyByWord(word string) *[]models.EMDictionary {
	d := models.EMDictionary{}
	e := d.FindManyByWord(word)

	return e
}
