import { PropertyItemObjectResponse } from "@notionhq/client/build/src/api-endpoints";

/**
 * Enum representing all supported property types for Notion-like properties.
 * This enum provides a type-safe and descriptive way to refer to property types throughout the codebase.
 *
 * @enum
 */
export enum PropertyType {
  Title = "title",
  RichText = "rich_text",
  Number = "number",
  Select = "select",
  MultiSelect = "multi_select",
  Date = "date",
  Checkbox = "checkbox",
  Url = "url",
  Email = "email",
  PhoneNumber = "phone_number",
  Formula = "formula",
  Relation = "relation",
  Rollup = "rollup",
  CreatedBy = "created_by",
}

/**
 * Represents a multi-select option value.
 */
export class PropertyMultiSelectValue {
  id: string;
  name: string;
}

/**
 * Union type representing all possible property values.
 */
export type PropertyValue =
  | string
  | number
  | boolean
  | Date
  | string[]
  | number[]
  | PropertyMultiSelectValue[];

/**
 * Base interface that all property implementations must satisfy.
 * This provides a consistent structure for working with Notion properties.
 *
 * @interface
 */
export interface Property {
  id: string;
  type: PropertyType;
  original: PropertyItemObjectResponse;

  get value(): PropertyValue;
}
