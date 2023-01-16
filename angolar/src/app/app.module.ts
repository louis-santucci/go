import {APP_INITIALIZER, NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppRoutingModule} from './app-routing.module';
import {AppComponent} from './app.component';
import {HttpClientModule} from "@angular/common/http";
import {LoginComponent} from './login/login.component';
import {HomeComponent} from './home/home.component';
import {PropertiesService} from "./services/properties.service";
import {HashLocationStrategy, LocationStrategy} from "@angular/common";
import {ToastrModule} from "ngx-toastr";
import { RedirectionTableComponent } from './components/redirection-table/redirection-table.component';
import {MatTableModule} from "@angular/material/table";
import {MatIconModule} from "@angular/material/icon";
import {BrowserAnimationsModule} from "@angular/platform-browser/animations";
import {MatButtonModule} from "@angular/material/button";
import {MatInputModule} from "@angular/material/input";
import {MatButtonToggleModule} from "@angular/material/button-toggle";
import {MatSortModule} from "@angular/material/sort";

const appInit = (propertiesService: PropertiesService) => () => propertiesService.loadConfig();

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    HomeComponent,
    RedirectionTableComponent,
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
    MatSortModule
  ],
  providers: [
    {provide: APP_INITIALIZER, useFactory: appInit, deps: [PropertiesService], multi: true},
    {provide: LocationStrategy, useClass: HashLocationStrategy}
  ],
  bootstrap: [AppComponent]
})
export class AppModule {
}
