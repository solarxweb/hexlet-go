package movies

import (
	"strings"	
	"fmt"
);

type FavMovie struct {
	Title string
	Yom int
	Genre []string
	Rating float64
}

func GetFavMovie() string {
	hp := &FavMovie{
		Title: "Harry Potter & Prisoner of Azkaban",
		Yom: 2004,
		Genre: []string{"fantasy", "adventures"},
		Rating: 8.3,
	};
	
	var builder strings.Builder;

	for i := 0; i < len(hp.Genre); i++ {
		builder.WriteString(hp.Genre[i]);
		if !(i == len(hp.Genre) - 1) {
			builder.WriteString(" ");
		}
	}

	gnr := builder.String()

	return fmt.Sprintf("Title: %s,\nGenre: %s,\nYear: %d\n", hp.Title, gnr, hp.Yom);
}