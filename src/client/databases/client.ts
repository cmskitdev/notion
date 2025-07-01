import { Configuration, DatabasesApi } from "@mateothegreat/notion-openapi";

export class DatabasesClient {
  private readonly client = new DatabasesApi(
    new Configuration({
      basePath: "https://api.notion.com",
      accessToken: "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",
      headers: {
        "Notion-Version": "2022-06-28",
        "Content-Type": "application/json",
      },
    })
  );

  async list(): Promise<void> {
    const response = await this.client.retrieveADatabase(
      "16ad7342-e571-80c4-a065-c7a1015871d3"
    );
    console.log(response);
    // const body = await response.raw.json();
    // console.log(inspect(body, { depth: null, colors: true }));
    // return body;
  }
}
