# ./docs

Welcome to the article-test/rest generated docs.

## Routes

<details>
<summary>`/article`</summary>

- [(*Cors).Handler-fm]()
- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article**
	- **/**
		- _GET_
			- [ListArticles]()
		- _POST_
			- [CreateArticle]()

</details>
<details>
<summary>`/article/{articleID}`</summary>

- [(*Cors).Handler-fm]()
- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article**
	- **/{articleID}**
		- [ArticleCtx]()
		- **/**
			- _DELETE_
				- [DeleteArticle]()
			- _GET_
				- [GetArticle]()
			- _PUT_
				- [UpdateArticle]()

</details>
<details>
<summary>`/article/{limit}/{offset}`</summary>

- [(*Cors).Handler-fm]()
- [RequestID]()
- [Logger]()
- [Recoverer]()
- [URLFormat]()
- [SetContentType.func1]()
- **/article**
	- **/{limit}/{offset}**
		- _GET_
			- [ListArticlesWithPage]()

</details>

Total # of routes: 3
