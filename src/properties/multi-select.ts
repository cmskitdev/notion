import { MultiSelectPropertyItemObjectResponse } from "@notionhq/client/build/src/api-endpoints";
import { Property, PropertyType, PropertyValue } from "./property";

/**
 * Represents a multi-select property in Notion.
 *
 * @class
 * @implements Property
 */
export class PropertyMultiSelect implements Property {
  id: string;
  type: PropertyType.MultiSelect;
  original: MultiSelectPropertyItemObjectResponse;

  constructor(original: MultiSelectPropertyItemObjectResponse) {
    this.id = original.id;
    this.type = PropertyType.MultiSelect;
    this.original = original;
  }

  get value(): PropertyValue {
    // return this.original.multi_select.map((option) => option.id);
    return this.original.multi_select;
  }
}

export class PropertyMultiSelectValue {
  id: string;
  name: string;
}
