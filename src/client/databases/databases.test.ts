import { test } from "vitest";
import { DatabasesClient } from "./client";

test("get database", async () => {
  try {
    const response = await new DatabasesClient().list();
    // console.log(inspect(response, { depth: null, colors: true }));
    // console.log(inspect(body, { depth: null, colors: true }));
    // console.log("Response body:", body);
    // expect(response).toBeDefined();
  } catch (error) {
    console.error(error);
  }
});
