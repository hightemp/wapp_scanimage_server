var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
import { defineConfig, loadEnv } from 'vite';
import vue from '@vitejs/plugin-vue';
export default defineConfig(function (_a) {
    var mode = _a.mode;
    // Load env from parent directory (project root) and current directory
    var rootEnv = loadEnv(mode, '../', '');
    var localEnv = loadEnv(mode, './', '');
    // Merge envs, local takes priority
    var env = __assign(__assign({}, rootEnv), localEnv);
    var serverPort = parseInt(env.VITE_PORT || env.SERVER_PORT || '3000');
    var backendPort = parseInt(env.VITE_BACKEND_PORT || env.SERVER_PORT || '8080');
    var backendHost = env.VITE_BACKEND_HOST || env.SERVER_HOST || 'localhost';
    var backendUrl = "http://".concat(backendHost, ":").concat(backendPort);
    return {
        plugins: [vue()],
        server: {
            port: serverPort,
            host: env.VITE_HOST || '0.0.0.0',
            proxy: {
                '/api': {
                    target: backendUrl,
                    changeOrigin: true
                },
                '/files': {
                    target: backendUrl,
                    changeOrigin: true
                }
            }
        },
        build: {
            outDir: '../cmd/server/dist',
            emptyOutDir: true
        }
    };
});
