import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react-swc'
import tsconfigPaths from 'vite-tsconfig-paths'
import { viteSingleFile } from 'vite-plugin-singlefile'
import { execSync } from 'child_process'

const commitHash = execSync('git rev-parse --short HEAD').toString()
const commitCount = execSync('git rev-list --all --count').toString()
const commitDate = execSync('git log -1 --format="%at" | xargs -I{} date -d @{} +%Y.%m.%d_%H:%M:%S').toString()

// https://vitejs.dev/config/
export default defineConfig({
  define: {
    __COMMIT_HASH: JSON.stringify(commitHash),
    __COMMIT_NUMBER: JSON.stringify(commitCount),
    __COMMIT_DATE: JSON.stringify(commitDate),
  },
  plugins: [react(), tsconfigPaths(), viteSingleFile()],
})
