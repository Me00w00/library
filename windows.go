package main

import (
	"fmt"
	"net/url"

	"fyne.io/fyne"
	"fyne.io/fyne/container"

	"fyne.io/fyne/widget"

	_ "modernc.org/sqlite"
)

// Окно регистрации
func window_show_Reg(app fyne.App, parent_window fyne.Window) fyne.Window {
	// Создаем новое окно для регистрации
	regWindow := app.NewWindow("Регистрация")
	regWindow.Resize(fyne.NewSize(400, 320))
	regWindow.Hide()
	win := widget.NewLabel("Регистрация")
	reg_lb := widget.NewLabel("")

	// Поля для ввода логина и пароля
	reglogin := widget.NewEntry()
	reglogin.SetPlaceHolder("login")

	regPass := widget.NewEntry()
	regPass.SetPlaceHolder("password")
	regPass.Password = true

	// Поле ввода email
	email := widget.NewEntry()
	email.SetPlaceHolder("e-mail")

	// Кнопка отправки данных регистрации
	regBtn := widget.NewButton("Зарегистрироваться", func() {
		result, txt := db_set_user(reglogin.Text, regPass.Text, email.Text)
		if result {
			regWindow.Hide()
			parent_window.Show()
		} else {
			reg_lb.SetText(txt)
		}

	})

	// Устанавливаем содержимое окна регистрации
	regWindow.SetContent(container.NewVBox(
		win,
		reglogin,
		regPass,
		email,
		regBtn,
		reg_lb,
		widget.NewButton("Назад", func() {
			regWindow.Hide()
			parent_window.Show()
		}),
	))
	return regWindow
}

/*ГЛАВНОЕ ОКНО*/
func window_show_Base(app fyne.App, login string) fyne.Window {
	books_name := []string{""}
	basewin := app.NewWindow("Library.ru")
	basewin.Resize(fyne.NewSize(520, 520))

	//basewin.Canvas().Size().Width
	//basewin.Content().Size().Width
	list := widget.NewList(
		func() int {
			return len(books_name)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("test")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(books_name[i])
		})
	//открываем окно с информацией о книге
	list.OnSelected = func(id widget.ListItemID) {
		information := window_show_book_info(app, basewin, books_name[id], login)
		information.Show()
		//basewin.Hide()
		//fmt.Println(books_name[id])
	}
	//basewin.Canvas().SetBackgroundColor(fyne.NewColor(0, 128, 255))
	search := widget.NewEntry()
	search.SetPlaceHolder("Поиск. Название или описание.")
	btn_sea := widget.NewButton("Поиск", func() {
		_, books_name = db_book_search(search.Text)

		list.Resize(fyne.NewSize(400, 600))
		list.MinSize()
		list.Refresh()
		//res.SetText(book_name)
	})

	new_book := window_show_new_book(app, basewin)
	btn_set_book := widget.NewButton("Добавить книгу", func() {
		new_book.Show()
		basewin.Hide()
		//открываем окно добавления книги
	})
	//показыввем кнопку только админам
	if login == "admin" {
		btn_set_book.Show()
	} else {
		btn_set_book.Hide()
	}
	office := window_show_office(app, basewin, login)
	btn_off := widget.NewButton("Личный кабинет", func() {
		office.Show()
		basewin.Hide()
	})

	//создаем форму с результатом поиска
	/*label1 := widget.NewLabel("Label 1")
	value1 := widget.NewLabel("Value")
	label2 := widget.NewLabel("Label 2")
	value2 := widget.NewLabel("Something")
	//grid := container.NewGridWrap(layout.NewFormLayout().Layout(), label1, value1, label2, value2)
	basewin.SetContent(grid)*/

	// Устанавливаем содержимое главного окна
	basewin.SetContent(container.NewVBox(
		search,
		btn_sea,
		btn_set_book,
		btn_off,
		list,
	))
	return basewin
}

// Окно добавления книги
func window_show_new_book(app fyne.App, window_show_Base fyne.Window) fyne.Window {
	new_book := app.NewWindow("Добавление")
	new_book.Hide()
	new_book.Resize(fyne.NewSize(400, 320))
	win := widget.NewLabel("Новая книга")
	book_lb := widget.NewLabel("")

	name_book := widget.NewEntry()
	name_book.SetPlaceHolder("Название")

	author := widget.NewEntry()
	author.SetPlaceHolder("Автор")

	br_content := widget.NewMultiLineEntry()
	br_content.SetPlaceHolder("Описание")
	br_content.Wrapping = fyne.TextWrapBreak

	date_release := widget.NewEntry()
	date_release.SetPlaceHolder("Дата релиза")

	link := widget.NewEntry()
	link.SetPlaceHolder("Ссылка")

	//link.SetPlaceHolder("ссылка")
	//link.Wrapping = fyne.TextWrapBreak

	// Кнопка отправки данных
	bookBtn := widget.NewButton("Добавить", func() {
		result, txt := db_set_book(name_book.Text, author.Text, br_content.Text, date_release.Text, link.Text)
		if result {
			new_book.Hide()
			window_show_Base.Show()
		} else {
			book_lb.SetText(txt)
		}

	})

	new_book.SetContent(container.NewVBox(
		win,
		name_book,
		author,
		br_content,
		date_release,
		link,
		bookBtn,
		widget.NewButton("Назад", func() {
			new_book.Hide()
			window_show_Base.Show()

		}),
		book_lb,
	))
	return new_book
}

/*
Окно личного кабинета
*/
func window_show_office(app fyne.App, window_show_Base fyne.Window, login string) fyne.Window {
	office := app.NewWindow("Личный кабинет")
	office.Resize(fyne.NewSize(400, 320))
	personal_account := widget.NewLabel("Личный кабинет")
	us_lb := widget.NewLabel("")
	name := widget.NewEntry()
	name.SetPlaceHolder("Имя")
	surname := widget.NewEntry()
	surname.SetPlaceHolder("Фамилия")
	name.SetText(name.Text)
	surname.SetText(surname.Text)
	save_btn := widget.NewButton("Сохранить", func() {
		result, txt := db_save_user_office(login, name.Text, surname.Text)
		if result {
			office.Hide()
			window_show_Base.Show()
		} else {
			us_lb.SetText(txt)
		}
	})
	office.SetContent(container.NewVBox(
		personal_account,
		name,
		surname,
		save_btn,
		us_lb,
		widget.NewButton("Назад", func() {
			office.Hide()
			window_show_Base.Show()
		}),
	))
	return office
}

/*про книгу*/
func window_show_book_info(app fyne.App, window_show_Base fyne.Window, book_name, login string) fyne.Window {
	information := app.NewWindow("Информация")

	//и щем книгу
	_, books := db_get_book_info(book_name)
	////////

	information.Resize(fyne.NewSize(400, 320))
	name := widget.NewLabel("")
	author := widget.NewLabel("")
	date_release := widget.NewLabel("")
	br_content := widget.NewLabel("")
	url, _ := url.ParseRequestURI("")
	link := widget.NewHyperlink("Ссылка на произведение", url)
	if len(books) == 0 {
		name.SetText("Не найдено")
		author.SetText("")
		date_release.SetText("")
		br_content.SetText("")
		link.SetText("")
	} else {
		name.SetText("Название: " + books[0].name)
		author.SetText("Автор: " + books[0].author)
		date_release.SetText("Дата релиза: " + books[0].date_release)
		br_content.SetText("Краткое содержание: " + books[0].br_content)
		link.SetText("Ссылка: " + books[0].link)
	}

	//link.SetText(link.Text)
	/////////////

	//ОТЗЫВЫ
	reviews := widget.NewLabel("Отзывы:")
	//подгружаем отзывы из базы

	//в цикле для каждого найденного отзыва выводим его в общее поле
	all_reviw_str := ""
	_, rev_arr := db_get_review(book_name)

	fmt.Println("книга '" + book_name + "'")
	fmt.Println("отзыв", rev_arr)

	for _, val := range rev_arr {
		all_reviw_str += val.user_name + ": " + val.text + "\n"
	}

	sep := widget.NewSeparator()
	edt_leave_review := widget.NewMultiLineEntry()
	edt_leave_review.SetText(all_reviw_str)
	write := window_show_write_review(app, book_name, login)
	btn_leave_review := widget.NewButton("Оставить отзыв", func() {
		write.Show()
	})

	information.SetContent(container.NewVBox(
		name,
		author,
		date_release,
		br_content,
		link,
		reviews,
		sep,
		edt_leave_review,
		btn_leave_review,
		widget.NewButton("Назад", func() {
			information.Hide()
			window_show_Base.Show()
		}),
	))
	return information
}

// Окно написания отзыва
func window_show_write_review(app fyne.App, book_name, login string) fyne.Window {
	write := app.NewWindow("Оставить отзыв")
	write.Resize(fyne.NewSize(400, 320))
	us_lb := widget.NewLabel("")
	name_surname := widget.NewLabel("Имя пользователя: " + login)
	text := widget.NewMultiLineEntry()
	text.SetPlaceHolder("Ваш отзыв...")
	text.Wrapping = fyne.TextWrapBreak
	send_btn := widget.NewButton("Отправить", func() {
		result, txt := db_write_a_review(book_name, login, text.Text)
		if result {
			write.Hide()
		} else {
			us_lb.SetText(txt)
		}
	})
	write.SetContent(container.NewVBox(
		name_surname,
		text,
		send_btn,
		us_lb,
		widget.NewButton("Назад", func() {
			write.Hide()
		}),
	))
	return write
}
