import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RedirectionCreationComponent } from './redirection-creation.component';

describe('RedirectionCreationComponent', () => {
  let component: RedirectionCreationComponent;
  let fixture: ComponentFixture<RedirectionCreationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RedirectionCreationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RedirectionCreationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
