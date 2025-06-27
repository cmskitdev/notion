import { Client } from "@notionhq/client";
import dotenv from "dotenv";
import { getDatabases } from "../src/databases/databases";

dotenv.config();

export const notion = new Client({
  auth: process.env.NOTION_API_KEY,
});

getDatabases().then((databases) => {
  console.log(databases);
});
