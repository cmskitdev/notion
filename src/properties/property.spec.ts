import { expect, test } from "vitest";
import { PropertyMultiSelect } from "./multi-select";
import { Properties } from "./properties";

const fixtures = {
  page: require("../../test/fixtures/page.json"),
};

test("PropertyMultiSelect", () => {
  const properties = new Properties(fixtures.page.properties);
  expect(properties.size).toBeGreaterThan(0);
  expect(properties.get("Tags")).toBeInstanceOf(PropertyMultiSelect);
});
