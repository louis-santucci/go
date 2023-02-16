import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RedirectionEditionComponent } from './redirection-edition.component';

describe('RedirectionEditionComponent', () => {
  let component: RedirectionEditionComponent;
  let fixture: ComponentFixture<RedirectionEditionComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RedirectionEditionComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RedirectionEditionComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
