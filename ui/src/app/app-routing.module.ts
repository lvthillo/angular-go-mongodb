import { HomeComponent } from './home/home.component';
import { RouterModule, Routes } from "@angular/router";
import { NgModule } from '@angular/core';
import { MovieComponent } from './movie/movie.component';

const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'movie', component: MovieComponent },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [ RouterModule ]
})
export class AppRoutingModule { }