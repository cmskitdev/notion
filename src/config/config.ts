import dotenv from "dotenv";
import fs from "fs";
import { save } from "../util/fs";

dotenv.config();

export type ConfigExcludeType = "database" | "page";

export class Config {
  databases: {
    [name: string]: {
      id: string;
      name: string;
    };
  };

  excludes: {
    pattern: string | RegExp;
    types: ConfigExcludeType[];
  }[];

  constructor(path: string) {
    const f = require(path);
    this.databases = {};
    this.excludes = f.excludes;
    for (const database of f.databases) {
      this.databases[database.name.toLowerCase()] = {
        id: database.id,
        name: database.name,
      };
    }
  }

  load?(path?: string) {
    if (path) {
      const f = require(path);
      this.databases = {};
      this.excludes = f.excludes;
      for (const database of f.databases) {
        this.databases[database.name.toLowerCase()] = {
          id: database.id,
          name: database.name,
        };
      }
    } else {
      const f = require(process.env.CONFIG ?? "./config.json");
      this.databases = {};
      for (const database of f.databases) {
        this.databases[database.name.toLowerCase()] = {
          id: database.id,
          name: database.name,
        };
      }
    }
  }

  save?(path: string) {
    save(path, this);
  }
}

export const getConfig = async (path: string): Promise<Config> => {
  if (!path) {
    throw new Error("CONFIG environment variable is not set!");
  }

  if (!fs.existsSync(`${process.cwd()}/${path}`)) {
    throw new Error(`Config file "${path}" does not exist!`);
  }

  const f = await import(`${process.cwd()}/${path}`);
  const databases: Config["databases"] = {};
  const excludes: Config["excludes"] = [];

  for (const database in f.databases) {
    databases[database] = {
      id: f.databases[database].id,
      name: f.databases[database].name,
    };
  }

  return { databases, excludes };
};

export const config = await getConfig(process.env.CONFIG!);

export const saveConfig = (config: Config) => {
  if (!process.env.CONFIG) {
    throw new Error("CONFIG environment variable is not set!");
  }

  if (!fs.existsSync(process.env.CONFIG!)) {
    throw new Error(`Config file "${process.env.CONFIG!}" does not exist!`);
  }

  const f = require(process.env.CONFIG!);

  f.databases = config.databases;

  save(process.env.CONFIG!, f);
};
