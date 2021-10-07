module.exports = function (config, env) {
    config.output = {
        ...config.output,
        filename: '[name].js',
        chunkFilename: '[name].js',
    };
    config.optimization.splitChunks = false;
    config.optimization.runtimeChunk = false;
    return config;
};