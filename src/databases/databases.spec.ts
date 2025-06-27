import { describe, expect, test } from "vitest";
import { config } from "../config/config";
import { debug } from "../util/debug";
import { getDatabases } from "./databases";

describe("databases", () => {
  debug("config", config);
  // for (const database of Object.values(config.databases)) {
  test(`getDatabases`, async () => {
    const response = await getDatabases();
    // console.log(
    //   inspect(response, { depth: null, colors: true, compact: false })
    // );
    expect(response).toBeDefined();
  });
  // }

  // test(`search: istio`, async () => {
  //   const response = await search("content", "minecraft");
  //   // console.log(response);
  //   // console.log(
  //   //   inspect(response, { depth: null, colors: true, compact: false })
  //   // );
  //   expect(response).toBeDefined();
  // });
});
