import dotenv from "dotenv";
import { describe, expect, test } from "vitest";
import { getConfig } from "./config";
dotenv.config();

describe("config", () => {
  test(`getConfig`, async () => {
    const response = await getConfig(process.env.CONFIG!);
    console.log(response);
    expect(response).toBeDefined();
  });

  // test("save", async () => {
  //   const config = new Config(process.env.CONFIG!);
  //   config.save(process.env.CONFIG!);
  // });
});
