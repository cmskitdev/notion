import { notion } from "../client/client";
import { debug } from "./debug";

/**
 * Interface for a Notion relation property.
 */
interface RelationProperty {
  id: string;
  type: "relation";
  relation: Array<{ id: string }>;
  has_more: boolean;
}

/**
 * Interface for a nested relation with full page data.
 */
interface NestedRelation {
  id: string;
  type: "relation";
  relation: Array<{ id: string; page?: any }>;
  has_more: boolean;
  nested_pages?: any[];
}

export const findRelationProperties = (properties: any) => {
  return Object.keys(properties).filter(
    (key) => properties[key]?.type === "relation"
  );
};

/**
 * Extracts and nests relation properties by fetching the full page data for each related page.
 *
 * @param properties - The properties object from a Notion page or database query result
 * @param relationPropertyName - The name of the relation property to nest (e.g., "Series")
 * @returns The properties object with the specified relation property nested with full page data
 */
export const resolveRelations = async (
  properties: Record<string, any>,
  relationPropertyName: string
): Promise<Record<string, any>> => {
  debug("resolveRelations", { relationPropertyName, properties });

  // Check if the property exists and is a relation type
  const relationProperty = properties[relationPropertyName];
  if (!relationProperty || relationProperty.type !== "relation") {
    throw new Error(
      `Property ${relationPropertyName} not found or not a relation type`
    );
  }

  const typedRelationProperty = relationProperty as RelationProperty;

  // If there are no relations, return as is
  if (
    !typedRelationProperty.relation ||
    typedRelationProperty.relation.length === 0
  ) {
    debug("No relations found", { relationPropertyName });
    return properties;
  }

  try {
    // Fetch full page data for each related page
    const nestedPages = await Promise.all(
      typedRelationProperty.relation.map(async (relation) => {
        try {
          const page = await notion.pages.retrieve({
            page_id: relation.id,
          });
          debug("Fetched related page", {
            pageId: relation.id,
            title: page.properties.Name.title[0].plain_text,
          });
          return page;
        } catch (error) {
          debug("Error fetching related page", { pageId: relation.id, error });
          return null;
        }
      })
    );

    // Filter out any failed fetches
    const validNestedPages = nestedPages.filter((page) => page !== null);

    // Create the nested relation property
    const nestedRelation: NestedRelation = {
      ...typedRelationProperty,
      nested_pages: validNestedPages,
      relation: typedRelationProperty.relation.map((relation, index) => ({
        ...relation,
        page: nestedPages[index],
      })),
    };

    // Return the properties with the nested relation
    return {
      ...properties,
      [relationPropertyName]: nestedRelation,
    };
  } catch (error) {
    debug("Error nesting relation property", { relationPropertyName, error });
    return properties;
  }
};

/**
 * Extracts and nests multiple relation properties.
 *
 * @param properties - The properties object from a Notion page or database query result
 * @param relationPropertyNames - Array of relation property names to nest
 * @returns The properties object with all specified relation properties nested
 */
export const nestMultipleRelationProperties = async (
  properties: Record<string, any>,
  relationPropertyNames: string[]
): Promise<Record<string, any>> => {
  let nestedProperties = { ...properties };

  for (const propertyName of relationPropertyNames) {
    nestedProperties = await resolveRelations(nestedProperties, propertyName);
  }

  return nestedProperties;
};

/**
 * Automatically detects and nests all relation properties in a properties object.
 *
 * @param properties - The properties object from a Notion page or database query result
 * @returns The properties object with all relation properties nested
 */
export const nestAllRelationProperties = async (
  properties: Record<string, any>
): Promise<Record<string, any>> => {
  const relationPropertyNames = Object.keys(properties).filter(
    (key) => properties[key]?.type === "relation"
  );

  debug("Found relation properties", { relationPropertyNames });

  return nestMultipleRelationProperties(properties, relationPropertyNames);
};
