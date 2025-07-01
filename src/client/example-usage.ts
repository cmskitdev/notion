import { config } from "dotenv";
import path from "path";
import { WorkspaceExporter } from "./workspace-exporter";

// Load environment variables
config();

/**
 * Example usage of the WorkspaceExporter
 *
 * This example demonstrates how to export an entire Notion workspace
 * to JSON files on disk.
 */
async function main() {
  // Ensure we have a Notion token
  const notionToken = process.env.NOTION_TOKEN;
  if (!notionToken) {
    console.error("Please set NOTION_TOKEN environment variable");
    process.exit(1);
  }

  // Create output directory path
  const outputDir = path.join(
    process.cwd(),
    "notion-export",
    new Date().toISOString().split("T")[0]
  );

  // Initialize the exporter with configuration
  const exporter = new WorkspaceExporter({
    token: notionToken,
    outputDir,
    includeArchived: false, // Skip archived content
    maxDepth: 10, // Maximum depth for recursive block fetching
    includeComments: true, // Include page comments
    rateLimitDelay: 100, // Delay between API calls in ms
  });

  console.log(`Starting Notion workspace export...`);
  console.log(`Output directory: ${outputDir}`);

  try {
    // Execute the export
    const result = await exporter.export();

    // Display results
    console.log("\nExport completed successfully!");
    console.log(`\nExport Summary:`);
    console.log(`- Users exported: ${result.usersCount}`);
    console.log(`- Databases exported: ${result.databasesCount}`);
    console.log(`- Pages exported: ${result.pagesCount}`);
    console.log(`- Blocks exported: ${result.blocksCount}`);
    console.log(`- Comments exported: ${result.commentsCount}`);
    console.log(
      `- Duration: ${(result.endTime.getTime() - result.startTime.getTime()) / 1000}s`
    );

    // Display any errors encountered
    if (result.errors.length > 0) {
      console.log(`\nErrors encountered (${result.errors.length}):`);
      result.errors.forEach((error) => {
        console.error(
          `- [${error.type}${error.id ? ` ${error.id}` : ""}]: ${error.error}`
        );
      });
    }

    console.log(`\nAll data saved to: ${outputDir}`);
  } catch (error) {
    console.error("Export failed:", error);
    process.exit(1);
  }
}

// Run the example if this file is executed directly
if (import.meta.url === `file://${process.argv[1]}`) {
  main().catch(console.error);
}
