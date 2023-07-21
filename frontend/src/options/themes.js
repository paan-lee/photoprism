import { $gettext, T } from "common/vm";

/* Theme Styles */

let themes = {
  default: {
    dark: true,
    sponsor: false,
    title: "Default",
    name: "default",
    colors: {
      application: "#2f3031",
      form: "#2f3031",
      card: "#232425",
      primary: "#9E7BEA",
      "primary-button": "#5F1DB7",
      "secondary-dark": "#7E4FE3",
      secondary: "#1c1d1e",
      "secondary-light": "#252627",
      accent: "#333",
      error: "#e57373",
      info: "#00acc1",
      success: "#4db6ac",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#4A464F",
      navigation: "#141417",
      "navigation-home": "#0e0f10",
    },
  },
  grayscale: {
    name: "grayscale",
    dark: true,
    sponsor: false,
    colors: {
      application: "#525252",
      form: "#525252",
      card: "#5e5e5e",
      primary: "#c8bdb1",
      "primary-button": "#726e69",
      "secondary-dark": "#c8bdb1",
      secondary: "#444",
      "secondary-light": "#5E5E5E",
      accent: "#333",
      error: "#e57373",
      info: "#5a94dd",
      success: "#26A69A",
      warning: "#e3d181",
      love: "#ef5350",
      remove: "#e35333",
      restore: "#64b5f6",
      album: "#ffab40",
      download: "#07bd9f",
      private: "#48bcd6",
      edit: "#48bcd6",
      share: "#0070a0",
      terminal: "#333333",
      navigation: "#353839",
      "navigation-home": "#212121",
    },
  },
  lavender: {
    name: "lavender",
    dark: false,
    sponsor: false,
    colors: {
      application: "#fafafa",
      form: "#fafafa",
      card: "#DFE0E8",
      primary: "#9ca2c9",
      "primary-button": "#6c6f84",
      "secondary-dark": "#475185",
      secondary: "#dee0ed",
      "secondary-light": "#eef0f6",
      accent: "#8c8c8c",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#333333",
      navigation: "#1b1e32",
      "navigation-home": "#121421",
    },
  },
  legacy: {
    name: "legacy",
    dark: false,
    sponsor: false,
    colors: {
      application: "#F5F5F5",
      form: "#F5F5F5",
      card: "#e0e0e0",
      primary: "#FFCA28",
      "primary-button": "#212121",
      "secondary-dark": "#212121",
      secondary: "#bdbdbd",
      "secondary-light": "#e0e0e0",
      accent: "#757575",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#333333",
      navigation: "#212121",
      "navigation-home": "#000000",
    },
  },
  mint: {
    dark: true,
    sponsor: false,
    title: "Mint",
    name: "mint",
    colors: {
      application: "#121212",
      form: "#121212",
      card: "#1e1e1e",
      primary: "#2bb14c",
      "primary-button": "#22903d",
      "secondary-dark": "#2bb14c",
      secondary: "#181818",
      "secondary-light": "#1f1f1f",
      accent: "#2bb14c",
      error: "#e57373",
      info: "#00acc1",
      success: "#4db6ac",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#181818",
      navigation: "#181818",
      "navigation-home": "#181818",
    },
  },
  nordic: {
    dark: false,
    sponsor: false,
    title: "Nordic",
    name: "nordic",
    colors: {
      application: "#f7f8fa",
      form: "#f7f8fa",
      card: "#ECEFF4",
      primary: "#4ca0b8",
      "primary-button": "#519fb6",
      "secondary-dark": "#4ca0b8",
      secondary: "#e2e7ee",
      "secondary-light": "#eceff4",
      accent: "#81A1C1",
      error: "#BF616A",
      info: "#88C0D0",
      success: "#8FBCBB",
      warning: "#f0d8a8",
      remove: "#BF616A",
      restore: "#81A1C1",
      album: "#EBCB8B",
      download: "#8FBCBB",
      private: "#88C0D0",
      edit: "#88C0D0",
      share: "#B48EAD",
      love: "#ef5350",
      terminal: "#4C566A",
      navigation: "#e7ebf1",
      "navigation-home": "#dde3eb",
    },
  },
  onyx: {
    name: "onyx",
    dark: false,
    sponsor: false,
    colors: {
      application: "#e5e4e2",
      form: "#e5e4e2",
      card: "#cdccca",
      primary: "#c8bdb1",
      "primary-button": "#353839",
      "secondary-dark": "#353839",
      secondary: "#a8a8a8",
      "secondary-light": "#cdccca",
      accent: "#656565",
      error: "#e57373",
      info: "#5a94dd",
      success: "#26A69A",
      warning: "#e3d181",
      love: "#ef5350",
      remove: "#e35333",
      restore: "#64b5f6",
      album: "#ffab40",
      download: "#07bd9f",
      private: "#48bcd6",
      edit: "#48bcd6",
      share: "#0070a0",
      terminal: "#333333",
      navigation: "#353839",
      "navigation-home": "#212121",
    },
  },
  abyss: {
    name: "abyss",
    dark: true,
    sponsor: true,
    colors: {
      application: "#202020",
      form: "#202020",
      card: "#242424",
      primary: "#814fd9",
      "primary-button": "#7e57c2",
      "secondary-dark": "#814fd9",
      secondary: "#111111",
      "secondary-light": "#1a1a1a",
      accent: "#090c10",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#9575cd",
      restore: "#64b5f6",
      album: "#7e57c2",
      download: "#673ab7",
      private: "#512da8",
      edit: "#4527a0",
      share: "#311b92",
      love: "#ef5350",
      terminal: "#333333",
      navigation: "#0d0d0d",
      "navigation-home": "#000000",
    },
  },
  carbon: {
    dark: true,
    sponsor: true,
    title: "Carbon",
    name: "carbon",
    colors: {
      application: "#16141c",
      form: "#16141c",
      card: "#292732",
      primary: "#8a6eff",
      "primary-button": "#53478a",
      "secondary-dark": "#7f63fd",
      secondary: "#0E0D12",
      "secondary-light": "#292733",
      accent: "#262238",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#292733",
      navigation: "#0E0D12",
      "navigation-home": "#0E0D12",
    },
  },
  gemstone: {
    name: "gemstone",
    dark: true,
    sponsor: true,
    colors: {
      application: "#2f2f31",
      form: "#2f2f31",
      card: "#2b2b2d",
      primary: "#AFB4D4",
      "primary-button": "#545465",
      "secondary-dark": "#9BA0C5",
      secondary: "#272727",
      "secondary-light": "#37373a",
      accent: "#333",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#4A464F",
      navigation: "#1C1C21",
      "navigation-home": "#131316",
    },
  },
  neon: {
    name: "neon",
    dark: true,
    sponsor: true,
    colors: {
      application: "#242326",
      form: "#242326",
      card: "#1b1a1c",
      primary: "#f44abf",
      "primary-button": "#890664",
      "secondary-dark": "#cc0d99",
      secondary: "#111111",
      "secondary-light": "#1a1a1a",
      accent: "#090c10",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#fece3e",
      love: "#fb4483",
      remove: "#9100a0",
      restore: "#5e33f8",
      album: "#6234b5",
      download: "#8d56eb",
      private: "#4749c8",
      edit: "#5658eb",
      share: "#5692eb",
      terminal: "#333333",
      navigation: "#0e0d0f",
      "navigation-home": "#000000",
    },
  },
  shadow: {
    name: "shadow",
    dark: true,
    sponsor: true,
    colors: {
      application: "#444",
      form: "#444",
      card: "#666666",
      primary: "#c4f1e5",
      "primary-button": "#74817d",
      "secondary-dark": "#c8e3e7",
      secondary: "#585858",
      "secondary-light": "#666",
      accent: "#333",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#333333",
      navigation: "#212121",
      "navigation-home": "#000000",
    },
  },
  vanta: {
    name: "vanta",
    dark: true,
    sponsor: true,
    colors: {
      application: "#212121",
      form: "#212121",
      card: "#1d1d1d",
      primary: "#04acaf",
      "primary-button": "#444444",
      "secondary-dark": "#04acaf",
      secondary: "#111111",
      "secondary-light": "#1a1a1a",
      accent: "#090c10",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#333333",
      navigation: "#0d0d0d",
      "navigation-home": "#000000",
    },
  },
  yellowstone: {
    name: "yellowstone",
    dark: true,
    sponsor: true,
    colors: {
      application: "#32312f",
      form: "#32312f",
      card: "#262524",
      primary: "#ffb700",
      "primary-button": "#54524e",
      "secondary-dark": "#ffb700",
      secondary: "#21201f",
      "secondary-light": "#262523",
      accent: "#333",
      error: "#e57373",
      info: "#00acc1",
      success: "#26A69A",
      warning: "#ffd740",
      remove: "#e57373",
      restore: "#64b5f6",
      album: "#ffab00",
      download: "#00bfa5",
      private: "#00b8d4",
      edit: "#00b8d4",
      share: "#9575cd",
      love: "#ef5350",
      terminal: "#464544",
      navigation: "#191817",
      "navigation-home": "#0c0c0b",
    },
  },
};

/* Available Themes */

let options = [
  {
    text: $gettext("Default"),
    value: "default",
    disabled: false,
  },
  {
    text: "Grayscale",
    value: "grayscale",
    disabled: false,
  },
  {
    text: "Lavender",
    value: "lavender",
    disabled: false,
  },
  {
    text: "Legacy",
    value: "legacy",
    disabled: false,
  },
  {
    text: "Mint",
    value: "mint",
    disabled: false,
  },
  {
    text: "Nordic",
    value: "nordic",
    disabled: false,
  },
  {
    text: "Onyx",
    value: "onyx",
    disabled: false,
  },
  {
    text: "Abyss",
    value: "abyss",
    disabled: false,
  },
  {
    text: "Carbon",
    value: "carbon",
    disabled: false,
  },
  {
    text: "Gemstone",
    value: "gemstone",
    disabled: false,
  },
  {
    text: "Neon",
    value: "neon",
    disabled: false,
  },
  {
    text: "Shadow",
    value: "shadow",
    disabled: false,
  },
  {
    text: "Vanta",
    value: "vanta",
    disabled: false,
  },
  {
    text: "Yellowstone",
    value: "yellowstone",
    disabled: false,
  },
];

/* Theme Functions */

// Returns a theme by name.
export const Get = (name) => {
  if (typeof themes[name] === "undefined") {
    return themes[options[0].value];
  }

  return themes[name];
};

// Adds or replaces a theme by name.
export const Set = (name, val) => {
  if (typeof themes[name] === "undefined") {
    options.push({
      text: val.title,
      value: val.name,
      disabled: false,
    });
  }

  themes[name] = val;
};

// Removes a theme by name.
export const Remove = (name) => {
  delete themes[name];
  const i = options.findIndex((el) => el.value === name);
  if (i > -1) {
    options.splice(i, 1);
  }
};

// Returns translated theme options.
export const Translated = () => {
  return options.map((v) => {
    if (v.disabled) {
      return null;
    }

    return {
      text: T(v.text),
      value: v.value,
    };
  });
};

export const Options = () => options;

export const SetOptions = (v) => (options = v);
