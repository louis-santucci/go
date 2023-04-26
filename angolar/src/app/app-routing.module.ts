import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {HomeComponent} from "./components/home/home.component";
import {RegisterComponent} from "./components/register/register.component";
import {LoginComponent} from "./components/login/login.component";
import {RedirectionCreationComponent} from "./components/redirection-creation/redirection-creation.component";
import {RedirectionEditionComponent} from "./components/redirection-edition/redirection-edition.component";
import {UserInfoComponent} from "./components/user-info/user-info.component";
import {AuthGuard} from "./guards/auth.guard";
import {ErrorNotFoundComponent} from "./components/errors/error-not-found.component";
import {ErrorUnauthorizedComponent} from "./components/errors/error-unauthorized.component";
import {HistoryComponent} from "./components/history/history.component";
import {UserEditionComponent} from "./components/user-edition/user-edition.component";

const routes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'login', component: LoginComponent},
  {path: 'history', component: HistoryComponent},
  {path: 'redirection/new', component: RedirectionCreationComponent, canActivate: [AuthGuard]},
  {path: 'redirection/edit/:id', component: RedirectionEditionComponent, canActivate: [AuthGuard]},
  {path: 'user/info', component: UserInfoComponent, canActivate: [AuthGuard]},
  {path: 'user/edit', component: UserEditionComponent, canActivate: [AuthGuard]},
  {path: 'error/notFound', component: ErrorNotFoundComponent},
  {path: 'error/unauthorized', component: ErrorUnauthorizedComponent},
  {path: '**', redirectTo: 'error/notFound'}
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
