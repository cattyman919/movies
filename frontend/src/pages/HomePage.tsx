import { useEffect, useState } from 'react';
import './HomePage.css';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;
const TMDB_IMAGE_BASE_URL = 'https://image.tmdb.org/t/p/w500';

interface Movie {
  id: number;
  title: string;
  poster_path: string;
  overview: string;
}

interface PopularMoviesResponse {
  results: Movie[];
}

export function HomePage() {
  const [movies, setMovies] = useState<Movie[]>([]);

  useEffect(() => {
    fetch(`${API_BASE_URL}/api/movies/popular`)
      .then((res) => res.json())
      .then((data: PopularMoviesResponse) => {
        setMovies(data.results);
      })
      .catch(console.error);
  }, []);

  return (
    <div className="movie-grid">
      {movies.map((movie) => (
        <div key={movie.id} className="movie-card">
          <img src={`${TMDB_IMAGE_BASE_URL}${movie.poster_path}`} alt={movie.title} />
          <div className="movie-info">
            <h3>{movie.title}</h3>
          </div>
        </div>
      ))}
    </div>
  );
}
