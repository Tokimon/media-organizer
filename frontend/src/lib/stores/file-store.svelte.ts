import type { api } from "$lib/wailsjs/go/models";
import { SvelteMap } from "svelte/reactivity";

const files = new SvelteMap<string, api.ApiFile>();

export default {
  files,

  addFiles(filesToAdd: api.ApiFile | api.ApiFile[]) {
    const fileArray = Array.isArray(filesToAdd) ? filesToAdd : [filesToAdd];

    for (const file of fileArray) files.set(file.path, file);
  },

  removeFiles(paths: string | string[]) {
    const pathArray = Array.isArray(paths) ? paths : [paths];

    for (const path of pathArray) files.delete(path);
  },

  removeDirectory(dirPath: string) {
    for (const file of files.values())
      if (file.path.startsWith(dirPath)) files.delete(file.path);
  },
};
