import { Get } from "@cmskit/common";
import { Client } from "@notionhq/client";
import dotenv from "dotenv";

dotenv.config();

export const notion = new Client({
  auth: process.env.NOTION_API_KEY,
});

type asdf = Get.Operation;
