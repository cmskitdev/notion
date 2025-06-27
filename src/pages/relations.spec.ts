import { describe, test } from "vitest";
import { trace } from "../util/debug";
import { getPage } from "./pages";

describe("relations", () => {
  // test("getPagesWithNestedSeries", async () => {
  //   const response = await getPagesWithNestedSeries("Content");
  //   // console.log(
  //   //   inspect(response, { depth: null, colors: true, compact: false })
  //   // );
  //   expect(response).toBeDefined();
  //   expect(response.results).toBeDefined();

  //   // Check if any pages have the Series property nested
  //   const pageWithSeries = response.results.find(
  //     (page: any) => page.properties?.Series?.type === "relation"
  //   );

  //   if (pageWithSeries) {
  //     expect(pageWithSeries.properties.Series).toHaveProperty("nested_pages");
  //     debug("Found page with nested Series", {
  //       pageWithSeries: pageWithSeries.properties.Series.length,
  //     });
  //   }
  // });

  test("getPageWithNestedSeries - specific page", async () => {
    // Using the page ID from the attached query.json
    const pageId = "1e8d7342-e571-8032-8d1e-de076836a018";

    const page = await getPage(pageId, {
      relations: ["Content Links"],
    });

    trace("page", await page.resolve());

    // debugNested(response);
    // console.log(
    //   inspect(response, { depth: null, colors: true, compact: false })
    // );
    // expect(response).toBeDefined();

    // if ((response as any).properties?.Series) {
    //   expect((response as any).properties.Series).toHaveProperty(
    //     "type",
    //     "relation"
    //   );
    //   // debug("Page Series property", {
    //   //   series: (response as any).properties.Series,
    //   // });
    // }
  });
});

//   test("getPagesWithAllNestedRelations", async () => {
//     const response = await getPagesWithAllNestedRelations("Content");
//     // console.log(
//     //   inspect(response, { depth: null, colors: true, compact: false })
//     // );
//     expect(response).toBeDefined();
//     expect(response.results).toBeDefined();

//     // Check if any pages have relation properties nested
//     const pageWithRelations = response.results.find((page: any) => {
//       const relationProps = Object.keys(page.properties || {}).filter(
//         (key) => page.properties[key]?.type === "relation"
//       );
//       return relationProps.length > 0;
//     });

//     if (pageWithRelations) {
//       debug("Found page with nested relations", {
//         relationProperties: Object.keys(pageWithRelations.properties).filter(
//           (key) => pageWithRelations.properties[key]?.type === "relation"
//         ),
//       });
//     }
//   });
// });

// /**
//  * Example demonstrating how to get and nest the Series relation property.
//  */
// const demonstrateSeriesNesting = async () => {
//   console.log("=== Demonstrating Series Relation Nesting ===\n");

//   try {
//     // Example 1: Get all pages from Content database with nested Series
//     console.log("1. Getting all pages with nested Series relations...");
//     const pagesWithSeries = await getPagesWithNestedSeries("Content");

//     console.log(`Found ${pagesWithSeries.results.length} pages`);

//     // Show the first page with a Series relation
//     const pageWithSeries = pagesWithSeries.results.find(
//       (page: any) => page.properties?.Series?.relation?.length > 0
//     );

//     if (pageWithSeries) {
//       console.log("\nFirst page with Series relation:");
//       // console.log(
//       //   inspect(pageWithSeries.properties.Series, {
//       //     depth: null,
//       //     colors: true,
//       //     compact: false,
//       //   })
//       // );
//     } else {
//       console.log("No pages found with Series relations");
//     }

//     // Example 2: Get a specific page with nested Series
//     console.log("\n2. Getting specific page with nested Series...");
//     const specificPageId = "1e5d7342-e571-8062-a13d-c535135f5c15"; // From your query.json
//     const specificPage = await getPageWithNestedSeries(specificPageId);

//     if ((specificPage as any).properties?.Series) {
//       console.log("\nSpecific page Series property:");
//       // console.log(
//       //   inspect((specificPage as any).properties.Series, {
//       //     depth: null,
//       //     colors: true,
//       //     compact: false,
//       //   })
//       // );
//     }

//     // Example 3: Manual nesting of a single property
//     console.log("\n3. Manual nesting example...");
//     const manualExample = {
//       Series: {
//         id: "%3FxsZ",
//         type: "relation",
//         relation: [{ id: "some-page-id-1" }, { id: "some-page-id-2" }],
//         has_more: false,
//       },
//     };

//     const nestedManualExample = await nestRelationProperty(
//       manualExample,
//       "Series"
//     );
//     console.log("\nManually nested Series property:");
//     // console.log(
//     //   inspect(nestedManualExample.Series, {
//     //     depth: null,
//     //     colors: true,
//     //     compact: false,
//     //   })
//     // );
//   } catch (error) {
//     console.error("Error demonstrating Series nesting:", error);
//   }
