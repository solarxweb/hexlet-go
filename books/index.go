package books;

import "fmt";

type Book struct {
	Title string
	Author string
}

func GetAuthorCreation(creations [5]Book, author string) ([]string, error) {
	creationsSlice := []Book{};
	if len(creations) == 0 {
		return []string{}, fmt.Errorf("length of books equal %d", len(creations));
	}
	for i := range creations {
		creationsSlice = append(creationsSlice, creations[i]);
	}

	filteredSlice := []string{};
	for _, item := range creationsSlice {
		if item.Author == author {
			filteredSlice = append(filteredSlice, item.Title)
		}
	}
	
	if len(filteredSlice) == 0 {
		return []string{}, fmt.Errorf("Creations of that author are doesnt exist in current array");
	}
	
	return filteredSlice, nil;


func (b Book) Description() string {
	return fmt.Sprintf("%s - %s", b.Title, b.Author);
}
