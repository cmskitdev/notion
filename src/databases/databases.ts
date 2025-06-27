import {
  DatabaseObjectResponse,
  GetDatabaseResponse,
} from "@notionhq/client/build/src/api-endpoints";
import { config } from "../config/config";
import { debug } from "../util/debug";
import { save } from "../util/fs";
import { Database } from "./database";
import { Client } from "@notionhq/client";

const notion = new Client();

export const getDatabases = async (): Promise<Database[]> => {
  const response = await notion.search({
    query: "",
    filter: {
      property: "object",
      value: "database",
    },
  });

  console.log(config.excludes);
  console.log(config.excludes);

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
  const response = await notion.retrieve({
    database_id: databaseId,
  });

  return response;
};

export const search = async (databaseName: string, query: string) => {
  debug("search", databaseName, query);
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

  debug("search", response);

  for (const result of response.results) {
    const page = await notion.pages.retrieve({
      page_id: result.id,
    });
    debug("page", page);
    save(`./schemas/page.json`, page);
  }

  return response;
};
