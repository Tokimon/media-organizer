import type { Snippet } from "svelte";

export type Stylable = {
  class?: string;
  style?: string;
};

export type BaseComponent = Stylable & {
  children: Snippet;
};
