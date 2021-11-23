const redisRStream = require("redis-rstream");

module.exports = (app) => {
  return async (req, res, next) => {
    const key = `${req.params.username}/${req.params.endUrl}`;
    res.header("Access-Control-Allow-Origin", "*");

    if (key.endsWith("m3u8")) {
      res.header("Content-Type", "application/vnd.apple.mpegurl");
      const m3u8 = await app.redisClient.get(key).catch((e) => {
        console.error(e);
        return null;
      });
      if (m3u8) return res.status(200).send(m3u8);
      else return res.status(400).end();
    } else if (key.endsWith("ts")) {
      res.header("Content-Type", "video/mp2t");
    }

    const exists = await app.redisClient.exists(key).catch((e) => {
      console.error(e);
      return null;
    });

    if (exists) redisRStream(app.redisClient, key, { highWaterMark: 512 * 1024 }).pipe(res);
    else res.status(404).end();
  };
};
