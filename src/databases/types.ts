/**
 * Arguments for creating an Annotations instance.
 */
export type AnnotationsArgs = {
  bold: boolean;
  italic: boolean;
  strikethrough: boolean;
  underline: boolean;
  code: boolean;
  color: string;
};

/**
 * Represents annotations for a text block.
 */
export class Annotations {
  /**
   * Whether the text is bold.
   */
  bold: boolean;

  /**
   * Whether the text is italic.
   */
  italic: boolean;

  /**
   * Whether the text has a strikethrough.
   */
  strikethrough: boolean;

  /**
   * Whether the text is underlined.
   */
  underline: boolean;

  /**
   * Whether the text is formatted as code.
   */
  code: boolean;

  /**
   * The color of the text.
   */
  color: string;

  /**
   * Creates an instance of Annotations.
   * @param {AnnotationsArgs} data - The data to create the instance from.
   */
  constructor(data: AnnotationsArgs) {
    this.bold = data.bold;
    this.italic = data.italic;
    this.strikethrough = data.strikethrough;
    this.underline = data.underline;
    this.code = data.code;
    this.color = data.color;
  }
}

/**
 * Arguments for creating a CreatedBy instance.
 */
export type CreatedByArgs = {
  object: string;
  id: string;
};

/**
 * Represents the user who created an item.
 */
export class CreatedBy {
  /**
   * The type of object, typically "user".
   */
  object: string;

  /**
   * The ID of the user.
   */
  id: string;

  /**
   * Creates an instance of CreatedBy.
   * @param {CreatedByArgs} data - The data to create the instance from.
   */
  constructor(data: CreatedByArgs) {
    this.object = data.object;
    this.id = data.id;
  }
}

/**
 * Arguments for creating a Database instance.
 */
export type DatabaseArgs = {
  object: string;
  id: string;
  cover: { type: string; external: ExternalArgs } | null;
  icon: IconArgs | null;
  created_time: string;
  created_by: CreatedByArgs;
  last_edited_by: CreatedByArgs;
  last_edited_time: string;
  title: TitleArgs[];
  description: DescriptionArgs[];
  is_inline: boolean;
  properties: PropertiesArgs;
  parent: ParentArgs;
  url: string;
  public_url: string | null;
  archived: boolean;
  in_trash: boolean;
};

/**
 * Represents a Notion database.
 */
export class Database {
  /**
   * The type of object, typically "database".
   */
  object: string;

  /**
   * The ID of the database.
   */
  id: string;

  /**
   * The cover image of the database, if any.
   */
  cover: { type: string; external: External } | null;

  /**
   * The icon of the database, if any.
   */
  icon: Icon | null;

  /**
   * The time the database was created.
   */
  createdTime: Date;

  /**
   * The user who created the database.
   */
  createdBy: CreatedBy;

  /**
   * The user who last edited the database.
   */
  lastEditedBy: CreatedBy;

  /**
   * The time the database was last edited.
   */
  lastEditedTime: Date;

  /**
   * The title of the database.
   */
  title: Title[];

  /**
   * The description of the database.
   */
  description: Description[];

  /**
   * Whether the database is inline.
   */
  isInline: boolean;

  /**
   * The properties of the database.
   */
  properties: Properties;

  /**
   * The parent of the database.
   */
  parent: Parent;

  /**
   * The URL of the database.
   */
  url: string;

  /**
   * The public URL of the database, if any.
   */
  publicUrl: string | null;

  /**
   * Whether the database is archived.
   */
  archived: boolean;

  /**
   * Whether the database is in the trash.
   */
  inTrash: boolean;

  /**
   * Creates an instance of Database.
   * @param {DatabaseArgs} data - The data to create the instance from.
   */
  constructor(data: DatabaseArgs) {
    this.object = data.object;
    this.id = data.id;
    this.cover = data.cover
      ? { type: data.cover.type, external: new External(data.cover.external) }
      : null;
    this.icon = data.icon ? new Icon(data.icon) : null;
    this.createdTime = new Date(data.created_time);
    this.createdBy = new CreatedBy(data.created_by);
    this.lastEditedBy = new CreatedBy(data.last_edited_by);
    this.lastEditedTime = new Date(data.last_edited_time);
    this.title = data.title.map((t) => new Title(t));
    this.description = data.description.map((d) => new Description(d));
    this.isInline = data.is_inline;
    this.properties = new Properties(data.properties);
    this.parent = new Parent(data.parent);
    this.url = data.url;
    this.publicUrl = data.public_url;
    this.archived = data.archived;
    this.inTrash = data.in_trash;
  }
}

/**
 * Arguments for creating a DatabaseEntry instance.
 */
export type DatabaseEntryArgs = {
  id: string;
  name?: string;
  database: DatabaseArgs;
};

/**
 * Represents a single entry in the list of databases.
 */
export class DatabaseEntry {
  /**
   * The ID of the database entry.
   */
  id: string;

  /**
   * The name of the database.
   */
  name?: string;

  /**
   * The database object.
   */
  database: Database;

  /**
   * Creates an instance of DatabaseEntry.
   * @param {DatabaseEntryArgs} data - The data to create the instance from.
   */
  constructor(data: DatabaseEntryArgs) {
    this.id = data.id;
    this.name = data.name;
    this.database = new Database(data.database);
  }
}

/**
 * Arguments for creating a DateProperty instance.
 */
export type DatePropertyArgs = {};

/**
 * Represents a date property.
 */
export class DateProperty {
  /**
   * Creates an instance of DateProperty.
   * @param {DatePropertyArgs} data - The data to create the instance from.
   */
  constructor(data: DatePropertyArgs) {}
}

/**
 * Arguments for creating a Description instance.
 */
export type DescriptionArgs = {
  type: string;
  text: TextArgs;
  annotations: AnnotationsArgs;
  plain_text: string;
  href: string | null;
};

/**
 * Represents a description element.
 */
export class Description {
  /**
   * The type of the description, typically "text".
   */
  type: string;

  /**
   * The text content of the description.
   */
  text: Text;

  /**
   * Annotations for the description text.
   */
  annotations: Annotations;

  /**
   * The plain text version of the description.
   */
  plainText: string;

  /**
   * A URL link associated with the description, if any.
   */
  href: string | null;

  /**
   * Creates an instance of Description.
   * @param {DescriptionArgs} data - The data to create the instance from.
   */
  constructor(data: DescriptionArgs) {
    this.type = data.type;
    this.text = new Text(data.text);
    this.annotations = new Annotations(data.annotations);
    this.plainText = data.plain_text;
    this.href = data.href;
  }
}

/**
 * Arguments for creating a DualProperty instance.
 */
export type DualPropertyArgs = {
  synced_property_name: string;
  synced_property_id: string;
};

/**
 * Represents a dual property in a relation.
 */
export class DualProperty {
  /**
   * The name of the synced property.
   */
  syncedPropertyName: string;

  /**
   * The ID of the synced property.
   */
  syncedPropertyId: string;

  /**
   * Creates an instance of DualProperty.
   * @param {DualPropertyArgs} data - The data to create the instance from.
   */
  constructor(data: DualPropertyArgs) {
    this.syncedPropertyName = data.synced_property_name;
    this.syncedPropertyId = data.synced_property_id;
  }
}

/**
 * Arguments for creating an External instance.
 */
export type ExternalArgs = {
  url: string;
};

/**
 * Represents an external link.
 */
export class External {
  /**
   * The URL of the external resource.
   */
  url: string;

  /**
   * Creates an instance of External.
   * @param {ExternalArgs} data - The data to create the instance from.
   */
  constructor(data: ExternalArgs) {
    this.url = data.url;
  }
}

/**
 * Arguments for creating a FilesProperty instance.
 */
export type FilesPropertyArgs = {};

/**
 * Represents a files property.
 */
export class FilesProperty {
  /**
   * Creates an instance of FilesProperty.
   * @param {FilesPropertyArgs} data - The data to create the instance from.
   */
  constructor(data: FilesPropertyArgs) {}
}

/**
 * Arguments for creating a Group instance.
 */
export type GroupArgs = {
  id: string;
  name: string;
  color: string;
  option_ids: string[];
};

/**
 * Represents a group of options in a status property.
 */
export class Group {
  /**
   * The ID of the group.
   */
  id: string;

  /**
   * The name of the group.
   */
  name: string;

  /**
   * The color of the group.
   */
  color: string;

  /**
   * The IDs of the options belonging to this group.
   */
  optionIds: string[];

  /**
   * Creates an instance of Group.
   * @param {GroupArgs} data - The data to create the instance from.
   */
  constructor(data: GroupArgs) {
    this.id = data.id;
    this.name = data.name;
    this.color = data.color;
    this.optionIds = data.option_ids;
  }
}

/**
 * Arguments for creating an Icon instance.
 */
export type IconArgs = {
  type: string;
  external?: ExternalArgs;
  emoji?: string;
};

/**
 * Represents an icon, which can be external or an emoji.
 */
export class Icon {
  /**
   * The type of the icon.
   */
  type: string;

  /**
   * The external icon details, if applicable.
   */
  external?: External;

  /**
   * The emoji character, if applicable.
   */
  emoji?: string;

  /**
   * Creates an instance of Icon.
   * @param {IconArgs} data - The data to create the instance from.
   */
  constructor(data: IconArgs) {
    this.type = data.type;
    if (data.external) {
      this.external = new External(data.external);
    }
    this.emoji = data.emoji;
  }
}

/**
 * Arguments for creating a MultiSelect instance.
 */
export type MultiSelectArgs = {
  options: OptionArgs[];
};

/**
 * Represents a multi-select property.
 */
export class MultiSelect {
  /**
   * The available options for the multi-select property.
   */
  options: Option[];

  /**
   * Creates an instance of MultiSelect.
   * @param {MultiSelectArgs} data - The data to create the instance from.
   */
  constructor(data: MultiSelectArgs) {
    this.options = data.options.map((option) => new Option(option));
  }
}

/**
 * Arguments for creating a NumberProperty instance.
 */
export type NumberPropertyArgs = {
  format: string;
};

/**
 * Represents a number property.
 */
export class NumberProperty {
  /**
   * The format of the number.
   */
  format: string;

  /**
   * Creates an instance of NumberProperty.
   * @param {NumberPropertyArgs} data - The data to create the instance from.
   */
  constructor(data: NumberPropertyArgs) {
    this.format = data.format;
  }
}

/**
 * Arguments for creating an Option instance.
 */
export type OptionArgs = {
  id: string;
  name: string;
  color: string;
  description: string | null;
};

/**
 * Represents an option in a select or multi-select property.
 */
export class Option {
  /**
   * The ID of the option.
   */
  id: string;

  /**
   * The name of the option.
   */
  name: string;

  /**
   * The color of the option.
   */
  color: string;

  /**
   * The description of the option, if any.
   */
  description: string | null;

  /**
   * Creates an instance of Option.
   * @param {OptionArgs} data - The data to create the instance from.
   */
  constructor(data: OptionArgs) {
    this.id = data.id;
    this.name = data.name;
    this.color = data.color;
    this.description = data.description;
  }
}

/**
 * Arguments for creating a Parent instance.
 */
export type ParentArgs = {
  type: string;
  workspace?: boolean;
  page_id?: string;
};

/**
 * Represents the parent of a database or page.
 */
export class Parent {
  /**
   * The type of the parent.
   */
  type: string;

  /**
   * Whether the parent is the workspace.
   */
  workspace?: boolean;

  /**
   * The ID of the parent page, if applicable.
   */
  pageId?: string;

  /**
   * Creates an instance of Parent.
   * @param {ParentArgs} data - The data to create the instance from.
   */
  constructor(data: ParentArgs) {
    this.type = data.type;
    this.workspace = data.workspace;
    this.pageId = data.page_id;
  }
}

/**
 * Arguments for creating a Properties instance.
 */
export type PropertiesArgs = {
  [key: string]: PropertyArgs;
};

/**
 * Represents a collection of properties in a database.
 */
export class Properties {
  [key: string]: Property;

  /**
   * Creates an instance of Properties.
   * @param {PropertiesArgs} data - The data to create the instance from.
   */
  constructor(data: PropertiesArgs) {
    for (const key in data) {
      if (Object.prototype.hasOwnProperty.call(data, key)) {
        this[key] = new Property(data[key]);
      }
    }
  }
}

/**
 * Arguments for creating a Property instance.
 */
export type PropertyArgs = {
  id: string;
  name: string;
  type: string;
  date?: DatePropertyArgs;
  relation?: RelationArgs;
  select?: SelectArgs;
  status?: StatusArgs;
  number?: NumberPropertyArgs;
  url?: UrlPropertyArgs;
  files?: FilesPropertyArgs;
  multi_select?: MultiSelectArgs;
  rich_text?: RichTextPropertyArgs;
  title?: TitlePropertyArgs;
  created_by?: {};
  created_time?: {};
};

/**
 * Represents a single property in a database.
 */
export class Property {
  /**
   * The ID of the property.
   */
  id: string;

  /**
   * The name of the property.
   */
  name: string;

  /**
   * The type of the property.
   */
  type: string;

  /**
   * The date details, if the property is a date type.
   */
  date?: DateProperty;

  /**
   * The relation details, if the property is a relation type.
   */
  relation?: Relation;

  /**
   * The select details, if the property is a select type.
   */
  select?: Select;

  /**
   * The status details, if the property is a status type.
   */
  status?: Status;

  /**
   * The number details, if the property is a number type.
   */
  number?: NumberProperty;

  /**
   * The URL details, if the property is a URL type.
   */
  url?: UrlProperty;

  /**
   * The files details, if the property is a files type.
   */
  files?: FilesProperty;

  /**
   * The multi-select details, if the property is a multi-select type.
   */
  multiSelect?: MultiSelect;

  /**
   * The rich text details, if the property is a rich text type.
   */
  richText?: RichTextProperty;

  /**
   * The title details, if the property is a title type.
   */
  title?: TitleProperty;

  /**
   * The created by details, if the property is a created_by type.
   */
  createdBy?: {};

  /**
   * The created time details, if the property is a created_time type.
   */
  createdTime?: {};

  /**
   * Creates an instance of Property.
   * @param {PropertyArgs} data - The data to create the instance from.
   */
  constructor(data: PropertyArgs) {
    this.id = data.id;
    this.name = data.name;
    this.type = data.type;
    if (data.date) {
      this.date = new DateProperty(data.date);
    }
    if (data.relation) {
      this.relation = new Relation(data.relation);
    }
    if (data.select) {
      this.select = new Select(data.select);
    }
    if (data.status) {
      this.status = new Status(data.status);
    }
    if (data.number) {
      this.number = new NumberProperty(data.number);
    }
    if (data.url) {
      this.url = new UrlProperty(data.url);
    }
    if (data.files) {
      this.files = new FilesProperty(data.files);
    }
    if (data.multi_select) {
      this.multiSelect = new MultiSelect(data.multi_select);
    }
    if (data.rich_text) {
      this.richText = new RichTextProperty(data.rich_text);
    }
    if (data.title) {
      this.title = new TitleProperty(data.title);
    }
    this.createdBy = data.created_by;
    this.createdTime = data.created_time;
  }
}

/**
 * Arguments for creating a Relation instance.
 */
export type RelationArgs = {
  database_id: string;
  type: string;
  dual_property: DualPropertyArgs;
};

/**
 * Represents a relation property.
 */
export class Relation {
  /**
   * The ID of the related database.
   */
  databaseId: string;

  /**
   * The type of the relation.
   */
  type: string;

  /**
   * The dual property details of the relation.
   */
  dualProperty: DualProperty;

  /**
   * Creates an instance of Relation.
   * @param {RelationArgs} data - The data to create the instance from.
   */
  constructor(data: RelationArgs) {
    this.databaseId = data.database_id;
    this.type = data.type;
    this.dualProperty = new DualProperty(data.dual_property);
  }
}

/**
 * Arguments for creating a RichTextProperty instance.
 */
export type RichTextPropertyArgs = {};

/**
 * Represents a rich text property.
 */
export class RichTextProperty {
  /**
   * Creates an instance of RichTextProperty.
   * @param {RichTextPropertyArgs} data - The data to create the instance from.
   */
  constructor(data: RichTextPropertyArgs) {}
}

/**
 * Arguments for creating a Select instance.
 */
export type SelectArgs = {
  options: OptionArgs[];
};

/**
 * Represents a select property.
 */
export class Select {
  /**
   * The available options for the select property.
   */
  options: Option[];

  /**
   * Creates an instance of Select.
   * @param {SelectArgs} data - The data to create the instance from.
   */
  constructor(data: SelectArgs) {
    this.options = data.options.map((option) => new Option(option));
  }
}

/**
 * Arguments for creating a Status instance.
 */
export type StatusArgs = {
  options: OptionArgs[];
  groups: GroupArgs[];
};

/**
 * Represents a status property.
 */
export class Status {
  /**
   * The available options for the status property.
   */
  options: Option[];

  /**
   * The groups of options for the status property.
   */
  groups: Group[];

  /**
   * Creates an instance of Status.
   * @param {StatusArgs} data - The data to create the instance from.
   */
  constructor(data: StatusArgs) {
    this.options = data.options.map((option) => new Option(option));
    this.groups = data.groups.map((group) => new Group(group));
  }
}

/**
 * Arguments for creating a Text instance.
 */
export type TextArgs = {
  content: string;
  link: string | null;
};

/**
 * Represents a text block.
 */
export class Text {
  /**
   * The content of the text block.
   */
  content: string;

  /**
   * A link associated with the text, if any.
   */
  link: string | null;

  /**
   * Creates an instance of Text.
   * @param {TextArgs} data - The data to create the instance from.
   */
  constructor(data: TextArgs) {
    this.content = data.content;
    this.link = data.link;
  }
}

/**
 * Arguments for creating a Title instance.
 */
export type TitleArgs = {
  type: string;
  text: TextArgs;
  annotations: AnnotationsArgs;
  plain_text: string;
  href: string | null;
};

/**
 * Represents a title element.
 */
export class Title {
  /**
   * The type of the title, typically "text".
   */
  type: string;

  /**
   * The text content of the title.
   */
  text: Text;

  /**
   * Annotations for the title text.
   */
  annotations: Annotations;

  /**
   * The plain text version of the title.
   */
  plainText: string;

  /**
   * A URL link associated with the title, if any.
   */
  href: string | null;

  /**
   * Creates an instance of Title.
   * @param {TitleArgs} data - The data to create the instance from.
   */
  constructor(data: TitleArgs) {
    this.type = data.type;
    this.text = new Text(data.text);
    this.annotations = new Annotations(data.annotations);
    this.plainText = data.plain_text;
    this.href = data.href;
  }
}

/**
 * Arguments for creating a TitleProperty instance.
 */
export type TitlePropertyArgs = {};

/**
 * Represents a title property.
 */
export class TitleProperty {
  /**
   * Creates an instance of TitleProperty.
   * @param {TitlePropertyArgs} data - The data to create the instance from.
   */
  constructor(data: TitlePropertyArgs) {}
}

/**
 * Arguments for creating a UrlProperty instance.
 */
export type UrlPropertyArgs = {};

/**
 * Represents a URL property.
 */
export class UrlProperty {
  /**
   * Creates an instance of UrlProperty.
   * @param {UrlPropertyArgs} data - The data to create the instance from.
   */
  constructor(data: UrlPropertyArgs) {}
}
