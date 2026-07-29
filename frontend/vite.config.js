import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  root: "./frontend",
  base: "/",
  publicDir: "public",
  build: {
    rollupOptions: {
      input: "/workspaces/wikispeedia/frontend/src/main.tsx",
      output: {
        entryFileNames: `assets/[name].js`,
        chunkFileNames: `assets/[name].js`,
        assetFileNames: `assets/[name].[ext]`,
      },
    },
    outDir: "../dist/frontend",
  },
  plugins: [react()],
});
