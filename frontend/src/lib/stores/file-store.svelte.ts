import { GetExtensions, GetFiles } from "$lib/wailsjs/go/api/Api";
import type { api, tools } from "$lib/wailsjs/go/models";
import { SvelteMap } from "svelte/reactivity";

class FileStore {
  files = new SvelteMap<string, api.ApiFile>();
  extensions = $state<Partial<tools.Extensions>>({});

  async update(updatedFileList: Record<string, api.ApiFile>) {
    const currPaths = new Set(this.files.keys());

    // Update entries in the list
    for (const file of Object.values(updatedFileList)) {
      this.files.set(file.path, file);
      currPaths.delete(file.path);
    }

    // Remove any paths that were not in the updated file list
    for (const path of currPaths) this.files.delete(path);
  }

  async init() {
    const [ext, files] = await Promise.all([GetExtensions(), GetFiles()]);
    this.update(files);
    this.extensions = ext;
  }
}

export default new FileStore();
