import svelte from 'rollup-plugin-svelte';
import commonjs from '@rollup/plugin-commonjs';
import resolve from '@rollup/plugin-node-resolve';
import livereload from 'rollup-plugin-livereload';
import { terser } from 'rollup-plugin-terser';
import styles from "rollup-plugin-styles";

const production = !process.env.ROLLUP_WATCH;

function serve() {
    let server;

    function toExit() {
        if (server) server.kill(0);
    }

    return {
        writeBundle() {
            if (server) return;
            server = require('child_process').spawn('npm', ['run', 'start', '--', '--dev'], {
                stdio: ['ignore', 'inherit', 'inherit'],
                shell: true
            });

            process.on('SIGTERM', toExit);
            process.on('exit', toExit);
        }
    };
}

export default {
    input: 'src/index.js',
    output: {
        sourcemap: !production,
        format: 'iife',
        file: 'public/app.js'
    },
    plugins: [
        svelte({ compilerOptions: { dev: !production } }),
        styles(),
        resolve({ browser: true, dedupe: ['svelte'] }),
        commonjs(),
        !production && serve(),
        !production && livereload('public'),
        production && terser()
    ],
    watch: {
        clearScreen: false
    }
};
