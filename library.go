package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"

	//"fyne.io/fyne/v2"
	"fyne.io/fyne/widget"

	_ "modernc.org/sqlite"
)

func main() {
	a := app.New()                     // Создает новое приложение
	loginWindow := a.NewWindow("Вход") // Заголовок окна
	loginWindow.Resize(fyne.NewSize(400, 320))
	tx := widget.NewLabel("Вход") // Текст, который будет написан
	lb := widget.NewLabel("")     // Метка для сообщений

	// Поле ввода логина
	login := widget.NewEntry()
	login.SetPlaceHolder("login") // Текст, подсказывающий, что вводить

	// Поле ввода пароля
	pass := widget.NewEntry()
	pass.SetPlaceHolder("password")
	pass.Password = true
	/* TEST*/
	login.SetText("admin")
	pass.SetText("1234")
	/***********************/

	regWindow := window_show_Reg(a, loginWindow)
	// Кнопка для регистрации
	btn_reg := widget.NewButton("Registration", func() {

		//выводи глав окнa
		regWindow.Show()
		loginWindow.Hide()
		//login.SetText(text)
		//pass.SetText("")

	})

	// Кнопка для авторизации
	btn_login := widget.NewButton("Login", func() {
		rst, rst_str := db_get_user(login.Text, pass.Text)
		if rst {
			//выводи глав окнa
			basewindow := window_show_Base(a, login.Text)
			loginWindow.Hide()
			basewindow.Show()
		} else {
			lb.SetText(rst_str)
		}
		//w.Hide()

	})

	// Устанавливаем содержимое основного окна
	loginWindow.SetContent(container.NewVBox(
		tx,
		login,
		pass,
		btn_login,
		btn_reg,
		lb,
	))

	loginWindow.ShowAndRun()
}

// к кнопке логин
// подключаемся к бд

//loginWindow.Hide()

/*


go env -w GO111MODULE=auto
go mod init main
go mod tidy

инструкция:  https://github.com/fyne-io/fyne
go get fyne.io/fyne/v2@latest


поставить GCC https://winlibs.com/#download-release
добавить в переменную системы PATH ....\mingw64\bin
go env -w "CGO_ENABLED=1


*/
