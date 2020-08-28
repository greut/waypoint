import { helper } from '@ember/component/helper';

const wordBreak = /(?:^|\s|-|\/)\S/g;
const replace = /_+|-+/g;

// componentName
export function componentName([component]: [string]): string {
  // Split into words based on common characters and uppercase
  // the first letter of each word
  let result = component.toLowerCase().replace(wordBreak, function (m) {
    return m.toUpperCase();
  });

  // Replace any separators that are not human readable
  return result.replace(replace, ' ');
}

export default helper(componentName);
