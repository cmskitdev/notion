#!/usr/bin/env node

import { Client } from "@notionhq/client";
import { promises as fs } from "fs";
import path from "path";
import { parentPort, workerData } from "worker_threads";

interface WorkerData {
  token: string;
  blockIds: string[];
  outputDir: string;
  maxDepth: number;
  rateLimitDelay: number;
}

interface WorkerResult {
  processedBlocks: number;
  errors: Array<{ blockId: string; error: string }>;
}

async function collectPaginatedAPI<
  T extends { next_cursor: string | null; results: any[] },
>(
  listFn: (args: any) => Promise<T>,
  firstPageArgs: any
): Promise<T["results"]> {
  const results: T["results"] = [];
  let nextCursor: string | undefined = firstPageArgs.start_cursor;

  do {
    const response = await listFn({
      ...firstPageArgs,
      start_cursor: nextCursor,
    });
    results.push(...response.results);
    nextCursor = response.next_cursor ?? undefined;
  } while (nextCursor);

  return results;
}

async function processBlocks(): Promise<WorkerResult> {
  const { token, blockIds, outputDir, maxDepth, rateLimitDelay }: WorkerData =
    workerData;

  const client = new Client({ auth: token });
  const exportedBlockIds = new Set<string>();
  const errors: Array<{ blockId: string; error: string }> = [];

  async function exportBlocks(blockId: string, depth: number): Promise<void> {
    if (depth >= maxDepth || exportedBlockIds.has(blockId)) {
      return;
    }

    exportedBlockIds.add(blockId);

    // Rate limiting
    await new Promise((resolve) => setTimeout(resolve, rateLimitDelay));

    try {
      const blocks = await collectPaginatedAPI(
        client.blocks.children.list.bind(client.blocks.children),
        { block_id: blockId }
      );

      // Save blocks
      await fs.writeFile(
        path.join(outputDir, "blocks", `${blockId}-children.json`),
        JSON.stringify(blocks, null, 2)
      );

      // Recursively export child blocks
      for (const block of blocks) {
        if (
          block.object === "block" &&
          "has_children" in block &&
          block.has_children
        ) {
          await exportBlocks(block.id, depth + 1);
        }
      }
    } catch (error) {
      errors.push({
        blockId,
        error: error instanceof Error ? error.message : String(error),
      });
    }
  }

  // Process all block IDs
  await Promise.all(blockIds.map((blockId) => exportBlocks(blockId, 0)));

  return {
    processedBlocks: exportedBlockIds.size,
    errors,
  };
}

// Run the worker
processBlocks()
  .then((result) => {
    parentPort?.postMessage({ success: true, result });
  })
  .catch((error) => {
    parentPort?.postMessage({
      success: false,
      error: error instanceof Error ? error.message : String(error),
    });
  });
