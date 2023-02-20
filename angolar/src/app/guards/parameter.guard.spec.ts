import { TestBed } from '@angular/core/testing';

import { ParameterGuard } from './parameter.guard';

describe('ParameterGuard', () => {
  let guard: ParameterGuard;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    guard = TestBed.inject(ParameterGuard);
  });

  it('should be created', () => {
    expect(guard).toBeTruthy();
  });
});
