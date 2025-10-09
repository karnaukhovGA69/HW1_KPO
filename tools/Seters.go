package tools

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetValidFood(app *tview.Application, menu tview.Primitive, callback func(int)) {
	var input *tview.InputField
	var pages *tview.Pages

	input = tview.NewInputField().
		SetLabel("Количество еды: ").
		SetFieldWidth(20).
		SetAcceptanceFunc(tview.InputFieldInteger)

	input.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			text := input.GetText()

			if text == "" {
				showError(app, pages, "Пустой ввод! Введите число.", input)
				return
			}

			value, err := strconv.Atoi(text)
			if err != nil {
				showError(app, pages, "Введите корректное число!", input)
				return
			}

			if value < 0 {
				showError(app, pages, "Число не может быть отрицательным!", input)
				return
			}

			// Успешный ввод
			callback(value)
			app.SetRoot(menu, true)

		} else if key == tcell.KeyEscape {
			app.SetRoot(menu, true)
		}
	})

	inputBox := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(input, 1, 0, true)

	inputBox.SetBorder(true).
		SetTitle("Ввод данных (ESC - отмена)").
		SetTitleAlign(tview.AlignCenter)

	pages = tview.NewPages().
		AddPage("input", inputBox, true, true)

	app.SetRoot(pages, true).SetFocus(input)
}

// Вспомогательная функция для показа ошибки
func showError(app *tview.Application, pages *tview.Pages, message string, returnToInput *tview.InputField) {
	modal := tview.NewModal().
		SetText("❌ " + message).
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.RemovePage("error")
			returnToInput.SetText("")
			app.SetFocus(returnToInput)
		})

	pages.AddPage("error", modal, true, true)
	app.SetFocus(modal)
}
