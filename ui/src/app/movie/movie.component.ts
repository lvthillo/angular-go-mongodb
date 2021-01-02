import { Component, OnInit } from '@angular/core';
import { MovieService, Movie } from '../movie.service';

@Component({
  selector: 'app-movie',
  templateUrl: './movie.component.html',
  styleUrls: ['./movie.component.css']
})
export class MovieComponent implements OnInit {

  activeMovies: Movie[];
  watchedMovies: Movie[];
  movieTitle: string;
  movieYear: number;
  movieWatched: boolean;

  constructor(private movieService: MovieService) { }

  ngOnInit() {
    this.getAll();
  }

  getAll() {
    this.movieService.getMovieList().subscribe((data: Movie[]) => {
      this.activeMovies = data.filter((a) => !a.watched);
      this.watchedMovies = data.filter((a) => a.watched);
    });
  }

  addMovie() {
    var newMovie : Movie = {
      title: this.movieTitle,
      _id: '',
      year: this.movieYear,
      watched: this.movieWatched
    };

    this.movieService.addMovie(newMovie).subscribe(() => {
      this.getAll();
      this.movieTitle = '';
    });
  }

  watchedMovie(movie: Movie) {
    this.movieService.watchedMovie(movie).subscribe(() => {
      this.getAll();
    });
  }

  deleteMovie(movie: Movie) {
    this.movieService.deleteMovie(movie).subscribe(() => {
      this.getAll();
    })
  }
}