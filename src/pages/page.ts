import {
  GetPageResponse,
  PageObjectResponse,
} from "@notionhq/client/build/src/api-endpoints";
import { trace } from "../util/debug";
import { Options } from "../util/options";
import { getPage } from "./pages";

export class Page {
  object: PageObjectResponse;
  options: Options;

  constructor(object: GetPageResponse, options?: Options) {
    this.object = object as PageObjectResponse;
    this.options = options;
  }

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
