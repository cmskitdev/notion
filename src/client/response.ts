import {
  BlockObjectResponse,
  DatabaseObjectResponse,
  PageObjectResponse,
  PartialBlockObjectResponse,
  PartialDatabaseObjectResponse,
  PartialPageObjectResponse,
} from "@notionhq/client/build/src/api-endpoints";

export type NotionObjectResponse = (
  | PartialPageObjectResponse
  | PageObjectResponse
  | PartialDatabaseObjectResponse
  | DatabaseObjectResponse
) & {
  content?: Array<PartialBlockObjectResponse | BlockObjectResponse>;
};
