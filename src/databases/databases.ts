import { notion } from "$lib/client";
import { Page } from "$lib/pages/page";
import {
  DatabaseObjectResponse,
  GetDatabaseResponse,
} from "@notionhq/client/build/src/api-endpoints";
import { config } from "../config/config";
import { save } from "../util/fs";
import { Database } from "./database";

export const getDatabases = async (): Promise<Database[]> => {
  const response = await notion.search({
    query: "",
    filter: {
      property: "object",
      value: "database",
    },
  });

  return response.results
    .filter((result) => {
      if (
        config.excludes.some((exclude) => exclude.types.includes("database"))
      ) {
        console.log(JSON.stringify(config.excludes));
        return !config.excludes.some((exclude) => {
          if (typeof exclude.pattern === "string") {
            const database = result as DatabaseObjectResponse;
            console.log(
              database.title[0].plain_text,
              database.title[0].plain_text.indexOf(exclude.pattern)
            );
            return database.title[0].plain_text.indexOf(exclude.pattern) !== -1;
          } else {
            return exclude.pattern.test(result.id);
          }
        });
      }
      return true;
    })
    .map((result) => new Database(result as DatabaseObjectResponse));
};

export const queryDatabase = async (
  databaseId: string
): Promise<GetDatabaseResponse> => {
  const response = await notion.databases.retrieve({
    database_id: databaseId,
  });

  return response;
};

export const search = async (
  databaseName: string,
  query: string
): Promise<Page[]> => {
  if (!config.databases[databaseName]) {
    throw new Error(`Database "${databaseName}" not found!`);
  }

  const response = await notion.databases.query({
    database_id: config.databases[databaseName].id,
    filter_properties: ["title"],
    filter: {
      property: "Excerpt",
      and: [
        {
          property: "Excerpt",
          rich_text: {
            contains: query,
          },
        },
      ],
    },
  });

  const results: SearchResult[] = [];

  for (const result of response.results) {
    const page = await notion.pages.retrieve({
      page_id: result.id,
    });
    results.push(new Page(page));
    save(`databases-search-${databaseName}-${result.id}`, page);
  }

  return results;
};
