import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {HTTP_INTERCEPTORS, HttpClientModule} from "@angular/common/http";
import {LoginComponent} from './components/login/login.component';
import {HomeComponent} from './home/home.component';
import {HashLocationStrategy, LocationStrategy} from "@angular/common";
import {ToastrModule} from "ngx-toastr";
import {RedirectionTableComponent} from './components/redirection-table/redirection-table.component';
import {MatTableModule} from "@angular/material/table";
import {MatIconModule} from "@angular/material/icon";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {MatButtonModule} from "@angular/material/button";
import {MatInputModule} from "@angular/material/input";
import {MatButtonToggleModule} from "@angular/material/button-toggle";
import {MatSortModule} from "@angular/material/sort";
import {AuthInterceptor} from "./interceptors/auth.interceptor";
import {RegisterComponent} from './components/register/register.component';
import {ReactiveFormsModule} from "@angular/forms";
import {MatToolbarModule} from "@angular/material/toolbar";
import { NavBarComponent } from './components/nav-bar/nav-bar.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent,
    RedirectionTableComponent,
    RegisterComponent,
    NavBarComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ToastrModule.forRoot(),
    MatTableModule,
    MatIconModule,
    BrowserAnimationsModule,
    MatButtonModule,
    MatInputModule,
    MatButtonToggleModule,
    MatSortModule,
    ReactiveFormsModule,
    MatToolbarModule,
  ],
  providers: [
    {provide: LocationStrategy, useClass: HashLocationStrategy},
    {provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true}
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
