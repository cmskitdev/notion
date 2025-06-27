import path from "path";
import { notion } from "../client/client";
import { config } from "../config/config";
import { save } from "../util/fs";
import { Options } from "../util/options";
import { Page } from "./page";

export const getPages = async (databaseName: string) => {
  const database = config.databases[databaseName];
  if (!database) {
    throw new Error(`Database "${databaseName}" not found!`);
  }

  const response = await notion.databases.query({
    database_id: database.id,
  });

  save(
    path.resolve(__dirname, `../../tmp/${databaseName}.pages.json`),
    response.results
  );

  return response.results;
};

export const getPage = async (
  pageId: string,
  options?: Options
): Promise<Page> => {
  const page = new Page(
    await notion.pages.retrieve({ page_id: pageId }),
    options
  );

  return page;
};
