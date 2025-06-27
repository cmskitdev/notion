import { PageObjectResponse } from "@notionhq/client/build/src/api-endpoints";
import { PropertyMultiSelect } from "./multi-select";
import { Property } from "./property";

/**
 * Mapping of property type strings to their corresponding class constructors.
 * This enables dynamic instantiation of property objects based on their type.
 */
export const PropertyTypeMapping = {
  multi_select: PropertyMultiSelect,
};

/**
 * A specialized Map implementation for managing Notion page properties.
 * This class provides type-safe access to properties with automatic instantiation
 * based on property types.
 *
 * @class
 * @extends Map<string, Property>
 */
export class Properties extends Map<string, Property> {
  constructor(
    properties: Pick<PageObjectResponse, "properties">["properties"]
  ) {
    super();
    for (const key in properties) {
      if (properties[key].type in PropertyTypeMapping) {
        this.set(
          key,
          new PropertyTypeMapping[
            properties[key].type as keyof typeof PropertyTypeMapping
          ](properties[key] as any)
        );
      }
    }
  }

  /**
   * Retrieves a property by its key name.
   *
   * @param key - The property key to retrieve
   * @throws Error if the property is not found
   */
  get(key: string): Property | undefined {
    return super.get(key);
  }

  /**
   * Returns all properties as a plain object.
   */
  get all(): Record<string, Property> {
    const result: Record<string, Property> = {};
    for (const [key, value] of this.entries()) {
      result[key] = value;
    }
    return result;
  }
}
