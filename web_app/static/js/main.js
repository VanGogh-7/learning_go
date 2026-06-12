// The browser calls the Go JSON endpoint, then builds movie cards from the response.
const loadButton = document.querySelector("#load-recommendations");
const results = document.querySelector("#recommendation-results");

if (loadButton && results) {
  loadButton.addEventListener("click", async () => {
    loadButton.disabled = true;
    results.innerHTML = "<p>Loading recommendations...</p>";

    try {
      const response = await fetch("/api/recommendations?genre=Sci-Fi");
      if (!response.ok) {
        throw new Error(`Request failed with status ${response.status}`);
      }

      const movies = await response.json();
      results.innerHTML = movies.map((movie) => `
        <article class="movie-card">
          <span class="genre">${movie.genre}</span>
          <h3>${movie.title}</h3>
          <p class="movie-meta">${movie.year} · Rating ${movie.rating.toFixed(1)}</p>
          <p>${movie.description}</p>
          <a href="/movie?id=${movie.id}">View details</a>
        </article>
      `).join("");
    } catch (error) {
      results.innerHTML = `<p class="error">Could not load recommendations: ${error.message}</p>`;
    } finally {
      loadButton.disabled = false;
    }
  });
}
