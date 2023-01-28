//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:

//  - Check out a book
//  - Check in a book

//

package main

import (
	"fmt"

	//* Use the `time` package from the standard library for check in/out times
	"time"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned

type Book struct {
	name           string
	numberOfCopies int
}

type Member struct {
	name string
}

type Library struct {
	books   []Book
	members []Member
	lend    LendInfo
}

type LendInfo struct {
	checkOutInfo     map[string]string
	returnedBookInfo map[string]string
	booksLended      []Book
	whoLendedWhich   map[string][]string
}

func findBook(book string, library *Library) *Book {
	for i := range library.books {
		if library.books[i].name == book {
			return &library.books[i]
		}
	}
	return nil
}

func checkBookStock(book string, library *Library) int {
	for i := range library.books {
		if library.books[i].numberOfCopies > 0 {
			return library.books[i].numberOfCopies
		}
	}
	return 0
}

func findMember(member string, library *Library) *Member {
	for i := range library.members {
		if library.members[i].name == member {
			return &library.members[i]
		}
	}
	return nil
}

func checkOutBook(memberName, bookName string, library *Library) {
	//Find the book in the library
	book := findBook(bookName, library)
	// Find the member in the library
	member := findMember(memberName, library)

	bookStock := checkBookStock(bookName, library)

	// make sure member is part of the library
	if member == nil {
		fmt.Println("Sorry, checkout not possible . Member has not been found")
		return
	}

	// Make sure the book was found
	if book == nil {
		fmt.Println("Sorry, checkout not possible . Book has not been found")
		return
	}

	// make sure it's in stock
	if bookStock < 1 {
		fmt.Println("Sorry, checkout not possible . Book not available in stock")
		return
	}

	// Update the number of copies
	updateNumberCheckOut(book)

	// Update lendInfo
	updateLendInfoCheckOut(book, member, library)

	fmt.Println("The following book", bookName, "has been succesfully checked out by", memberName, "at", time.Now().Format(time.RFC1123))

}

func updateLendInfoCheckOut(book *Book, member *Member, library *Library) {

	timeRent := time.Now().Format(time.RFC1123)

	library.lend.checkOutInfo[book.name] = timeRent

	library.lend.booksLended = append(library.lend.booksLended, *book)

	library.lend.whoLendedWhich[member.name] = append(library.lend.whoLendedWhich[book.name], book.name)
}

func updateLendInfoReturn(book *Book, member *Member, library *Library) {

	delete(library.lend.whoLendedWhich, member.name)

	timeRent := time.Now().Format(time.RFC1123)

	library.lend.returnedBookInfo[book.name] = timeRent

	library.lend.booksLended = library.lend.booksLended[:len(library.lend.booksLended)-1]

}

func updateNumberCheckOut(book *Book) {
	book.numberOfCopies -= 1
}

func showBooksCheckedOut(library *Library) {
	fmt.Println("The following books have been checked out : ", library.lend.booksLended)
}

func returnBook(memberName, bookName string, library *Library) {

	book := findBook(bookName, library)
	// Find the member in the library
	member := findMember(memberName, library)

	// Make sure it's in stock
	if member == nil {
		fmt.Println("Sorry, checkout not possible . Member has not been found")
		return
	}

	// Make sure the book was found
	if book == nil {
		fmt.Println("Sorry, checkout not possible . Book has not been found")
		return
	}

	// Update the number of copies
	updateNumberReturn(book)

	// Update lendInfo
	updateLendInfoReturn(book, member, library)

	fmt.Println("The following book", bookName, "has been succesfully returned by", memberName, "at", time.Now().Format(time.RFC1123))
}

func updateNumberReturn(book *Book) {
	book.numberOfCopies += 1
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {

	lendBooks := LendInfo{checkOutInfo: make(map[string]string), returnedBookInfo: make(map[string]string), booksLended: make([]Book, 0, 10), whoLendedWhich: make(map[string][]string)}

	//  - Add at least 4 books and at least 3 members to the library
	peterPanBook := Book{name: "Peter Pan", numberOfCopies: 10}
	harryPotterBook := Book{name: "Harry Potter", numberOfCopies: 20}
	maryPoppinsBook := Book{name: "Mary Poppins", numberOfCopies: 5}
	lionKingBook := Book{name: "Lion King", numberOfCopies: 30}

	member1 := Member{name: "João"}
	member2 := Member{name: "Roberta"}
	member3 := Member{name: "Emília"}

	booksLibrary := make([]Book, 0, 10)
	booksLibrary = append(booksLibrary, peterPanBook, harryPotterBook, maryPoppinsBook, lionKingBook)

	members := make([]Member, 0, 10)
	members = append(members, member1, member2, member3)

	myLibrary := Library{books: booksLibrary, members: members, lend: lendBooks}

	//  - Print out initial library information, and after each change
	//* There must only ever be one copy of the library in memory at any time
	fmt.Println("Showing library : ")
	fmt.Println("This library contains the following books : ", myLibrary.books)
	fmt.Println("This library contains the following members :", myLibrary.members)
	showBooksCheckedOut(&myLibrary)
	fmt.Println()
	fmt.Println()
	fmt.Println("*********************************************************************************************************************")
	fmt.Println("Checking book out....")
	checkOutBook("João", "Peter Pan", &myLibrary)
	fmt.Println()
	fmt.Println()
	fmt.Println("*********************************************************************************************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println("Checking 4 books out....")
	checkOutBook("João", "Harry Potter", &myLibrary)
	checkOutBook("Pedro", "Harry Potter", &myLibrary)
	checkOutBook("Pedro", "Machado de Assis", &myLibrary)
	checkOutBook("João", "Machado de Assis", &myLibrary)
	fmt.Println()
	fmt.Println()
	fmt.Println("*********************************************************************************************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println("This library contains the following books : ", myLibrary.books)
	fmt.Println("This library contains the following members :", myLibrary.members)
	fmt.Println()
	fmt.Println()
	fmt.Println("*********************************************************************************************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println("Showings checked out books...")
	showBooksCheckedOut(&myLibrary)
	fmt.Println()
	fmt.Println()
	fmt.Println("*********************************************************************************************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println("Returning 1 book...")
	returnBook("João", "Peter Pan", &myLibrary)
	fmt.Println()
	fmt.Println()
	fmt.Println("*********************************************************************************************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println("This library contains the following books : ", myLibrary.books, "\n And the following members :", myLibrary.members)

}
