import fs from "fs";
import path from "path";

export const save = (filename: string, content: object) => {
  const tmpDir = path.resolve(__dirname, "../../tmp");
  if (!fs.existsSync(tmpDir)) {
    fs.mkdirSync(tmpDir);
  }
  fs.writeFileSync(
    path.resolve(tmpDir, `${new Date().toUTCString()}-${filename}.json`),
    JSON.stringify(content, null, 2)
  );
};

export const load = (filename: string) => {
  const tmpDir = path.resolve(__dirname, "../../tmp");
  if (!fs.existsSync(tmpDir)) {
    throw new Error(`tmp directory does not exist: ${tmpDir}`);
  }
  return JSON.parse(fs.readFileSync(path.resolve(tmpDir, filename), "utf8"));
};
