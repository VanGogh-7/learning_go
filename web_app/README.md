# Go Movie Recommender

A small educational movie recommendation website built with Go's standard library. It demonstrates routing, request handling, HTML templates, static files, form submissions, and JSON responses.

## Run the project

From the `web_app` folder:

```bash
go run .
```

Then open <http://localhost:8081>.

## Project folders

- `main.go`: creates the router, connects routes to handlers, serves static files, and starts the server.
- `handlers/`: contains HTTP handler functions for pages, forms, and the JSON API.
- `models/`: contains the `Movie` and `User` data types plus in-memory movie data.
- `templates/`: contains the shared layout and individual page templates.
- `static/css/`: contains the website styles.
- `static/js/`: contains browser-side JavaScript.

## Browser URLs

- <http://localhost:8081/> - home page and JavaScript fetch demo
- <http://localhost:8081/login> - example GET and POST form
- <http://localhost:8081/profile?user=Alex> - profile using a query parameter
- <http://localhost:8081/movies> - server-rendered movie list
- <http://localhost:8081/movie?id=1> - one movie and related recommendations
- <http://localhost:8081/api/recommendations?genre=Sci-Fi> - JSON API response

## How frontend JSON fetching works

The home page contains a button with the ID `load-recommendations`. The script in `static/js/main.js` listens for a click and calls:

```javascript
fetch("/api/recommendations?genre=Sci-Fi")
```

The `RecommendationsAPI` handler reads the `genre` query parameter, filters the in-memory movies, and uses `encoding/json` to return them as JSON. JavaScript converts that JSON into movie cards and adds them to the page without a full page reload.
