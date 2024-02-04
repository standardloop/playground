"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.createRedisInstance = void 0;
const ioredis_1 = __importDefault(require("ioredis"));
const configuration = {
    redis: {
        host: process.env.REDIS_HOST,
        password: process.env.REDIS_PASSWORD,
        port: process.env.REDIS_PORT,
    },
};
function getRedisConfiguration() {
    return configuration.redis;
}
function createRedisInstance() {
    let redis;
    try {
        redis = new ioredis_1.default();
        return redis;
    }
    catch (e) {
        return null;
    }
}
exports.createRedisInstance = createRedisInstance;
//     config = getRedisConfiguration()
// ) {
//     try {
//         const options: RedisOptions = {
//             host: config.host,
//             lazyConnect: true,
//             showFriendlyErrorStack: true,
//             enableAutoPipelining: true,
//             maxRetriesPerRequest: 3,
//             retryStrategy: (times: number) => {
//                 if (times > 3) {
//                     throw new Error(`[Redis] Could not connect after ${times} attempts`);
//                 }
//                 return Math.min(times * 200, 1000);
//             },
//         };
//         if (config.port) {
//             options.port = Number(config.port);
//         }
//         if (config.password) {
//             options.password = config.password;
//         }
//         const redis = new Redis(options);
//         redis.on('error', (error: unknown) => {
//             console.warn('[Redis] Error connecting', error);
//         });
//         return redis;
//     } catch (e) {
//         throw new Error(`[Redis] Could not create a Redis instance`);
//     }
// }
