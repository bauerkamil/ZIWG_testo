export interface IComboboxProps<T> {
  items: T[];
  keyPath: string;
  valuePath?: string;
  getItemValue?: (item: T) => string;
  selectedItem: T | undefined;
  onItemSelected: (value: T | undefined) => void;
}
