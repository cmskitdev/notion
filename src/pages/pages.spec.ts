import { inspect } from "util";
import { describe, expect, test } from "vitest";
import { search } from "../databases/databases";
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

  test(`search: istio`, async () => {
    const response = await search("Content", "minecraft");
    console.log(
      inspect(Object.keys(response), {
        depth: null,
        colors: true,
        compact: false,
      })
    );
    expect(response).toBeDefined();
    save(`pages-search-content-minecraft`, response);
    // save("search-content-istio.json", response);
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
