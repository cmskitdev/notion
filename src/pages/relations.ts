import { notion } from "../client/client";
import { config } from "../config/config";
import { debug } from "../util/debug";
import { nestAllRelationProperties, resolveRelations } from "../util/relations";
import { getPage } from "./pages";

export const findRelationProperties = (properties: any) => {
  return Object.keys(properties).filter(
    (key) => properties[key]?.type === "relation"
  );
};

/**
 * Retrieves pages from a database and nests the Series relation property.
 *
 * @param databaseName - The name of the database to query
 * @returns Database query results with Series relations nested
 */
export const getPagesWithNestedSeries = async (databaseName: string) => {
  const database = config.databases[databaseName];
  if (!database) {
    throw new Error(`Database "${databaseName}" not found!`);
  }

  debug("Querying database with nested Series", {
    databaseName,
    databaseId: database.id,
  });

  const response = await notion.databases.query({
    database_id: database.id,
  });

  // Nest the Series relation property for each page
  const resultsWithNestedSeries = await Promise.all(
    response.results.map(async (page: any) => {
      if (page.properties) {
        const nestedProperties = await resolveRelations(
          page.properties,
          "Series"
        );
        return {
          ...page,
          properties: nestedProperties,
        };
      }
      return page;
    })
  );

  return {
    ...response,
    results: resultsWithNestedSeries,
  };
};

/**
 * Retrieves a specific page and nests the Series relation property.
 *
 * @param pageId - The ID of the page to retrieve
 * @returns Page with Series relation nested
 */
export const getPageWithNestedSeries = async (
  pageId: string,
  relationalPropertyNames: string[]
) => {
  const page = await getPage(pageId, { resolveRelations: true });

  // for (const relationalPropertyName of relationalPropertyNames) {
  //   const nestedProperties = await resolveRelations(
  //     page.properties,
  //     relationalPropertyName
  //   );

  //   // page.properties[relationalPropertyName] = nestedProperties;
  // }

  // page.properties = properties;

  return page;
};

/**
 * Retrieves pages from a database and nests all relation properties.
 *
 * @param databaseName - The name of the database to query
 * @returns Database query results with all relations nested
 */
export const getPagesWithAllNestedRelations = async (databaseName: string) => {
  const database = config.databases[databaseName];
  if (!database) {
    throw new Error(`Database "${databaseName}" not found!`);
  }

  debug("Querying database with all nested relations", {
    databaseName,
    databaseId: database.id,
  });

  const response = await notion.databases.query({
    database_id: database.id,
  });

  // Nest all relation properties for each page
  const resultsWithNestedRelations = await Promise.all(
    response.results.map(async (page: any) => {
      if (page.properties) {
        const nestedProperties = await nestAllRelationProperties(
          page.properties
        );
        return {
          ...page,
          properties: nestedProperties,
        };
      }
      return page;
    })
  );

  return {
    ...response,
    results: resultsWithNestedRelations,
  };
};
