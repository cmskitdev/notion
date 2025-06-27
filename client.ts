import { Client as NotionClient } from "@notionhq/client";
import { CommonClient, GetType } from "@cmskit/common";

export class Client implements CommonClient {
  private notion: NotionClient;

  constructor() {
    this.notion = new NotionClient({
      auth: process.env.NOTION_API_KEY,
    });
  }

  async search(query: string, type: GetType) {
    return this.notion.search({
      query,
      filter: {
        property: "object",
        value: type,
      },
    });
  }

  async get(id: string, type: GetType) {
    switch (type) {
      case "page":
        return this.notion.pages.retrieve({
          page_id: id,
        });
      case "database":
        return this.notion.databases.retrieve({
          database_id: id,
        });
    }
  }

}

export const client = new Client();