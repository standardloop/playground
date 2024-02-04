"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const url_1 = require("url");
const next_1 = __importDefault(require("next"));
const express_1 = __importDefault(require("express"));
const redis_1 = require("./lib/redis");
const port = parseInt(process.env.PORT || "3000", 10);
const dev = process.env.NODE_ENV !== "production";
const app = (0, next_1.default)({ dev });
const handle = app.getRequestHandler();
const redis = (0, redis_1.createRedisInstance)();
app.prepare().then(() => {
    const server = (0, express_1.default)();
    server.get("/_next/*", (req, res) => {
        handle(req, res);
    });
    server.get('*', async (req, res) => {
        if (redis !== null) {
            console.log("will use redis path");
            await renderAndCache(req, res);
        }
        else {
            await handle(req, res);
        }
    });
    server.listen(port, (err) => {
        if (err)
            throw err;
        console.log(`> Ready http://localhost:${port}`);
    });
});
function getCacheKey(req) {
    return `${req.path}`;
}
async function renderAndCache(req, res) {
    const key = getCacheKey(req);
    const pageForYou = await redis.get(key);
    const parsedUrl = (0, url_1.parse)(req.url, true);
    const { pathname, query } = parsedUrl;
    if (pageForYou) {
        console.log(`serving from cache ${key}`);
        res.setHeader('x-cache', 'HIT');
        res.send(pageForYou);
        return;
    }
    else {
        console.log(`key ${key} not found, rendering`);
        const html = await app.renderToHTML(req, res, pathname, query);
        const MAX_AGE = 10000; // 10 second
        const EXPIRY_MS = `PX`; // milliseconds
        await redis.set(key, html, EXPIRY_MS, MAX_AGE);
        res.setHeader('x-cache', 'MISS');
        res.end(html);
    }
}
