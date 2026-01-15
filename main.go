package main

import (
	"fmt"
	"encoding/json"
	"github.com/solarxweb/hexlet-go/greeting"
	"github.com/solarxweb/hexlet-go/movies"
	"github.com/solarxweb/hexlet-go/banking"
	"github.com/solarxweb/hexlet-go/books"
	"github.com/solarxweb/hexlet-go/articles"
	"github.com/fatih/color"
	// "strings"
)

type IntStr interface {
	string | int
}

func ReverseSlice[T IntStr](xs []T) []T {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs;
}

func main() {
	color.Green(greeting.Hello());

	color.Red(movies.GetFavMovie());

	account := banking.Account{
		ID: 1337,
		Owner: "rmnk",
		Balance: 315.55,
	}

	banking.Deposit(&account, 2050.5);

	nosov := books.NewBook("Незнайка", "Носов Н.")
	chukovsky := books.NewBook("Незнайка", "Носов Н.")

	color.Cyan(chukovsky.Description());
	color.Cyan(nosov.Description());

	finalAtcl := articles.Article{};
	article := articles.Article{ Title: "Создание soccertmr.ru", Tags: []string{"Футбол", "Тутаев", "Чемпионат"}};

	a,_ := json.Marshal(article);
	color.Magenta(string(a));

	json.Unmarshal([]byte(a), &finalAtcl);
	fmt.Println(finalAtcl.Title);


	res := ReverseSlice([]int{2, 4, 6, 8, 10, 12, 1488});
	res2 := ReverseSlice([]string{"don", "ya", "gon"})


	color.Yellow("%d", res)
	color.Blue("%s", res2)
};

