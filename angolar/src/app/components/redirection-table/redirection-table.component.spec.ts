import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RedirectionTableComponent } from './redirection-table.component';

describe('RedirectionTableComponent', () => {
  let component: RedirectionTableComponent;
  let fixture: ComponentFixture<RedirectionTableComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RedirectionTableComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RedirectionTableComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
