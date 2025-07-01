import { inspect } from "util";
import { test } from "vitest";
import { SearchClient } from "./search/client";

test("search", async () => {
  try {
    const response = await new SearchClient().search("Content");
    console.log(inspect(response, { depth: null, colors: true }));
    // console.log(inspect(body, { depth: null, colors: true }));
    // console.log("Response body:", body);
    // expect(response).toBeDefined();
  } catch (error) {
    console.error(error);
  }
});
