const path = require("path");
const exec = require("child_process").exec;
const del = require("del");
const gulp = require("gulp");

function installUserInterface(cb) {
  const script = exec("yarn install", {
    cwd: path.resolve(__dirname, "ui"),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    cb(error);
  });
  return script;
}

function buildUserInterface(cb) {
  const script = exec("yarn build", {
    cwd: path.resolve(__dirname, "ui"),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    cb(error);
  });
  return script;
}

function removeUserInterfaceBuild() {
  return del([
    "ui/build",
  ]);
}

function removeUserInterfaceNodeModules() {
  return del([
    "ui/node_modules",
  ]);
}

function installElectron(cb) {
  const script = exec("yarn install", {
    cwd: path.resolve(__dirname, "electron"),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    cb(error);
  });
  return script;
}

function buildElectron(cb) {
  const script = exec("yarn build", {
    cwd: path.resolve(__dirname, "electron"),
  });
  script.stdout.pipe(process.stdout);
  script.stderr.pipe(process.stderr);
  script.stderr.on("error", (error) => {
    cb(error);
  });
  return script;
}

function removeElectronBuild() {
  return del([
    "electron/build",
  ]);
}

function removeElectronNodeModules() {
  return del([
    "electron/node_modules",
  ]);
}

//function javascript() {
//  const tsProject = ts.createProject("ultimate/ui/tsconfig.json");
//
//  return gulp.src("ultimate/ui/src/**/*.tsx")
//    .pipe(sourcemaps.init())
//    .pipe(tsProject())
//    .pipe(babel({
//      presets: [
//        ["@babel/preset-env", {
//          targets: {
//            "electron": "5"
//          }
//        }],
//        ["@babel/preset-react", {
//          targets: {
//            "electron": "gg5"
//          }
//        }],
//      ]
//    }))
//    .pipe(uglify())
//    .pipe(concat("bundle.min.js"))
//    .pipe(sourcemaps.write("."))
//    .pipe(gulp.dest("build/ui"));
//}
//
//function assets() {
//  return gulp.src("ultimate/ui/public/*")
//    .pipe(gulp.dest("build/ui"));
//}

const install = gulp.parallel(
  installUserInterface,
  installElectron,
)
const build = gulp.parallel(
  buildUserInterface,
  buildElectron
);
const clean = gulp.parallel(
  removeUserInterfaceBuild,
  removeElectronBuild
);
const uninstall = gulp.parallel(
  removeUserInterfaceNodeModules,
  removeElectronNodeModules
);

module.exports = {
  "install": install,
  "build": build,
  "clean": clean,
  "uninstall": uninstall,
};