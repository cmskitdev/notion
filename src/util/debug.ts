import { inspect } from "util";

const colors = {
  reset: "\x1b[0m",
  bright: "\x1b[1m",
  dim: "\x1b[2m",
  red: "\x1b[31m",
  green: "\x1b[32m",
  yellow: "\x1b[33m",
  blue: "\x1b[34m",
  magenta: "\x1b[35m",
  cyan: "\x1b[36m",
  white: "\x1b[37m",
  gray: "\x1b[90m",
};

export const debug = (msg: any, ...args: any[]) => {
  console.log(
    msg,
    inspect(args, {
      depth: null,
      colors: true,
      compact: false,
    })
  );
};

export const trace = (label: string, data: any) => {
  console.log(`${colors.bright}${colors.cyan}🔍 ${label}${colors.reset}`, data);
};

/**
 * Recursively prints nested arrays and objects with indentation and colors.
 *
 * @param data - The data structure to print (arrays, objects, primitives)
 * @param label - Optional label to display before the data
 * @param indent - Current indentation level (used internally for recursion)
 */
export const debugNested = (
  data: any,
  label?: string,
  indent: number = 0
): void => {
  const indentStr = "  ".repeat(indent);

  // Print label if provided
  if (label && indent === 0) {
    console.log(`${colors.bright}${colors.cyan}🔍 ${label}${colors.reset}`);
  }

  // Handle null and undefined
  if (data === null) {
    console.log(`${indentStr}${colors.gray}null${colors.reset}`);
    return;
  }
  if (data === undefined) {
    console.log(`${indentStr}${colors.gray}undefined${colors.reset}`);
    return;
  }

  // Handle primitives
  if (typeof data === "string") {
    console.log(`${indentStr}${colors.green}"${data}"${colors.reset}`);
    return;
  }
  if (typeof data === "number") {
    console.log(`${indentStr}${colors.yellow}${data}${colors.reset}`);
    return;
  }
  if (typeof data === "boolean") {
    console.log(`${indentStr}${colors.magenta}${data}${colors.reset}`);
    return;
  }

  // Handle arrays
  if (Array.isArray(data)) {
    if (data.length === 0) {
      console.log(`${indentStr}${colors.gray}[]${colors.reset}`);
      return;
    }

    console.log(
      `${indentStr}${colors.blue}[${colors.reset} ${colors.dim}(${data.length} items)${colors.reset}`
    );
    data.forEach((item, index) => {
      console.log(`${indentStr}  ${colors.gray}${index}:${colors.reset}`);
      debugNested(item, undefined, indent + 2);
    });
    console.log(`${indentStr}${colors.blue}]${colors.reset}`);
    return;
  }

  // Handle objects
  if (typeof data === "object") {
    const keys = Object.keys(data);
    if (keys.length === 0) {
      console.log(`${indentStr}${colors.gray}{}${colors.reset}`);
      return;
    }

    console.log(
      `${indentStr}${colors.cyan}{${colors.reset} ${colors.dim}(${keys.length} properties)${colors.reset}`
    );
    keys.forEach((key) => {
      console.log(`${indentStr}  ${colors.white}${key}:${colors.reset}`);
      debugNested(data[key], undefined, indent + 2);
    });
    console.log(`${indentStr}${colors.cyan}}${colors.reset}`);
    return;
  }

  // Fallback for other types
  console.log(
    `${indentStr}${colors.red}${typeof data}: ${String(data)}${colors.reset}`
  );
};
