package search

type Search_result struct {
	Type     string
	Total    int
	Per_page int
	Order    int
	Results  []interface{} //List[T]
}
