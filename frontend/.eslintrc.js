module.exports = {
    plugins: ['react', 'prettier'],
    extends: ['eslint:recommended', 'plugin:react/recommended'],
    parserOptions: {
        ecmaVersion: 2017,
        sourceType: 'module',
    },
    env: { browser: true, node: true, es6: true },
    globals: {
        WebSocket: true,
    },
    rules: {
        'prettier/prettier': 'error',
        semi: ['error', 'never'],
        quotes: ['error', 'single'],
        'no-debugger': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'no-console': process.env.NODE_ENV === 'production' ? 'error' : 'off',
        'react/prop-types': [0],
    },
}
