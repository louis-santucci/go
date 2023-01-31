export class EventData {
  private _name: string;
  private _value: any;

  constructor(name: string, value: any) {
    this._name = name;
    this._value = value;
  }

  get name(): string {
    return this._name;
  }

  get value(): any {
    return this._value;
  }
}
