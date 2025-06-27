export const findProperties = (properties: any, type: string) => {
  return Object.keys(properties).filter(
    (key) => properties[key]?.type === type
  );
};
