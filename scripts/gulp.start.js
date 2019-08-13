const path = require("path");
const exec = require("child_process").exec;

function startUI() {
  const script = exec("yarn start", {
    cwd: path.resolve(process.cwd(), "ui"),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    done(error);
  });
  return script;
}

function startElectron() {
  const script = exec("yarn start", {
    cwd: path.resolve(process.cwd(), "electron"),
    env: {
      "NODE_ENV": "development",
    },
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    done(error);
  });
  return script;
}

function startGateway() {
  const cmd = `"services/gateway/build/gateway"`;
  const script = exec(cmd, {
    cwd: path.resolve(process.cwd()),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    done(error);
  });
  return script;
}

function startNyaa() {
  const cmd = `"services/nyaa/build/nyaa"`;
  const script = exec(cmd, {
    cwd: path.resolve(process.cwd()),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    done(error);
  });
  return script;
}

function startMyAnimeList() {
  const cmd = `"services/myanimelist/build/myanimelist"`;
  const script = exec(cmd, {
    cwd: path.resolve(process.cwd()),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    done(error);
  });
  return script;
}

function startYoutubeDL() {
  const cmd = `"services/youtube-dl/dist/youtube-dl/youtube-dl"`;
  const script = exec(cmd, {
    cwd: path.resolve(process.cwd()),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    done(error);
  });
  return script;
}

function startAtlas() {
  const cmd = `"services/atlas/build/atlas"`;
  const script = exec(cmd, {
    cwd: path.resolve(process.cwd()),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    done(error);
  });
  return script;
}

module.exports = {
  startUI,
  startElectron,
  startGateway,
  startNyaa,
  startMyAnimeList,
  startYoutubeDL,
  startAtlas,
};