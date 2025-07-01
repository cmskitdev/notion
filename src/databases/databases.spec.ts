import {
  DatabaseObjectResponse,
  PageObjectResponse,
  RichTextItemResponse,
} from "@notionhq/client/build/src/api-endpoints";
import { clearTimeout, setTimeout as nodeSetTimeout } from "timers";
import { describe, test } from "vitest";
import { notion } from "../client";
import { getBasePath, save } from "../util/fs";
import { toSafeName } from "../util/strings";

describe("databases", () => {
  test(`getDatabases`, async () => {
    const databases = await notion.search({
      query: "",
      filter: {
        property: "object",
        value: "database",
      },
      page_size: 100,
    });

    for (const database of databases.results) {
      if (isDatabaseResponse(database) && database.title.length > 0) {
        const response = await notion.databases.query({
          database_id: database.id,
          page_size: 100,
        });
        const databaseName = toSafeName(database.title[0].plain_text);
        save(
          getBasePath("tmp/dump", `${databaseName}/database.json`),
          response.results
        );

        for (const page of response.results as PageObjectResponse[]) {
          if (
            isTitleProperty(page.properties.Name) &&
            page.properties.Name.title.length > 0
          ) {
            const title = page.properties.Name.title[0].plain_text;
            const safeFileName = toSafeName(title);
            console.log("saving:", `${databaseName}/${safeFileName}.json`);
            // Truncate filename if too long to avoid ENAMETOOLONG error
            const maxLength = 255; // Max filename length on most filesystems
            const extension = ".json";
            const truncatedFileName =
              safeFileName.slice(0, maxLength - extension.length) + extension;
            save(
              getBasePath("tmp/dump", `${databaseName}/${truncatedFileName}`),
              page
            );
          }
        }
      }
    }
  });
});

function isDatabaseResponse(response: any): response is DatabaseObjectResponse {
  return response.object === "database";
}

function isTitleProperty(
  property: PageObjectResponse["properties"][string]
): property is {
  type: "title";
  title: Array<RichTextItemResponse>;
  id: string;
} {
  return property.type === "title";
}

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => nodeSetTimeout(resolve, ms));
}

class ConcurrencyLimiter {
  private running = 0;
  private queue: Array<() => void> = [];

  constructor(private maxConcurrent: number) {}

  async run<T>(fn: () => Promise<T>, timeout?: number): Promise<T> {
    while (this.running >= this.maxConcurrent) {
      await new Promise<void>((resolve) => this.queue.push(resolve));
    }

    this.running++;

    try {
      // Add timeout wrapper if specified
      if (timeout) {
        const operation = fn();
        const result = await Promise.race([
          operation,
          new Promise<T>((_, reject) => {
            const timeoutId = nodeSetTimeout(() => {
              reject(new Error("Operation timed out"));
            }, timeout);
            // Clean up timeout if operation completes
            operation.finally(() => clearTimeout(timeoutId));
          }),
        ]);
        return result;
      }
      return await fn();
    } finally {
      this.running--;
      const next = this.queue.shift();
      if (next) next();
    }
  }
}
