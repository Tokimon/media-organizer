import type { IconNames } from "$lib/icons/icons";
import fileStore from "$lib/stores/file-store.svelte.js";

export default (ext: string): IconNames => {
  const { Image, Video, Audio, Book, Code, Data, Settings, Vector } =
    fileStore.extensions;

  if (Image?.Web.includes(ext) || Image?.Other.includes(ext))
    return "image_square";
  if (Image?.Camera.includes(ext)) return "camera_roll";
  if (Video?.includes(ext)) return "play_circle";
  if (Audio?.includes(ext)) return "volume_up";
  if (Vector?.includes(ext)) return "vector";
  if (Book?.includes(ext)) return "book";
  if (Code?.includes(ext)) return "code";
  if (Data?.includes(ext)) return "data";
  if (Settings?.includes(ext)) return "settings_square";

  return "unknown_square";
};
