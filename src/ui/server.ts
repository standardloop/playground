import { parse } from "url";
import next from "next";
import express from "express";
import { createRedisInstance } from "./lib/redis"


const port = parseInt(process.env.PORT || "3000", 10);
const dev = process.env.NODE_ENV !== "production";
const app = next({ dev });
const handle = app.getRequestHandler();
const redis = createRedisInstance();

app.prepare().then(() => {
    const server = express();
    server.get("/_next/*", (req, res) => {
        handle(req, res);
    });

    server.get('*', async (req, res) => {
        if (redis !== null) {
            console.log("will use redis path")
            await renderAndCache(req, res);
        } else {
            await handle(req, res);
        }
    });
    server.listen(port, (err) => {
        if (err) throw err;
        console.log(`> Ready http://localhost:${port}`)
    })
});


function getCacheKey(req) {
    return `${req.path}`
}

async function renderAndCache(req, res) {
    const key = getCacheKey(req);
    const pageForYou = await redis.get(key);

    const parsedUrl = parse(req.url, true)
    const { pathname, query } = parsedUrl

    if (pageForYou) {
        console.log(`serving from cache ${key}`);
        res.setHeader('x-cache', 'HIT');
        res.send(pageForYou);
        return
    } else {
        console.log(`key ${key} not found, rendering`);
        const html = await app.renderToHTML(req, res, pathname, query);
        const MAX_AGE = 10000; // 10 second
        const EXPIRY_MS = `PX`; // milliseconds
        await redis.set(key, html, EXPIRY_MS, MAX_AGE);
        res.setHeader('x-cache', 'MISS');
        res.end(html);
    }
}
