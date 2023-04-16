/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export type Page = List | Detail | Form;
export type Action =
  | {
      /**
       * The type of the action.
       */
      type: "copy";
      /**
       * The title of the action.
       */
      title?: string;
      /**
       * The text to copy.
       */
      text: string;
      /**
       * The inputs to show when the action is run.
       */
      inputs?: Input[];
      /**
       * The key used as a shortcut.
       */
      key?: string;
    }
  | {
      /**
       * The type of the action.
       */
      type: "open";
      /**
       * The title of the action.
       */
      title?: string;
      /**
       * The inputs to show when the action is run.
       */
      inputs?: Input[];
      /**
       * The key used as a shortcut.
       */
      key?: string;
      /**
       * The target to open.
       */
      target: string;
    }
  | {
      /**
       * The type of the action.
       */
      type: "exit";
      /**
       * The title of the action.
       */
      title?: string;
      /**
       * The key used as a shortcut.
       */
      key?: string;
    }
  | {
      /**
       * The type of the action.
       */
      type: "reload";
      /**
       * The title of the action.
       */
      title?: string;
      /**
       * The key used as a shortcut.
       */
      key?: string;
    }
  | {
      /**
       * The type of the action.
       */
      type: "run";
      /**
       * The title of the action.
       */
      title?: string;
      /**
       * The inputs to show when the action is run.
       */
      inputs?: Input[];
      /**
       * The key used as a shortcut.
       */
      key?: string;
      command: Command;
      onSuccess?: "reload" | "copy" | "open" | "exit" | "push";
    }
  | {
      /**
       * The type of the action.
       */
      type: "push";
      /**
       * The title of the action.
       */
      title?: string;
      /**
       * The key used as a shortcut.
       */
      key?: string;
      /**
       * The inputs to show when the action is run.
       */
      inputs?: Input[];
      page: string;
    };
export type Input =
  | {
      /**
       * The name of the input.
       */
      name: string;
      /**
       * The title of the input.
       */
      title: string;
      /**
       * The type of the input.
       */
      type: "textfield";
      /**
       * The placeholder of the input.
       */
      placeholder?: string;
      /**
       * The default value of the input.
       */
      default?: string;
      /**
       * Whether the input should be secure.
       */
      secure?: boolean;
    }
  | {
      /**
       * The name of the input.
       */
      name: string;
      /**
       * The title of the input.
       */
      title: string;
      /**
       * The type of the input.
       */
      type: "checkbox";
      /**
       * The default value of the input.
       */
      default?: boolean;
      /**
       * The label of the input.
       */
      label?: string;
      /**
       * The text substitution to use when the input is true.
       */
      trueSubstitution?: string;
      /**
       * The text substitution to use when the input is false.
       */
      falseSubstitution?: string;
    }
  | {
      /**
       * The name of the input.
       */
      name: string;
      /**
       * The title of the input.
       */
      title: string;
      /**
       * The type of the input.
       */
      type: "textarea";
      /**
       * The placeholder of the input.
       */
      placeholder?: string;
      /**
       * The default value of the input.
       */
      default?: string;
    }
  | {
      /**
       * The name of the input.
       */
      name: string;
      /**
       * The title of the input.
       */
      title: string;
      /**
       * The type of the input.
       */
      type: "dropdown";
      /**
       * The items of the input.
       */
      items: {
        /**
         * The title of the item.
         */
        title: string;
        /**
         * The value of the item.
         */
        value: string;
      }[];
      /**
       * The default value of the input.
       */
      default?: string;
    };
export type Command =
  | string
  | [string, ...string[]]
  | {
      args?: string[];
      input?: string;
      dir?: string;
    };
/**
 * The preview to show in the detail view.
 */
export type Preview =
  | {
      type: "static";
      /**
       * The text to show in the preview.
       */
      text: string;
      /**
       * The language of the preview text.
       */
      language?: string;
    }
  | {
      type: "dynamic";
      /**
       * The language of the preview text.
       */
      language?: string;
      command: Command;
    };

export interface List {
  /**
   * The type of the response.
   */
  type: "list";
  /**
   * The title of the page.
   */
  title?: string;
  emptyView?: {
    /**
     * The text to show when the list is empty.
     */
    text: string;
    /**
     * The actions to show when the list is empty.
     */
    actions?: Action[];
  };
  /**
   * Whether to show the preview on the right side of the list.
   */
  showPreview?: boolean;
  /**
   * The items in the list.
   */
  items?: Listitem[];
}
export interface Listitem {
  /**
   * The title of the item.
   */
  title: string;
  /**
   * The id of the item.
   */
  id?: string;
  /**
   * The subtitle of the item.
   */
  subtitle?: string;
  preview?: Preview;
  /**
   * The accessories to show on the right side of the item.
   */
  accessories?: string[];
  /**
   * The actions attached to the item.
   */
  actions?: Action[];
}
/**
 * A detail view displayign a preview and actions.
 */
export interface Detail {
  /**
   * The type of the response.
   */
  type: "detail";
  /**
   * The title of the page.
   */
  title?: string;
  preview: Preview;
  /**
   * The actions attached to the detail view.
   */
  actions?: Action[];
}
export interface Form {
  /**
   * The type of the response.
   */
  type: "form";
  /**
   * The title of the page.
   */
  title?: string;
  submitAction: Action;
}