const config = require("../config/config.json");
const { createClient } = require("redis");
const express = require("express");
const app = express();
app.disable("x-powered-by");
app.listen(config.port, async () => {
  console.log(`Angelthump REDIS Replication listening on port ${config.port}!`);

  app.redisClient = createClient({
    socket: config.redis.useUnixSocket
      ? {
          path: config.redis.unix,
        }
      : {
          host: config.redis.hostname,
        },
  });

  await app.redisClient
    .connect()
    .then(() => {
      console.info("Redis client connected.");
    })
    .catch((e) => console.error(e));
});
const cache = require("./cache");

app.get("/hls/:username/:endUrl", cache(app));
