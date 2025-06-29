import { PageBase } from "@cmskit/common";
import {
  GetPageResponse,
  PageObjectResponse,
} from "@notionhq/client/build/src/api-endpoints";
import { trace } from "../util/debug";
import { Options } from "../util/options";
import { getPage } from "./pages";

export class Page<T> implements PageBase<T> {
  object: PageObjectResponse;
  options: Options;
  properties: Map<string, string | number | boolean | Date>;

  constructor(object: GetPageResponse, options?: Options) {
    const v = object as PageObjectResponse;
    this.object = v;
    this.options = options;
    this.properties = new Map(Object.entries(v.properties));
  }
  metadata: { id: string };
  frontmatter: {
    title: string;
    description: string;
    url: string;
    created: Date;
    updated: Date;
    author: string;
  };
  children: PageBase<T>[];
  content: string;

  get id() {
    return this.object.id;
  }

  async resolve(options?: Options) {
    const result = {
      page: this.id,
      children: [] as any[],
    };

    if (this.options?.relations) {
      const relationPromises = this.options.relations.flatMap((relation) => {
        const relationProperty = this.object.properties[relation];

        if (relationProperty?.type !== "relation") {
          trace(`Property ${relation} is not a relation type`, {
            type: relationProperty?.type,
          });
          return [];
        }

        trace(`resolving relations: ${relation}`, {
          length: relationProperty.relation.length,
        });

        return relationProperty.relation.map(async (rel) => {
          trace("relation", rel);
          const page = await getPage(rel.id);
          trace("page", page);
          return page;
        });
      });

      const resolvedPages = await Promise.all(relationPromises);
      result.children.push(...resolvedPages);
    }

    // debugNested(findProperties(this.object.properties, "relation"));
    return;
  }
}
