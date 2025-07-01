export const toSafeName = (str: string): string => {
  return (
    str
      .toLowerCase()
      // Replace spaces with hyphens
      .replace(/\s+/g, "-")
      // Remove special characters and punctuation
      .replace(/[^a-z0-9-]/g, "")
      // Remove multiple consecutive hyphens
      .replace(/-+/g, "-")
      // Remove leading/trailing hyphens
      .replace(/^-+|-+$/g, "")
  );
};
