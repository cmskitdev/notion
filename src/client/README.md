# Notion Workspace Exporter

A comprehensive TypeScript library for exporting entire Notion workspaces to JSON files, with full support for databases, pages, blocks, users, and comments.

## Features

- **Complete Workspace Export**: Export all content from your Notion workspace
- **Recursive Block Fetching**: Automatically fetches nested blocks up to a configurable depth
- **Type-Safe**: Full TypeScript support with comprehensive type definitions
- **Error Handling**: Graceful error handling with detailed error reporting
- **Rate Limiting**: Built-in rate limiting to respect Notion API limits
- **Configurable**: Flexible options for archived content, comments, and more
- **100% Test Coverage**: Comprehensive test suite using Vitest

## Installation

```bash
npm install @cmskit/notion
```

## Usage

### Basic Usage

```typescript
import { WorkspaceExporter } from '@cmskit/notion';

const exporter = new WorkspaceExporter({
  token: process.env.NOTION_TOKEN,
  outputDir: './notion-export',
});

const result = await exporter.export();
console.log(`Exported ${result.pagesCount} pages and ${result.blocksCount} blocks`);
```

### Advanced Configuration

```typescript
const exporter = new WorkspaceExporter({
  token: process.env.NOTION_TOKEN,
  outputDir: './notion-export',
  includeArchived: true,        // Include archived pages and databases
  maxDepth: 15,                 // Maximum depth for recursive block fetching
  includeComments: true,        // Export page comments
  rateLimitDelay: 200,          // Milliseconds between API calls
});
```

### Export Result

The export method returns a detailed result object:

```typescript
interface ExportResult {
  usersCount: number;
  databasesCount: number;
  pagesCount: number;
  blocksCount: number;
  commentsCount: number;
  startTime: Date;
  endTime: Date;
  errors: Array<{
    type: string;
    id?: string;
    error: string;
  }>;
}
```

## Output Structure

The exporter creates the following directory structure:

```
output-dir/
├── users/
│   ├── all-users.json
│   ├── user-1.json
│   └── user-2.json
├── databases/
│   ├── database-1.json
│   └── database-2.json
├── pages/
│   ├── page-1.json
│   └── page-2.json
├── blocks/
│   ├── page-1-children.json
│   └── block-1-children.json
└── comments/
    └── page-1-comments.json
```

## API Documentation

### `WorkspaceExporter`

The main class for exporting Notion workspaces.

#### Constructor

```typescript
constructor(config: WorkspaceExporterConfig)
```

#### Configuration Options

| Option | Type | Default | Description |
|--------|------|---------|-------------|
| `token` | `string` | Required | Notion API integration token |
| `outputDir` | `string` | Required | Directory where exported data will be saved |
| `includeArchived` | `boolean` | `false` | Whether to include archived pages and databases |
| `maxDepth` | `number` | `10` | Maximum depth for recursive block fetching |
| `includeComments` | `boolean` | `true` | Whether to include comments on pages |
| `rateLimitDelay` | `number` | `100` | Delay in milliseconds between API calls |

#### Methods

##### `export(): Promise<ExportResult>`

Exports the entire workspace. Returns a promise that resolves with export statistics and any errors encountered.

## Error Handling

The exporter handles errors gracefully and continues processing when possible. All errors are collected and returned in the result object:

```typescript
const result = await exporter.export();

if (result.errors.length > 0) {
  console.log('Errors encountered:');
  result.errors.forEach(error => {
    console.error(`[${error.type}] ${error.id}: ${error.error}`);
  });
}
```

## Rate Limiting

The exporter includes built-in rate limiting to avoid hitting Notion API limits. The default delay is 100ms between API calls, but this can be configured:

```typescript
const exporter = new WorkspaceExporter({
  token: process.env.NOTION_TOKEN,
  outputDir: './notion-export',
  rateLimitDelay: 200, // 200ms between API calls
});
```

## Testing

The library includes comprehensive tests with 100% code coverage:

```bash
# Run tests
npm test

# Run tests with coverage
npm run test:coverage

# Run tests in watch mode
npm run test:watch
```

## Examples

### Export with Error Handling

```typescript
import { WorkspaceExporter } from '@cmskit/notion';

async function exportWorkspace() {
  const exporter = new WorkspaceExporter({
    token: process.env.NOTION_TOKEN,
    outputDir: './exports/notion-' + Date.now(),
  });

  try {
    const result = await exporter.export();
    
    console.log('Export completed:', {
      pages: result.pagesCount,
      databases: result.databasesCount,
      duration: (result.endTime.getTime() - result.startTime.getTime()) / 1000 + 's',
    });

    if (result.errors.length > 0) {
      console.warn(`Completed with ${result.errors.length} errors`);
    }
  } catch (error) {
    console.error('Export failed:', error);
    process.exit(1);
  }
}
```

### Selective Export

```typescript
// Export only non-archived content without comments
const exporter = new WorkspaceExporter({
  token: process.env.NOTION_TOKEN,
  outputDir: './notion-export-minimal',
  includeArchived: false,
  includeComments: false,
});
```

## Limitations

- The Notion API has rate limits. The exporter includes delays to respect these limits.
- Very large workspaces may take considerable time to export.
- Binary content (files, images) URLs are exported but the actual files are not downloaded.

## Contributing

Contributions are welcome! Please ensure all tests pass and maintain 100% code coverage.

## License

MIT 