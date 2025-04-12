package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

// создаем структуру чтобы коректно вытаскивать поля из базы данных
type book_s struct {
	id           string
	name         string
	author       string
	date_release string
	br_content   string
	link         string
}
type reviw_s struct {
	id        string
	user_name string
	text      string
}

// подключаем к баззе данных чтобы искать книги
func db_book_search(search_str string) (bool, []string) {
	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	if err != nil {
		return false, []string{err.Error()}
	}
	//делаем запрос, чтобы нашел юзера(с логином и паролем) и если не находит то выводить ошибку
	//query
	rows, err := db.Query("SELECT name FROM book WHERE LOWER(name)  LIKE LOWER(\"%" + search_str + "%\") OR LOWER(br_content) LIKE LOWER(\"%" + search_str + "%\");")
	if err != nil {
		return false, []string{err.Error()}
	}

	//работа с результатом запрса из бд
	result := []string{}
	res_rmp := ""
	for rows.Next() {
		if err = rows.Scan(&res_rmp); err != nil {
			return false, []string{err.Error()}
		}
		result = append(result, res_rmp)
	}

	if err = rows.Err(); err != nil {
		return false, []string{err.Error()}
	}

	if err = db.Close(); err != nil {
		return false, []string{err.Error()}
	}

	if len(result) == 0 {
		return false, []string{"Книга не найдена"}
	} else {
		return true, result
	}
}

// подключаемся к базе чтобы админ добавлял книги
func db_set_book(name_book, author, br_content, date_release, link string) (bool, string) {
	if len(name_book) < 2 {
		return false, "Название слишком короткое"
	}
	if len(author) < 2 {
		return false, "Напишите автора"
	}
	if len(br_content) < 2 {
		return false, "Напишите хотя бы пару строк для описания"
	}
	if len(date_release) < 10 {
		return false, "Напишите дату релиза"
	}
	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	if err != nil {
		return false, err.Error()
	}
	//делаем чтобы добавляло данные в базу, потом выводило что все ок, а если что то не так то выводило ошибку
	if _, err = db.Exec("INSERT INTO book (name, author, br_content, date_release, link) VALUES ('" + name_book + "', '" + author + "', '" + br_content + "', '" + date_release + "', '" + link + "');"); err != nil {
		return false, err.Error()
	}
	if err = db.Close(); err != nil {
		return false, err.Error()
	}
	return true, name_book
}

// Авторизация
func db_get_user(login, pass string) (bool, string) {
	if len(login) < 4 || len(pass) < 6 {
		return false, "Login, password to short."
	}
	pass_hash := GetMD5Hash(pass)
	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	if err != nil {
		return false, err.Error()
	}
	//делаем запрос, чтобы нашел юзера(с логином и паролем) и если не находит то выводить ошибку
	//query
	rows, err := db.Query("select login from user where login = '" + login + "'  and password = '" + pass_hash + "';")
	if err != nil {
		return false, err.Error()
	}
	result := ""
	//перебираем результат из БД
	for rows.Next() {
		if err = rows.Scan(&result); err != nil {
			return false, err.Error()
		}

	}

	if err = rows.Err(); err != nil {
		return false, err.Error()
	}

	if err = db.Close(); err != nil {
		return false, err.Error()
	}

	if result == "" {
		return false, "User not found."
	} else {
		return true, "Users are logined."
	}
}

// хеширование. берет строку(пароль), генерирует хеш и возвращает ее(строку) или ошибку
// возвращаем хеш
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// к кнопке регистрации
// подключаемся к базе данных
func db_set_user(login, pass, email string) (bool, string) {
	if len(login) <= 4 || len(pass) <= 6 || len(email) <= 7 {
		return false, "Login, password, email to short."
	}

	pass_hash := GetMD5Hash(pass)

	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	if err != nil {
		return false, err.Error()
	}
	//делаем чтобы добавляло данные в базу, потом выводило что все ок, а если что то не так то выводило ошибку
	if _, err = db.Exec("INSERT INTO user (login, password, email) VALUES ('" + login + "',  '" + pass_hash + "', '" + email + "');"); err != nil {
		return false, err.Error()
	}
	if err = db.Close(); err != nil {
		return false, err.Error()
	}
	return true, login
}

// Личный кабинет
func db_save_user_office(login, name, surname string) (bool, string) {

	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	if err != nil {
		return false, err.Error()
	}
	//делаем чтобы добавляло данные в базу, потом выводило что все ок, а если что то не так то выводило ошибку
	if _, err = db.Exec("UPDATE user SET name = '" + name + "', surname = '" + surname + "' WHERE login = '" + login + "';"); err != nil {
		return false, err.Error()
	}
	if err = db.Close(); err != nil {
		return false, err.Error()
	}
	return true, name
}

// инфо книги
func db_get_book_info(name string) (bool, []book_s) {

	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	result := []book_s{}
	if err != nil {
		return false, result
	}
	//делаем чтобы добавляло данные в базу, потом выводило что все ок, а если что то не так то выводило ошибку
	rows, err := db.Query("select id, name, author, date_release, br_content, link from book where name = '" + name + "';")
	if err != nil {
		return false, result
	}

	//перебираем результат из БД
	for rows.Next() {
		book := book_s{}
		if err = rows.Scan(&book.id, &book.name, &book.author, &book.date_release, &book.br_content, &book.link); err != nil {
			return false, result
		}
		result = append(result, book)
	}

	if err = rows.Err(); err != nil {
		return false, result
	}

	if err = db.Close(); err != nil {
		return false, result
	}

	return true, result

}

// Отзывы
func db_write_a_review(book, user, text string) (bool, string) {
	if len(text) < 2 {
		return false, "Напишите хотя бы одно слово:()"
	}

	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	if err != nil {
		return false, err.Error()
	}
	//делаем чтобы добавляло данные в базу, потом выводило что все ок, а если что то не так то выводило ошибку
	if _, err = db.Exec("INSERT INTO feeback (book_name, user_name, text) VALUES ('" + book + "', '" + user + "', '" + text + "');"); err != nil {
		return false, err.Error()
	}
	if err = db.Close(); err != nil {
		return false, err.Error()
	}
	return true, text
}

// Отзывы под книгой
func db_get_review(book string) (bool, []reviw_s) {

	//connect tp db
	db, err := sql.Open("sqlite", "library.sqlite")
	result := []reviw_s{}
	if err != nil {
		return false, result
	}
	//делаем чтобы добавляло данные в базу, потом выводило что все ок, а если что то не так то выводило ошибку
	rows, err := db.Query("select user_name, text from feeback where book_name = \"" + book + "\";")
	if err != nil {
		return false, result
	}

	//перебираем результат из БД
	for rows.Next() {
		review := reviw_s{}
		if err = rows.Scan(&review.user_name, &review.text); err != nil {
			return false, result
		}
		result = append(result, review)
	}

	if err = rows.Err(); err != nil {
		return false, result
	}

	if err = db.Close(); err != nil {
		return false, result
	}

	return true, result

}
