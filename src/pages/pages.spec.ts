import { beforeAll, describe, expect, test } from "vitest";
import { search, SearchResult } from "../databases/databases";
import { save } from "../util/fs";

describe("pages", () => {
  // test(`getPages`, async () => {
  //   const pages = await getPages("Series");

  //   expect(pages).toBeDefined();
  //   expect(pages.length).toBeGreaterThan(0);

  //   for (const page of pages) {
  //     expect(page).toBeInstanceOf(Page);
  //   }
  // });

  let artifact: SearchResult;

  beforeAll(async () => {
    artifact = await search("Content", "minecraft");
  });

  test(`search: istio`, async () => {
    expect(artifact).toBeDefined();
    save(`pages-search-content-minecraft`, artifact);
  });

  test("resolve relational properties", async () => {
    for (const page of artifact.results) {
      const resolvedPage = await resolveAllRelations(page);
      expect(resolvedPage).toBeDefined();
      save(`pages-search-content-minecraft-resolved`, resolvedPage);
    }
  });

  // test(`getPagesWithNestedSeries`, async () => {
  //   const pages = await getPagesWithNestedSeries("Content");
  //   expect(pages).toBeDefined();
  // });

  // test(`getPageWithNestedSeries`, async () => {
  //   const page = await getPageWithNestedSeries(
  //     "1e5d7342-e571-804b-9fd2-d97fd12ae897",
  //     ["Series"]
  //   );
  //   expect(page).toBeDefined();
  // });
});
