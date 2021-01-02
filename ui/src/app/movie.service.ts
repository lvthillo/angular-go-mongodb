import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../environments/environment';

@Injectable()
export class MovieService {
  constructor(private httpClient: HttpClient) {}

  /*getMovieList() {
    return this.httpClient.get(environment.gateway + '/movies');
  }*/

  getMovieList() {
    return this.httpClient.get<Movie[]>(environment.gateway + '/movie');
  }

  addMovie(movie: Movie) {
    return this.httpClient.post(environment.gateway + '/movie', movie);
  }

  watchedMovie(movie: Movie) {
    return this.httpClient.put(environment.gateway + '/movie', movie);
  }

  deleteMovie(movie: Movie) {
    return this.httpClient.delete(environment.gateway + '/movie/' + movie._id);
  }
}

export class Movie {
  _id: string;
  title: string;
  year: number;
  watched: boolean;
}