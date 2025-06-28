import fs from "fs";
import path from "path";

export const save = (filename: string, content: object) => {
  const f = `${new Date().toISOString()}-${filename}.json`;
  const tmpDir = path.resolve(__dirname, "../tmp");
  console.log(tmpDir);
  console.log(`saving ${f} to ${path.resolve(tmpDir, f)}..`);
  if (!fs.existsSync(tmpDir)) {
    console.log(`creating tmp directory: ${tmpDir}..`);
    fs.mkdirSync(tmpDir);
  }
  console.log(`writing ${f} to ${tmpDir} (${path.resolve(tmpDir, f)})..`);
  fs.writeFileSync(path.resolve(tmpDir, f), JSON.stringify(content, null, 2));
};

export const load = (filename: string) => {
  const tmpDir = path.resolve(__dirname, "../tmp");
  if (!fs.existsSync(tmpDir)) {
    throw new Error(`tmp directory does not exist: ${tmpDir}`);
  }
  return JSON.parse(fs.readFileSync(path.resolve(tmpDir, filename), "utf8"));
};
