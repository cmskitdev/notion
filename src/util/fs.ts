import fs from "fs";

export const save = (path: string, content: object) => {
  fs.writeFileSync(path, JSON.stringify(content, null, 2));
};

export const load = (path: string) => {
  return JSON.parse(fs.readFileSync(path, "utf8"));
};
