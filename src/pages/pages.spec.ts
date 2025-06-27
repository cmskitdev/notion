import { inspect } from "util";
import { describe, expect, test } from "vitest";
import { config } from "../config/config";
import { debug } from "../util/debug";
import { getPages } from "./pages";
import {
  getPagesWithNestedSeries,
  getPageWithNestedSeries,
  nestRelationProperty,
} from "./relations";

describe("pages", () => {
  debug("config", config);
  // for (const database of Object.values(config.databases)) {
  test(`getPages`, async () => {
    const response = await getPages("Series");
    console.log(
      inspect(response, { depth: null, colors: true, compact: false })
    );
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

// Get pages with nested Series relations
const pages = await getPagesWithNestedSeries("Content");

// Get a specific page with nested Series
const page = await getPageWithNestedSeries("page-id");

// Manually nest a relation property
const nestedProps = await nestRelationProperty(properties, "Series");
