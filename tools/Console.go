package tools

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var items []MenuItem

func WriteConsole() {
	if items == nil {
		items = ReadFile("ru")
	}
	app := tview.NewApplication()
	i := 1
	menu := tview.NewList()
	for _, item := range items {
		displayText := fmt.Sprintf("%d) %s", i, item.Text)
		menu.AddItem(displayText, "", 0, func() { SearchFunction(item.Key, app, menu) })
		i++
	}
	menu.AddItem(fmt.Sprintf("%d) [red::bu]ВЫХОД ❌", len(items)+1), "", 0, func() {
		app.Stop()
	})
	menu.SetBorder(true).SetTitle("================ZOO================").SetTitleAlign(tview.AlignCenter)
	if err := app.SetRoot(menu, true).Run(); err != nil {
		panic(err)
	}
}

func SearchFunction(keyFunction string, app *tview.Application, menu *tview.List) {
	switch keyFunction {
	case "add":
		// TODO: функция для добавления животного
	case "allAnimal":
		// TODO: функция для отображения всех животных
	case "handAnimal":
		// TODO: функция для отображения контактных животных
	case "allZoo":
		// TODO: функция для базы данных зоопарка
	case "changeLanguage":
		changeLanguage(app, menu)
	default:
		// TODO: обработка неизвестного ключа
	}
}

func changeLanguage(app *tview.Application, menu *tview.List) {
	// создаём поле для ввода языка
	input := tview.NewInputField().
		SetLabel("Введите язык ('ru', 'en', 'ch'): ").
		SetFieldWidth(10)

	input.SetDoneFunc(func(key tcell.Key) {
		language := input.GetText()
		items = ReadFile(language)

		// очищаем старое меню
		menu.Clear()
		for i, item := range items {
			displayText := fmt.Sprintf("%d) %s", i+1, item.Text)
			// нужно зафиксировать переменную в лямбде
			currentKey := item.Key
			menu.AddItem(displayText, "", 0, func() { SearchFunction(currentKey, app, menu) })
		}
		menu.AddItem(fmt.Sprintf("%d) [red::bu]ВЫХОД ❌", len(items)+1), "", 0, func() {
			app.Stop()
		})

		// возвращаемся обратно к меню
		app.SetRoot(menu, true).SetFocus(menu)
	})

	// показываем поле ввода вместо меню
	app.SetRoot(input, true).SetFocus(input)
}
