import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

export const getBasePath = (...additionalPath: string[]) => {
  return path.resolve(__dirname, "../../", ...additionalPath);
};

export const save = (p: string, content: object) => {
  // const dir = path.dirname(p);
  const dir = path.dirname(p);
  console.log(dir);
  if (!fs.existsSync(dir)) {
    fs.mkdirSync(dir, {
      recursive: true,
    });
  }
  fs.writeFileSync(p, JSON.stringify(content, null, 2));
};

export const load = (p: string) => {
  if (!fs.existsSync(p)) {
    throw new Error(`file does not exist: ${p}`);
  }
  return JSON.parse(fs.readFileSync(p, "utf8"));
};
