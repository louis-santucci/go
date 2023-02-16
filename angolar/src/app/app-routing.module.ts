import {NgModule} from '@angular/core';
import {RouterModule, Routes} from '@angular/router';
import {HomeComponent} from "./components/home/home.component";
import {RegisterComponent} from "./components/register/register.component";
import {LoginComponent} from "./components/login/login.component";
import {RedirectionCreationComponent} from "./components/redirection-creation/redirection-creation.component";
import {RedirectionEditionComponent} from "./components/redirection-edition/redirection-edition.component";
import {UserInfoComponent} from "./components/user-info/user-info.component";
import {AuthGuard} from "./guards/auth.guard";

const routes: Routes = [
  {path: '', component: HomeComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'login', component: LoginComponent},
  {path: 'redirection/new', component: RedirectionCreationComponent, canActivate: [AuthGuard]},
  {path: 'redirection/edit/:id', component: RedirectionEditionComponent, canActivate: [AuthGuard]},
  {path: 'user/info', component: UserInfoComponent, canActivate: [AuthGuard]},
];


@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule {
}
