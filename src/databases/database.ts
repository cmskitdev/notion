import { DatabaseObjectResponse } from "@notionhq/client/build/src/api-endpoints";

export class Database {
  public id: string;
  public name: string;

  constructor(private readonly database: DatabaseObjectResponse) {
    this.id = database.id;
    if (!database.in_trash) {
      if (database.is_inline) {
        // console.log(database.parent);
        // getPage(database.parent.page_id).then((page) => {
        //   console.log(page);
        // });
      } else {
        this.name = database.title[0].plain_text;
      }
    }
  }
}
