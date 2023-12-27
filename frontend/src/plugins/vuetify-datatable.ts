// This type exists at https://github.com/vuetifyjs/vuetify/blob/master/packages/vuetify/src/labs/VDataTable/composables/sort.ts#L29
export type SortItem = { key: string; order?: boolean | "asc" | "desc" };

// This is the un-reffed version of what is passed into useOptions() at https://github.com/vuetifyjs/vuetify/blob/master/packages/vuetify/src/labs/VDataTable/composables/options.ts#L9
export type Options = {
  page: number;
  itemsPerPage: number;
  sortBy: readonly SortItem[];
  groupBy: readonly SortItem[];
  search: string | undefined;
};
