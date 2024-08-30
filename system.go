package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

type Book struct {
    ID     int
    Title  string
    Author string
}

type Movie struct {
    ID     int
    Title  string
    Director string
}

var books []Book
var movies []Movie
var nextBookID int = 1
var nextMovieID int = 1

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Println("\nSistema de Gerenciamento")
        fmt.Println("1. Adicionar Livro")
        fmt.Println("2. Listar Livros")
        fmt.Println("3. Buscar Livro")
        fmt.Println("4. Remover Livro")
        fmt.Println("5. Adicionar Filme")
        fmt.Println("6. Listar Filmes")
        fmt.Println("7. Buscar Filme")
        fmt.Println("8. Remover Filme")
        fmt.Println("9. Sair")
        fmt.Print("Escolha uma opção: ")
        
        option, _ := reader.ReadString('\n')
        option = strings.TrimSpace(option)

        switch option {
        case "1":
            addBook(reader)
        case "2":
            listBooks()
        case "3":
            searchBook(reader)
        case "4":
            removeBook(reader)
        case "5":
            addMovie(reader)
        case "6":
            listMovies()
        case "7":
            searchMovie(reader)
        case "8":
            removeMovie(reader)
        case "9":
            fmt.Println("Saindo...")
            return
        default:
            fmt.Println("Opção inválida! Tente novamente.")
        }
    }
}

func addBook(reader *bufio.Reader) {
    fmt.Print("Digite o título do livro: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)

    fmt.Print("Digite o autor do livro: ")
    author, _ := reader.ReadString('\n')
    author = strings.TrimSpace(author)

    book := Book{
        ID:     nextBookID,
        Title:  title,
        Author: author,
    }
    books = append(books, book)
    nextBookID++

    fmt.Println("Livro adicionado com sucesso!")
}

func listBooks() {
    if len(books) == 0 {
        fmt.Println("Nenhum livro cadastrado.")
        return
    }

    fmt.Println("\nLista de Livros:")
    for _, book := range books {
        fmt.Printf("ID: %d, Título: %s, Autor: %s\n", book.ID, book.Title, book.Author)
    }
}

func searchBook(reader *bufio.Reader) {
    fmt.Print("Digite o título ou autor para buscar: ")
    query, _ := reader.ReadString('\n')
    query = strings.TrimSpace(query)

    found := false
    for _, book := range books {
        if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) ||
           strings.Contains(strings.ToLower(book.Author), strings.ToLower(query)) {
            fmt.Printf("Encontrado - ID: %d, Título: %s, Autor: %s\n", book.ID, book.Title, book.Author)
            found = true
        }
    }

    if !found {
        fmt.Println("Nenhum livro encontrado.")
    }
}

func removeBook(reader *bufio.Reader) {
    fmt.Print("Digite o ID do livro a ser removido: ")
    var id int
    fmt.Scanf("%d\n", &id)

    for i, book := range books {
        if book.ID == id {
            books = append(books[:i], books[i+1:]...)
            fmt.Println("Livro removido com sucesso!")
            return
        }
    }

    fmt.Println("Livro não encontrado.")
}

func addMovie(reader *bufio.Reader) {
    fmt.Print("Digite o título do filme: ")
    title, _ := reader.ReadString('\n')
    title = strings.TrimSpace(title)

    fmt.Print("Digite o diretor do filme: ")
    director, _ := reader.ReadString('\n')
    director = strings.TrimSpace(director)

    movie := Movie{
        ID:       nextMovieID,
        Title:    title,
        Director: director,
    }
    movies = append(movies, movie)
    nextMovieID++

    fmt.Println("Filme adicionado com sucesso!")
}

func listMovies() {
    if len(movies) == 0 {
        fmt.Println("Nenhum filme cadastrado.")
        return
    }

    fmt.Println("\nLista de Filmes:")
    for _, movie := range movies {
        fmt.Printf("ID: %d, Título: %s, Diretor: %s\n", movie.ID, movie.Title, movie.Director)
    }
}

func searchMovie(reader *bufio.Reader) {
    fmt.Print("Digite o título ou diretor para buscar: ")
    query, _ := reader.ReadString('\n')
    query = strings.TrimSpace(query)

    found := false
    for _, movie := range movies {
        if strings.Contains(strings.ToLower(movie.Title), strings.ToLower(query)) ||
           strings.Contains(strings.ToLower(movie.Director), strings.ToLower(query)) {
            fmt.Printf("Encontrado - ID: %d, Título: %s, Diretor: %s\n", movie.ID, movie.Title, movie.Director)
            found = true
        }
    }

    if !found {
        fmt.Println("Nenhum filme encontrado.")
    }
}

func removeMovie(reader *bufio.Reader) {
    fmt.Print("Digite o ID do filme a ser removido: ")
    var id int
    fmt.Scanf("%d\n", &id)

    for i, movie := range movies {
        if movie.ID == id {
            movies = append(movies[:i], movies[i+1:]...)
            fmt.Println("Filme removido com sucesso!")
            return
        }
    }

    fmt.Println("Filme não encontrado.")
}
