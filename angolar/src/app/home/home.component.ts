import {Component, OnDestroy, OnInit} from '@angular/core';
import {RedirectionService} from "../services/redirection.service";
import {Redirection} from "../models/redirection";
import {Subscription} from "rxjs";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit, OnDestroy{

  private redirectionMap?: Map<string, Redirection>;

  private redirectionListSubscription?: Subscription;

  public constructor(private redirectionService: RedirectionService) {
  }

  public ngOnInit() {
    this.redirectionListSubscription = this.redirectionService.getRedirectionMapObservable()
      .subscribe(map => this.redirectionMap = map);
    this.redirectionService.getRedirections();
  }

  public ngOnDestroy() {
    this.redirectionListSubscription?.unsubscribe();
  }
}