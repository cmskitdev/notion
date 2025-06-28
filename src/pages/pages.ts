import { notion } from "$lib/client";
import { GetPageResponse } from "@notionhq/client";
import { config } from "../config/config";
import { Options } from "../util/options";
import { Page } from "./page";

export const getPages = async (databaseName: string): Promise<Page[]> => {
  const database = config.databases[databaseName];
  if (!database) {
    throw new Error(`Database "${databaseName}" not found!`);
  }

  const response = await notion.databases.query({
    database_id: database.id,
  });

  return response.results.map((page) => new Page(page as GetPageResponse));
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
